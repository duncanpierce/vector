package constraintsExt

type (
	Vector[E any] interface {
		~[2]E | ~[4]E | ~[8]E | ~[16]E | ~[32]E | ~[64]E
	}
)
