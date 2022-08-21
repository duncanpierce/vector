package vec2

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
	"math"
)

func Add[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] + b[i]
	})
	return
}

func Sub[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] - b[i]
	})
	return
}

func Div[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] / b[i]
	})
	return
}

func Mul[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] * b[i]
	})
	return
}

func FMA[Element constraintsExt.Number](m Bool, a, b, c [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i]*b[i] + c[i]
	})
	return
}

func Neg[Element constraintsExt.Number](m Bool, a [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = -a[i]
	})
	return
}

func Abs[Element constraintsExt.Number](m Bool, a [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		v := a[i]
		if v < 0 {
			v = -v
		}
		r[i] = v
	})
	return
}

func Max[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		v := a[i]
		if b[i] > v {
			v = b[i]
		}
		r[i] = v
	})
	return
}

func Min[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		v := a[i]
		if b[i] < v {
			v = b[i]
		}
		r[i] = v
	})
	return
}

func IsInf[Element constraints.Float](m Bool, a [Length]Element, sign int) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, math.IsInf(float64(a[i]), sign))
	})
	return
}

func IsNaN[Element constraints.Float](m Bool, a [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, math.IsNaN(float64(a[i])))
	})
	return
}
