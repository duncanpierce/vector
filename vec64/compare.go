package vec64

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Less[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] < b[i])
	})
	return
}

func LessOrEqual[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] <= b[i])
	})
	return
}

func Greater[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] > b[i])
	})
	return
}

func GreaterOrEqual[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] >= b[i])
	})
	return
}

func Equal[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] == b[i])
	})
	return
}

func NotEqual[T constraintsExt.Number](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] != b[i])
	})
	return
}
