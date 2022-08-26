package mpi

type (
	Range []DataSource

	DataSource interface {
		AlignmentOffset() int
		Len() int
		Load(index int, count int)
	}

	Process[State any] func(state *State)

	// LoopBody is a function that executes the main body of the loop. `state` has been preloaded with data and the number of items is passed in `nItems`.
	// The function should update the state, which it is free to use to accumulate values over different invocations of the loop.
	LoopBody[State any] func(nItems int, state *State)

	Lanes struct {
		// opaque
	}

	Partitioner func(r Range) Partition

	// Partition describes the strategy for partitioning the input data into groups for faster processing.
	Partition struct {
		// HeadLen is the number of items to be processed in the head of the loop
		HeadLen int

		// Iterations is the number of full blocks to be processed in the main body of the loop. Each element in the returned slice indicates a partition of the data that may be processed in an independent goroutine.
		Iterations []int

		// TailLen is the number of items to be processed in the tail of the loop
		TailLen int
	}
)

func For[State any](init Process[State], r Range, body Process[State]) {
	var state State
	if init != nil {
		init(&state)
	}
	panic("implement this")
}
