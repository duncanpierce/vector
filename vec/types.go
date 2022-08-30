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
	}

	ScalableVector[E any] interface{}

	Vector[E any] interface {
		FixedVector[E] | ScalableVector[E]
		slicer[E]
	}

	Mask struct {
		zero bool
		m    uint64
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

type (
	Bool1 struct {
		mask uint64
	}

	Bool2 struct {
		mask uint64
	}

	Bool4 struct {
		mask uint64
	}

	Bool8 struct {
		mask uint64
	}

	Bool16 struct {
		mask uint64
	}

	Bool32 struct {
		mask uint64
	}

	Bool64 struct {
		mask uint64
	}

	FixedBool interface {
		*Bool1 | *Bool2 | *Bool4 | *Bool8 | *Bool16 | *Bool32 | *Bool64
	}

	ScalableBool interface {
	}

	Bool interface {
		FixedBool | ScalableBool
		lanes() lanes
	}

	lanes struct {
		mask   *uint64
		nLanes int
	}
)

var (
	_ Bool = &Bool1{}
	_ Bool = &Bool2{}
	_ Bool = &Bool4{}
	_ Bool = &Bool8{}
	_ Bool = &Bool16{}
	_ Bool = &Bool32{}
	_ Bool = &Bool64{}
)

func (b *Bool1) lanes() lanes {
	return lanes{&b.mask, 1}
}

func (b *Bool2) lanes() lanes {
	return lanes{&b.mask, 2}
}

func (b *Bool4) lanes() lanes {
	return lanes{&b.mask, 4}
}

func (b *Bool8) lanes() lanes {
	return lanes{&b.mask, 8}
}

func (b *Bool16) lanes() lanes {
	return lanes{&b.mask, 16}
}

func (b *Bool32) lanes() lanes {
	return lanes{&b.mask, 32}
}

func (b *Bool64) lanes() lanes {
	return lanes{&b.mask, 64}
}

func (l lanes) set(lane int, active bool) {
	bit := uint64(1 << lane)
	mask := (*l.mask) &^ bit
	if active {
		mask |= bit
	}
	*l.mask = mask
}
