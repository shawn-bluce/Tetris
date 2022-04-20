package gamePlay

func MoveToLeft() {
	if !TouchLeftBlockOrWall() {
		for index := range CurrentBlock.BlockList {
			CurrentBlock.BlockList[index].X -= BasicLength
		}
	}
}

func MoveToRight() {
	if !TouchRightBlockOrWall() {
		for index := range CurrentBlock.BlockList {
			CurrentBlock.BlockList[index].X += BasicLength
		}
	}
}
