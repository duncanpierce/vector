package vec

import "github.com/duncanpierce/vector/constraintsExt"

func Broadcast[E any, Z constraintsExt.Vector[E]](z *Z, v E) {
	for i := 0; i < len(*z); i++ {
		(*z)[i] = v
	}
}

func Replicate[E any, Z, X constraintsExt.Vector[E]](z *Z, x *X) {
	xSlice := unsafeSlice[E](x)
	zSlice := unsafeSlice[E](z)
	if len(xSlice) > len(zSlice) {
		panic("vector being replicated cannot have more elements than the one it is being copied to")
	}
	for len(zSlice) > 0 {
		copy(zSlice, xSlice)
		zSlice = zSlice[len(xSlice):]
	}
}

//func Broadcast[E any, Small, Z constraintsExt.Vector[E], X constraintsExt.VectorBroadcast[E]](z *Z, x *X) {
//	xSlice, _ := unsafeSliceBroadcast[E, Small](x)
//	zSlice := unsafeSlice[E](z)
//	if len(xSlice) > len(zSlice) {
//		panic("vector being broadcast cannot have more elements than the one it is being broadcast to")
//	}
//	for len(zSlice) > 0 {
//		copy(zSlice, xSlice)
//		zSlice = zSlice[len(xSlice):]
//	}
//}
