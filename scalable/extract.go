package scalable

func ReturnSlice[Element any](b Bunch[Element]) (result []Element) {
	b.ForRange(func(value Element) {
		result = append(result, value)
	})
	return
}

func ReturnChan[Element any](c chan<- Element, b Bunch[Element]) {
	b.ForRange(func(value Element) {
		c <- value
	})
	return
}
