---
layout: home

title: Gradient
order: 7

---

The last canvas primitive type is `canvas.Gradient` which is used
to draw a gradient from one colour to another in various patterns.
You can create linear or radial gradients using this type.

To create a gradient you need a start and end colour - every colour
in between is calculated by the driver. In this example we use 
`color.Transparent` to show how a gradient (or any other type) could
use an alpha value to be semi-transparent over the content behind.

That's it for our tour of the canvas elements in Fyne. Next take some 
time to learn about the various `Widget`s available.