package mpi

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/scalable"
)

type (
	// Slice is a DataSource that reads values from a slice and stores them in a scalable.Bunch
	Slice[T any] struct {
		Var   *scalable.Bunch[T]
		Slice []T
	}

	// ActiveLanes is a DataSource that stores the number of active lanes in a scalable.Predicate
	ActiveLanes struct {
		Var *scalable.Predicate
	}

	Counter[T constraintsExt.Number] struct {
		Var    *scalable.Bunch[T]
		buffer []T
	}

	ScalarCounter[T constraintsExt.Number] struct {
		Var *T
	}
)

var _ DataSource = &Slice[int]{}
var _ DataSource = &ActiveLanes{}
var _ DataSource = &Counter[int]{}
var _ DataSource = &ScalarCounter[int]{}

func (f Slice[T]) AlignmentOffset() int {
	//TODO implement me
	panic("implement me")
}

func (f Slice[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (f Slice[T]) Load(index int, count int) {
	//(*f.Var).Load(f.Slice[index:index+count])
	//TODO implement me
	panic("implement me")
}

func (a ActiveLanes) AlignmentOffset() int {
	return -1
}

func (a ActiveLanes) Len() int {
	return -1
}

func (a ActiveLanes) Load(index int, count int) {
	*a.Var = scalable.Predicate{} // TODO represent required number of lanes
}

func (c ScalarCounter[T]) AlignmentOffset() int {
	return -1
}

func (c ScalarCounter[T]) Len() int {
	return -1
}

func (c ScalarCounter[T]) Load(index int, count int) {
	*c.Var += T(count)
}

func (c Counter[T]) AlignmentOffset() int {
	return -1
}

func (c Counter[T]) Len() int {
	return -1
}

func (c Counter[T]) Load(index int, count int) {
	// Load stores index+0 .. index+count in c.Var vector
	panic("implement me")
}
