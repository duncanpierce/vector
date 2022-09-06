package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
)

// ConvertComplex returns a vector with all elements converted to another complex numeric type.
func ConvertComplex[To, From constraints.Complex, Z constraintsExt.Vector[To], X constraintsExt.Vector[From]](z *Z, x *X) {
	unary(z, nil, x, func(a From) To {
		return To(a)
	})
}

func Complex128[Z constraintsExt.Vector[complex128], XY constraintsExt.Vector[float64]](z *Z, x, y *XY) {
	binaryConvert(z, nil, x, y, func(a, b float64) complex128 {
		return complex(a, b)
	})
}

func Complex64[Z constraintsExt.Vector[complex64], XY constraintsExt.Vector[float32]](z *Z, x, y *XY) {
	binaryConvert(z, nil, x, y, func(a, b float32) complex64 {
		return complex(a, b)
	})
}

func Real64[Z constraintsExt.Vector[float64], X constraintsExt.Vector[complex128]](z *Z, x *X) {
	unaryConvert(z, nil, x, func(a complex128) float64 {
		return real(a)
	})
}

func Imag64[Z constraintsExt.Vector[float64], X constraintsExt.Vector[complex128]](z *Z, x *X) {
	unaryConvert(z, nil, x, func(a complex128) float64 {
		return imag(a)
	})
}

func Real32[Z constraintsExt.Vector[float32], X constraintsExt.Vector[complex64]](z *Z, x *X) {
	unaryConvert(z, nil, x, func(a complex64) float32 {
		return real(a)
	})
}

func Imag32[Z constraintsExt.Vector[float32], X constraintsExt.Vector[complex64]](z *Z, x *X) {
	unaryConvert(z, nil, x, func(a complex64) float32 {
		return imag(a)
	})
}
