package vec

import (
	"reflect"
	"testing"
)

func TestCompareFloat(t *testing.T) {
	x := Vec2[float64]{1, 2}
	y := Vec2[float64]{0, 4}
	z := Lanes[float64]{}
	LessThan[float64](&z, &x, &y)
	t.Fail()
}

func TestEqualComplex(t *testing.T) {
	x := [2]complex128{1 + 3i, 2}
	y := [2]complex128{1 + 3i, 4}
	z := Lanes[complex128]{}
	Equal[complex128](&z, &x, &y)
	t.Fail()
}

func TestAdd(t *testing.T) {
	x := Vec4[float64]{1, 2, 3, 4}
	y := Vec4[float64]{4, 0, -2, -4}
	z := Vec4[float64]{}

	Add[float64](&z, &x, &y)

	if !reflect.DeepEqual(z, Vec4[float64]{5, 2, 1, 0}) {
		t.Fail()
	}
}

func TestAddBroadcastY(t *testing.T) {
	x := Vec4[float64]{1, 2, 3, 4}
	y := Broadcast[float64](10.0)
	z := Vec4[float64]{}

	Add[float64](&z, &x, y)

	if !reflect.DeepEqual(z, Vec4[float64]{5, 2, 1, 0}) {
		t.Fail()
	}
}

func TestAddComplex(t *testing.T) {
	x := Vec2[complex128]{1 + 2i, 2 + 3i}
	y := Vec2[complex128]{0 + 1i, 4 + 5i}
	z := Vec2[complex128]{}
	Add[complex128](&z, &x, &y)

	if !reflect.DeepEqual(z, Vec2[complex128]{1 + 3i, 6 + 8i}) {
		t.Fail()
	}
}

//func TestSliceScalarBroadcast(t *testing.T) {
//	broadcast := Broadcast[int]{99}
//	slice, isBroadcast := unsafeSliceBroadcast[int, [16]int](&broadcast)
//	if !isBroadcast {
//		t.Fail()
//	}
//	if len(slice) != 1 {
//		t.Fail()
//	}
//	if slice[0] != 99 {
//		t.Fail()
//	}
//}

//func TestSliceVectorBroadcast(t *testing.T) {
//	vector := [4]int{2, 3, 4, 5}
//	slice, isBroadcast := unsafeSliceBroadcast[int, [4]int](&vector)
//	if isBroadcast {
//		t.Fail()
//	}
//	if len(slice) != len(vector) {
//		t.Fail()
//	}
//	if slice[0] != 2 {
//		t.Fail()
//	}
//}

func TestSliceVector(t *testing.T) {
	vector := Vec4[int]{2, 3, 4, 5}
	slice, _ := vector.slice()
	if len(slice) != len(vector) {
		t.Fail()
	}
	if slice[0] != 2 {
		t.Fail()
	}
}

//func TestReplicate(t *testing.T) {
//	small := Vec2[int]{2, 3}
//	big := Vec8[int]{}
//	Replicate[int](&big, &small)
//
//	if !reflect.DeepEqual(big, [8]int{2, 3, 2, 3, 2, 3, 2, 3}) {
//		t.Fail()
//	}
//}

//func TestReplicateScalar(t *testing.T) {
//	big := Vec8[int]{}
//	small := Vec1[int]{99}
//	Replicate[int](&big, &small)
//
//	if !reflect.DeepEqual(big, [8]int{99, 99, 99, 99, 99, 99, 99, 99}) {
//		t.Fail()
//	}
//}

//func TestReplicatePanicsIfBigger(t *testing.T) {
//	small := Vec2[int]{2, 3}
//	big := Vec8[int]{}
//
//	defer func() {
//		if r := recover(); r == nil {
//			t.Fail()
//		}
//	}()
//	Replicate[int](&small, &big)
//}

//func TestReplicateSameSize(t *testing.T) {
//	small := Vec2[int]{2, 3}
//	alsoSmall := Vec2[int]{}
//
//	Replicate[int](&alsoSmall, &small)
//	if !reflect.DeepEqual(alsoSmall, [2]int{2, 3}) {
//		t.Fail()
//	}
//}

func TestConvert(t *testing.T) {
	from := [4]int{1, 2, 3, 4}
	to := [4]float32{}
	Convert[float32, int](&to, &from)

	if !reflect.DeepEqual(to, [4]float32{1, 2, 3, 4}) {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	from := Vec4[int]{1, 2, 3, 4}
	to := Vec4[int]{}
	Copy[int](&to, &from)

	if !reflect.DeepEqual(to, Vec4[int]{1, 2, 3, 4}) {
		t.Fail()
	}
}

func TestCopyScalar(t *testing.T) {
	big := Vec8[int]{}
	scalar := Broadcast(99)
	Copy[int](&big, scalar)

	if !reflect.DeepEqual(big, Vec8[int]{99, 99, 99, 99, 99, 99, 99, 99}) {
		t.Fail()
	}
}
