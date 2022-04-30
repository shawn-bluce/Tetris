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
var LastMoveTime int64 = 0
var LastOperateTime int64 = 0
var OperateTimeInterval int64 = 500000
var MoveTimeInterval int64 = 50000
var GameStatus string = "wait"

func addBlockToStack() {
	for subBlockIndex := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[subBlockIndex]
		x := int(block.X-5) / 30
		y := int(block.Y-5) / 30
		fmt.Println("add", x, y)
		ExistsBlockMap[x][y].Exists = true
		ExistsBlockMap[x][y].Color = block.Color
	}
	for y := 0; y < 20; y++ {
		fmt.Println(
			ExistsBlockMap[0][y].Exists, ExistsBlockMap[1][y].Exists, ExistsBlockMap[2][y].Exists, ExistsBlockMap[3][y].Exists, ExistsBlockMap[4][y].Exists,
			ExistsBlockMap[5][y].Exists, ExistsBlockMap[6][y].Exists, ExistsBlockMap[7][y].Exists, ExistsBlockMap[8][y].Exists, ExistsBlockMap[9][y].Exists,
		)
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
