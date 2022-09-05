package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
)

// ConvertComplex returns a vector with all elements converted to another complex numeric type.
func ConvertComplex[To, From constraints.Complex, Z constraintsExt.Vector[To], X constraintsExt.Vector[From]](z *Z, x *X) {
	unary(z, nil, x, func(a From) To {
		return To(a)
	})
}
