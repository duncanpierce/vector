package vec

import (
	"golang.org/x/exp/constraints"
)

// ConvertComplex returns a vector with all elements converted to another complex numeric type.
func ConvertComplex[To, From constraints.Complex, Z Vector[To], X Vector[From]](z Z, x X) {
	unary[To, From](nil, z, x, func(x From) To {
		return To(x)
	})
}
