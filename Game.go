package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	numRows int32 = 20
	numCols int32 = 10
	cellSize int32 = 30
)

type Game struct {
	Grid Grid
	Blocks []BlockInterface
	GameOver bool
	CurrentBlock BlockInterface
	NextBlock BlockInterface
  	totalScore int32
}

func (g *Game) Initialize() {
	g.Grid.Initialize(cellSize, numRows, numCols)
	g.GameOver = false

	g.Blocks = g.GetAllBlocks()
	g.CurrentBlock = g.GetRandomBlock()
	g.CurrentBlock.Initialize(cellSize)
	g.NextBlock = g.GetRandomBlock()
	g.NextBlock.Initialize(cellSize)
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
	blocks := g.GetAllBlocks()

	randomIndex := rand.Intn(len(blocks))
	return blocks[randomIndex]
}

func (g *Game) Draw() {
	g.Grid.Draw()
	g.CurrentBlock.Draw()

	g.drawScore()
	g.drawNextBlock()
}

func (g *Game) DoBlinking(rowsToBlink []int) {
	if len(rowsToBlink) == 0 {
		return
	}

	gridCopy := g.Grid.GetGridCopy()

	for i := 0; i < 2; i++ {
		for i := range rowsToBlink {
			g.Grid.ClearRow(int32(rowsToBlink[i]))
		}

		rl.BeginDrawing()
		g.Grid.Draw()
		rl.EndDrawing()

		rl.WaitTime(0.2)

		g.Grid.SetGrid(gridCopy)
		fmt.Print() 														// TODO: Remove

		rl.BeginDrawing()
		g.Grid.Draw()
		rl.EndDrawing()

		rl.WaitTime(0.2)
	}
}

func (g *Game) drawNextBlock() {
	const (
		paddingTop int32 = 20
		paddingLeft int32 = 3

		boxWidth int32 = 150
		boxHeight int32 = 120

		positionX int32 = windowWidth - panelWidth / 2 - 65
		positionY int32 = 160

		fontSize int32 = 20
	)

	rl.DrawRectangle(positionX, positionY, boxWidth, boxHeight, rl.Black)
	rl.DrawRectangleLines(positionX, positionY, boxWidth, boxHeight, rl.White)

	rl.DrawText("NEXT", positionX + 40 + paddingLeft, positionY + paddingTop, fontSize, rl.White)

	var offsetX int32 = 7
	var offsetY int32 = 9
	g.NextBlock.Move(offsetX, offsetY)

	if b, ok := g.NextBlock.(*IBlock); ok {
		var x int32 = -10
		var y int32 = 20

		cells := b.GetCellPositions()
		for _, square := range cells {
			rl.DrawRectangle(square.Column * b.cellSize + 1 + x, square.Row * b.cellSize + 1 + y, b.cellSize - 1, b.cellSize - 1, b.colors[b.Id])
		}
	} else if b, ok := g.NextBlock.(*OBlock); ok {
		var x int32 = -15
		var y int32 = 0

		cells := b.GetCellPositions()
		for _, square := range cells {
			rl.DrawRectangle(square.Column * b.cellSize + 1 + x, square.Row * b.cellSize + 1 + y, b.cellSize - 1, b.cellSize - 1, b.colors[b.Id])
		}
	} else {
		g.NextBlock.Draw()
	}

	g.NextBlock.Move(-offsetX, -offsetY)
}

func (g *Game) drawScore() {
	const (
		paddingTop int32 = 20
		paddingLeft int32 = 3
		scoreWidth int32 = 120
		scoreHeight int32 = 120

		positionX int32 = windowWidth - panelWidth / 2 - 50
		positionY int32 = 20

		fontSize int32 = 20
	)

	rl.DrawRectangle(positionX, positionY, scoreWidth, scoreHeight, rl.Black)
	rl.DrawRectangleLines(positionX, positionY, scoreWidth, scoreHeight, rl.White)

	rl.DrawText("SCORE", positionX + 20 + paddingLeft, positionY + paddingTop, fontSize, rl.White)
	rl.DrawText(fmt.Sprintf("%0*d", 8, g.totalScore), positionX + 10 + paddingLeft, positionY + 50 + paddingTop, fontSize, rl.White)
}

func (g *Game) IsBlockOutside() bool {
	tiles := g.CurrentBlock.GetCellPositions()
	for _, tile := range tiles {
		if g.Grid.IsCellOutside(tile.Row, tile.Column) {
			return true
		}
	}

	return false
}

func (g *Game) RotateBlock() {
	g.CurrentBlock.Rotate()
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.RotateUndo()
	}
}

func (g *Game) LockBlock() {
	tiles := g.CurrentBlock.GetCellPositions()
	for _, tile := range tiles {
		g.Grid.grid[tile.Row][tile.Column] = g.CurrentBlock.GetId()
	}

	g.CurrentBlock = g.NextBlock
	g.NextBlock = g.GetRandomBlock()
	g.NextBlock.Initialize(cellSize)

	rowsToClear := g.Grid.GetRowsToClear()
	g.DoBlinking(rowsToClear)

	g.totalScore += g.Grid.ClearFullRows() * 100

	if !g.BlockFits() {
		g.GameOver = true
	}
}

func (g *Game) BlockFits() bool {
	tiles := g.CurrentBlock.GetCellPositions()
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

	g.CurrentBlock.Move(0, -1)
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.Move(0, 1)
	}
}


func (g *Game) MoveBlockRight() {
	if g.GameOver {
		return
	}

	g.CurrentBlock.Move(0, 1)
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.Move(0, -1)
	}
}

func (g *Game) MoveBlockDown() bool {
	if g.GameOver {
		return false
	}

	g.CurrentBlock.Move(1, 0)
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.Move(-1, 0)
		g.LockBlock()
		return true
	}

	return false
}

