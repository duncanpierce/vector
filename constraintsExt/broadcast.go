package constraintsExt

type (
	Broadcast[E any, V Vector[E]] struct {
		Replicated V
	}

	Scalar[E any] struct {
		Value E
	}

	VectorBroadcast[E any] interface {
		//Vector[E] | Broadcast[E, V]
		Vector[E] | Scalar[E]
	}
)
