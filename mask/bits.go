package mask

type Bits struct {
	v uint64
}

const (
	size = 64
	all  = 0xffffffffffffffff
)

func All() (r Bits) {
	return Bits{all}
}

func One() (r Bits) {
	return Bits{1}
}

func (m Bits) True(i int) bool {
	if m.v&(1<<i) != 0 {
		return true
	}
	return false
}

func (m Bits) Set(i int, b bool) (r Bits) {
	r.v = m.v &^ (1 << i)
	if b {
		r.v |= 1 << i
	}
	return
}

func (m Bits) And(n Bits) (r Bits) {
	r.v = m.v & n.v
	return
}

func (m Bits) AndNot(n Bits) (r Bits) {
	r.v = m.v &^ n.v
	return
}

func (m Bits) Or(n Bits) (r Bits) {
	r.v = m.v | n.v
	return
}

func (m Bits) Xor(n Bits) (r Bits) {
	r.v = m.v ^ n.v
	return
}

func (m Bits) Not() (r Bits) {
	r.v = m.v ^ all
	return
}

func (m Bits) ForTrue(n int, f func(i int)) {
	if n > size {
		n = size
	}
	for i := 0; i < n; i++ {
		if m.True(i) {
			f(i)
		}
	}
}

func (m Bits) For(n int, f func(i int, c bool)) {
	if n > size {
		n = size
	}
	for i := 0; i < n; i++ {
		f(i, m.True(i))
	}
}
