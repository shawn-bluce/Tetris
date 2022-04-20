package graphics

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type subBlock struct {
	X float64
	Y float64
}

type Block struct {
	Name       string
	BlockPartA subBlock
	BlockPartB subBlock
	BlockPartC subBlock
	BlockPartD subBlock
	Color      color.RGBA
}

func GetSquareBlock() Block {
	return Block{
		Name:       "square",
		BlockPartA: subBlock{125, 5},
		BlockPartB: subBlock{155, 5},
		BlockPartC: subBlock{125, 35},
		BlockPartD: subBlock{155, 35},
		Color:      colornames.Red,
	}
}
