package scalable

type (
	Range  []Ranger
	Ranger interface {
		AlignmentOffset() int
		Len() int
		Load(lanes Predicate)
	}
	Store  []Storer
	Storer interface {
	}
	Body[State any] func(lanesToProcess Predicate, vars *State)
)

/*
idea: take a single loop body, if repeat is needed, do lanes.For(...) or for lanes.Remaining() {...}
e.g. for lanes.Remaining() { ... todo = x.IsDuplicate() ... lanes = todo }
*/
func For2[State any](r Range, b Body[State]) {
	var lanes Predicate
	// TODO align data, set up loop head, calculate number of all-lane iterations
	var remaining int
	for remaining > 0 {
		for _, ranger := range r {
			ranger.Load(lanes)

		}
	}
	// TODO if remaining lanes, do loop tail
}
