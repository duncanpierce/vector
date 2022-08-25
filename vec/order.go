package vec

import (
	"fmt"
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
)

func LessThan[E constraints.Ordered, V constraintsExt.Vector[E]](z *Lanes[E], x, y *V) {
	xc := any(x).(*[2]E)
	yc := any(y).(*[2]E)
	fmt.Printf("first cell %v", (*xc)[0] < (*yc)[0])
	// TODO implement this
}