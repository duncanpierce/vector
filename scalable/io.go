package scalable

type (
	Reader[Element any] interface {
		Read() (value Element, ok bool)
	}

	Writer[Element any] interface {
		Write(value Element)
	}

	sliceReader[Element any] struct {
		slice *[]Element
	}

	sliceWriter[Element any] struct {
		slice *[]Element
	}

	chanReader[Element any] struct {
		channel <-chan Element
	}

	chanWriter[Element any] struct {
		channel chan<- Element
	}
)

func FromSlice[Element any](s *[]Element) Reader[Element] {
	return sliceReader[Element]{s}
}

func ToSlice[Element any](s *[]Element) Writer[Element] {
	return sliceWriter[Element]{s}
}

func FromChan[Element any](c <-chan Element) Reader[Element] {
	return chanReader[Element]{c}
}

func ToChan[Element any](c chan<- Element) Writer[Element] {
	return chanWriter[Element]{c}
}

func (s sliceReader[Element]) Read() (value Element, ok bool) {
	if len(*s.slice) < 1 {
		return
	}
	value = (*s.slice)[0]
	ok = true
	*s.slice = (*s.slice)[1:]
	return
}

func (s sliceWriter[Element]) Write(value Element) {
	*s.slice = append(*s.slice, value)
}

func (c chanWriter[Element]) Write(value Element) {
	c.channel <- value
}

func (c chanReader[Element]) Read() (value Element, ok bool) {
	value, ok = <-c.channel
	return
}

func Extract[Element any](b Bunch[Element]) (result []Element) {
	b.Store(ToSlice(&result))
	return
}
