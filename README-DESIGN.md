# Design decisions

## Goals and assumptions

* Aim to lift common Go scalar operations to vectors
  * Including those from standard libraries
* A simple API that does not expose the complexities of the underlying CPU architecture
* Should not require understanding of the underlying CPU vector instructions or be aware of vector size limits
* Sympathetic to hardware
  * Support for common hardware vector lengths - currently small powers of 2
  * Assume better performance from mixing nearby lanes than far-away lanes, and best performance from operating within lane
* Open to extension in future
  * Adding new instructions
  * Adding new vector lengths
* API should be free of side effects
* Favour returning a result over passing in a destination pointer
* Focus on vector lengths that have been implemented in hardware - currently up to 512 bits (e.g. AVX512, Fugaku supercomputer implements Arm SVE at 512 bits)
* Assume Go compiler can emulate longer vectors on vector hardware by ganging registers or looping
  * Aim for straightforward integration of hardware instructions with larger/wider operations that require fallback implementations in Go
* Assume Go compiler can learn to fuse broadcast, memory and blend operations
  * No need to support the full range of addressing modes, blending, broadcast and architectural registers in the API
* Utilise existing Go types
  * float16, bfloat16, etc. out of scope of this API
* Utilise existing Go golang.org/x/exp/constraints where possible
  * Propose adding `Number interface { constraints.Float | constraints.Integer }`
* Utilise Go arrays to represent vectors
  * Allows direct access to array elements
    * Avoids adding additional operations to insert/extract a scalar value from a vector
  * Simplifies future vector size extensions by decoupling vector sizes from types declared in other vector packages
* Simple naming that follows Go library/intrinsics names where possible (particularly the math and bits packages)
  * Except that we use opportunities that generics open up to reduce API call variants
* Where there is no Go name, adopt the name used by AVX512
  * Except where there is a simpler name (e.g. `VPCONFLICTD` detects duplicates, so might be better named `IsDuplicate`)
* Names are generally singular (e.g. `IsNaN` rather than `AreNaNs`) for consistency with existing APIs
* Utilise generic types to abstract over different integer/float sizes
  * i.e. favour `Add[T constraints.Number](a,b)` over `AddInt64`, `AddInt32`, `AddFloat64`, etc.
  * Avoids explosion in the number variants of basic operations
* Utilise packages to represent different vector lengths
  * i.e. favour `vec4.Add(a,b)` over `vec.Add4(a,b)`, `vec.Add8(a,b)`
  * Avoids explosion in the number variants of basic operations in the documentation
  * All operations on vectors of a specific length are grouped together and documented in a package
* Control flow/predication is expressed through mask/condition vectors
* Mask vectors should be opaque in order to support future vector length extension
  * e.g. Arm SVE permits up to 2048 bit vector supporting byte operations, which implies 256-bit masks
* Mask vectors are a separate type

  
## Status

I am not very familiar with programming vector CPUs so it is likely that common idioms are not well supported.
I have tried to build understanding by researching the Intel AVX instruction set flavours and reading general details of Arm SVE.
It is possible I have misunderstood the semantics of some instructions and implemented the fallbacks incorrectly.

No effort has been made to make fallback implementations efficient. The aim is for simple
and succinct code that keeps the focus on the API.

There are likely many opportunities to express Go fallback implementations in terms of smaller
vectors. Go compiler intrinsics could implement specific operations like `vec2.Add` in hardware,
allowing unsupported operations like `vec4.Add` to gang 2 operations using shorter vectors.
This would probably be the best way to structure fallback implementations of the different vector sizes.

Because I am generating the different vector sizes packages, `vec2.Deinterlace` returns `[1]T` instead of `T`.


## To do

* Vector.Slice() to return a slice of an otherwise-opaque vector's elements
* Mix
* Horizontal instructions and naming convention
  * "Horizontal" "Across" "All" "Vector" "Lanes" "Element(s)"
* Loop tail and alignment operations
* Should there be a neighbour swap based on mask bits? (1 = swap, 0 = leave)
* Swap(mask, a,b *[N]T) ? 
* VP2INTERSECT?
* AppendAndPermute (2-source table based permute in AVX512)
* Absolute difference? (SABD, UABD, FABD in SVE)
* Saturating arithmetic
* Dot product (SVE and AVX define partial integer dot product)
* Improve documentation by substituting `Length` with exact length
* Vector address generation
  * How to implement safely? Could return N x ptr to slice element
* Interlace and other ops may be over-constrained - no reason not to allow pointer elements
  * May need better generic type constraints to allow number + pointer
* Extended multiply and divide (i.e. 32b x 32b -> 64b)
* Complex number operations?
* Interlace/Deinterlace could be Zip/Unzip
* BlendAdd, BlendSub etc - m = false: return a, m = true: return a+b
  * Emulates x += 10 (vs Add which emulates x+10)
  * Make consistent with Add, Sub...
  * Maybe `BlendAdd(m, &a, b)`? Or `AddBlend`?
    * Should `Blend` also take a pointer to assign?
