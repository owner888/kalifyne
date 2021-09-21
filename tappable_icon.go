package main

import (
    "log"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    //"fyne.io/fyne/v2/app"
    //"fyne.io/fyne/v2/theme"
)

type tappableIcon struct {
    widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
    icon := &tappableIcon{}
    icon.ExtendBaseWidget(icon)
    icon.SetResource(res)

    return icon
}

// 点击
func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
    log.Println("I have been tapped")
}

// 长按
func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
    log.Println("I have been TappedSecondary")
}

//func main() {
    //a := app.New()
    //w := a.NewWindow("Tappable")
    //w.SetContent(newTappableIcon(theme.FyneLogo()))
    //w.ShowAndRun()
//}
