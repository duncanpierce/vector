package constraintsExt

type (
	Scalar[E any] struct {
		Value E
	}

	VectorBroadcast[E any] interface {
		Vector[E] | Scalar[E]
	}
)
