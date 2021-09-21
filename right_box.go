package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    //"fyne.io/fyne/v2/widget"
)

func makeRightBox(tableBox fyne.CanvasObject, centerBox *fyne.Container, _ fyne.Window) fyne.CanvasObject {

    box := container.NewHSplit(tableBox, centerBox)
    box.SetOffset(0.2)
    return box
}

