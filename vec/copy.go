package vec

func Copy[E any, Z, X Vector[E]](z Z, x X) {
	unary[E](z, x, func(x E) E {
		return x
	})
}

//func Replicate[E any, Z, X FixedVector[E]](z Z, x X) {
//	xSlice, _ := x.slice()
//	zSlice, _ := z.slice()
//	if len(xSlice) > len(zSlice) {
//		panic("vector being replicated cannot have more elements than the one it is being copied to")
//	}
//	for len(zSlice) > 0 {
//		copy(zSlice, xSlice)
//		zSlice = zSlice[len(xSlice):]
//	}
//}
