package vector

import (
	"testing"
	"vector/condition"
	"vector/vec2"
)

func TestAdd(t *testing.T) {
	v := vec2.Add(condition.All(), [2]int{3, 4}, [2]int{5, 6})
	if v[0] != 8 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 10 {
		t.Errorf("got %v", v[1])
	}
}

func TestMax(t *testing.T) {
	v := vec2.Max(condition.All(), [2]int{3, 4}, [2]int{2, 6})
	if v[0] != 3 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 6 {
		t.Errorf("got %v", v[1])
	}
}
