package gamePlay

import "fmt"

// 不要看了，没用脑子写的

func TouchLeftBlockOrWall() bool {
	for index := range CurrentBlock.BlockList {
		if CurrentBlock.BlockList[index].X <= MinXPosition {
			return true
		}
	}
	for blockIndex := range ExistsBlockList {
		for currentIndex := range CurrentBlock.BlockList {
			currentSubBlock := CurrentBlock.BlockList[currentIndex]
			judgeBlockList := ExistsBlockList[blockIndex].BlockList
			for existsIndex := range judgeBlockList {
				if currentSubBlock.X == judgeBlockList[existsIndex].X+BasicLength && currentSubBlock.Y == judgeBlockList[existsIndex].Y {
					return true
				}
			}
		}
	}
	return false
}

func TouchRightBlockOrWall() bool {
	for index := range CurrentBlock.BlockList {
		if CurrentBlock.BlockList[index].X >= MaxXPosition-BasicLength {
			fmt.Println("touch the wall")
			return true
		}
	}
	for blockIndex := range ExistsBlockList {
		for currentIndex := range CurrentBlock.BlockList {
			currentSubBlock := CurrentBlock.BlockList[currentIndex]
			judgeBlockList := ExistsBlockList[blockIndex].BlockList
			for existsIndex := range judgeBlockList {
				if currentSubBlock.X == judgeBlockList[existsIndex].X-BasicLength && currentSubBlock.Y == judgeBlockList[existsIndex].Y {
					fmt.Println("touch the block")
					return true
				}
			}
		}
	}
	fmt.Println("can move")
	return false
}

func TouchBottomBlockOrWall() bool {

	for index := range CurrentBlock.BlockList {
		if CurrentBlock.BlockList[index].Y+BasicLength >= MaxYPosition {
			return true
		}
	}

	for blockIndex := range ExistsBlockList {
		for currentIndex := range CurrentBlock.BlockList {
			currentSubBlock := CurrentBlock.BlockList[currentIndex]
			judgeBlockList := ExistsBlockList[blockIndex].BlockList
			for existsIndex := range judgeBlockList {
				if currentSubBlock.X == judgeBlockList[existsIndex].X && currentSubBlock.Y+BasicLength >= judgeBlockList[existsIndex].Y {
					return true
				}
			}
		}
	}

	return false
}

func IsGameOver() bool {
	for index := range CurrentBlock.BlockList {
		if CurrentBlock.BlockList[index].Y == MinYPosition && TouchBottomBlockOrWall() {
			return true
		}
	}
	return false
}
