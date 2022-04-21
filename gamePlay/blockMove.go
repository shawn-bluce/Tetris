package gamePlay

func MoveLeft() {
	if !TouchLeftBlockOrWall() {
		for index := range CurrentBlock.BlockList {
			CurrentBlock.BlockList[index].X -= BasicLength
		}
	}
}

func MoveRight() {
	if !TouchRightBlockOrWall() {
		for index := range CurrentBlock.BlockList {
			CurrentBlock.BlockList[index].X += BasicLength
		}
	}
}

func MoveDown() {
	if !TouchBottomBlockOrWall() {
		for index := range CurrentBlock.BlockList {
			CurrentBlock.BlockList[index].Y += BasicLength
		}
	}
}
