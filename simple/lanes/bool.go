package lanes

import "github.com/duncanpierce/vector/constraintsExt"

type Bool struct {
	m uint64
}

const allBits = 0xffffffffffffffff

func IsActive(m *Bool, i int) bool {
	if m == nil {
		return true
	}
	return (m.m)&(1<<i) != 0
}

func SetActive(m *Bool, i int, b bool) {
	bit := uint64(1 << i)
	m.m &^= bit
	if b {
		m.m |= bit
	}
}

func RangeActive[E any, Z constraintsExt.Vector[E]](z *Z, m *Bool, f func(i, j int)) {
	j, l := 0, len(*z)
	for i := 0; i < l; i++ {
		if IsActive(m, i) {
			f(i, j)
			j++
		}
	}
}

func RangeInactive[E any, Z constraintsExt.Vector[E]](z *Z, m *Bool, f func(i, j int)) {
	j, l := 0, len(*z)
	for i := 0; i < l; i++ {
		if !IsActive(m, i) {
			f(i, j)
			j++
		}
	}
}

func RangeAll[E any, Z constraintsExt.Vector[E]](z *Z, m *Bool, f func(i, j int, b bool)) {
	j, l := 0, len(*z)
	for i := 0; i < l; i++ {
		active := IsActive(m, i)
		f(i, j, active)
		if active {
			j++
		}
	}
}

func IsAny(b *Bool) bool {
	if b == nil {
		return true
	}
	return b.m != 0
}

func IsNone(b *Bool) bool {
	if b == nil {
		return false
	}
	return b.m == 0
}

func And(z, x, y *Bool) {
	(*z).m = (*x).m & (*y).m
}

func AndNot(z, x, y *Bool) {
	(*z).m = (*x).m &^ (*y).m
}

func Or(z, x, y *Bool) {
	(*z).m = (*x).m | (*y).m
}

func Xor(z, x, y *Bool) {
	(*z).m = (*x).m ^ (*y).m
}

func Not(z, x *Bool) {
	(*z).m = (*x).m ^ allBits
}

func All() *Bool {
	return &Bool{allBits}
}
