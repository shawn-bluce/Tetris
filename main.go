package main

import (
	"Tetris/gamePlay"
	"Tetris/gameUI"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	// 每个 tick 都会调用该函数
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 每一帧都会调用该函数
	gameUI.DrawUI(screen)
	gamePlay.DrawGame(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 520, 610
}

func main() {
	ebiten.SetWindowSize(530, 620)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}