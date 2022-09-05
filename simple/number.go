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
