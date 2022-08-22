package fluentvec

import "testing"

func TestArithmetic(t *testing.T) {
	a := &Vector16[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b := Broadcast16(10)
	r := &Vector16[int]{}
	r.All().Mul(a, b).Blend(None()).Add(a, b)
	r.Blend(All())

	plain := [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	Assign16(&plain).Add(a, b)
}
