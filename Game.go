package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	numRows int32 = 20
	numCols int32 = 10
	cellSize int32 = 30
)

var totalScore int32 = 0

var CurrentBlock BlockInterface
var NextBlock BlockInterface

type Game struct {
	Grid Grid
	Blocks []BlockInterface
	GameOver bool
}


func (g *Game) Initialize() {
	g.Grid.Initialize(cellSize, numRows, numCols)
	g.GameOver = false

	g.Blocks = g.GetAllBlocks()
	CurrentBlock = g.GetRandomBlock()
	CurrentBlock.Initialize(cellSize)
	NextBlock = g.GetRandomBlock()
	NextBlock.Initialize(cellSize)
}


func (g *Game) GetAllBlocks() []BlockInterface {
	return []BlockInterface{
		&IBlock{},
		&JBlock{},
		&LBlock{},
		&OBlock{},
		&SBlock{},
		&TBlock{},
		&ZBlock{},
	}
}

func (g *Game) GetRandomBlock() BlockInterface {
	rand.Seed(time.Now().UnixNano())
	blocks := g.GetAllBlocks()

	randomIndex := rand.Intn(len(blocks))
	return blocks[randomIndex]
}



func (g *Game) Draw() {
	g.Grid.Draw()
	CurrentBlock.Draw()
	g.Grid.drawScore()
}

func (g *Game) IsBlockOutside() bool {
	tiles := CurrentBlock.GetCellPositions()
	for _, tile := range tiles {
		if g.Grid.IsCellOutside(tile.Row, tile.Column) {
			return true
		}
	}

	return false
}

func (g *Game) RotateBlock() {
	CurrentBlock.Rotate()
	if g.IsBlockOutside() || !g.BlockFits() {
		CurrentBlock.RotateUndo()
	}
}

func (g *Game) LockBlock() {
	tiles := CurrentBlock.GetCellPositions()
	for _, tile := range tiles {
		g.Grid.grid[tile.Row][tile.Column] = CurrentBlock.GetId()
	}

	CurrentBlock = NextBlock
	NextBlock = g.GetRandomBlock()
	NextBlock.Initialize(cellSize)

	totalScore += g.Grid.ClearFullRows() * 100

	if !g.BlockFits() {
		g.GameOver = true
	}
}

func (g *Game) BlockFits() bool {
	tiles := CurrentBlock.GetCellPositions()
	for _, tile := range tiles {
		if !g.Grid.IsCellEmpty(tile.Row, tile.Column) {
			return false
		}
	}

	return true
}

func (g *Game) HandleInput() {
	keyPressed := rl.GetKeyPressed()
	switch keyPressed {
	case rl.KeyLeft:
		g.MoveBlockLeft()
	case rl.KeyRight:
		g.MoveBlockRight()
	case rl.KeyDown:
		g.MoveBlockDown()
	case rl.KeyUp:
		g.RotateBlock()
	case rl.KeySpace:
		for !g.MoveBlockDown() {
		}
	}
}

func (g *Game) MoveBlockLeft() {
	if g.GameOver {
		return
	}

	CurrentBlock.Move(0, -1)
	if g.IsBlockOutside() || !g.BlockFits() {
		CurrentBlock.Move(0, 1)
	}
}


func (g *Game) MoveBlockRight() {
	if g.GameOver {
		return
	}

	CurrentBlock.Move(0, 1)
	if g.IsBlockOutside() || !g.BlockFits() {
		CurrentBlock.Move(0, -1)
	}
}

func (g *Game) MoveBlockDown() bool {
	if g.GameOver {
		return false
	}

	CurrentBlock.Move(1, 0)
	if g.IsBlockOutside() || !g.BlockFits() {
		CurrentBlock.Move(-1, 0)
		g.LockBlock()
		return true
	}

	return false
}

