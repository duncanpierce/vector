package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
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
