package vec

//func New[T any]() (b VecN[T]) {
//	var el T
//	size := uint64(unsafe.Sizeof(el))
//	elementSize := 1 << (64 - bits.LeadingZeros64(size) - 1)
//	if bits.OnesCount64(size) != 1 {
//		// Go to next size up if size isn't a power of 2
//		elementSize *= 2
//	}
//	vectorLength := runtimeExt.VectorLenBytes() / elementSize
//	switch vectorLength {
//	case 2:
//		b = &VecN[T, [2]T]{}
//	case 4:
//		b = &bunch[T, [4]T]{}
//	case 8:
//		b = &bunch[T, [8]T]{}
//	case 16:
//		b = &bunch[T, [16]T]{}
//	case 32:
//		b = &bunch[T, [32]T]{}
//	default:
//		b = &bunch[T, [64]T]{}
//	}
//	return
//}
