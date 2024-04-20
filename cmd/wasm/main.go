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

type Game struct {
	board      *conway.Board
	pixels     []byte
	currentTPS int
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.board.RandomizeBoard()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.board.TogglePause()
	}

	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 {
		g.currentTPS -= 2
		if g.currentTPS <= 0 {
			g.currentTPS = 2
		}
		ebiten.SetTPS(g.currentTPS)
	}

	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 {
		g.currentTPS += 2
		if g.currentTPS > 256 {
			g.currentTPS = 256
		}
		ebiten.SetTPS(g.currentTPS)
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
