package main

import (
    "fmt"
    "net/url"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    //"fyne.io/fyne/v2/container"
    //"fyne.io/fyne/v2/widget"
)

// Settings gives access to user interfaces to control Fyne settings
type Settings struct {
	fyneSettings app.SettingsSchema

	preview *canvas.Image
	colors  []fyne.CanvasObject
}

// NewSettings returns a new settings instance with the current configuration loaded
func NewSettings() *Settings {
    s := &Settings{}
    //s.load()

    return s
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
    newItem := fyne.NewMenuItem("New", nil)
    otherItem := fyne.NewMenuItem("Other", nil)
    otherItem.ChildMenu = fyne.NewMenu(
        "",
        fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
        fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") }),
    )
    newItem.ChildMenu = fyne.NewMenu(
        "",
        fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") }),
        fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") }),
        otherItem,
    )
    settingsItem := fyne.NewMenuItem("Settings", func() {
        w := a.NewWindow("Fyne Settings")
        //w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
        //w.SetContent(NewSettings().LoadAppearanceScreen(w))
        w.Resize(fyne.NewSize(480, 480))
        w.Show()
    })

    cutItem := fyne.NewMenuItem("Cut", func() {
        shortcutFocused(&fyne.ShortcutCut{
            Clipboard: w.Clipboard(),
        }, w)
    })
    copyItem := fyne.NewMenuItem("Copy", func() {
        shortcutFocused(&fyne.ShortcutCopy{
            Clipboard: w.Clipboard(),
        }, w)
    })
    pasteItem := fyne.NewMenuItem("Paste", func() {
        shortcutFocused(&fyne.ShortcutPaste{
            Clipboard: w.Clipboard(),
        }, w)
    })
    findItem := fyne.NewMenuItem("Find", func() { fmt.Println("Menu Find") })

    helpMenu := fyne.NewMenu("Help",
    fyne.NewMenuItem("Documentation", func() {
        u, _ := url.Parse("https://developer.fyne.io")
        _ = a.OpenURL(u)
    }),
    fyne.NewMenuItem("Support", func() {
        u, _ := url.Parse("https://fyne.io/support/")
        _ = a.OpenURL(u)
    }),
    fyne.NewMenuItemSeparator(),
    fyne.NewMenuItem("Sponsor", func() {
        u, _ := url.Parse("https://fyne.io/sponsor/")
        _ = a.OpenURL(u)
    }))
    file := fyne.NewMenu("File", newItem)
    if !fyne.CurrentDevice().IsMobile() {
        file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
    }
    mainMenu := fyne.NewMainMenu(
        // a quit item will be appended to our first menu
        file,
        fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
        helpMenu,
    )

    return mainMenu
}
