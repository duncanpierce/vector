package scalable

func (z *bunch[Element, Array]) Mul(x, y Bunch[Element]) Bunch[Element] {
	// TODO how to deal with misalignment here?
	// We can assume element sizes are the same because the element types are forced to be the same
	return z
}
