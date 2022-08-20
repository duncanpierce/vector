package vec8

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Gather[T constraintsExt.Number](m Bool, p [N]*T) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = *p[i]
	})
	return
}

func GatherSlice[T constraintsExt.Number](m Bool, s []T, index [N]int) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = s[index[i]]
	})
	return
}

func Scatter[T constraintsExt.Number](m Bool, p [N]*T, a [N]T) {
	m.ForTrue(func(i int) {
		*p[i] = a[i]
	})
	return
}

func ScatterSlice[T constraintsExt.Number](m Bool, s []T, index [N]int, a [N]T) {
	m.ForTrue(func(i int) {
		s[index[i]] = a[i]
	})
	return
}
