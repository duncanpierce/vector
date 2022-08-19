package vec2

import (
	"vector/condition"
	"vector/constraintsExt"
)

func Gather[T constraintsExt.Number](m condition.Mask, p [N]*T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = *p[i]
	})
	return
}

func GatherSlice[T constraintsExt.Number](m condition.Mask, s []T, index [N]int) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = s[index[i]]
	})
	return
}

func Scatter[T constraintsExt.Number](m condition.Mask, p [N]*T, a [N]T) {
	m.ForTrue(N, func(i int) {
		*p[i] = a[i]
	})
	return
}

func ScatterSlice[T constraintsExt.Number](m condition.Mask, s []T, index [N]int, a [N]T) {
	m.ForTrue(N, func(i int) {
		s[index[i]] = a[i]
	})
	return
}
