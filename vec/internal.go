package vec

import "fmt"

type elements[E any] struct {
	slice     []E
	broadcast bool
}

const allBits = uint64(0xFFFFFFFFFFFFFFFF)

func (x elements[E]) readIndex(i int) E {
	if x.broadcast {
		return x.slice[0]
	} else {
		return x.slice[i]
	}
}

func (z elements[E]) masked(m *Mask, f func(index, count int)) {
	mask, zero := allBits, false
	var zeroValue E
	if m != nil {
		mask, zero = m.m, m.zero
	}
	l := len(z.slice)
	for index, count := 0, 0; index < l; index++ {
		if mask&(1<<index) != 0 {
			f(index, count)
			count++
		} else if zero {
			z.slice[index] = zeroValue
		}
	}
}

func unary[Z, X any](m *Mask, z Vector[Z], x Vector[X], f func(x X) Z) {
	xEl, zEl := x.elements(), z.elements()
	compatible[Z, X]("z", zEl, "x", xEl)
	zEl.masked(m, func(index, count int) {
		zEl.slice[index] = f(xEl.readIndex(index))
	})
}

func binary[Z, XY any](m *Mask, z Vector[Z], x, y Vector[XY], f func(x XY, y XY) Z) {
	xEl, yEl, zEl := x.elements(), y.elements(), z.elements()
	compatible[Z, XY]("z", zEl, "x", xEl)
	compatible[Z, XY]("z", zEl, "y", yEl)
	zEl.masked(m, func(index, count int) {
		zEl.slice[index] = f(xEl.readIndex(index), yEl.readIndex(index))
	})
}

func binaryBool[XY any](z Bool, x, y Vector[XY], f func(x, y XY) bool) {
	xEl, yEl, zEl := x.elements(), y.elements(), z.lanes()
	compatibleBool[XY]("z", zEl, "x", xEl)
	compatibleBool[XY]("z", zEl, "y", yEl)
	for i := 0; i < zEl.nLanes; i++ {
		zEl.set(i, f(xEl.readIndex(i), yEl.readIndex(i)))
	}
}

func unaryBool[X any](z Bool, x Vector[X], f func(x X) bool) {
	xEl, zEl := x.elements(), z.lanes()
	compatibleBool[X]("z", zEl, "x", xEl)
	for i := 0; i < zEl.nLanes; i++ {
		zEl.set(i, f(xEl.readIndex(i)))
	}
}

func compatible[Z, X any](zIdent string, z elements[Z], xIdent string, x elements[X]) {
	if x.broadcast {
		if len(x.slice) > len(z.slice) {
			panic(fmt.Sprintf("broadcastable value %s cannot be longer than vector %s", xIdent, zIdent))
		}
	} else {
		if len(x.slice) != len(z.slice) {
			panic(fmt.Sprintf("vector %s must be the same length as vector %s", xIdent, zIdent))
		}
	}
}

func compatibleBool[X any](zIdent string, z lanes, xIdent string, x elements[X]) {
	if x.broadcast {
		if len(x.slice) > z.nLanes {
			panic(fmt.Sprintf("broadcastable value %s cannot be longer than bool %s", xIdent, zIdent))
		}
	} else {
		if len(x.slice) != z.nLanes {
			panic(fmt.Sprintf("vector %s must be the same length as bool %s", xIdent, zIdent))
		}
	}
}
