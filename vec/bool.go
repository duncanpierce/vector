package vec

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
