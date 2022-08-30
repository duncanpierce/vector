package vec

func IsSet[X Bool](x X, lane int) bool {
	l := x.lanes()
	if lane < l.nLanes && (*l.mask&(1<<lane) != 0) {
		return true
	}
	return false
}
