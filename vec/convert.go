package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
)

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[To, From constraintsExt.ConvertableNumber, Z Vector[To], X Vector[From]](z Z, x X) {
	unary[To, From](z, x, func(x From) To {
		return To(x)
	})
}

// ConvertComplex returns a vector with all elements converted to another complex numeric type.
func ConvertComplex[To, From constraints.Complex, Z Vector[To], X Vector[From]](z Z, x X) {
	unary[To, From](z, x, func(x From) To {
		return To(x)
	})
}
