package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Add[E constraintsExt.Number](z, x, y Vector[E]) {
	xSlice, xBroad := x.Slice()
	ySlice, yBroad := y.Slice()
	zSlice, _ := z.Slice()
	assignable[E](zSlice, xSlice, xBroad, "x")
	assignable[E](zSlice, ySlice, yBroad, "y")
	// TODO implement this somehow
}
