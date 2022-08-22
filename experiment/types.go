package experiment

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/vec16"
)

type (
	Vector[Element any] interface {
		[vec16.Length]Element
	}

	NumberPointer[T constraintsExt.Number] interface {
		*T
	}

	Pointer[Element any] interface {
		*Element
	}

	NumberOrPointer[T constraintsExt.Number] interface {
		constraintsExt.Number | *T
	}

	AnyVector[Element constraintsExt.Number] interface {
		[16]Element | [8]Element | [4]Element | [2]Element
	}

	AnyNumericVector interface {
		[16]int | [8]int | [4]int | [2]int
		// | [16]float64 | [8]float64 | [4]float64 | [2]float64 // unfortunately this is disallowed by https://go.dev/ref/spec#Index_expressions "The element types of all types in P's type set must be identical." :-(
		//[16]float32 | [8]float32 | [4]float32 | [2]float32
	}
)

// Doesn't compile:
//func Copy[T Pointer[constraintsExt.Number], V Vector[T]](m Bool, a V) (result V) {
//	m.ForTrue(func(i int) {
//		result[i] = a[i]
//	})
//	return
//}

func Copy2[V constraintsExt.Number](m vec16.Bool, a [2]V) (result [2]V) {
	m.ForTrue(func(i int) {
		result[i] = a[i]
	})
	return
}

func Copy3[V Vector[int]](m vec16.Bool, a V) (result V) {
	m.ForTrue(func(i int) {
		result[i] = a[i]
	})
	return
}

func Add[Element constraintsExt.Number, V AnyVector[Element]](a, b V) (r V) {
	for i := 0; i < len(r); i++ {
		r[i] = a[i] + b[i]
	}
	return
}

func Add2[V AnyNumericVector](a, b V) (r V) {
	for i := 0; i < len(r); i++ {
		r[i] = a[i] + b[i]
	}
	return
}
