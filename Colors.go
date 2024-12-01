package main

import rl "github.com/gen2brain/raylib-go/raylib"

var DarkGrey = rl.Color{R: 26, G: 31, B: 40, A: 255}
var Red = rl.Color{R: 255, G: 0, B: 0, A: 255}
var Green = rl.Color{R: 0, G: 255, B: 0, A: 255}
var Blue = rl.Color{R: 0, G: 0, B: 255, A: 255}
var Yellow = rl.Color{R: 255, G: 255, B: 0, A: 255}
var Cyan = rl.Color{R: 0, G: 255, B: 255, A: 255}
var Magenta = rl.Color{R: 255, G: 0, B: 255, A: 255}
var Purple = rl.Color{R: 128, G: 0, B: 128, A: 255}
var LightGrey = rl.Color{R: 150, G: 150, B: 150, A: 255}

func GetCellColors() []rl.Color {
	return []rl.Color{
		DarkGrey,
		Green,
		Red,
		Blue,
		Yellow,
		Cyan,
		Magenta,
		Purple,
		LightGrey,
	}
}