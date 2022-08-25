package vec

func (z *Vector1[E]) Mul(x, y *Vector1[E]) {
	z[0] = x[0] * y[0]
}

func (z *Vector2[E]) Mul(x, y *Vector2[E]) {
	binaryOp[E, Vector2[E], Vector1[E]](z, x, y, func(zl, xl, yl *Vector1[E]) {
		zl.Mul(xl, yl)
	})
}

func (z *Vector4[E]) Mul(x, y *Vector4[E]) {
	binaryOp[E, Vector4[E], Vector2[E]](z, x, y, func(zl, xl, yl *Vector2[E]) {
		zl.Mul(xl, yl)
	})
}

func (z *Vector8[E]) Mul(x, y *Vector8[E]) {
	binaryOp[E, Vector8[E], Vector4[E]](z, x, y, func(zl, xl, yl *Vector4[E]) {
		zl.Mul(xl, yl)
	})
}

func (z *Vector16[E]) Mul(x, y *Vector16[E]) {
	binaryOp[E, Vector16[E], Vector8[E]](z, x, y, func(zl, xl, yl *Vector8[E]) {
		zl.Mul(xl, yl)
	})
}

func (z *Vector32[E]) Mul(x, y *Vector32[E]) {
	binaryOp[E, Vector32[E], Vector16[E]](z, x, y, func(zl, xl, yl *Vector16[E]) {
		zl.Mul(xl, yl)
	})
}

func (z *Vector64[E]) Mul(x, y *Vector64[E]) {
	binaryOp[E, Vector64[E], Vector32[E]](z, x, y, func(zl, xl, yl *Vector32[E]) {
		zl.Mul(xl, yl)
	})
}
