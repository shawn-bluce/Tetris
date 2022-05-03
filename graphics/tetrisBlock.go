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
	BlockList         [4]SubBlock
	Color             color.RGBA
	AllDirectionsList [4][4][2]float64
	DirectionId       int
}

var colorList [7]color.RGBA = [7]color.RGBA{
	colornames.Lightyellow, colornames.Lightgreen, colornames.Lightblue, colornames.Lightgray,
	colornames.Lightsalmon, colornames.Lightslategray, colornames.Lightpink,
}

var blockList [7]string = [7]string{"O", "Z", "RZ", "L", "RL", "T", "I"}

func GenerateNewBlock() Block {
	index := rand.Int() % 7
	thisColor := colorList[index]
	thisBlockName := blockList[index]
	block := Block{
		Color: thisColor,
	}
	block.DirectionId = rand.Int() % 4

	allDirectionsList := [4][4][2]float64{}

	switch thisBlockName {
	case "O":
		for i := 0; i < 4; i++ {
			allDirectionsList[i][0] = [2]float64{0, 0}
			allDirectionsList[i][1] = [2]float64{1, 0}
			allDirectionsList[i][2] = [2]float64{0, 1}
			allDirectionsList[i][3] = [2]float64{1, 1}
		}
	case "Z":
		allDirectionsList[0][0] = [2]float64{0, 0}
		allDirectionsList[0][1] = [2]float64{1, 0}
		allDirectionsList[0][2] = [2]float64{1, 1}
		allDirectionsList[0][3] = [2]float64{2, 1}

		allDirectionsList[1][0] = [2]float64{1, 0}
		allDirectionsList[1][1] = [2]float64{0, 1}
		allDirectionsList[1][2] = [2]float64{1, 1}
		allDirectionsList[1][3] = [2]float64{0, 2}

		allDirectionsList[2][0] = allDirectionsList[0][0]
		allDirectionsList[2][1] = allDirectionsList[0][1]
		allDirectionsList[2][2] = allDirectionsList[0][2]
		allDirectionsList[2][3] = allDirectionsList[0][3]

		allDirectionsList[3][0] = allDirectionsList[1][0]
		allDirectionsList[3][1] = allDirectionsList[1][1]
		allDirectionsList[3][2] = allDirectionsList[1][2]
		allDirectionsList[3][3] = allDirectionsList[1][3]
	case "RZ":
		allDirectionsList[0][0] = [2]float64{1, 0}
		allDirectionsList[0][1] = [2]float64{2, 0}
		allDirectionsList[0][2] = [2]float64{0, 1}
		allDirectionsList[0][3] = [2]float64{1, 1}

		allDirectionsList[1][0] = [2]float64{0, 0}
		allDirectionsList[1][1] = [2]float64{1, 0}
		allDirectionsList[1][2] = [2]float64{1, 1}
		allDirectionsList[1][3] = [2]float64{1, 2}

		allDirectionsList[2][0] = allDirectionsList[0][0]
		allDirectionsList[2][1] = allDirectionsList[0][1]
		allDirectionsList[2][2] = allDirectionsList[0][2]
		allDirectionsList[2][3] = allDirectionsList[0][3]

		allDirectionsList[3][0] = allDirectionsList[1][0]
		allDirectionsList[3][1] = allDirectionsList[1][1]
		allDirectionsList[3][2] = allDirectionsList[1][2]
		allDirectionsList[3][3] = allDirectionsList[1][3]
	case "L":
		allDirectionsList[0][0] = [2]float64{0, 0}
		allDirectionsList[0][1] = [2]float64{0, 1}
		allDirectionsList[0][2] = [2]float64{0, 2}
		allDirectionsList[0][3] = [2]float64{1, 2}

		allDirectionsList[1][0] = [2]float64{0, 0}
		allDirectionsList[1][1] = [2]float64{1, 0}
		allDirectionsList[1][2] = [2]float64{2, 0}
		allDirectionsList[1][3] = [2]float64{0, 1}

		allDirectionsList[2][0] = [2]float64{0, 0}
		allDirectionsList[2][1] = [2]float64{1, 0}
		allDirectionsList[2][2] = [2]float64{1, 1}
		allDirectionsList[2][3] = [2]float64{1, 2}

		allDirectionsList[3][0] = [2]float64{2, 0}
		allDirectionsList[3][1] = [2]float64{0, 1}
		allDirectionsList[3][2] = [2]float64{1, 1}
		allDirectionsList[3][3] = [2]float64{2, 1}
	case "RL":
		allDirectionsList[0][0] = [2]float64{1, 0}
		allDirectionsList[0][1] = [2]float64{1, 1}
		allDirectionsList[0][2] = [2]float64{1, 2}
		allDirectionsList[0][3] = [2]float64{0, 2}

		allDirectionsList[1][0] = [2]float64{0, 0}
		allDirectionsList[1][1] = [2]float64{0, 1}
		allDirectionsList[1][2] = [2]float64{1, 1}
		allDirectionsList[1][3] = [2]float64{2, 1}

		allDirectionsList[2][0] = [2]float64{0, 0}
		allDirectionsList[2][1] = [2]float64{1, 0}
		allDirectionsList[2][2] = [2]float64{0, 1}
		allDirectionsList[2][3] = [2]float64{0, 2}

		allDirectionsList[3][0] = [2]float64{0, 0}
		allDirectionsList[3][1] = [2]float64{1, 0}
		allDirectionsList[3][2] = [2]float64{2, 0}
		allDirectionsList[3][3] = [2]float64{2, 1}
	case "T":
		allDirectionsList[0][0] = [2]float64{1, 0}
		allDirectionsList[0][1] = [2]float64{0, 1}
		allDirectionsList[0][2] = [2]float64{1, 1}
		allDirectionsList[0][3] = [2]float64{2, 1}

		allDirectionsList[1][0] = [2]float64{1, 0}
		allDirectionsList[1][1] = [2]float64{1, 1}
		allDirectionsList[1][2] = [2]float64{2, 1}
		allDirectionsList[1][3] = [2]float64{1, 2}

		allDirectionsList[2][0] = [2]float64{0, 0}
		allDirectionsList[2][1] = [2]float64{1, 0}
		allDirectionsList[2][2] = [2]float64{2, 0}
		allDirectionsList[2][3] = [2]float64{1, 1}

		allDirectionsList[3][0] = [2]float64{1, 0}
		allDirectionsList[3][1] = [2]float64{0, 1}
		allDirectionsList[3][2] = [2]float64{1, 1}
		allDirectionsList[3][3] = [2]float64{1, 2}
	case "I":
		allDirectionsList[0][0] = [2]float64{0, 0}
		allDirectionsList[0][1] = [2]float64{1, 0}
		allDirectionsList[0][2] = [2]float64{2, 0}
		allDirectionsList[0][3] = [2]float64{3, 0}

		allDirectionsList[1][0] = [2]float64{1, 0}
		allDirectionsList[1][1] = [2]float64{1, 1}
		allDirectionsList[1][2] = [2]float64{1, 2}
		allDirectionsList[1][3] = [2]float64{1, 3}

		allDirectionsList[2][0] = allDirectionsList[0][0]
		allDirectionsList[2][1] = allDirectionsList[0][1]
		allDirectionsList[2][2] = allDirectionsList[0][2]
		allDirectionsList[2][3] = allDirectionsList[0][3]

		allDirectionsList[3][0] = allDirectionsList[1][0]
		allDirectionsList[3][1] = allDirectionsList[1][1]
		allDirectionsList[3][2] = allDirectionsList[1][2]
		allDirectionsList[3][3] = allDirectionsList[1][3]
	}

	block.AllDirectionsList = allDirectionsList
	block.BlockList[0] = SubBlock{(block.AllDirectionsList[block.DirectionId][0][0]+4)*30 + 5, block.AllDirectionsList[block.DirectionId][0][1]*30 + 5, true, thisColor}
	block.BlockList[1] = SubBlock{(block.AllDirectionsList[block.DirectionId][1][0]+4)*30 + 5, block.AllDirectionsList[block.DirectionId][1][1]*30 + 5, true, thisColor}
	block.BlockList[2] = SubBlock{(block.AllDirectionsList[block.DirectionId][2][0]+4)*30 + 5, block.AllDirectionsList[block.DirectionId][2][1]*30 + 5, true, thisColor}
	block.BlockList[3] = SubBlock{(block.AllDirectionsList[block.DirectionId][3][0]+4)*30 + 5, block.AllDirectionsList[block.DirectionId][3][1]*30 + 5, true, thisColor}
	return block
}
