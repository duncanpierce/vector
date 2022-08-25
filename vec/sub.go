package vec

//import (
//	"github.com/duncanpierce/vector/constraintsExt"
//)
//
//func Sub[E constraintsExt.Number, V constraintsExt.NumberVector[E]](z, x, y *V) {
//	// Would be much simpler and maybe more efficient if we could write `switch E.(type) {...}`
//	switch sizedZ := any(z).(type) {
//	case *[1]E:
//		(*z)[0] = (*x)[0] - (*y)[0]
//	case *[2]E:
//		sizedX := any(x).(*[2]E)
//		sizedY := any(y).(*[2]E)
//		// opportunity for vector intrinsic here
//		binaryOp[E, [2]E, [1]E](sizedZ, sizedX, sizedY, func(zn, xn, yn *[1]E) {
//			Sub[E, [1]E](zn, xn, yn)
//		})
//	default:
//		sizedX := any(x).(*[2]E)
//		sizedY := any(y).(*[2]E)
//		// opportunity for vector intrinsic here
//		binaryOp[E, [2]E, [1]E](sizedZ, sizedX, sizedY, func(zn, xn, yn *[1]E) {
//			Sub[E, [1]E](zn, xn, yn)
//		})
//
//	}
//}
