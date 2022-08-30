package vec

func Zeroing[B Bool](mask B) *Mask {
	return &Mask{
		zero: true,
		m:    *mask.lanes().mask,
	}
}

func Blending[B Bool](mask B) *Mask {
	return &Mask{
		zero: false,
		m:    *mask.lanes().mask,
	}
}
