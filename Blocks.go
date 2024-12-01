package main

// import rl "github.com/gen2brain/raylib-go/raylib"


type LBlock struct {
	Block
	
}

func (lb *LBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 1
	lb.Cells = map[int][]Point{
		0: {{0, 2}, {1, 0}, {1, 1}, {1, 2}},
		1: {{0, 1}, {1, 1}, {2, 1}, {2, 2}},
		2: {{1, 0}, {1, 1}, {1, 2}, {2, 0}},
		3: {{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	}
	lb.Move(0, 3)
}

type JBlock struct {
	Block
	
}

func (lb *JBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 2
	lb.Cells = map[int][]Point{
		0: {{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		1: {{0, 1}, {0, 2}, {1, 1}, {2, 1}},
		2: {{1, 0}, {1, 1}, {1, 2}, {2, 2}},
		3: {{0, 1}, {1, 1}, {2, 1}, {2, 0}},
	}
	lb.Move(0, 3)
}

type IBlock struct {
	Block
	
}

func (lb *IBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 3
	lb.Cells = map[int][]Point{
		0: {{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		1: {{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		2: {{2, 0}, {2, 1}, {2, 2}, {2, 3}},
		3: {{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	}
	lb.Move(-1, 3)
}

type OBlock struct {
	Block
	
}

func (lb *OBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 4
	lb.Cells = map[int][]Point{
		0: {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}
	lb.Move(0, 4)
}


type SBlock struct {
	Block
	
}

func (lb *SBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 5
	lb.Cells = map[int][]Point{
		0: {{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		1: {{0, 1}, {1, 1}, {1, 2}, {2, 2}},
		2: {{1, 1}, {1, 2}, {2, 0}, {2, 1}},
		3: {{0, 0}, {1, 0}, {1, 1}, {2, 1}},
	}
	lb.Move(0, 3)
}

type TBlock struct {
	Block
	
}

func (lb *TBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 6
	lb.Cells = map[int][]Point{
		0: {{0, 1}, {1, 0}, {1, 1}, {1, 2}},
		1: {{0, 1}, {1, 1}, {1, 2}, {2, 1}},
		2: {{1, 0}, {1, 1}, {1, 2}, {2, 1}},
		3: {{0, 1}, {1, 0}, {1, 1}, {2, 1}},
	}
	lb.Move(0, 3)
}

type ZBlock struct {
	Block
	
}

func (lb *ZBlock) Initialize(cellSize int32) {
	lb.Block.Initialize(cellSize)

	lb.Id = 7
	lb.Cells = map[int][]Point{
		0: {{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		1: {{0, 2}, {1, 1}, {1, 2}, {2, 1}},
		2: {{1, 0}, {1, 1}, {2, 1}, {2, 2}},
		3: {{0, 1}, {1, 0}, {1, 1}, {2, 0}},
	}
	lb.Move(0, 3)
}
