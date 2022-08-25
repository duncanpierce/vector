package vec

import "github.com/duncanpierce/vector/constraintsExt"

func Scalar[E any](v E) *constraintsExt.Scalar[E] {
	return &constraintsExt.Scalar[E]{v}
}
