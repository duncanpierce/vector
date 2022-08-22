package experiment

import (
	"reflect"
	"testing"
)

func TestVectorGenerics(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	b := [4]int{10, 11, 12, 13}
	r := Add[int, [4]int](a, b)
	if !reflect.DeepEqual(r, [4]int{11, 13, 15, 17}) {
		t.Errorf("it doesn't work!")
	}
}
