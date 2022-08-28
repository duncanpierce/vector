package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Add[E constraintsExt.Number](z, x, y Vector[E]) {
	binary[E](z, x, y, func(x, y E) E {
		return x + y
	})
}
