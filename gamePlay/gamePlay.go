package gamePlay

import (
	"Tetris/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

func DrawBlock(screen *ebiten.Image, block graphics.Block) {
	ebitenutil.DrawRect(screen, block.BlockPartA.X, block.BlockPartA.Y, block.BlockPartA.Width, block.BlockPartA.Height, colornames.Green)
}

var value = 0.0

func DrawGame(screen *ebiten.Image) {
	square := graphics.GetSquareBlock()
	if value > 300 {
		value -= 300
	} else {
		value += 2
	}
	square.BlockPartA.Y += value
	DrawBlock(screen, square)
}
