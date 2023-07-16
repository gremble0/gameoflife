package main

import (
    "image/color"

    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    const ROWS = 20
    const COLUMNS = 20

    app := app.New()
    window := app.NewWindow("Game of life")

    window.SetContent(widget.NewLabel("Game of Life"))

    var boxes [COLUMNS][ROWS]*canvas.Rectangle
    grid := container.New(layout.NewGridLayout(COLUMNS))

    for i := 0; i < ROWS; i++ {
        for j := 0; j < COLUMNS; j++ {
	    currentBox := canvas.NewRectangle(color.White)
	    boxes[j][i] = currentBox
	    grid.Add(currentBox)
        }
    }

    window.SetContent(grid)

    window.ShowAndRun()
}
