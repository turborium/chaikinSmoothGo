package main

import (
	"image/color"
	"strconv"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	points = []rl.Vector2{
		{X: 40, Y: 40},
		{X: 200, Y: 130},
		{X: 300, Y: 360},
		{X: 500, Y: 340},
	}
)

func chaikinSmooth(points []rl.Vector2) []rl.Vector2 {
	result := make([]rl.Vector2, 0)

	result = append(result, points[0])
	for i := 0; i < len(points)-1; i++ {
		prev := points[i]
		next := points[i+1]

		a := rl.Vector2{
			X: prev.X*3/4 + next.X*1/4,
			Y: prev.Y*3/4 + next.Y*1/4,
		}

		b := rl.Vector2{
			X: prev.X*1/4 + next.X*3/4,
			Y: prev.Y*1/4 + next.Y*3/4,
		}

		result = append(result, a)
		result = append(result, b)
	}
	result = append(result, points[len(points)-1])

	return result
}

func main() {
	divisionCount := 3

	rl.SetConfigFlags(rl.FlagWindowHighdpi | rl.FlagMsaa4xHint)

	rl.InitWindow(640, 480, "CHAIKIN’S ALGORITHM")

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// ---

		// рисуем линии между точками
		for i := 0; i < len(points)-1; i++ {
			rl.DrawLine(
				int32(points[i].X), int32(points[i].Y),
				int32(points[i+1].X), int32(points[i+1].Y),
				color.RGBA{255, 255, 255, 100},
			)
		}

		// рисуем сглаженные линии
		newPoints := points
		for i := 0; i < divisionCount; i++ {
			newPoints = chaikinSmooth(newPoints)
		}
		for i := 0; i < len(newPoints)-1; i++ {
			rl.DrawLineEx(
				rl.Vector2{X: newPoints[i].X, Y: newPoints[i].Y},
				rl.Vector2{X: newPoints[i+1].X, Y: newPoints[i+1].Y},
				2,
				color.RGBA{255, 255, 255, 255},
			)
		}

		// рисуем сглаженные точки
		for _, point := range newPoints {
			rl.DrawCircle(int32(point.X), int32(point.Y), 5, rl.Orange)
		}

		// рисуем базовые точки
		for _, point := range points {
			rl.DrawCircle(int32(point.X), int32(point.Y), 7, rl.Red)
		}

		// ---

		// рисуем GUI
		divisionCount = int(gui.SliderBar(
			rl.Rectangle{X: 30, Y: 0, Width: 200, Height: 30},
			"min",
			"max",
			float32(divisionCount),
			0,
			6) + 0.5)

		rl.DrawText(
			"Division Count: "+strconv.Itoa(divisionCount),
			300,
			0,
			20,
			rl.White,
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
