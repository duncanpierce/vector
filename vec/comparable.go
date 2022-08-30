package vec

func Equal[E comparable, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a == b
	})
}

func NotEqual[E comparable, Z Bool, XY Vector[E]](z Z, x, y XY) {
	binaryBool[E](z, x, y, func(a, b E) bool {
		return a != b
	})
}
