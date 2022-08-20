package vec64

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/mask"
)

func Gather[T constraintsExt.Number](m mask.Bits, p [N]*T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = *p[i]
	})
	return
}

func GatherSlice[T constraintsExt.Number](m mask.Bits, s []T, index [N]int) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = s[index[i]]
	})
	return
}

func Scatter[T constraintsExt.Number](m mask.Bits, p [N]*T, a [N]T) {
	m.ForTrue(N, func(i int) {
		*p[i] = a[i]
	})
	return
}

func ScatterSlice[T constraintsExt.Number](m mask.Bits, s []T, index [N]int, a [N]T) {
	m.ForTrue(N, func(i int) {
		s[index[i]] = a[i]
	})
	return
}
