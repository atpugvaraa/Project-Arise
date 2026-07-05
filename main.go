package main

import . "github.com/gen2brain/raylib-go/raylib"

func main() {
	InitWindow(800, 600, "Arise")
	defer CloseWindow()

	SetTargetFPS(60)

	for !WindowShouldClose() {
		BeginDrawing()
		ClearBackground(RayWhite)
		DrawText("Hello, World!", 10, 10, 20, Black)
		EndDrawing()
	}
}
