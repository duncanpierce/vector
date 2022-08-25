package constraintsExt

import "golang.org/x/exp/constraints"

type (
	// Number is a type constraint matching all numbers, including bytes and runes. Candidate for inclusion in constraints package.
	Number interface {
		constraints.Float | constraints.Integer | constraints.Complex
	}
)
