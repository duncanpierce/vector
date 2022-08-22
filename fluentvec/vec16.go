package fluentvec

import "github.com/duncanpierce/vector/constraintsExt"

type (
	Vector16[Element constraintsExt.Number] [16]Element

	Masked16[Element constraintsExt.Number] struct {
		vector    *Vector16[Element]
		predicate Predicate
		setToZero bool
	}
)

func (v *Vector16[Element]) Blend(p Predicate) Masked16[Element] {
	return Masked16[Element]{vector: v, predicate: p, setToZero: false}
}

func (v *Vector16[Element]) Zero(p Predicate) Masked16[Element] {
	return Masked16[Element]{vector: v, predicate: p, setToZero: true}
}

func (v *Vector16[Element]) All() Masked16[Element] {
	return Masked16[Element]{vector: v, predicate: All(), setToZero: false}
}

func (v Masked16[Element]) Blend(p Predicate) Masked16[Element] {
	return Masked16[Element]{vector: v.vector, predicate: p, setToZero: false}
}

func (v Masked16[Element]) Zero(p Predicate) Masked16[Element] {
	return Masked16[Element]{vector: v.vector, predicate: p, setToZero: true}
}

func (v Masked16[Element]) All() Masked16[Element] {
	return Masked16[Element]{vector: v.vector, predicate: All(), setToZero: false}
}

/////

func Assign16[Element constraintsExt.Number](v *[16]Element) Masked16[Element] {
	return Masked16[Element]{vector: (*Vector16[Element])(v), predicate: All(), setToZero: false}
}

/////

func Broadcast16[Element constraintsExt.Number](v Element) *Vector16[Element] {
	r := Vector16[Element]{}
	for i := range r {
		r[i] = v
	}
	return &r
}

/////

func (v Masked16[Element]) ForRange(f func(i int)) Masked16[Element] {
	if v.setToZero {
		*v.vector = Vector16[Element]{}
	}
	v.predicate.ForRange(16, f)
	return v
}

func (v Masked16[Element]) Add(a, b *Vector16[Element]) Masked16[Element] {
	return v.ForRange(func(i int) {
		v.vector[i] = a[i] + b[i]
	})
}

func (v Masked16[Element]) Sub(a, b *Vector16[Element]) Masked16[Element] {
	return v.ForRange(func(i int) {
		v.vector[i] = a[i] - b[i]
	})
}

func (v Masked16[Element]) Mul(a, b *Vector16[Element]) Masked16[Element] {
	return v.ForRange(func(i int) {
		v.vector[i] = a[i] * b[i]
	})
}

func (v Masked16[Element]) Div(a, b *Vector16[Element]) Masked16[Element] {
	return v.ForRange(func(i int) {
		v.vector[i] = a[i] / b[i]
	})
}
