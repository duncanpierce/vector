package mpi

import (
	"github.com/duncanpierce/vector/scalable"
	"testing"
)

type ThreeVars struct {
	c     int
	lanes scalable.Predicate
	x     scalable.Bunch[int32]
	y     scalable.Bunch[int32]
	z     scalable.Bunch[int64]
}

func TestMpi(t *testing.T) {
	xs := []int32{1, 2, 3, 4, 5}
	ys := []int32{2, 4, 6, 8, 10}

	For[ThreeVars](
		func(vars *ThreeVars) {
			// don't really need any setup so could have passed nil func
		},
		Range{
			Slice[int32]{
				Var: nil, // TODO - need a pointer and we don't have a ThreeVars here
			},
		},
		func(vars *ThreeVars) {
			// TODO mask to vars.lanes
			vars.z.Mul(vars.x, vars.y)
		},
	)
}
