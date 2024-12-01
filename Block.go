package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BlockInterface interface {
	Initialize(cellSize int32)
	GetId() int32
	Draw()
	Move(rows, cols int32)
	GetCellPositions() []Point
	Rotate()
	RotateUndo()
}

type Block struct {
	Id int32
	Cells map[int][]Point

	cellSize int32
	rotationState int
	colors []rl.Color

	rowOffset, columnOffset int32
}

func (b *Block) Initialize(cellSize int32) {
	b.cellSize = cellSize
	b.colors = GetCellColors()
}

func (b *Block) GetId() int32 {
	return b.Id
}

func (b *Block) Draw() {
	cells := b.GetCellPositions()
	for _, square := range cells {
		rl.DrawRectangle(square.Column * b.cellSize + 1, square.Row * b.cellSize + 1, b.cellSize - 1, b.cellSize - 1, b.colors[b.Id])
	}
}

func (b *Block) Move(rows, cols int32) {
	b.rowOffset += rows
	b.columnOffset += cols
}

func (b *Block) GetCellPositions() []Point {
	tiles := b.Cells[b.rotationState]
	tiles_new := make([]Point, len(tiles))
	for i, tile := range tiles {
		tiles_new[i] = Point{tile.Row + b.rowOffset, tile.Column + b.columnOffset}
	}

	return tiles_new
}

func (b *Block) Rotate() {
	b.rotationState++
	if b.rotationState >= len(b.Cells) {
		b.rotationState = 0
	}
}

func (b *Block) RotateUndo() {
	b.rotationState--
	if b.rotationState < 0 {
		b.rotationState = len(b.Cells) - 1
	}
}