* Zero/SetZero might be better named Copy
* SVE-like mask operators
  * Need to understand these better
  * BreakBefore, BreakAfter
* Generate mask blocks, e.g. Even() == Group(1), Group(2)
  * Could be used for Interlace provided we validate the mask has N/2 set bits
  * Could be used to Partition() into 2 slices - like Deinterlace but without
* `fluent.New16[int]` actually returns Masked16 around a new vector so you can `.Blend(p).Add(a,b)`
* **Investigate `Ranger[Element]` constraint** to match `<-chan Element` and `[]Element`
  * If it works, have a single `Consume` method
* Panic/error if 2+ Bunch elements can't fit in vector size
* How do we deal with Convert()? It will generally return a different vector length for a different type.
  * Could return a slice of bunches?
  * Could return a special Bunch that can hold more than 1 Bunch inside and hope multiple conversions don't blow up the vector size too much
    * One optimisation would be when converting the special Bunch, aim to pack a full result vector even if it means converting multiple embedded input bunches
    * Could represent as a bunch of bunches?
    * If Element constrained to largest native (64 or 128 bits), worst case is converting from 64 int8 Bunch (el size=1) to 64 complex128 Bunch (el size=16)
      * So it will fit inside a bunch[16]
* Parallel LoadCorresponding
  * Loads each Bunch in parallel, stopping when the shortest one is full or the first Reader is empty
* `Truncate` a vector to get a smaller size (ie. split and discard)

## Disadvantages

* Fallback Go implementations are currently duplicated across different vector sizes
* Pushes the effort of coping with baroque vector instruction sets onto the Go compiler
  * But I believe this is the correct place for it
* Different packages for each vector size mean the user has to navigate between packages
  * However, each vector size supports the same operations, so understanding `vec2` enables understanding of `vec4`, etc.


## Specific decisions

* There was originally 1 mask type (opaque, with 64 bits)
  * Masks are now reimplemented at the vector size and called `vecN.Bool`
    * Pro: removes the need for masks to be opaque in order to support future vector length extension
    * Con: duplicates all the mask operations across each vector size package
    * Con: may make it harder to implement in-vector conditional logic across differing vector sizes (need to resize mask vector)
    * Pro: we have `Any()` to detect set bits, if mask follows vector size, we can also have `All()`
    * Pro: no longer have to pass `Length` to `ForEach()`
* `Broadcast` should not accept a mask parameter
  * Debatable decision: simplifies the vecN packages but might not work with the fluent style package I'm working on
  * The assumption is that tracking the value of a mask would make it harder for the compiler to fuse a broadcast with another operation
  * Using broadcast without a mask is simpler/shorter for the user
  * The fallback implementation of broadcast may lose some performance through unnecessarily copying the scalar value across the full array


## Design variants

* It is possible to reduce the width of type supported by longer vectors in sympathy with hardware
  * Can be done by having narrower constraints on types within the package
    * e.g. `type Number interface { Float32 | Integer32 }`
    * This is still open to extension in future by relaxing the type constraints
    * Disadvantage is not reusing the standard `constraints` package and duplicating constraints in each vector size package
    * Another disadvantage is that widening code to a larger vector size can make it uncompilable for reasons that seem arbitrary to the user
* It looks possible within Go's generic type system to include pointers within some operations
  * Type constraints get a bit complicated when pointers and underlying values mix
  * It is difficult to cleanly express pointers to pointers because recursive types not allowed
