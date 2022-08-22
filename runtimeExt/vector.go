package runtimeExt

// VectorLenBytes returns the width of the CPU's widest supported SIMD vector measured in bytes.
func VectorLenBytes() int {
	return 64
}
