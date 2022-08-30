package vec

import "golang.org/x/exp/constraints"

func Less[E constraints.Ordered, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a < b
	})
}

func LessOrEqual[E constraints.Ordered, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a <= b
	})
}

func Greater[E constraints.Ordered, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a > b
	})
}

func GreaterOrEqual[E constraints.Ordered, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a >= b
	})
}
