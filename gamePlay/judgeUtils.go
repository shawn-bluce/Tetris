package gamePlay

func TouchLeftBlockOrWall() bool {
	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		x := int((block.X - 5) / 30)
		y := int((block.Y - 5) / 30)
		if x == 0 || (x-1 >= 0 && y >= 0 && ExistsBlockMap[x-1][y].Exists) {
			return true
		}
	}
	return false
}

func TouchRightBlockOrWall() bool {
	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		x := int((block.X - 5) / 30)
		y := int((block.Y - 5) / 30)
		if x == 9 || (x+1 <= 9 && y >= 0 && ExistsBlockMap[x+1][y].Exists) {
			return true
		}
	}
	return false
}

func TouchBottomBlockOrWall() bool {
	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		x := int((block.X - 5) / 30)
		y := int((block.Y - 5) / 30)
		if y == 19 || (y+1 > 0 && y+1 <= 19 && ExistsBlockMap[x][y+1].Exists) {
			return true
		}
	}
	return false
}

func IsGameOver() bool {
	for i := 0; i < 10; i++ {
		if ExistsBlockMap[i][0].Exists {
			return true
		}
	}
	return false
}
