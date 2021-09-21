package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func tableScreen(_ fyne.Window) fyne.CanvasObject {
    a := fyne.CurrentApp()

    table := a.Preferences().String(preferenceCurrentTutorial)
    //columns := db.Schema().ListColumns("user")
    //t.Logf("jsonStr = %v\n", database.FormatJSON(columns))

    fmt.Printf("table = %v\n", table)
    return container.NewCenter(container.NewVBox(
        widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
        widget.NewLabel("canvasScreen"),
    ))
}

