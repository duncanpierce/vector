package predicate

import "testing"

func TestPattern1(t *testing.T) {
	m := &Bool{}
	SetBlocks(m, 1)
	if m.m != 0b0101010101010101010101010101010101010101010101010101010101010101 {
		t.Fail()
	}
}

func TestPattern2(t *testing.T) {
	m := &Bool{}
	SetBlocks(m, 2)
	if m.m != 0b0011001100110011001100110011001100110011001100110011001100110011 {
		t.Fail()
	}
}

func TestPattern4(t *testing.T) {
	m := &Bool{}
	SetBlocks(m, 4)
	if m.m != 0b0000111100001111000011110000111100001111000011110000111100001111 {
		t.Fail()
	}
}

func TestPattern64(t *testing.T) {
	m := &Bool{}
	SetBlocks(m, 64)
	if m.m != 0b1111111111111111111111111111111111111111111111111111111111111111 {
		t.Fail()
	}
}

func TestFirst(t *testing.T) {
	m := &Bool{}
	SetFirst(m, 10)
	if m.m != 0b0000000000000000000000000000000000000000000000000000001111111111 {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	m := &Bool{}
	SetLast(m, 10)
	if m.m != 0b1111111111000000000000000000000000000000000000000000000000000000 {
		t.Fail()
	}
}
