package vec8

type Bool struct {
	v uint64
}

const (
	allBits = (1 << Length) - 1
)

func Full() (r Bool) {
	return Bool{allBits}
}

func Last(n int) (r Bool) {
	return Bool{(allBits << (Length - n)) & allBits}
}

func First(n int) (r Bool) {
	return Bool{allBits >> (Length - n)}
}

func (m Bool) True(i int) bool {
	if m.v&(1<<i) != 0 {
		return true
	}
	return false
}

func (m Bool) Set(i int, b bool) (r Bool) {
	r.v = m.v &^ (1 << i)
	if b {
		r.v |= 1 << i
	}
	return
}

func (m Bool) And(n Bool) (r Bool) {
	r.v = m.v & n.v
	return
}

func (m Bool) AndNot(n Bool) (r Bool) {
	r.v = m.v &^ n.v
	return
}

func (m Bool) Or(n Bool) (r Bool) {
	r.v = m.v | n.v
	return
}

func (m Bool) Xor(n Bool) (r Bool) {
	r.v = m.v ^ n.v
	return
}

func (m Bool) Not() (r Bool) {
	r.v = m.v ^ allBits
	return
}

func (m Bool) Any() bool {
	return m.v != 0
}

func (m Bool) All() bool {
	return m.v == allBits
}

func (m Bool) ShiftLeft(i int) (r Bool) {
	r.v = m.v << i
	return
}

func (m Bool) ShiftRight(i int) (r Bool) {
	r.v = m.v >> i
	return
}

func (m Bool) ForTrue(f func(i int)) {
	for i := 0; i < Length; i++ {
		if m.True(i) {
			f(i)
		}
	}
}

func (m Bool) For(f func(i int, c bool)) {
	for i := 0; i < Length; i++ {
		f(i, m.True(i))
	}
}
