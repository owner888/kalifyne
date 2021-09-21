package main

import (
	"fmt"
    "log"
	//"image/color"
	//"net/url"
	//"time"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeToolbar(_ fyne.Window) fyne.CanvasObject {
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
            //w.SetContent(help)
            log.Println("Display help")
        }),
    )
    return container.NewVBox(toolbar)
}

func makeButtonTab(_ fyne.Window) fyne.CanvasObject {
	disabled := widget.NewButton("Disabled", func() {})
	disabled.Disable()

	shareItem := fyne.NewMenuItem("Share via", nil)
	shareItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Twitter", func() { fmt.Println("context menu Share->Twitter") }),
		fyne.NewMenuItem("Reddit", func() { fmt.Println("context menu Share->Reddit") }),
	)
	menuLabel := newContextMenuButton("tap me for pop-up menu with submenus", fyne.NewMenu("",
		fyne.NewMenuItem("Copy", func() { fmt.Println("context menu copy") }),
		shareItem,
	))

	return container.NewVBox(
		widget.NewButton("Button (text only)", func() { fmt.Println("tapped text button") }),
		widget.NewButtonWithIcon("Button (text & leading icon)", theme.ConfirmIcon(), func() { fmt.Println("tapped text & leading icon button") }),
		&widget.Button{
			Alignment: widget.ButtonAlignLeading,
			Text:      "Button (leading-aligned, text only)",
			OnTapped:  func() { fmt.Println("tapped leading-aligned, text only button") },
		},
		&widget.Button{
			Alignment:     widget.ButtonAlignTrailing,
			IconPlacement: widget.ButtonIconTrailingText,
			Text:          "Button (trailing-aligned, text & trailing icon)",
			Icon:          theme.ConfirmIcon(),
			OnTapped:      func() { fmt.Println("tapped trailing-aligned, text & trailing icon button") },
		},
		disabled,
		layout.NewSpacer(),
		layout.NewSpacer(),
		menuLabel,
		layout.NewSpacer(),
	)
}


type contextMenuButton struct {
	widget.Button
	menu *fyne.Menu
}

func (b *contextMenuButton) Tapped(e *fyne.PointEvent) {
	widget.ShowPopUpMenuAtPosition(b.menu, fyne.CurrentApp().Driver().CanvasForObject(b), e.AbsolutePosition)
}

func newContextMenuButton(label string, menu *fyne.Menu) *contextMenuButton {
	b := &contextMenuButton{menu: menu}
	b.Text = label

	b.ExtendBaseWidget(b)
	return b
}
