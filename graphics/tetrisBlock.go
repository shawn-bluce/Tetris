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

func GenerateNewBlock() Block {
	thisColor := colorList[rand.Int()%9]
	block := Block{
		Color: thisColor,
	}
	block.BlockList[0] = SubBlock{125, 5, true, thisColor}
	block.BlockList[1] = SubBlock{155, 5, true, thisColor}
	block.BlockList[2] = SubBlock{125, 35, true, thisColor}
	block.BlockList[3] = SubBlock{155, 35, true, thisColor}
	return block
}
