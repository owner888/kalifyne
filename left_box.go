package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

func makeLeftBox(c *fyne.Container, w fyne.Window) fyne.CanvasObject {
    // 生成默认的右侧表格列表
    makeSplitBox(c, w)

    homeBtn := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
        c.Objects = []fyne.CanvasObject{
            welcomeScreen(w),
        }
        c.Refresh()

        fmt.Println("tapped home")
    })
    cancelBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
        c.Objects = []fyne.CanvasObject{
            makeToolbar(w),
        }
        c.Refresh()

        fmt.Println("tapped home")
    })
    settingsBtn := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
        //a := fyne.CurrentApp()
        //fmt.Printf("%v\n", a)
        fmt.Println("tapped home")
    })

    // Max Box 类似 html div z-index 层叠样式表，后面的元素会叠加到前面元素的上面
    left := container.NewMax(
        //canvas.NewRectangle(color.Black), // 黑色背景，不好用，因为icon无法设置颜色为白色
        container.NewVBox(
            homeBtn,
            cancelBtn,
            layout.NewSpacer(),
            settingsBtn,
        ),
    )

    return container.NewHBox(left, widget.NewSeparator())
}

