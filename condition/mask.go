package condition

type Mask struct {
	v uint64
}

const (
	size = 64
	all  = 0xffffffffffffffff
)

func All() (r Mask) {
	return Mask{all}
}

func One() (r Mask) {
	return Mask{1}
}

func (m Mask) True(i int) bool {
	if m.v&(1<<i) != 0 {
		return true
	}
	return false
}

func (m Mask) Set(i int, b bool) (r Mask) {
	r.v = m.v &^ (1 << i)
	if b {
		r.v |= 1 << i
	}
	return
}

func (m Mask) And(n Mask) (r Mask) {
	r.v = m.v & n.v
	return
}

func (m Mask) AndNot(n Mask) (r Mask) {
	r.v = m.v &^ n.v
	return
}

func (m Mask) Or(n Mask) (r Mask) {
	r.v = m.v | n.v
	return
}

func (m Mask) Xor(n Mask) (r Mask) {
	r.v = m.v ^ n.v
	return
}

func (m Mask) Not(n Mask) (r Mask) {
	r.v = m.v ^ all
	return
}

func (m Mask) ForTrue(n int, f func(i int)) {
	if n > size {
		n = size
	}
	for i := 0; i < n; i++ {
		if m.True(i) {
			f(i)
		}
	}
}

func (m Mask) For(n int, f func(i int, c bool)) {
	if n > size {
		n = size
	}
	for i := 0; i < n; i++ {
		f(i, m.True(i))
	}
}
