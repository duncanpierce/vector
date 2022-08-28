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
	compatible[Z, X](zEl, xEl, "x")
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i))
	}
}

func binary[Z, XY any](z Vector[Z], x, y Vector[XY], f func(x, y XY) Z) {
	xEl, yEl, zEl := x.elements(), y.elements(), z.elements()
	compatible[Z, XY](zEl, xEl, "x")
	compatible[Z, XY](zEl, yEl, "y")
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i), yEl.readIndex(i))
	}
}

func compatible[Z, X any](z elements[Z], x elements[X], xIdent string) {
	if x.broadcast {
		if len(x.slice) > len(z.slice) {
			panic(fmt.Sprintf("broadcastable value %s cannot be longer than vector z", xIdent))
		}
	} else {
		if len(x.slice) != len(z.slice) {
			panic(fmt.Sprintf("vector %s must be the same length as vector z", xIdent))
		}
	}
}
