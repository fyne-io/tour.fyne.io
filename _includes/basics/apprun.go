package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	go showAnother(myApp)
	myWindow.ShowAndRun()
	// myWindow.Show(); myApp.Run()
	tidyUp()
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown Later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Show()
}

func tidyUp() {
	fmt.Println("Exited")
}
