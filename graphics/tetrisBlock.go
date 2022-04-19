package graphics

type subBlock struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Block struct {
	Name       string
	BlockPartA subBlock
	BlockPartB subBlock
}

func GetSquareBlock() Block {
	return Block{
		Name:       "square",
		BlockPartA: subBlock{120, 0, 60, 60},
		BlockPartB: subBlock{0, 0, 0, 0},
	}
}
