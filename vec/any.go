package vec

func Copy[E any, Z, X Vector[E]](z Z, x X) {
	unary[E, E](z, x, func(x E) E {
		return x
	})
}
