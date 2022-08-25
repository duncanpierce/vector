package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"reflect"
	"testing"
)

func TestCompareFloat(t *testing.T) {
	x := [2]float64{1, 2}
	y := [2]float64{0, 4}
	z := Lanes[float64]{}
	LessThan[float64](&z, &x, &y)
}

func TestEqualComplex(t *testing.T) {
	x := [2]complex128{1 + 3i, 2}
	y := [2]complex128{1 + 3i, 4}
	z := Lanes[complex128]{}
	Equal[complex128](&z, &x, &y)
}

func TestLen(t *testing.T) {
	x := [2]float64{1, 2}
	y := [2]float64{0, 4}
	z := [2]float64{}
	combineNumbers2[float64](&z, &x, &y)
}

func TestAdd(t *testing.T) {
	x := [2]float64{1, 2}
	y := [2]float64{0, 4}
	z := [2]float64{}
	Add[float64](&z, &x, &y)
}

func TestAddComplex(t *testing.T) {
	x := [2]complex128{1, 2}
	y := [2]complex128{0, 4}
	z := [2]complex128{}
	Add[complex128](&z, &x, &y)
}

func TestSliceScalarBroadcast(t *testing.T) {
	broadcast := constraintsExt.Scalar[int]{99}
	slice, isBroadcast := unsafeSliceBroadcast[int, [16]int](&broadcast)
	if !isBroadcast {
		t.Fail()
	}
	if len(slice) != 1 {
		t.Fail()
	}
	if slice[0] != 99 {
		t.Fail()
	}
}

func TestSliceVectorBroadcast(t *testing.T) {
	vector := [4]int{2, 3, 4, 5}
	slice, isBroadcast := unsafeSliceBroadcast[int, [4]int](&vector)
	if isBroadcast {
		t.Fail()
	}
	if len(slice) != len(vector) {
		t.Fail()
	}
	if slice[0] != 2 {
		t.Fail()
	}
}

func TestSliceVector(t *testing.T) {
	vector := [4]int{2, 3, 4, 5}
	slice := unsafeSlice[int, [4]int](&vector)
	if len(slice) != len(vector) {
		t.Fail()
	}
	if slice[0] != 2 {
		t.Fail()
	}
}

func TestReplicate(t *testing.T) {
	small := [2]int{2, 3}
	big := [8]int{}
	Replicate[int, [2]int](&big, &small)

	if !reflect.DeepEqual(big, [8]int{2, 3, 2, 3, 2, 3, 2, 3}) {
		t.Fail()
	}
}
