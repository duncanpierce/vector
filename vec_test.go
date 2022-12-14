package vector

import (
	"github.com/duncanpierce/vector/vec16"
	"github.com/duncanpierce/vector/vec2"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	v := vec2.Add(vec2.SelectAll(), [2]int{3, 4}, [2]int{5, 6})
	if v[0] != 8 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 10 {
		t.Errorf("got %v", v[1])
	}
}

func TestMax(t *testing.T) {
	v := vec2.Max(vec2.SelectAll(), [2]int{3, 4}, [2]int{2, 6})
	if v[0] != 3 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 6 {
		t.Errorf("got %v", v[1])
	}
}

func TestMask(t *testing.T) {
	v := vec2.Add(vec2.SelectFirst(1), [2]int{1, 2}, [2]int{3, 4})
	if v[0] != 4 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 0 {
		t.Errorf("got %v", v[1])
	}

	v = vec2.Add(vec2.SelectFirst(1).Not(), [2]int{1, 2}, [2]int{3, 4})
	if v[0] != 0 {
		t.Errorf("got %v", v[0])
	}
	if v[1] != 6 {
		t.Errorf("got %v", v[1])
	}
}

func TestInterlace(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := [8]int{9, 10, 11, 12, 13, 14, 15, 16}

	r := vec16.Interlace(1, a, b)
	if !reflect.DeepEqual(r, [16]int{1, 9, 2, 10, 3, 11, 4, 12, 5, 13, 6, 14, 7, 15, 8, 16}) {
		t.Errorf("failed at interlace 1")
	}

	r = vec16.Interlace(2, a, b)
	if !reflect.DeepEqual(r, [16]int{1, 2, 9, 10, 3, 4, 11, 12, 5, 6, 13, 14, 7, 8, 15, 16}) {
		t.Errorf("failed at interlace 2")
	}

	r = vec16.Interlace(4, a, b)
	if !reflect.DeepEqual(r, [16]int{1, 2, 3, 4, 9, 10, 11, 12, 5, 6, 7, 8, 13, 14, 15, 16}) {
		t.Errorf("failed at interlace 4")
	}

	r = vec16.Interlace(8, a, b)
	if !reflect.DeepEqual(r, [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}) {
		t.Errorf("failed at interlace 8")
	}
}

func TestDeinterlace(t *testing.T) {
	r, s := vec16.Deinterlace(1, [16]int{1, 9, 2, 10, 3, 11, 4, 12, 5, 13, 6, 14, 7, 15, 8, 16})
	checkDeinterlace(t, r, s)

	r, s = vec16.Deinterlace(2, [16]int{1, 2, 9, 10, 3, 4, 11, 12, 5, 6, 13, 14, 7, 8, 15, 16})
	checkDeinterlace(t, r, s)

	r, s = vec16.Deinterlace(4, [16]int{1, 2, 3, 4, 9, 10, 11, 12, 5, 6, 7, 8, 13, 14, 15, 16})
	checkDeinterlace(t, r, s)

	r, s = vec16.Deinterlace(8, [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	checkDeinterlace(t, r, s)
}

func checkDeinterlace(t *testing.T, r, s [8]int) {
	if !reflect.DeepEqual(r, [8]int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("failed on r")
	}
	if !reflect.DeepEqual(s, [8]int{9, 10, 11, 12, 13, 14, 15, 16}) {
		t.Errorf("failed on s")
	}
}

func TestFirstN(t *testing.T) {
	a := [16]int{}
	b := [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	mask := vec16.SelectFirst(3)
	r := vec16.Blend(mask, a, b)
	if !reflect.DeepEqual(r, [16]int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("unexpectedly got %v", r)
	}
}

func TestLastN(t *testing.T) {
	a := [16]int{}
	b := [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	mask := vec16.SelectLast(3)
	r := vec16.Blend(mask, a, b)
	if !reflect.DeepEqual(r, [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 15, 16}) {
		t.Errorf("unexpectedly got %v", r)
	}
}

func TestHorizontalMax(t *testing.T) {
	a := [16]int{1, 2, 6, 10, 44, 3, 17, 19, 23, 33, 12, 2, 6, 43, 13, 16}
	r, ok := vec16.ElementMax(vec16.SelectAll(), a)
	if !ok {
		t.Errorf("should be ok")
	}
	if r != 44 {
		t.Errorf("wrong value returned %v", r)
	}
}

func TestMaskedHorizontalMax(t *testing.T) {
	a := [16]int{1, 2, 6, 10, 44, 3, 17, 19, 23, 33, 12, 2, 6, 43, 13, 16}
	r, ok := vec16.ElementMax(vec16.SelectFirst(4), a)
	if !ok {
		t.Errorf("should be ok")
	}
	if r != 10 {
		t.Errorf("wrong value returned %v", r)
	}
}

func TestNoneMaskedHorizontalMax(t *testing.T) {
	a := [16]int{1, 2, 6, 10, 44, 3, 17, 19, 23, 33, 12, 2, 6, 43, 13, 16}
	_, ok := vec16.ElementMax(vec16.SelectNone(), a)
	if ok {
		t.Errorf("should not be ok")
	}
}
