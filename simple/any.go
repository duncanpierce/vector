package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
)

func Copy[E any, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary(z, m, x, func(x E) E {
		return x
	})
}
