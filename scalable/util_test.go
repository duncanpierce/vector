package scalable

import "testing"

type (
	FailWriter[Element any] struct {
		t *testing.T
	}
)

var _ Writer[int] = FailWriter[int]{}

func (f FailWriter[Element]) Write(value Element) {
	f.t.Errorf("should not have written a value but got %v", value)
}
