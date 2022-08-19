package vec2

import (
	"vec/condition"
	"vec/constraintsExt"
)

func Less[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] < b[i])
	})
	return
}

func LessOrEqual[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] <= b[i])
	})
	return
}

func Greater[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] > b[i])
	})
	return
}

func GreaterOrEqual[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] >= b[i])
	})
	return
}

func Equal[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] == b[i])
	})
	return
}

func NotEqual[T constraintsExt.Number](m condition.Mask, a, b [N]T) (r condition.Mask) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] != b[i])
	})
	return
}
