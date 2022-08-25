package vec

import "github.com/duncanpierce/vector/constraintsExt"

func Add[E constraintsExt.Number, V constraintsExt.Vector[E], VB constraintsExt.VectorBroadcast[E]](z *V, x, y *VB) {
	// TODO implement this somehow
}

func (z *Vector1[E]) Add(x, y *Vector1[E]) {
	z[0] = x[0] + y[0]
}

func (z *Vector2[E]) Add(x, y *Vector2[E]) {
	binaryOp[E, Vector2[E], Vector1[E]](z, x, y, func(zl, xl, yl *Vector1[E]) {
		zl.Add(xl, yl)
	})
}

func (z *Vector4[E]) Add(x, y *Vector4[E]) {
	binaryOp[E, Vector4[E], Vector2[E]](z, x, y, func(zl, xl, yl *Vector2[E]) {
		zl.Add(xl, yl)
	})
}

func (z *Vector8[E]) Add(x, y *Vector8[E]) {
	binaryOp[E, Vector8[E], Vector4[E]](z, x, y, func(zl, xl, yl *Vector4[E]) {
		zl.Add(xl, yl)
	})
}

func (z *Vector16[E]) Add(x, y *Vector16[E]) {
	binaryOp[E, Vector16[E], Vector8[E]](z, x, y, func(zl, xl, yl *Vector8[E]) {
		zl.Add(xl, yl)
	})
}

func (z *Vector32[E]) Add(x, y *Vector32[E]) {
	binaryOp[E, Vector32[E], Vector16[E]](z, x, y, func(zl, xl, yl *Vector16[E]) {
		zl.Add(xl, yl)
	})
}

func (z *Vector64[E]) Add(x, y *Vector64[E]) {
	binaryOp[E, Vector64[E], Vector32[E]](z, x, y, func(zl, xl, yl *Vector32[E]) {
		zl.Add(xl, yl)
	})
}
