package main

import (
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

func (gr *Grid) GetGridCopy() [][]int32 {
	gridCopy := make([][]int32, len(gr.grid))
	for i, row := range gr.grid {
		gridCopy[i] = make([]int32, len(row))
		copy(gridCopy[i], row)
	}

	return gridCopy
}

func (gr *Grid) SetGrid(newGrid [][]int32) {
	for i, row := range newGrid {
		copy(gr.grid[i], row)
	}
}

func (gr *Grid) IsCellEmpty(row, column int32) bool {
	return gr.grid[row][column] == 0
}

func (gr *Grid) Draw() {
    for row := range gr.numRows {
        for column := range gr.numCols {
            cellValue := gr.grid[row][column]

            rl.DrawRectangle(
                column*gr.cellSize+1, 
                row*gr.cellSize+1, 
                gr.cellSize-1, 
                gr.cellSize-1, 
                gr.colors[cellValue],
            )
       }
    }
}


func (gr *Grid) ClearRow(row int32) {
	for col := range gr.grid[0] {
		gr.grid[row][col] = 0
	}
}

func (gr *Grid) GetRowsToClear() []int {
	rows := []int {}

	for i := range gr.grid {
		if gr.IsRowFull(int32(i)) {
			rows = append(rows, i)
		}
	}

	return rows
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
