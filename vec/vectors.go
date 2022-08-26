package vec

type (
	// Vec1 is a vector with 1 element.
	Vec1[E any] [1]E

	// Vec2 is a vector with 2 elements.
	Vec2[E any] [2]E

	// Vec4 is a vector with 4 elements.
	Vec4[E any] [4]E

	// Vec8 is a vector with 8 elements.
	Vec8[E any] [8]E

	// Vec16 is a vector with 16 elements.
	Vec16[E any] [16]E

	// Vec32 is a vector with 32 elements.
	Vec32[E any] [32]E

	// Vec64 is a vector with 64 elements.
	Vec64[E any] [64]E

	// Scalar holds a scalar value that can be used as a vector with the same value in all lanes.
	Scalar[E any] [1]E

	// Vector is an interface that represents any fixed-length or scalable vector or a scalar value.
	Vector[E any] interface {
		Slice() (slice []E, broadcast bool)
	}

	// FixedVector is a type constraint that matches any fixed-length vector.
	FixedVector[E any] interface {
		Vec1[E] | Vec2[E] | Vec4[E] | Vec8[E] | Vec16[E] | Vec32[E] | Vec64[E] | Scalar[E]
	}

	FixedVectorPointer[E any] interface {
		*Vec1[E] | *Vec2[E] | *Vec4[E] | *Vec8[E] | *Vec16[E] | *Vec32[E] | *Vec64[E] | *Scalar[E]
		Vector[E]
	}
)

var (
	_ Vector[int] = &Vec1[int]{}
	_ Vector[int] = &Vec2[int]{}
	_ Vector[int] = &Vec4[int]{}
	_ Vector[int] = &Vec8[int]{}
	_ Vector[int] = &Vec16[int]{}
	_ Vector[int] = &Vec32[int]{}
	_ Vector[int] = &Vec64[int]{}
	_ Vector[int] = Scalar[int]{}
)

func (s Scalar[E]) Slice() (slice []E, broadcast bool) {
	return s[:], true
}

func (v *Vec1[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec2[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec4[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec8[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec16[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec32[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}

func (v *Vec64[E]) Slice() (slice []E, broadcast bool) {
	return (*v)[:], false
}
