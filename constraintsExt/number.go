package constraintsExt

import "golang.org/x/exp/constraints"

type (
	// Ideally, this would be added to constraints package
	Number interface {
		constraints.Float | constraints.Integer
	}
)
