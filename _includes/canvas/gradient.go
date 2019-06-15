package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")

	gradient := canvas.NewLinearGradient(color.White, color.Transparent,
		canvas.GradientDirectionHorizontal)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}