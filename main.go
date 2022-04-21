package main

import (
	"Tetris/gamePlay"
	"Tetris/gameUI"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
)

type Game struct{}

func (g *Game) Update() error {
	// call this function ever tick
	now := time.Now().UnixMicro()
	if now > gamePlay.LastOperateTime+gamePlay.MoveTimeInterval && gamePlay.GameStatus != "stop" {
		gamePlay.GetGameInput()
		gamePlay.LastOperateTime = now
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// call this function ever frame
	if gamePlay.IsGameOver() {
		gameUI.DrawGameOverUI(screen)
		gamePlay.GameStatus = "stop"
		return
	}
	gameUI.DrawUI(screen)
	gamePlay.GameMainFunction(screen)
	gamePlay.DrawGameLive(screen)
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
