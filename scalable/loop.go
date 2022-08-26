package scalable

type (
	Looper interface {
		// DataAlignment returns offset from the best alignment for the underlying data. If there is no underlying data or alignment makes no difference, returns -1.
		// Vectors of loop range data will be chosen to minimise misaligned memory load penalties. Non-portable, even within a CPU family.
		DataAlignment() (offset int)

		// DataLen returns the number of items in the data source, if any. If this Looper can contain data items but has none, returns 0. If this Looper cannot contain data items, returns -1.
		DataLen() (l int)

		// Initialize is called before the first iteration of the loop in order to allow loop setup steps to be performed.
		Initialize()

		// Load 'length' data items from the data source in preparation for executing one step of the loop body or bodies.
		// Load is not called again before repeating the loop body for unprocessed lanes.
		Load(length int)

		// Iterate is called for each iteration of the loop. If err is not nil, For will terminate as soon as possible, returning the error.
		// If unprocessedLanes.IsEmpty() is false, Iterate will be called again with the remaining lanes. Iterate should not disturb the contents of
		// any lanes other than lanesToProcess. Iterate should not assume that active lanes in lanesToProcess are consecutive.
		Iterate(lanesToProcess Predicate) (unprocessedLanes Predicate, err error)

		// Store is called after Iterate returns no further lanes to process. Any data from the loop should be written to data sources.
		Store()

		// Finalize is called after the last iteration of the loop in order to allow loop finalization steps to be performed.
		Finalize()
	}

	ParallelLooper interface {
		// DataAlignment returns offset from the best alignment for the underlying data. If there is no underlying data or alignment makes no difference, returns -1.
		// Vectors of loop range data will be chosen to minimise misaligned memory load penalties. Non-portable, even within a CPU family.
		DataAlignment() (offset int)

		// DataLen returns the number of items in the data source, if any. If this Looper can contain data items but has none, returns 0. If this Looper cannot contain data items, returns -1.
		DataLen() (l int)

		// PartitionData(minSize, nWays)
		// TODO

		// InitializePartition is called before the first iteration of the loop in order to allow loop setup steps to be performed.
		InitializePartition(partition int)

		// Load 'length' data items from the data source in preparation for executing one step of the loop body or bodies.
		// Load is not called again before repeating the loop body for unprocessed lanes.
		Load(length int)

		// IteratePartition is called for each iteration of the loop. If err is not nil, For will terminate as soon as possible, returning the error.
		// If unprocessedLanes.IsEmpty() is false, IteratePartition will be called again with the remaining lanes. IteratePartition should not disturb the contents of
		// any lanes other than lanesToProcess or any partitions other than partition. IteratePartition should not assume that active lanes in lanesToProcess are consecutive.
		// TODO can we do away with induction variables which couple loop iterations, so we can run each loop body fully independently using processor farm? don't need to partition data then (but do still require var x[]vec.Var instead of var x vec.Var)
		// In fact, we must do away with induction variables or results will be undefined
		IteratePartition(partition int, lanesToProcess Predicate) (unprocessedLanes Predicate, err error)

		// Store is called after Iterate returns no further lanes to process. Any data from the loop should be written to data sources.
		Store()

		// Finalize is called after the last iteration of the loop in order to allow loop finalization steps to be performed.
		Finalize()
	}
)

func For(steps ...[]Looper) error {
	panic("implement this")
}

// ForParallel uses a sync.Pool of worker goroutines to process
func ForParallel(minPartitionLen, maxWorkers int, steps ...[]ParallelLooper) error {
	panic("implement this")
}
