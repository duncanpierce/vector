package vec64

import (
	"math/bits"
)

type Bool struct {
	v uint64
}

const (
	allBits = (1 << Length) - 1
)

func All() (r Bool) {
	return Bool{allBits}
}

func None() (r Bool) {
	return Bool{0}
}

func Last(n int) (r Bool) {
	return Bool{(allBits << (Length - n)) & allBits}
}

func First(n int) (r Bool) {
	return Bool{allBits >> (Length - n)}
}

func (m Bool) IsTrue(i int) bool {
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

func (m Bool) FirstTrue() (i int, ok bool) {
	firstSet := bits.TrailingZeros64(m.v)
	if firstSet < Length {
		return firstSet, true
	}
	return 0, false
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

func (m Bool) AnyTrue() bool {
	return m.v != 0
}

func (m Bool) AllTrue() bool {
	return m.v == allBits
}

func (m Bool) CountTrue() int {
	return bits.OnesCount64(m.v & allBits)
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
		if m.IsTrue(i) {
			f(i)
		}
	}
}

func (m Bool) For(f func(i int, c bool)) {
	for i := 0; i < Length; i++ {
		f(i, m.IsTrue(i))
	}
}

func ReduceElements[Element any](m Bool, a [Length]Element, f func(x, y Element) Element) (result Element, ok bool) {
	var firstTrue int
	firstTrue, ok = m.FirstTrue()
	if !ok {
		return
	}
	result = a[firstTrue]
	restTrue := m.Set(firstTrue, false)
	restTrue.ForTrue(func(i int) {
		result = f(result, a[i])
	})
	return
}
