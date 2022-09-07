package predicate

type Bool struct {
	m uint64
}

const (
	allBits         = 0xffffffffffffffff
	alternatingBits = 0x5555555555555555
)

func IsSet(m *Bool, i int) bool {
	if m == nil {
		return true
	}
	return (m.m)&(1<<i) != 0
}

func Set(m *Bool, i int, b bool) {
	bit := uint64(1 << i)
	m.m &^= bit
	if b {
		m.m |= bit
	}
}

func SetAll(m *Bool) {
	m.m = allBits
}

func SetNone(m *Bool) {
	m.m = 0
}

func SetAlternating(m *Bool) {
	m.m = alternatingBits
}

func SetBlocks(m *Bool, w int) {
	mask := uint64(0)
	for j, bit, set, i := w, uint64(1), true, 0; i < 64; i, j, bit = i+1, j-1, bit<<1 {
		if j <= 0 {
			j, set = w, !set
		}
		if set {
			mask |= bit
		}
	}
	m.m = mask
}

func RangeActive(l int, m *Bool, f func(i, j int)) {
	j := 0
	for i := 0; i < l; i++ {
		if IsSet(m, i) {
			f(i, j)
			j++
		}
	}
}

func RangeInactive(l int, m *Bool, f func(i, j int)) {
	j := 0
	for i := 0; i < l; i++ {
		if !IsSet(m, i) {
			f(i, j)
			j++
		}
	}
}

func RangeAll(l int, m *Bool, f func(i int, b bool)) {
	for i := 0; i < l; i++ {
		active := IsSet(m, i)
		f(i, active)
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
