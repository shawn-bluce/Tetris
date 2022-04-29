package gamePlay

import (
	"Tetris/graphics"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"time"
)

var CurrentBlock graphics.Block
var ExistsBlockMap [10][20]graphics.SubBlock
var BasicLength float64 = 30
var MinXPosition float64 = 5
var MaxXPosition float64 = 305
var MinYPosition float64 = 5
var MaxYPosition float64 = 605
var LastMoveTime int64 = 0
var LastOperateTime int64 = 0
var OperateTimeInterval int64 = 500000
var MoveTimeInterval int64 = 50000
var GameStatus string = "wait"
var CurrentLastPositionX float64 = 0
var CurrentLastPositionY float64 = 0

func addBlockToStack() {
	for subBlockIndex := range CurrentBlock.BlockList {
		for i := range ExistsBlockMap {
			for j := range ExistsBlockMap[i] {
				if !ExistsBlockMap[i][j].Exists {
					ExistsBlockMap[i][j].X = CurrentBlock.BlockList[subBlockIndex].X
					ExistsBlockMap[i][j].Y = CurrentBlock.BlockList[subBlockIndex].Y
					ExistsBlockMap[i][j].Exists = true
					ExistsBlockMap[i][j].Color = CurrentBlock.Color
				}
			}
		}
	}
}

func DrawGameLive(screen *ebiten.Image) {
	for x := range ExistsBlockMap {
		for y := range ExistsBlockMap[x] {
			block := ExistsBlockMap[x][y]
			if block.Exists {
				ebitenutil.DrawRect(screen, block.X-1, block.Y, BasicLength+1, BasicLength+1, colornames.Black)
				ebitenutil.DrawRect(screen, block.X, block.Y+1, BasicLength-1, BasicLength-1, block.Color)
			}
		}
	}

	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		if block.X != CurrentLastPositionX || block.Y != CurrentLastPositionY {
			fmt.Println("current: (", block.X, ",", block.Y, ")")
			CurrentLastPositionX = block.X
			CurrentLastPositionY = block.Y
		}
		ebitenutil.DrawRect(screen, block.X-1, block.Y, BasicLength+1, BasicLength+1, colornames.Black)
		ebitenutil.DrawRect(screen, block.X, block.Y+1, BasicLength-1, BasicLength-1, block.Color)
	}
}

func moveLinesDown() {

}

func cleanLines() {

}

func GameMainFunction(screen *ebiten.Image) {
	if !CurrentBlock.BlockList[0].Exists {
		CurrentBlock = graphics.GenerateNewBlock()
	}

	if !TouchBottomBlockOrWall() {
		now := time.Now().UnixMicro()
		if now > LastMoveTime+OperateTimeInterval {
			for index := range CurrentBlock.BlockList {
				CurrentBlock.BlockList[index].Y += BasicLength
			}
			LastMoveTime = now
		}
	} else {
		fmt.Println("current block is bottom")
		addBlockToStack()
		moveLinesDown()
		cleanLines()
		CurrentBlock = graphics.Block{}
		CurrentBlock.BlockList[0].Exists = false
	}
}

func WaitRetry() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("retry game")
		ExistsBlockMap = [10][20]graphics.SubBlock{}
		GameStatus = "gaming"
	}
}

func GetGameInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {

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
