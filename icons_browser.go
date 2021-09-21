package main

import (
    //"fmt"
    _ "embed"
    _ "unsafe"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

var icons map[fyne.ThemeIconName]fyne.Resource    

var (
    // go:embed icons.go
    srcIcons string
)

func main5() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Fyne icons browser")

    myWindow.SetContent(iconScreen(myWindow))
    myWindow.Resize(fyne.NewSize(640, 460))
    myWindow.ShowAndRun()
}

func iconScreen(_ fyne.Window) fyne.CanvasObject {
    txtSource := widget.NewMultiLineEntry()
    txtSource.SetText(srcIcons)

    txt := widget.NewEntry()
    c  := container.NewGridWrap(fyne.Size{50, 50})
    cc := container.NewVScroll(container.NewVBox(txt, c))
    for _, icon := range icons {
        // c.Add(widget.NewButtonWithIcon(icn.icon)
        btn := widget.NewButtonWithIcon("", icon, nil)
        btn.OnTapped = func() {
            // fmt.Println(btn.Icon.Name())
            txt.SetText(btn.Icon.Name())
        }

        c.Add(btn)
    }

    tabs := container.NewAppTabs(
        container.NewTabItem("main",   cc),
        container.NewTabItem("source", txtSource),
        container.NewTabItemWithIcon("source", theme.CancelIcon(), txtSource),
    )

    tabs.SetTabLocation(container.TabLocationTop)
    return tabs
}
