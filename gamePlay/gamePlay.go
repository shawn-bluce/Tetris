package gamePlay

import (
	"Tetris/graphics"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var currentBlock graphics.Block
var basicLength float64 = 30
var minXPosition float64 = 5
var maxXPosition float64 = 305
var minYPosition float64 = 5
var maxYPosition float64 = 605
var existsBlockList [1024]graphics.Block

func drawBlock(screen *ebiten.Image, block graphics.Block) {
	ebitenutil.DrawRect(screen, block.BlockPartA.X, block.BlockPartA.Y, basicLength, basicLength, block.Color)
	ebitenutil.DrawRect(screen, block.BlockPartB.X, block.BlockPartB.Y, basicLength, basicLength, block.Color)
	ebitenutil.DrawRect(screen, block.BlockPartC.X, block.BlockPartC.Y, basicLength, basicLength, block.Color)
	ebitenutil.DrawRect(screen, block.BlockPartD.X, block.BlockPartD.Y, basicLength, basicLength, block.Color)
}

func currentBlockIsAtBottom() bool {
	if (currentBlock.BlockPartA.Y+basicLength >= maxYPosition) || (currentBlock.BlockPartB.Y+basicLength >= maxYPosition) || (currentBlock.BlockPartC.Y+basicLength >= maxYPosition) || (currentBlock.BlockPartD.Y+basicLength >= maxYPosition) {
		return true
	} else {

		for index := range existsBlockList {
			if (currentBlock.BlockPartA.X == existsBlockList[index].BlockPartA.X) && (currentBlock.BlockPartA.Y+basicLength >= existsBlockList[index].BlockPartA.Y) ||
				(currentBlock.BlockPartA.X == existsBlockList[index].BlockPartB.X) && (currentBlock.BlockPartA.Y+basicLength >= existsBlockList[index].BlockPartB.Y) ||
				(currentBlock.BlockPartA.X == existsBlockList[index].BlockPartC.X) && (currentBlock.BlockPartA.Y+basicLength >= existsBlockList[index].BlockPartC.Y) ||
				(currentBlock.BlockPartA.X == existsBlockList[index].BlockPartD.X) && (currentBlock.BlockPartA.Y+basicLength >= existsBlockList[index].BlockPartD.Y) ||
				(currentBlock.BlockPartB.X == existsBlockList[index].BlockPartA.X) && (currentBlock.BlockPartB.Y+basicLength >= existsBlockList[index].BlockPartA.Y) ||
				(currentBlock.BlockPartB.X == existsBlockList[index].BlockPartB.X) && (currentBlock.BlockPartB.Y+basicLength >= existsBlockList[index].BlockPartB.Y) ||
				(currentBlock.BlockPartB.X == existsBlockList[index].BlockPartC.X) && (currentBlock.BlockPartB.Y+basicLength >= existsBlockList[index].BlockPartC.Y) ||
				(currentBlock.BlockPartB.X == existsBlockList[index].BlockPartD.X) && (currentBlock.BlockPartB.Y+basicLength >= existsBlockList[index].BlockPartD.Y) ||
				(currentBlock.BlockPartC.X == existsBlockList[index].BlockPartA.X) && (currentBlock.BlockPartC.Y+basicLength >= existsBlockList[index].BlockPartA.Y) ||
				(currentBlock.BlockPartC.X == existsBlockList[index].BlockPartB.X) && (currentBlock.BlockPartC.Y+basicLength >= existsBlockList[index].BlockPartB.Y) ||
				(currentBlock.BlockPartC.X == existsBlockList[index].BlockPartC.X) && (currentBlock.BlockPartC.Y+basicLength >= existsBlockList[index].BlockPartC.Y) ||
				(currentBlock.BlockPartC.X == existsBlockList[index].BlockPartD.X) && (currentBlock.BlockPartC.Y+basicLength >= existsBlockList[index].BlockPartD.Y) ||
				(currentBlock.BlockPartD.X == existsBlockList[index].BlockPartA.X) && (currentBlock.BlockPartD.Y+basicLength >= existsBlockList[index].BlockPartA.Y) ||
				(currentBlock.BlockPartD.X == existsBlockList[index].BlockPartB.X) && (currentBlock.BlockPartD.Y+basicLength >= existsBlockList[index].BlockPartB.Y) ||
				(currentBlock.BlockPartD.X == existsBlockList[index].BlockPartC.X) && (currentBlock.BlockPartD.Y+basicLength >= existsBlockList[index].BlockPartC.Y) ||
				(currentBlock.BlockPartD.X == existsBlockList[index].BlockPartD.X) && (currentBlock.BlockPartD.Y+basicLength >= existsBlockList[index].BlockPartD.Y) {
				return true
			}
		}

		return false
	}
}

func addBlockToStack() {
	for index := range existsBlockList {
		if existsBlockList[index].Name == "" {
			existsBlockList[index] = currentBlock
			fmt.Println("加入列表")
			break
		}
	}
}

func generateNewBlock() graphics.Block {
	square := graphics.GetSquareBlock()
	return square
}

func DrawGameLive(screen *ebiten.Image) {
	// 首先绘制现有的方块
	for index := range existsBlockList {
		if existsBlockList[index].Name != "" {
			drawBlock(screen, existsBlockList[index])
		}
	}

	// 然后绘制当前操作的方块的位置
	drawBlock(screen, currentBlock)
}

func GameMainFunction(screen *ebiten.Image) {
	if !currentBlockIsAtBottom() {
		currentBlock.BlockPartA.Y += 3
		currentBlock.BlockPartB.Y += 3
		currentBlock.BlockPartC.Y += 3
		currentBlock.BlockPartD.Y += 3
	} else {
		fmt.Println("触底")
		addBlockToStack()
		currentBlock = generateNewBlock()
	}
}
