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
	
	panelWidth int32 = 200
	windowWidth int32 = numCols*cellSize + panelWidth
	windowHeight int32 = numRows*cellSize

	FPS int32 = 60
)

type Game struct {
	Grid Grid
	Blocks []BlockInterface
	CurrentBlock BlockInterface
	NextBlock BlockInterface
  	totalScore int32
	gameOver bool
	background rl.Texture2D
}

func (g *Game) Initialize(bg rl.Texture2D) {
	g.Grid.Initialize(cellSize, numRows, numCols)
	g.gameOver = false
	g.background = bg
	g.Blocks = g.GetAllBlocks()
	g.CurrentBlock = g.GetRandomBlock()
	g.CurrentBlock.Initialize(cellSize)
	g.NextBlock = g.GetRandomBlock()
	g.NextBlock.Initialize(cellSize)
}

func (g *Game) Run() {
	g.HandleInput()

	time := g.getSpeed(g.totalScore)
	if (EventTriggered(time)) {
		g.MoveBlockDown()
	}
}

var lastUpdateTime float64 = 0;
func EventTriggered(interval float64) bool {
	currentTime := rl.GetTime()
	if (currentTime - lastUpdateTime) >= interval {
		lastUpdateTime = currentTime
		return true
	}

	return false
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
	g.DrawBackground()
	g.Grid.Draw()
	g.CurrentBlock.Draw()
	g.drawScore()
	g.drawNextBlock()


	if g.gameOver {
		g.DrawGameOver()
	}
}

func (g *Game) DrawBackground() {
	rl.DrawTexture(g.background, 0, 0, LightGrey)
	rl.DrawRectangle(0, 0, windowWidth - panelWidth, windowHeight, LightGrey)
	rl.DrawRectangle(windowWidth - panelWidth, 0, 2, windowHeight, rl.Blue)
}

func (g *Game) DrawGameOver() {
	rl.DrawRectangle(numCols * cellSize / 2 - 45, numRows * cellSize / 2 - 10, 230, 80, rl.Black)
	rl.DrawText("GAME OVER", numCols * cellSize / 2, numRows * cellSize / 2, 20, rl.Green)
	rl.DrawText("Press R to Restart", windowWidth/2-130, windowHeight/2+40, 20, rl.Green)
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


func (g *Game) getSpeed(score int32) float64 {
	switch score {
	case 400:
		return 0.3
	case 800:
		return 0.2
	case 1200:
		return 0.1
	}

	return 0.4
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
		g.gameOver = true
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

var inputTimer int32 = 0
func (g *Game) HandleInput() {
	if inputTimer > 0 {
		inputTimer--
	} else {
		if rl.IsKeyDown(rl.KeyLeft) {
			g.MoveBlockLeft()
			inputTimer = 5
		}
		if rl.IsKeyDown(rl.KeyRight) {
			g.MoveBlockRight()
			inputTimer = 5
		}
		if rl.IsKeyDown(rl.KeyDown) {
			g.MoveBlockDown()
			inputTimer = 5
		}
	}


	if rl.IsKeyPressed(rl.KeyUp) {
		g.RotateBlock()
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		for !g.MoveBlockDown() {
		}
	}
	
	if g.gameOver && rl.IsKeyPressed(rl.KeyR) {
		g.Initialize(g.background)
	}
}


func (g *Game) MoveBlockLeft() {
	if g.gameOver {
		return
	}

	g.CurrentBlock.Move(0, -1)
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.Move(0, 1)
	}
}


func (g *Game) MoveBlockRight() {
	if g.gameOver {
		return
	}

	g.CurrentBlock.Move(0, 1)
	if g.IsBlockOutside() || !g.BlockFits() {
		g.CurrentBlock.Move(0, -1)
	}
}

func (g *Game) MoveBlockDown() bool {
	if g.gameOver {
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

