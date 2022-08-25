package vec

import (
	"fmt"
	"github.com/duncanpierce/vector/constraintsExt"
	"unsafe"
)

func Split[E constraintsExt.Number, Vector constraintsExt.Vector[E], HalfVector constraintsExt.Vector[E]](x *Vector) (x0, x1 *HalfVector) {
	validSplit[E, Vector, HalfVector]()
	return unsafeSplit[E, Vector, HalfVector](x)
}

func validSplit[E constraintsExt.Number, Vector constraintsExt.Vector[E], HalfVector constraintsExt.Vector[E]]() {
	var (
		big        Vector
		little     HalfVector
		bigSize    = unsafe.Sizeof(big)
		littleSize = unsafe.Sizeof(little)
	)
	if littleSize != bigSize/2 {
		panic("can only split a vector into 2 equal-sized halves")
	}
}

func unsafeSplit[E constraintsExt.Number, Vector constraintsExt.Vector[E], HalfVector constraintsExt.Vector[E]](x *Vector) (x0, x1 *HalfVector) {
	var (
		little     HalfVector
		littleSize = unsafe.Sizeof(little)
	)
	x0 = (*HalfVector)(unsafe.Pointer(x))
	x1 = (*HalfVector)(unsafe.Add(unsafe.Pointer(x), littleSize))
	return
}

func binaryOp[E constraintsExt.Number, Big, Little constraintsExt.Vector[E]](z, x, y *Big, f func(zl, xl, yl *Little)) {
	validSplit[E, Big, Little]()
	z0, z1 := unsafeSplit[E, Big, Little](z)
	x0, x1 := unsafeSplit[E, Big, Little](x)
	y0, y1 := unsafeSplit[E, Big, Little](y)
	f(z0, x0, y0)
	f(z1, x1, y1)
}

func combineNumbers2[E any, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](z *V, x, y *VB, op ...[]func(c, a, b *any)) {
	//xSlice, xBroadcast := unsafeSliceBroadcast[E,V](x)
	fmt.Printf("vector len %v", len(*z))
}
