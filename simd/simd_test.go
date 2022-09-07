package simd

import (
	"github.com/duncanpierce/vector/predicate"
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

	m := predicate.Bool{}
	Equal[int](&m, nil, &z, &[4]int{11, 11, 11, 11})

	if !predicate.IsSet(&m, 0) {
		t.Fail()
	}
	if predicate.IsSet(&m, 1) {
		t.Fail()
	}
	if !predicate.IsSet(&m, 2) {
		t.Fail()
	}
	if predicate.IsSet(&m, 3) {
		t.Fail()
	}

	Broadcast(&z, &m, 99)
	if !reflect.DeepEqual(z, [4]int{99, 12, 99, 14}) {
		t.Fail()
	}
}
