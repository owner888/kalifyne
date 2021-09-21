package main

import (
    "log"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    //"fyne.io/fyne/v2/layout"
)


func main2() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Toolbar Widget")
    button := widget.NewButtonWithIcon("Hi", theme.CancelIcon(), nil)

    content2 := container.NewBorder(nil, nil, nil, button, widget.NewLabel("Content"))

    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
            log.Println("New document")
        }),
        widget.NewToolbarSeparator(),
        widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
        widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
        widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
        widget.NewToolbarSpacer(),
        widget.NewToolbarAction(theme.HelpIcon(), func() {
            myWindow.SetContent(content2)
            log.Println("Display help")
        }),
    )

    content1 := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
    myWindow.SetContent(content1)
    myWindow.ShowAndRun()
}
