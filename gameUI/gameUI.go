package gameUI

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	_ "image/png"
)

func DrawUI(screen *ebiten.Image) {
	// gray grid
	var x, y float64
	x = 5
	y = 5
	for x <= 305 {
		ebitenutil.DrawLine(screen, x, 5, x, 605, colornames.Gray)
		x += 30
	}
	for y <= 605 {
		ebitenutil.DrawLine(screen, 5, y, 305, y, colornames.Gray)
		y += 30
	}

	// game box
	frameColor := colornames.Red
	ebitenutil.DrawLine(screen, 5, 5, 455, 5, frameColor)
	ebitenutil.DrawLine(screen, 5, 605, 455, 605, frameColor)
	ebitenutil.DrawLine(screen, 5, 5, 5, 605, frameColor)
	ebitenutil.DrawLine(screen, 305, 5, 305, 605, frameColor)
	ebitenutil.DrawLine(screen, 455, 5, 455, 605, frameColor)
}

func DrawGameOverUI(screen *ebiten.Image) {
	file, _, err := ebitenutil.NewImageFromFile("media/gameover.png")
	if err != nil {
		return
	}
	screen.DrawImage(file, nil)
}
