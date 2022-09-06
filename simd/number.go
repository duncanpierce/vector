package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simd/lanes"
)

func Add[E constraintsExt.Number, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		return a + b
	})
}

func Sub[E constraintsExt.Number, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		return a - b
	})
}

func Mul[E constraintsExt.Number, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		return a * b
	})
}

// Note this is named FMA in math package.
func MulAdd[E constraintsExt.Number, WXYZ constraintsExt.Vector[E]](z *WXYZ, m *lanes.Bool, w, x, y *WXYZ) {
	ternary[E](z, m, w, x, y, func(a, b, c E) E {
		return a*b + c
	})
}

func Div[E constraintsExt.Number, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		return a / b
	})
}

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[To, From constraintsExt.ConvertableNumber, Z constraintsExt.Vector[To], X constraintsExt.Vector[From]](z *Z, x *X) {
	unary(z, nil, x, func(x From) To {
		return To(x)
	})
}

func Neg[E constraintsExt.Number, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary(z, nil, x, func(x E) E {
		return -x
	})
}