* You can *almost* fold all combinations of vector length N of type T down to a single (verbose) type constraint:
  * i.e. `type Vector interface { [2]int32 | [2]int64 ... | [4]int32 | [4]int64 ...`
  * But fallback implementations in Go trip over this rule:
    * https://go.dev/ref/spec#Index_expressions, specifically "The element types of all types in P's type set must be identical."
    * This prevents indexing the elements of `Vector` if they can be of different types
  * In practical terms, it seems better not to have types with specific (typically power-of-2) vector lengths
    * Makes it easier to express fallbacks of vector length N operations in terms of operations on vectors of length N/2
    * The compiler substitutes hardware instructions where they are available
      * e.g. recognise `Add[[2]uint64](a, b)` has a native instruction, which `Add[[4]uint64](a, b)` can execute twice
      * Similar to the approach in [Zig](https://zig.news/michalz/fast-multi-platform-simd-math-library-in-zig-2adn) linked from [proposal #53171](https://github.com/golang/go/issues/53171)
* Another variant which is possible but inconvenient is to define fallback as `Add[Element,Vector[Element]](a, b)` where `type Vector[E any] interface { [2]E | [4]E | [8]...}`
  * This works but automatic type inference isn't smart enough to allow type parameters to be elided at the call site:
    * So you have to write `Add[int,[4]int](a, b)` instead of `Add(a, b)`
    * TODO needs reconfirming - discovered this during a number of related generics experiments
    * As with previous bullet point, I prefer array types at hardware-sympathetic vector lengths where it is possible to write fallbacks of length N in terms of intrinsics of length N/2 
      * So not sure I would advocate this route even if type inference were better
* Could adopt https://pkg.go.dev/math/big convention of assigning the method receiver
  * This would require named Vector types and we might as well collapse all the vector types into 1 package
    * i.e. replace `x := vec2.Add(m, a, b)` with `x := Vec2{}; x.Add(m, a, b)`
    * Would we want to pass blend/zero parameter in this case?
      * i.e. `x.Add(m, vector.Blend, a, b)`
    * One alternative is extend "fluent" API style
      * e.g. `x.Blend(m).Add(a,b)`
      * `x.Add(a,b)` (not masked) and `x.Zero(m).Add(a,b)` zeroed instead of blended
    * Horizontal ops could be...
      * e.g. `x.Reduce().Max()`
      * or generally: `x.Reduce().ForRange(func(i int, v Element){...})`
        * Unclear if it's good/safe to pass element index `i`
  * Note math/big passes around pointers (required for receivers to allow assignment but we might pass vectors as values to operations like Add)
  * A vector as an opaque type (like big.Float) is in sympathy with Arm's SVE design
    * SVE more or less requires an opaque vector
    * SVE makes interesting and extensive use of predication to hide the vector length
      * Not an immediately intuitive model (to me)
      * There are hints that their primitives could allow lifting to code that is almost auto-vectorized in flavour
      * This might provide an intuitive model in which "looping" is captured by pulling vector-sized chunks from an input slice, processing using SIMD operations, then reducing to final output in a loop tail
    * SVE mask/predicate registers introduce a complication not present in x86:
      * in x86 mask bits are aligned to vector element index - easy to use across element sizes
      * in SVE mask bits are aligned to vector bytes
        * so a predicate formed from comparing int16s can't be used in int32 ops without a spread/unpack operation on the predicate
        * to cope:
          * define type-width-specific mask types (which I don't favour)
          * normalise masks after every operation
            * i.e. pack to x86-format after every op and unpack before every op
            * Could often be elided in practice as you will often work on the same width types (e.g. float64)
          * hide the masks by bundling them with the vectors
            * might work in a Fluent-style API
            * might be awkward to mask a vector using another vector, rather than a mask
            * might need to panic when combining vectors with different masks? (or you'd potentially be subtly losing data)
  * Potentially get partway to auto-vectorization with this approach
    * Assuming you can range over a vector and lift Go's scalar operations to vector equivalents
    * Also require loop head and tail in many cases, e.g. memory alignment in head + tail, final reduction in tail
    * Alignment may not be achievable when combining 2 or more slices and it's not clear SVE requires it
* Allow non-power-of-2 vector sizes and have `binaryOp`, `split` et al inspect array length in binary to determine how to split up vector into smaller ops
  * 15 should be sized up to 16
    * On CPU which can blend and doesn't read/write/fault a masked-off memory access, we could still store 15 elements but safely operate on 16
    * If we can't use raw Go arrays, makes more sense to have a Vec15 type which is [16]E internally
  * 9 should be sized down to 8 plus a scalar op
  * Probably best assume a given width should be half full or more

## Go language suggestions

* Improve type inference so type parameters of e.g. `Add[E Number, V[E] Vector](z, x, y *V)` can be inferred
  * **Edit** I've just realised you only need to provide enough type params to allow the rest to be worked out!
    * So it only requires `Add[float64](z, x, y)`, rather than `Add[float64,[8]float64](z, x, y)` as I originally thought 
      * Big improvement to readability
* Add ability to switch on type parameter, e.g. `switch V.(type) { case [2]E:`...
  * This should allow several same-type parameters (e.g. `(z, x, y *V)`) to be interpreted as a concrete type after the `case`
  * In some instances this could avoid a lot of type assertions
  * Compiler can check cases are possible, which can't be done with `switch any(v).(type)`
* In some cases it would help to be able to embed a type param in a constraint:
  * `type Broadcast[E any, V Vector[E]] V` <- `V` can only be struct field
    * This stops us taking `len` of values that include broadcast, even though they are all arrays
  * Similarly `type VectorScalar[E any] interface { Vector[E] | E }` <- `E` must be wrapped in a struct
* Puzzle:
  * Given `Broadcast[E any, V Vector[E]] struct { Replicated V }`
  * And `VectorBroadcast[E any, V Vector[E]] interface { Vector[E] | Broadcast[E, V] }`
  * Where I have a value of type `*Broadcast[E, V]` type-asserted from a `VectorBroadcast[E, V]`
  * I'm not sure why I can find the `len` of `Broadcast.Replicated` but I can't slice it
    * Go knows it has type `V` but it doesn't seem to know that is also `Vector[E]` and therefore sliceable