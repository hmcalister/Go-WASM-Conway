package conway

import (
	"errors"
	"math/rand"
	"slices"
)

type Board struct {
	width              int
	height             int
	currentStatusArray []statusType
	nextStatusArray    []statusType
	paused             bool
}

func NewBoard(width, height int) (*Board, error) {
	if width < 1 || height < 1 {
		return nil, errors.New("board must have width and height larger than one")
	}

	board := &Board{
		width:              width,
		height:             height,
		currentStatusArray: make([]statusType, width*height),
		nextStatusArray:    make([]statusType, width*height),
		paused:             false,
	}

	return board, nil
}

func (b *Board) RandomizeBoard() {
	const RANDOM_ACTIVE_PROPORTION = 0.3
	for i := range b.currentStatusArray {
		if rand.Float32() < RANDOM_ACTIVE_PROPORTION {
			b.currentStatusArray[i] = status_ACTIVE
		} else {
			b.currentStatusArray[i] = status_INACTIVE
		}
	}
}

func (b *Board) TogglePause() {
	b.paused = !b.paused
}

func (b *Board) Draw(pixels []byte) {
	var pixelValue [4]byte
	for cellIndex, cellStatus := range b.currentStatusArray {
		switch cellStatus {
		case status_ACTIVE:
			pixelValue = [4]byte{0xFF, 0xFF, 0xFF, 0xFF}
		case status_INACTIVE:
			pixelValue = [4]byte{0x00, 0x00, 0x00, 0x00}
		}

		targetPixelStartIndex := 4 * cellIndex
		_ = slices.Replace(pixels, targetPixelStartIndex, targetPixelStartIndex+4, pixelValue[0], pixelValue[1], pixelValue[2], pixelValue[3])
	}
}

// Convert a coordinate to an index into a status array.
func (b *Board) coordinateToIndex(c coordinate) int {
	return b.width*c.y + c.x
}

// Convert an index into a status array to a coordinate.
func (b *Board) indexToCoordinate(i int) coordinate {
	return coordinate{
		x: i % b.width,
		y: i / b.width,
	}
}

// Get the active neighbor count of an index.
func (b *Board) getActiveNeighborCount(currentIndex int) int {
	currentCoordinate := b.indexToCoordinate(currentIndex)
	activeCount := 0

	for _, neighborCoordinateDelta := range neighborCoordinateDeltas {
		neighborCoordinate := coordinate{
			currentCoordinate.x + neighborCoordinateDelta.x,
			currentCoordinate.y + neighborCoordinateDelta.y,
		}
		if neighborCoordinate.x < 0 || neighborCoordinate.x >= b.width || neighborCoordinate.y < 0 || neighborCoordinate.y >= b.height {
			continue
		}

		neighborIndex := b.coordinateToIndex(neighborCoordinate)
		activeCount += int(b.currentStatusArray[neighborIndex])
	}

	return activeCount
}

// Update the board state by a single step.
func (b *Board) NextState() {
	if b.paused {
		return
	}

	// First, process the state using the functional implementation of Conway's game in processState
	//
	// This places the next state into b.nextStatusArray
	processState(b.currentStatusArray, b.nextStatusArray, b.getActiveNeighborCount)

	// Then, simply swap the current and next status array. We therefore recycle the memory (reducing allocations)
	//
	// This is similar to a frame buffer.
	b.currentStatusArray, b.nextStatusArray = b.nextStatusArray, b.currentStatusArray
}
