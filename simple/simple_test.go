package simple

import (
	"github.com/duncanpierce/vector/simple/lanes"
	"reflect"
	"testing"
)

func TestAddBroadcast(t *testing.T) {
	x := [4]int{1, 2, 1, 4}
	y, z := [4]int{}, [4]int{}
	Broadcast(&y, nil, 10)
	Add[int](&z, nil, &x, &y)

	if !reflect.DeepEqual(z, [4]int{11, 12, 11, 14}) {
		t.Fail()
	}

	m := lanes.Bool{}
	Equal[int](&m, nil, &z, &[4]int{11, 11, 11, 11})

	if !lanes.IsSet(&m, 0) {
		t.Fail()
	}
	if lanes.IsSet(&m, 1) {
		t.Fail()
	}
	if !lanes.IsSet(&m, 2) {
		t.Fail()
	}
	if lanes.IsSet(&m, 3) {
		t.Fail()
	}

	Broadcast(&z, &m, 99)
	if !reflect.DeepEqual(z, [4]int{99, 12, 99, 14}) {
		t.Fail()
	}
}
