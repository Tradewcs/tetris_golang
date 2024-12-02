package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	const FPS int32 = 60

	rl.InitWindow(windowWidth, windowHeight, "TETRIS")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	bg := rl.LoadTexture("img/bg.jpg")
	defer rl.UnloadTexture(bg)


	var game Game
	game.Initialize(bg)
	
	for !rl.WindowShouldClose() {
		game.Run()

		rl.BeginDrawing()
		
		game.Draw()

		rl.EndDrawing()
	}
}

