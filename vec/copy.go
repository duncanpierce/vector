package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"unsafe"
)

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

func Len[E any, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](vb *VB) (length int, isBroadcast bool) {
	slice, b := unsafeSliceBroadcast[E, V](vb)
	return len(slice), b
}

func Replicate[E any, Small, Big constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](v *Big, vb *VB) {
	vbSlice, _ := unsafeSliceBroadcast[E, Small](vb)
	vSlice := unsafeSlice[E](v)
	for len(vSlice) > 0 {
		copy(vSlice, vbSlice)
		vSlice = vSlice[len(vbSlice):]
	}
}
