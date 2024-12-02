package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	panelWidth int32 = 200
	windowWidth int32 = numCols*cellSize + panelWidth
	windowHeight int32 = numRows*cellSize
)

var lastUpdateTime float64 = 0;
func EventTriggered(interval float64) bool {
	currentTime := rl.GetTime()
	if (currentTime - lastUpdateTime) >= interval {
		lastUpdateTime = currentTime
		return true
	}

	return false
}

func main() {
	var game Game
	game.Initialize()

	rl.InitWindow(windowWidth, windowHeight, "TETRIS")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	bg := rl.LoadTexture("img/bg.jpg")
	defer rl.UnloadTexture(bg)
	
	for !rl.WindowShouldClose() {
		if game.GameOver {
			if rl.IsKeyPressed(rl.KeyR) {
				game.Initialize()
			}
		} else {
			game.HandleInput()
			if (EventTriggered(0.6)) {
				game.MoveBlockDown()
			}
		}

		rl.BeginDrawing()

		// println(rl.GetTime())
		
		rl.DrawTexture(bg, 0, 0, LightGrey)
		rl.DrawRectangle(0, 0, windowWidth - panelWidth, windowHeight, LightGrey)
		rl.DrawRectangle(windowWidth - panelWidth, 0, 2, windowHeight, rl.Blue)
		game.Draw()

		if game.GameOver {
			rl.DrawRectangle(numCols * cellSize / 2 - 45, numRows * cellSize / 2 - 10, 230, 80, rl.Black)
			rl.DrawText("GAME OVER", numCols * cellSize / 2, numRows * cellSize / 2, 20, rl.Green)
			rl.DrawText("Press R to Restart", windowWidth/2-130, windowHeight/2+40, 20, rl.Green)
		}

		rl.EndDrawing()
	}
}

