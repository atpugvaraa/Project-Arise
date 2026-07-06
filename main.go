package main

import (
	panels "project-arise/src/Panels"

	. "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	InitWindow(1080, 720, "Arise")
	defer CloseWindow()

	status := panels.NewStatusPanel(1, 0, 100, 0, 0)

	SetTargetFPS(60)

	for !WindowShouldClose() {
		screenWidth := float32(GetScreenWidth())
		screenHeight := float32(GetScreenHeight())

		const size = 100

		rectWidth := float32(size) * 3
		rectX := (screenWidth - rectWidth) * 0.5
		rectY := (screenHeight - size) * 0.25

		// Title
		label := "Project: Arise"
		const fontSize = 40
		textWidth := float32(MeasureText(label, fontSize))

		textX := rectX + (rectWidth-textWidth)*0.5
		textY := rectY + (size-fontSize)*0.5

		BeginDrawing()
		ClearBackground(RayWhite)
		DrawRectangle(int32(rectX), int32(rectY), int32(rectWidth), int32(size), Black)
		DrawText(label, int32(textX), int32(textY), fontSize, RayWhite)
		DrawText(status.Title, int32(textX), int32(textY+100), fontSize, Black)

		for i, line := range status.Lines {
			DrawText(line, int32(textX), int32(textY+200+float32(i)*fontSize), fontSize, Black)
		}
		EndDrawing()
	}
}
