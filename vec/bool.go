package vec

type (
	Lanes[E any] struct {
		// Arm SVE motivates Lanes[E] being different from Bool.
		// On AVX512, 1 bit corresponds to 1 element within vector register.
		// Arm SVE it may be wider because it each bit corresponds to 1 byte of vector register.
		mask [1]uint64
	}

	Bool struct {
		// mask has enough bits for each lane of the largest supported vector size
		mask uint64
	}

	Predicate interface{}
)

var (
	_ Predicate = Lanes[int]{}
	_ Predicate = Bool{}
)
