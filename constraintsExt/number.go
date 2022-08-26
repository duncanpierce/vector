package constraintsExt

import "golang.org/x/exp/constraints"

type (
	// Number is a type constraint matching all numbers, including bytes and runes. Candidate for inclusion in constraints package.
	Number interface {
		constraints.Float | constraints.Integer | constraints.Complex
	}

	// OrderedNumber is a type constraint matching all ordered numbers, including bytes and runes but not complex numbers, which are not ordered. Candidate for inclusion in constraints package.
	OrderedNumber interface {
		constraints.Float | constraints.Integer
	}

	// ConvertableNumber is a type constraint matching all numbers which can be converted to another ConvertableNumber. This includes bytes and runes but not complex numbers. Candidate for inclusion in constraints package.
	ConvertableNumber interface {
		constraints.Float | constraints.Integer
	}
)
