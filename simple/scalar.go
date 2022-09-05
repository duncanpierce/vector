package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
)

func Broadcast[E any, Z constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x E) {
	lanes.Range[E](z, m, func(i, j int) {
		(*z)[i] = x
	})
}
