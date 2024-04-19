package conway

var (
	// Give all of the delta's (offsets) to get the eight adjacent neighbors of a Coordinate in the Cartesian plane.
	// Note the lack of {0,0}, meaning a cell is not a neighbor to itself.
	neighborCoordinateDeltas []coordinate = []coordinate{
		{-1, -1},
		{-1, +0},
		{-1, +1},
		{+0, -1},
		{+0, +1},
		{+1, -1},
		{+1, +0},
		{+1, +1},
	}
)

type coordinate struct {
	x int
	y int
}
