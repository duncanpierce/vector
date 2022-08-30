package vec

func Set[Z Bool](z Z, lane int, value bool) {
	l := z.lanes()
	if lane < l.nLanes {
		*l.mask &^= 1 << lane
		if value {
			*l.mask |= 1 << lane
		}
	}
}

func SetAll[Z Bool](z Z) {
	l := z.lanes()
	for i := 0; i < l.nLanes; i++ {
		Set(z, i, true)
	}
}

func SetNone[Z Bool](z Z) {
	l := z.lanes()
	for i := 0; i < l.nLanes; i++ {
		Set(z, i, false)
	}
}

func SetFirst[Z Bool](z Z) {
	Set(z, 0, true)
}

func SetLast[Z Bool](z Z) {
	Set(z, z.lanes().nLanes-1, true)
}

func IsSet[X Bool](x X, lane int) bool {
	l := x.lanes()
	if lane < l.nLanes && (*l.mask&(1<<lane) != 0) {
		return true
	}
	return false
}

func Any[X Bool](x X) bool {
	return Count(x) != 0
}

func All[X Bool](x X) bool {
	return Count(x) == x.lanes().nLanes
}

func None[X Bool](x X) bool {
	return Count(x) == 0
}

func Count[X Bool](x X) int {
	l := x.lanes()
	n := 0
	for i := 0; i < l.nLanes; i++ {
		if IsSet(x, i) {
			n++
		}
	}
	return n
}
