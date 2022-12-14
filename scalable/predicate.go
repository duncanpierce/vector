package scalable

import "math/bits"

type (
	// Predicate records which elements of a vector are active and hold usable values.
	// To support Arm SVE, this requires re-striding when changing vector element widths. See README-DESIGN.md.
	Predicate struct {
		mask uint64 // could be a Go array, must be long enough to hold 1 bit for each byte in runtimeExt.VectorSizeBytes()
	}
)

func (p Predicate) ForActive(f func(index int)) {
	for i := 0; i < 64; i++ {
		if p.mask&(1<<i) != 0 {
			f(i)
		}
	}
}

func (p Predicate) Set(i int, b bool) (result Predicate) {
	result.mask = p.mask &^ (1 << i)
	if b {
		result.mask |= 1 << i
	}
	return
}

func (p Predicate) Count() int {
	return bits.OnesCount64(p.mask)
}
