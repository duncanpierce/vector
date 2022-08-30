package vec

import "fmt"

type elements[E any] struct {
	slice     []E
	broadcast bool
}

func (x elements[E]) readIndex(i int) E {
	if x.broadcast {
		return x.slice[0]
	} else {
		return x.slice[i]
	}
}

func unary[Z, X any](z Vector[Z], x Vector[X], f func(x X) Z) {
	xEl, zEl := x.elements(), z.elements()
	compatible[Z, X]("z", zEl, "x", xEl)
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i))
	}
}

func binary[Z, XY any](z Vector[Z], x, y Vector[XY], f func(x, y XY) Z) {
	xEl, yEl, zEl := x.elements(), y.elements(), z.elements()
	compatible[Z, XY]("z", zEl, "x", xEl)
	compatible[Z, XY]("z", zEl, "y", yEl)
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i), yEl.readIndex(i))
	}
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
