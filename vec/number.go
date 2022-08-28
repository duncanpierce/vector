package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[To, From constraintsExt.ConvertableNumber, Z Vector[To], X Vector[From]](z Z, x X) {
	unary[To, From](z, x, func(x From) To {
		return To(x)
	})
}

func Add[E constraintsExt.Number](z, x, y Vector[E]) {
	binary(z, x, y, func(x, y E) E {
		return x + y
	})
}

func Sub[E constraintsExt.Number](z, x, y Vector[E]) {
	binary(z, x, y, func(x, y E) E {
		return x - y
	})
}

func Mul[E constraintsExt.Number](z, x, y Vector[E]) {
	binary(z, x, y, func(x, y E) E {
		return x * y
	})
}

func Div[E constraintsExt.Number](z, x, y Vector[E]) {
	binary(z, x, y, func(x, y E) E {
		return x / y
	})
}
