package vec

func Broadcast[E any](x E) *Scalar[E] {
	return &Scalar[E]{x}
}

//func Replicate[E any, Z, X FixedVectorPointer[E]](z Z, x X) {
//	xSlice, _ := x.Slice()
//	zSlice, _ := z.Slice()
//	if len(xSlice) > len(zSlice) {
//		panic("vector being replicated cannot have more elements than the one it is being copied to")
//	}
//	for len(zSlice) > 0 {
//		copy(zSlice, xSlice)
//		zSlice = zSlice[len(xSlice):]
//	}
//}

func Copy[E any, Z, X FixedVectorPointer[E]](z Z, x X) {
	xSlice, xBroad := x.Slice()
	zSlice, _ := z.Slice()
	assignable[E](zSlice, xSlice, xBroad, "x")
	if xBroad {
		for len(zSlice) > 0 {
			copy(zSlice, xSlice)
			zSlice = zSlice[len(xSlice):]
		}
	} else {
		copy(zSlice, xSlice)
	}
}
