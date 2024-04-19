package conway

type statusType int

const (
	status_INACTIVE statusType = 0
	status_ACTIVE   statusType = 1
)

var (
	// Define the state machine for Conway's game of life.
	//
	// The first index is the current state of the cell, while the second is the number of active neighbors the cell has.
	// The default value of this lookup table implements the classic Conway rules. For example, an inactive cell
	// with exactly three active neighbors becomes active, so `conwayStateMachine[0][3]` is `ACTIVE`.
	conwayStateMachine [][]statusType = [][]statusType{
		{status_INACTIVE, status_INACTIVE, status_INACTIVE, status_ACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE},
		{status_INACTIVE, status_INACTIVE, status_ACTIVE, status_ACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE, status_INACTIVE},
	}
)

// Given a current status array, calculate the next state of the game and place this into the nextStatusArray (so memory can be reused).
//
// neighborCountingFunction takes an index (from the status array) and returns the number of active neighbors of that index.
// We pass a function here so the conway game can be purely functional, having no state.
func processState(currentStatusArray []statusType, nextStatusArray []statusType, activeNeighborCountingFunction func(int) int) {
	for index, currentStatus := range currentStatusArray {
		neighborCount := activeNeighborCountingFunction(index)
		nextStatusArray[index] = conwayStateMachine[currentStatus][neighborCount]
	}
}
