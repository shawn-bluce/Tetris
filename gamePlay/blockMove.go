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

func RotateBlock() {
	offsetX := (CurrentBlock.BlockList[0].X - 5) / 30
	offsetY := (CurrentBlock.BlockList[0].Y - 5) / 30

	nextX := [4]float64{}
	nextY := [4]float64{}
	nextDirectionId := (CurrentBlock.DirectionId + 1) % 4

	for i := 0; i < 4; i++ {
		//CurrentBlock.BlockList[i].X = ((CurrentBlock.AllDirectionsList[CurrentBlock.DirectionId][i][0] + offsetX) * 30) + 5
		//CurrentBlock.BlockList[i].Y = ((CurrentBlock.AllDirectionsList[CurrentBlock.DirectionId][i][1] + offsetY) * 30) + 5
		nextX[i] = ((CurrentBlock.AllDirectionsList[nextDirectionId][i][0] + offsetX) * 30) + 5
		nextY[i] = ((CurrentBlock.AllDirectionsList[nextDirectionId][i][1] + offsetY) * 30) + 5

		if nextX[i] >= 305 || nextY[i] >= 605 {
			return
		}
	}

	CurrentBlock.DirectionId = nextDirectionId
	for i := 0; i < 4; i++ {
		CurrentBlock.BlockList[i].X = nextX[i]
		CurrentBlock.BlockList[i].Y = nextY[i]
	}
}
