package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	grid [][]int32

	colors []rl.Color

	numRows int32
	numCols int32
	cellSize int32
}

func (gr *Grid) Initialize(cellSize, numRows, numCols int32) {
	gr.cellSize = cellSize
	gr.colors = GetCellColors()

	gr.numRows = numRows
	gr.numCols = numCols

	gr.grid = make([][]int32, gr.numRows)
	for i := range gr.grid {
		gr.grid[i] = make([]int32, gr.numCols)
	}
}

func (gr *Grid) IsCellEmpty(row, column int32) bool {
	return gr.grid[row][column] == 0
}

func (gr *Grid) drawScore() {
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
	rl.DrawText(fmt.Sprintf("%0*d", 8, totalScore), positionX + 10 + paddingLeft, positionY + 50 + paddingTop, fontSize, rl.White)
}

func (gr *Grid) drawNextBlock() {
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
	NextBlock.Move(offsetX, offsetY)

	if b, ok := NextBlock.(*IBlock); ok {
		var x int32 = -10
		var y int32 = 20

		cells := b.GetCellPositions()
		for _, square := range cells {
			rl.DrawRectangle(square.Column * b.cellSize + 1 + x, square.Row * b.cellSize + 1 + y, b.cellSize - 1, b.cellSize - 1, b.colors[b.Id])
		}
	} else if b, ok := NextBlock.(*OBlock); ok {
		var x int32 = -15
		var y int32 = 0

		cells := b.GetCellPositions()
		for _, square := range cells {
			rl.DrawRectangle(square.Column * b.cellSize + 1 + x, square.Row * b.cellSize + 1 + y, b.cellSize - 1, b.cellSize - 1, b.colors[b.Id])
		}
	} else {
		NextBlock.Draw()
	}

	NextBlock.Move(-offsetX, -offsetY)

	// NextBlock.Move(-positionX, -positionY)
}

func (gr *Grid) Draw() {
	for row := range gr.numRows {
		for column := range gr.numCols {
			cellValue := gr.grid[row][column]
			rl.DrawRectangle(column * gr.cellSize + 1, row * gr.cellSize + 1, gr.cellSize - 1, gr.cellSize - 1, gr.colors[cellValue])
		}
	}

	gr.drawScore()
	gr.drawNextBlock()
}

func (gr *Grid) IsCellOutside(row, colum int32) bool {
	return row < 0 || colum < 0 || row >= numRows || colum >= numCols
}

func (gr *Grid) IsRowFull(row int32) bool {
	for col := range gr.grid[0] {
		if gr.grid[row][col] == 0 {
			return false
		}
	}

	return true
}

func (gr *Grid) ClearRow(row int32) {
	save_row := make([]int32, len(gr.grid[0]))
	copy(save_row, gr.grid[row])

	for i := 0; i < 3; i++ {

		for col := range gr.grid[0] {
			gr.grid[row][col] = 0
		}

		rl.BeginDrawing()
		gr.Draw()
		rl.EndDrawing()

		if i < 2 {
			time.Sleep(time.Millisecond * 300)
			gr.grid[row] = save_row
			rl.BeginDrawing()
			gr.Draw()
			rl.EndDrawing()
			time.Sleep(time.Millisecond * 300)
		}
	}

	for col := range gr.grid[0] {
		gr.grid[row][col] = 0
	}
}

func (gr *Grid) MoveRowDown(row, numRows int32) {
	for col := range gr.grid[0] {
		gr.grid[row + numRows][col] = gr.grid[row][col]
		gr.grid[row][col] = 0
	}
}

func (gr *Grid) ClearFullRows() int32 {
	var completed int32 = 0
	for row := numRows - 1; row >= 0; row-- {
		if gr.IsRowFull(row) {
			gr.ClearRow(row)
			completed++
		} else if completed > 0 {
			gr.MoveRowDown(row, completed)
		}
	}

	return completed
}

