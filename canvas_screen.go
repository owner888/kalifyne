package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func canvasScreen(_ fyne.Window) fyne.CanvasObject {

    return container.NewCenter(container.NewVBox(
        widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
        widget.NewLabel("canvasScreen"),
    ))
}

