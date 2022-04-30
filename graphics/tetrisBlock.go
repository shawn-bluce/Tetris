package graphics

import (
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
)

type SubBlock struct {
	X      float64
	Y      float64
	Exists bool
	Color  color.RGBA
}

type Block struct {
	BlockList [4]SubBlock
	Color     color.RGBA
}

var colorList [9]color.RGBA = [9]color.RGBA{
	colornames.Red, colornames.Yellow, colornames.Blue,
	colornames.Tan, colornames.Cornflowerblue, colornames.Deeppink,
	colornames.Peru, colornames.Orange, colornames.Blueviolet,
}

var blockList [7]string = [7]string{"S", "Z", "RZ", "L", "RL", "T", "I"}

func GenerateNewBlock() Block {
	thisColor := colorList[rand.Int()%9]
	thisBlockName := blockList[rand.Int()%7]
	block := Block{
		Color: thisColor,
	}

	blockList := [4][2]float64{}

	switch thisBlockName {
	case "S":
		blockList[0] = [2]float64{0, 0}
		blockList[1] = [2]float64{1, 0}
		blockList[2] = [2]float64{0, 1}
		blockList[3] = [2]float64{1, 1}
	case "Z":
		blockList[0] = [2]float64{0, 0}
		blockList[1] = [2]float64{1, 0}
		blockList[2] = [2]float64{1, 1}
		blockList[3] = [2]float64{2, 1}
	case "RZ":
		blockList[0] = [2]float64{1, 0}
		blockList[1] = [2]float64{2, 0}
		blockList[2] = [2]float64{0, 1}
		blockList[3] = [2]float64{1, 1}
	case "L":
		blockList[0] = [2]float64{0, 0}
		blockList[1] = [2]float64{0, 1}
		blockList[2] = [2]float64{0, 2}
		blockList[3] = [2]float64{1, 2}
	case "RL":
		blockList[0] = [2]float64{1, 0}
		blockList[1] = [2]float64{1, 1}
		blockList[2] = [2]float64{2, 0}
		blockList[3] = [2]float64{2, 1}
	case "T":
		blockList[0] = [2]float64{1, 0}
		blockList[1] = [2]float64{0, 1}
		blockList[2] = [2]float64{1, 1}
		blockList[3] = [2]float64{2, 1}
	case "I":
		blockList[0] = [2]float64{0, 0}
		blockList[1] = [2]float64{1, 0}
		blockList[2] = [2]float64{2, 0}
		blockList[3] = [2]float64{3, 0}
	}

	block.BlockList[0] = SubBlock{(blockList[0][0]+4)*30 + 5, blockList[0][1]*30 + 5, true, thisColor}
	block.BlockList[1] = SubBlock{(blockList[1][0]+4)*30 + 5, blockList[1][1]*30 + 5, true, thisColor}
	block.BlockList[2] = SubBlock{(blockList[2][0]+4)*30 + 5, blockList[2][1]*30 + 5, true, thisColor}
	block.BlockList[3] = SubBlock{(blockList[3][0]+4)*30 + 5, blockList[3][1]*30 + 5, true, thisColor}
	return block
}
