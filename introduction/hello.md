---
layout: home

title: Introduction
order: 1
code: |
  package main

  import (
  	"fyne.io/fyne/widget"
  	"fyne.io/fyne/app"
  )

  func main() {
  	app := app.New()

  	w := app.NewWindow("Hello")
  	w.SetContent(widget.NewVBox(
  		widget.NewLabel("Hello Fyne!"),
  		widget.NewButton("Quit", func() {
  			app.Quit()
  		}),
  	))

  	w.ShowAndRun()
  }

redirect_from:
  - /
  - /introduction/

---

Welcome to a tour of the [Fyne GUI Toolkit](https://fyne.io/).

The tour is divded into lists of areas that you can access through the menu
items at the top, right menu of this page.

Throughout this tour you will find a series of descriptions of Fyne features
and examples code associated with the content.

You can navigate though the pages of this tour using the navigation arrows
at the bottom of each page.

When you are ready to get started click on the [right arrow](/introduction/golang) below.
