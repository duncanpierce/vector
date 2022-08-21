package vec8

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
)

func Gather[Element constraintsExt.Number](m Bool, p [Length]*Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = *p[i]
	})
	return
}

func GatherSlice[Element constraintsExt.Number, Index constraints.Integer](m Bool, s []Element, index [Length]Index) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = s[index[i]]
	})
	return
}

func Scatter[Element constraintsExt.Number](m Bool, p [Length]*Element, a [Length]Element) {
	m.ForTrue(func(i int) {
		*p[i] = a[i]
	})
	return
}

func ScatterSlice[Element constraintsExt.Number, Index constraints.Integer](m Bool, s []Element, index [Length]Index, a [Length]Element) {
	m.ForTrue(func(i int) {
		s[index[i]] = a[i]
	})
	return
}
