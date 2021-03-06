package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fixed Grid Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	//	myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
