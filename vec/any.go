package vec

func Copy[E any, Z, X Vector[E]](z Z, x X, m *Mask) {
	unary[E, E](m, z, x, func(x E) E {
		return x
	})
}
