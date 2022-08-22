package fluentvec

type (
	Predicate uint64
)

func (p Predicate) IsTrue(i int) bool {
	if p&(1<<i) != 0 {
		return true
	}
	return false
}

func (p Predicate) ForRange(n int, f func(i int)) {
	for i := 0; i < n; i++ {
		if p.IsTrue(i) {
			f(i)
		}
	}
}

func None() Predicate {
	return 0
}

func All() Predicate {
	return 0xffffffffffffffff
}
