package graphics

import (
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
)

type subBlock struct {
	X      float64
	Y      float64
	Exists bool
}

type Block struct {
	Name      string
	BlockList [4]subBlock
	Color     color.RGBA
}

var colorList [9]color.RGBA = [9]color.RGBA{
	colornames.Red, colornames.Yellow, colornames.Blue,
	colornames.Tan, colornames.Cornflowerblue, colornames.Deeppink,
	colornames.Peru, colornames.Orange, colornames.Blueviolet,
}

func GetSquareBlock() Block {
	block := Block{
		Name:  "square",
		Color: colorList[rand.Int()%9],
	}
	block.BlockList[0] = subBlock{125, 5, true}
	block.BlockList[1] = subBlock{155, 5, true}
	block.BlockList[2] = subBlock{125, 35, true}
	block.BlockList[3] = subBlock{155, 35, true}
	return block
}
