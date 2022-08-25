package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"unsafe"
)

//func Len[E any, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E, V]](vb *VB) (l int, isBroadcast bool) {
//	switch vec := any(vb).(type) {
//	case *V:
//		return len(*vec), false
//	case *constraintsExt.Broadcast[E, V]:
//		return len((*vec).Replicated), true
//	default:
//		panic("unreachable code")
//	}
//}

func unsafeSlice[E any, V constraintsExt.Vector[E]](v *V) (slice []E) {
	// TODO not sure why I can't return (*v)[:]
	return unsafe.Slice((*E)(unsafe.Pointer(v)), len(*v))
}

func unsafeSliceBroadcast[E any, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](vb *VB) (slice []E, isBroadcast bool) {
	switch vec := any(vb).(type) {
	case *constraintsExt.Scalar[E]:
		slice = unsafe.Slice((*E)(unsafe.Pointer(&(*vec).Value)), 1)
		isBroadcast = true
	case *V:
		slice = unsafeSlice[E, V](vec)
		isBroadcast = false
	default:
		panic("unreachable code")
	}
	return
}

func Replicate[E any, V2, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](v *V, vb *VB) {
	vbSlice, _ := unsafeSliceBroadcast[E, V2](vb)
	vSlice := unsafeSlice[E](v)
	for len(vSlice) > 0 {
		copy(vSlice, vbSlice)
		vSlice = vSlice[len(vbSlice):]
	}
}
