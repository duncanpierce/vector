package scalable

func ReturnSlice[Element any](b Bunch[Element]) (result []Element) {
	ProduceSlice(&result, b)
	return
}

func ProduceSlice[Element any](s *[]Element, b Bunch[Element]) {
	b.ForRange(func(value Element) {
		*s = append(*s, value)
	})
	return
}

func ProduceChan[Element any](c chan<- Element, b Bunch[Element]) {
	b.ForRange(func(value Element) {
		c <- value
	})
	return
}
