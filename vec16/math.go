package vec16

import (
	"golang.org/x/exp/constraints"
	"math"
	"vector/constraintsExt"
	"vector/mask"
)

func Add[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] + b[i]
	})
	return
}

func Sub[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] - b[i]
	})
	return
}

func Div[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] / b[i]
	})
	return
}

func Mul[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] * b[i]
	})
	return
}

func FMA[T constraintsExt.Number](m mask.Bits, a, b, c [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i]*b[i] + c[i]
	})
	return
}

func Neg[T constraintsExt.Number](m mask.Bits, a [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = -a[i]
	})
	return
}

func Abs[T constraintsExt.Number](m mask.Bits, a [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		v := a[i]
		if v < 0 {
			v = -v
		}
		r[i] = v
	})
	return
}

func Max[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		v := a[i]
		if b[i] > v {
			v = b[i]
		}
		r[i] = v
	})
	return
}

func Min[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		v := a[i]
		if b[i] < v {
			v = b[i]
		}
		r[i] = v
	})
	return
}

func IsInf[T constraints.Float](m mask.Bits, a [N]T, sign int) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, math.IsInf(float64(a[i]), sign))
	})
	return
}

func IsNaN[T constraints.Float](m mask.Bits, a [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, math.IsNaN(float64(a[i])))
	})
	return
}
