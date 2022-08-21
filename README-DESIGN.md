# Design decisions

## Goals and assumptions

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
* Control flow is expressed through mask/condition vectors
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

* SwapLanes/SwapHorizontal
* Horizontal instructions and naming convention
  * "Horizontal" "Across" "All" "Vector" "Lanes"
  * e.g. SumAcross
* Loop tail and alignment operations
* Should there be a neighbour swap based on mask bits? (1 = swap, 0 = leave)
* Swap(mask, a,b *[N]T) ? 
* VP2INTERSECT?

## Disadvantages

* Fallback Go implementations are currently duplicated across different vector sizes
* Pushes the effort of coping with baroque vector instruction sets onto the Go compiler
  * But I believe this is the correct place for it
* Different packages for each vector size mean the user has to navigate between packages
  * However, each vector size supports the same operations, so understanding `vec2` enables understanding of `vec4`, etc.
  * User needs to be made aware of the different vector sizes in order to navigate the documentation
    * i.e. how do you discover `vec4` if you are reading the documentation of `vec2`?


## Specific decisions

* There is 1 mask type
  * Masks could instead be implemented at each vector size
    * e.g. replacing (`mask.Bits` with `vec2.Mask`, `vec4.Mask`, etc.)
    * Pro: removes the need for masks to be opaque in order to support future vector length extension
    * Con: duplicates all the mask operations across each vector size package
    * Con: may make it harder to implement in-vector conditional logic across differing vector sizes (need to resize mask vector)
    * Pro: we have `Any()` to detect set bits, if mask follows vector size, we can also have `All()`
    * Pro: would no longer have to pass `N` to `Bits.ForEach()`
* `Broadcast` should not accept a mask parameter
  * The assumption is that tracking the value of a mask would make it harder for the compiler to fuse a broadcast with another operation
  * Using broadcast without a mask is simpler/shorter for the user
  * The fallback implementation of broadcast may lose some performance through unnecessarily copying the scalar value across the full array
