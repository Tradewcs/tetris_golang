package main

import rl "github.com/gen2brain/raylib-go/raylib"

var DarkGrey = rl.Color{R: 26, G: 31, B: 40, A: 255}

var Orange = rl.Color{248, 160, 47, 255}
var Blue = rl.Color{0, 13, 234, 255}
var Cyan = rl.Color{0, 240, 239, 255}
var Yellow = rl.Color{242, 240, 61, 255}
var Green = rl.Color{0, 240, 54, 255}
var Purple = rl.Color{165, 15, 235, 255}
var Red = rl.Color{252, 5, 32, 255}

var LightGrey = rl.Color{R: 150, G: 150, B: 150, A: 255}

func GetCellColors() []rl.Color {
	return []rl.Color{
		DarkGrey,
		Orange,
		Blue,
		Cyan,
		Yellow,
		Green,
		Purple,
		Red,
		LightGrey,
	}
}