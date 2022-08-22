package scalable

type (
	Reader[Element any] interface {
		Read() (value Element, ok bool)
	}

	Writer[Element any] interface {
		Write(value Element)
	}

	Slice[Element any] struct {
		Elements []Element
	}

	Chan[Element any] struct {
		Elements chan Element
	}
)

var _ Reader[int] = &Slice[int]{}
var _ Writer[int] = &Slice[int]{}
var _ Reader[int] = &Chan[int]{}
var _ Writer[int] = &Chan[int]{}

func (s *Slice[Element]) Write(value Element) {
	s.Elements = append(s.Elements, value)
}

func (s *Slice[Element]) Read() (value Element, ok bool) {
	if len(s.Elements) < 1 {
		return
	}
	value = s.Elements[0]
	ok = true
	s.Elements = s.Elements[1:]
	return
}

func (c *Chan[Element]) Write(value Element) {
	c.Elements <- value
}

func (c *Chan[Element]) Read() (value Element, ok bool) {
	value, ok = <-c.Elements
	return
}

func Extract[Element any](b Bunch[Element]) (result []Element) {
	dest := &Slice[Element]{}
	b.Store(dest)
	return dest.Elements
}
