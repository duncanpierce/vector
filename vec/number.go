package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[To, From constraintsExt.ConvertableNumber, Z Vector[To], X Vector[From]](z Z, x X) {
	unary[To, From](nil, z, x, func(x From) To {
		return To(x)
	})
}

func Add[E constraintsExt.Number](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x + y
	})
}

func Sub[E constraintsExt.Number](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x - y
	})
}

func Mul[E constraintsExt.Number](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x * y
	})
}

func Div[E constraintsExt.Number](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x / y
	})
}

func Neg[E constraintsExt.Number](z, x Vector[E], m *Mask) {
	unary(m, z, x, func(x E) E {
		return -x
	})
}

func Abs[E constraintsExt.OrderedNumber](z, x Vector[E], m *Mask) {
	unary(m, z, x, func(x E) E {
		if x < E(0) {
			return -x
		} else {
			return x
		}
	})
}
