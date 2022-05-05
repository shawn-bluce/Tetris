package gamePlay

import (
	"Tetris/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"time"
)

var CurrentBlock graphics.Block
var NextBlock graphics.Block
var ExistsBlockMap [10][20]graphics.SubBlock
var BasicLength float64 = 30
var LastMoveTime int64 = 0
var LastOperateTime int64 = 0
var OperateTimeInterval int64 = 100000
var MoveTimeInterval int64 = 5000 * 20
var GameStatus string = "wait"

func addBlockToStack() {
	for subBlockIndex := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[subBlockIndex]
		x := int(block.X-5) / 30
		y := int(block.Y-5) / 30
		if 0 <= x && x <= 9 && 0 <= y && y <= 19 {
			ExistsBlockMap[x][y].Exists = true
			ExistsBlockMap[x][y].Color = block.Color
		}
	}
}

func DrawGameLive(screen *ebiten.Image) {
	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			if ExistsBlockMap[x][y].Exists {
				positionX := float64(x*30 + 5)
				positionY := float64(y*30 + 5)
				ebitenutil.DrawRect(screen, positionX-1, positionY, BasicLength+1, BasicLength+1, colornames.Black)
				ebitenutil.DrawRect(screen, positionX, positionY+1, BasicLength-1, BasicLength-1, ExistsBlockMap[x][y].Color)
			}
		}
	}

	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		ebitenutil.DrawRect(screen, block.X-1, block.Y, BasicLength+1, BasicLength+1, colornames.Black)
		ebitenutil.DrawRect(screen, block.X, block.Y+1, BasicLength-1, BasicLength-1, block.Color)
	}

	for i := range NextBlock.BlockList {
		block := NextBlock.BlockList[i]
		ebitenutil.DrawRect(screen, block.X+226, block.Y+61, BasicLength-1, BasicLength-1, block.Color)
	}
}

func cleanLines() {
	for y := 19; y > 0; y-- { // from bottom to top
		count := 0
		for x := 0; x < 10; x++ {
			if ExistsBlockMap[x][y].Exists {
				count++
			}
		}

		if count == 10 { // line need clean
			for x := 0; x < 10; x++ {
				ExistsBlockMap[x][y].Exists = false
			}

			for subY := y; subY > 0; subY-- {
				for x := 0; x < 10; x++ {
					ExistsBlockMap[x][subY].Exists = ExistsBlockMap[x][subY-1].Exists
					ExistsBlockMap[x][subY].Color = ExistsBlockMap[x][subY-1].Color
				}
			}
			y++ // move down, retry
		}
	}
}

func GameMainFunction(screen *ebiten.Image) {
	if !CurrentBlock.BlockList[0].Exists {
		CurrentBlock = NextBlock
		NextBlock = graphics.GenerateNewBlock()
	}

	if !TouchBottomBlockOrWall() {
		now := time.Now().UnixMicro()
		if now > LastMoveTime+MoveTimeInterval {
			for index := range CurrentBlock.BlockList {
				CurrentBlock.BlockList[index].Y += BasicLength
			}
			LastMoveTime = now
		}
	} else {
		addBlockToStack()
		cleanLines()
		CurrentBlock = graphics.Block{}
		for i := 0; i < 4; i++ {
			CurrentBlock.BlockList[i].Exists = false
			CurrentBlock.BlockList[i].X = -10
			CurrentBlock.BlockList[i].Y = -10
		}
	}
}

func WaitStart() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		ExistsBlockMap = [10][20]graphics.SubBlock{}
		for i := 0; i < 10; i++ {
			for j := 0; j < 20; j++ {
				ExistsBlockMap[i][j].Exists = false
				ExistsBlockMap[i][j].X = -10
				ExistsBlockMap[i][j].Y = -10
				ExistsBlockMap[i][j].Color = colornames.Black
			}
		}
		GameStatus = "gaming"
	}
}

func GetGameInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		RotateBlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		MoveRight()
	}
}
