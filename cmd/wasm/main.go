package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hmcalister/Golang-WASM-Conway/cmd/wasm/conway"
)

const (
	GAME_WIDTH  = 256
	GAME_HEIGHT = 256
)

var (
	UPDATE_PERIODS []int = []int{1, 2, 4, 8, 16, 32}
)

type Game struct {
	board  *conway.Board
	pixels []byte

	gameTPS           int
	updatePeriodIndex int
	updateCounter     int
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.board.RandomizeBoard()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.board.TogglePause()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.updatePeriodIndex += 1
		if g.updatePeriodIndex >= len(UPDATE_PERIODS) {
			g.updatePeriodIndex = len(UPDATE_PERIODS) - 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.updatePeriodIndex -= 1
		if g.updatePeriodIndex < 0 {
			g.updatePeriodIndex = 0
		}

	}

	g.board.NextState()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(g.pixels)
	screen.WritePixels(g.pixels)

	// tpsMsg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	// ebitenutil.DebugPrint(screen, tpsMsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

func main() {
	conwayBoard, err := conway.NewBoard(GAME_WIDTH, GAME_HEIGHT)
	if err != nil {
		log.Fatalf("error during creating board state %v", err.Error())
	}
	conwayBoard.RandomizeBoard()

	game := &Game{
		board:      conwayBoard,
		currentTPS: 16,
		pixels:            make([]byte, 4*GAME_WIDTH*GAME_HEIGHT),
	}

	ebiten.SetWindowSize(2*GAME_WIDTH, 2*GAME_HEIGHT)
	ebiten.SetWindowTitle("Conway's Game Of Life")
	ebiten.SetTPS(game.currentTPS)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("error during run Conway %v", err.Error())
	}
}
