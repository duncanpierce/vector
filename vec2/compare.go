package vec2

import (
	"golang.org/x/exp/constraints"
)

func Less[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] < b[i])
	})
	return
}

func LessOrEqual[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] <= b[i])
	})
	return
}

func Greater[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] > b[i])
	})
	return
}

func GreaterOrEqual[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] >= b[i])
	})
	return
}

func Equal[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] == b[i])
	})
	return
}

func NotEqual[Element constraints.Ordered](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] != b[i])
	})
	return
}
