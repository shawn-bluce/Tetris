package gamePlay

import (
	"Tetris/graphics"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
)

var CurrentBlock graphics.Block
var BasicLength float64 = 30
var MinXPosition float64 = 5
var MaxXPosition float64 = 305
var MinYPosition float64 = 5
var MaxYPosition float64 = 605
var ExistsBlockList [1024]graphics.Block
var lastMoveTime int64 = 0

func drawBlock(screen *ebiten.Image, block graphics.Block) {

	for index := range block.BlockList {
		ebitenutil.DrawRect(screen, block.BlockList[index].X, block.BlockList[index].Y, BasicLength, BasicLength, block.Color)
	}
}

func addBlockToStack() {
	for index := range ExistsBlockList {
		if ExistsBlockList[index].Name == "" {
			ExistsBlockList[index] = CurrentBlock
			//fmt.Println("take current block append to ExistsBlockList")
			break
		}
	}
}

func generateNewBlock() graphics.Block {
	//fmt.Println("generate a new block to replace current block")
	square := graphics.GetSquareBlock()
	return square
}

func DrawGameLive(screen *ebiten.Image) {
	// 首先绘制现有的方块
	for index := range ExistsBlockList {
		if ExistsBlockList[index].Name != "" {
			drawBlock(screen, ExistsBlockList[index])
		}
	}

	// 然后绘制当前操作的方块的位置
	drawBlock(screen, CurrentBlock)
}

func GameMainFunction(screen *ebiten.Image) {
	if !TouchBottomBlockOrWall() {
		now := time.Now().UnixMicro()
		if now > lastMoveTime+500000 {
			for index := range CurrentBlock.BlockList {
				CurrentBlock.BlockList[index].Y += BasicLength
			}
			lastMoveTime = now
		}
	} else {
		//fmt.Println("current block is bottom")
		addBlockToStack()
		CurrentBlock = generateNewBlock()
	}
}

func GetGameInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		fmt.Println("do not done.")
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		fmt.Println("do not done.")
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		MoveToLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		MoveToRight()
	}
}
