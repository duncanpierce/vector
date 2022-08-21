package vec4

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Less[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] < b[i])
	})
	return
}

func LessOrEqual[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] <= b[i])
	})
	return
}

func Greater[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] > b[i])
	})
	return
}

func GreaterOrEqual[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] >= b[i])
	})
	return
}

func Equal[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] == b[i])
	})
	return
}

func NotEqual[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i] != b[i])
	})
	return
}
