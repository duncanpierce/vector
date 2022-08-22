package scalable

type (
	Reader[Element any] interface {
		Read() (value Element, ok bool)
	}

	Writer[Element any] interface {
		Write(value Element)
	}

	ReaderWriter[Element any] interface {
		Reader[Element]
		Writer[Element]
	}

	sliceIo[Element any] struct {
		Elements []Element
	}

	chanIo[Element any] struct {
		Elements chan Element
	}
)

var _ Reader[int] = &sliceIo[int]{}
var _ Writer[int] = &sliceIo[int]{}
var _ Reader[int] = &chanIo[int]{}
var _ Writer[int] = &chanIo[int]{}

func Slice[Element any](s []Element) ReaderWriter[Element] {
	return &sliceIo[Element]{s}
}

func Chan[Element any](c chan Element) ReaderWriter[Element] {
	return &chanIo[Element]{c}
}

func (s *sliceIo[Element]) Write(value Element) {
	s.Elements = append(s.Elements, value)
}

func (s *sliceIo[Element]) Read() (value Element, ok bool) {
	if len(s.Elements) < 1 {
		return
	}
	value = s.Elements[0]
	ok = true
	s.Elements = s.Elements[1:]
	return
}

func (c *chanIo[Element]) Write(value Element) {
	c.Elements <- value
}

func (c *chanIo[Element]) Read() (value Element, ok bool) {
	value, ok = <-c.Elements
	return
}

func Extract[Element any](b Bunch[Element]) (result []Element) {
	dest := &sliceIo[Element]{}
	b.Store(dest)
	return dest.Elements
}
