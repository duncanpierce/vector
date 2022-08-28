package vec

import "fmt"

type elements[E any] struct {
	slice     []E
	broadcast bool
}

func (z elements[E]) assignableFrom(x elements[E], xIdent string) {
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

func (x elements[E]) readIndex(i int) E {
	if x.broadcast {
		return x.slice[0]
	} else {
		return x.slice[i]
	}
}

func unary[E any](z, x Vector[E], f func(x E) E) {
	xEl, zEl := x.elements(), z.elements()
	zEl.assignableFrom(xEl, "x")
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i))
	}
}

func binary[E any](z, x, y Vector[E], f func(x, y E) E) {
	xEl, yEl, zEl := x.elements(), y.elements(), z.elements()
	zEl.assignableFrom(xEl, "x")
	zEl.assignableFrom(yEl, "y")
	for i := 0; i < len(zEl.slice); i++ {
		zEl.slice[i] = f(xEl.readIndex(i), yEl.readIndex(i))
	}
}
