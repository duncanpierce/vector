package vec2

import (
	"vector/constraintsExt"
	"vector/mask"
)

func Less[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] < b[i])
	})
	return
}

func LessOrEqual[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] <= b[i])
	})
	return
}

func Greater[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] > b[i])
	})
	return
}

func GreaterOrEqual[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] >= b[i])
	})
	return
}

func Equal[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] == b[i])
	})
	return
}

func NotEqual[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r mask.Bits) {
	m.ForTrue(N, func(i int) {
		r.Set(i, a[i] != b[i])
	})
	return
}
