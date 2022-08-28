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

	// broadcast holds a single value that can be used as a vector with the same value in all lanes.
	broadcast[E any] [1]E

	// A slicer is a value that can be sliced.
	slicer[E any] interface {
		elements() elements[E]
	}

	// FixedVector is a type constraint that matches any fixed-length vector or scalar.
	FixedVector[E any] interface {
		*Vec1[E] | *Vec2[E] | *Vec4[E] | *Vec8[E] | *Vec16[E] | *Vec32[E] | *Vec64[E] | *broadcast[E]
		//slicer[E]
	}

	ScalableVector[E any] interface {
		//slicer[E]
	}

	Vector[E any] interface {
		FixedVector[E] | ScalableVector[E]
		slicer[E]
	}
)

var (
	_ slicer[int] = &Vec1[int]{}
	_ slicer[int] = &Vec2[int]{}
	_ slicer[int] = &Vec4[int]{}
	_ slicer[int] = &Vec8[int]{}
	_ slicer[int] = &Vec16[int]{}
	_ slicer[int] = &Vec32[int]{}
	_ slicer[int] = &Vec64[int]{}
	_ slicer[int] = broadcast[int]{}
)

func (s broadcast[E]) elements() elements[E] {
	return elements[E]{s[:], true}
}

func (v *Vec1[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec2[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec4[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec8[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec16[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec32[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}

func (v *Vec64[E]) elements() elements[E] {
	return elements[E]{(*v)[:], false}
}
