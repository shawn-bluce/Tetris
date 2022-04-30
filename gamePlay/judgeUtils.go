package gamePlay

// 不要看了，没用脑子写的

func TouchLeftBlockOrWall() bool {
	for i := range CurrentBlock.BlockList {
		block := CurrentBlock.BlockList[i]
		x := int((block.X - 5) / 30)
		y := int((block.Y - 5) / 30)
		if x == 0 || ExistsBlockMap[x-1][y].Exists {
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
		if x == 9 || ExistsBlockMap[x+1][y].Exists {
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
		// todo: 仍然有bug，是这里的判断出了问题
		if y == 19 || ExistsBlockMap[x][y+1].Exists {
			return true
		}
	}
	return false
}

func IsGameOver() bool {
	//for index := range CurrentBlock.BlockList {
	//	if CurrentBlock.BlockList[index].Y == MinYPosition && TouchBottomBlockOrWall() {
	//		return true
	//	}
	//}
	return false
}
