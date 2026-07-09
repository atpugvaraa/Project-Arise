package main

import (
	panels "project-arise/src/Panels"

	gui "github.com/gen2brain/raylib-go/raygui"
	. "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1080
	screenHeight = 720

	showTitleFor = 3.0
)

var (
	running         = true
	showsStatusMenu bool
	showTitle       bool = true

	currentTime float32 = 0
)

func init() {
	InitWindow(screenWidth, screenHeight, "Project: Arise")
	SetExitKey(0)
	SetTargetFPS(60)
}

func quit() {
	running = !WindowShouldClose()
}

func input() {
	if IsKeyPressed(KeyEscape) {
		showsStatusMenu = !showsStatusMenu
	}
}

func update() {
	currentTime += GetFrameTime()

	if currentTime <= showTitleFor {
		showTitle = true
	} else {
		showTitle = false
	}
}

func render() {
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

	// Panels
	status := panels.NewStatusPanel(1, 0, 100, 0, 0)

	BeginDrawing()
	ClearBackground(RayWhite)
	DrawFPS(10, 10)
	gui.Button(Rectangle{X: rectX, Y: rectY + 600, Width: rectWidth, Height: size}, "Button")

	// Title
	if showTitle {
		DrawRectangle(int32(rectX), int32(rectY), int32(rectWidth), int32(size), Black)
		DrawText(label, int32(textX), int32(textY), fontSize, RayWhite)
	}

	// Status Menu
	if showsStatusMenu {
		DrawText(status.Title, int32(textX), int32(textY+100), fontSize, Black)

		for i, line := range status.Lines {
			DrawText(line, int32(textX), int32(textY+200+float32(i)*fontSize), fontSize, Black)
		}
	}

	EndDrawing()
}

func main() {
	for running {
		input()
		update()
		render()
	}

	quit()
}
