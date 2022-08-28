package vec

import (
	"golang.org/x/exp/constraints"
)

func LessThan[E constraints.Ordered, V FixedVector[E]](z *Lanes[E], x, y V) {
	//xc := any(x).(*[2]E)
	//yc := any(y).(*[2]E)
	// TODO implement this
}
