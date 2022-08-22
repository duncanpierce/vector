package constraintsExt

import "golang.org/x/exp/constraints"

type (
	// Number is a type constraint matching constraints.Float and constraints.Integer. Candidate for adding to constraints package.
	Number interface {
		constraints.Float | constraints.Integer
	}
)
