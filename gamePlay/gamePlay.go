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
var ExistsBlockList [1024]graphics.Block
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

func drawBlock(screen *ebiten.Image, block graphics.Block) {

	for index := range block.BlockList {
		subBlock := block.BlockList[index]
		if subBlock.Exists {
			ebitenutil.DrawRect(screen, subBlock.X-1, subBlock.Y, BasicLength+1, BasicLength+1, colornames.Black)
			ebitenutil.DrawRect(screen, subBlock.X, subBlock.Y+1, BasicLength-1, BasicLength-1, block.Color)
		}
	}
}

func addBlockToStack() {
	for index := range ExistsBlockList {
		if ExistsBlockList[index].Name == "" {
			ExistsBlockList[index] = CurrentBlock
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
	for index := range ExistsBlockList {
		if ExistsBlockList[index].Name != "" {
			drawBlock(screen, ExistsBlockList[index])
		}
	}

	// 然后绘制当前操作的方块的位置
	drawBlock(screen, CurrentBlock)
}

func moveLinesDown() {
	var list = [20]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for index := range ExistsBlockList {
		for subIndex := range ExistsBlockList[index].BlockList {
			positionY := int(ExistsBlockList[index].BlockList[subIndex].Y-5) / 30
			if positionY > 5 && ExistsBlockList[index].BlockList[subIndex].Exists {
				list[positionY]++
			}
		}
	}

	for i := 0; i < 4; i++ {
		for index := range list {
			lineNumber := 19 - index
			if list[lineNumber] == 10 {
				for j := range ExistsBlockList {
					for k := range ExistsBlockList[j].BlockList {
						if int(ExistsBlockList[j].BlockList[k].Y) <= lineNumber*30+5 && ExistsBlockList[j].BlockList[k].Exists {
							ExistsBlockList[j].BlockList[k].Y += BasicLength
						}
					}
				}
				break
			}
		}
	}
}

func cleanLines() {
	for i := 0; i < 10; i++ {
		lineNumberPositionY := int(MaxYPosition) - int(BasicLength)*i
		count := 0
		for blockIndex := range ExistsBlockList {
			for subBLockIndex := range ExistsBlockList[blockIndex].BlockList {
				if int(ExistsBlockList[blockIndex].BlockList[subBLockIndex].Y) == lineNumberPositionY {
					count++
				}
			}

			if count == 10 {
				for interBlockIndex := range ExistsBlockList {
					for subBLockIndex := range ExistsBlockList[interBlockIndex].BlockList {
						if int(ExistsBlockList[interBlockIndex].BlockList[subBLockIndex].Y) == lineNumberPositionY {
							ExistsBlockList[interBlockIndex].BlockList[subBLockIndex].Exists = false
						}
					}
				}
			}
		}
	}
}

func GameMainFunction(screen *ebiten.Image) {
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
		CurrentBlock = generateNewBlock()
	}
}

func WaitRetry() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("retry game")
		CurrentBlock = generateNewBlock()
		for index := range ExistsBlockList {
			ExistsBlockList[index] = graphics.Block{}
		}
		GameStatus = "gaming"
	}
}

func GetGameInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		fmt.Println("block rotate (not finish)")
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		fmt.Println("move down")
		MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("move left")
		MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("move right")
		MoveRight()
	}
}
