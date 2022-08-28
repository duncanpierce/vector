package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

func Equal[E comparable, V constraintsExt.Vector[E]](z *Lanes[E], x, y *V) {
	//xc := any(x).(*[2]E)
	//yc := any(y).(*[2]E)
	//fmt.Printf("first cell %v", (*xc)[0] == (*yc)[0])
	// TODO implement this
}
