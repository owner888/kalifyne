package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/owner888/kaligo/database"
    sqlite "github.com/owner888/kaligo/database/driver/sqlite"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
    //kaliGoDB database.DB

    // Tutorials defines the metadata for each tutorial
    Tutorials = map[string]Tutorial{
        "welcome": {"Welcome", "Welcome Intor", welcomeScreen},
        "canvas": {"Canvas", "See the canvas capabilities.", canvasScreen},
    }

    // TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
        "":            {"welcome", "canvas"},
        //"collections": {"welcome", "canvas"},
	}
)

func init() {
    kaliGoDB, err := database.Open(sqlite.Open("./test.db"))
    if err != nil {
        panic(err)
    }

    tables := kaliGoDB.Schema().ListTables()
    fmt.Printf("jsonStr = %v\n", database.FormatJSON(tables))

    for _, v := range tables {
        Tutorials[v] = Tutorial{v, v, tableScreen}
    }
    TutorialIndex[""] = tables
}

func makeSplitBox(c *fyne.Container, w fyne.Window) fyne.CanvasObject {
	a := fyne.CurrentApp()

    title := widget.NewLabel("Component name")
    intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
    intro.Wrapping = fyne.TextWrapWord
    setTutorial := func(t Tutorial) {
        if fyne.CurrentDevice().IsMobile() {
            child := a.NewWindow(t.Title)
            topWindow = child
            child.SetContent(t.View(topWindow))
            child.Show()
            child.SetOnClosed(func() {
                topWindow = w
            })
            return
        }

        title.SetText(t.Title)
        intro.SetText(t.Intro)

        c.Objects = []fyne.CanvasObject{t.View(w)}
        c.Refresh()
    }

    //tutorial := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
    //if fyne.CurrentDevice().IsMobile() {
        //w.SetContent(makeNav(setTutorial, false))
    //} else {
        //split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
        //split.SetOffset(0.2)
        //w.SetContent(split)
    //}

    splitBox := container.NewHSplit(makeNav(setTutorial, true), c)
    splitBox.SetOffset(0.2)

    return splitBox
}

func makeNav(setTutorial func(tutorial Tutorial), loadPrevious bool) fyne.CanvasObject {
    a := fyne.CurrentApp()

    tree := &widget.Tree{
        ChildUIDs: func(uid string) []string {
            return TutorialIndex[uid]
        },
        IsBranch: func(uid string) bool {
            children, ok := TutorialIndex[uid]

            return ok && len(children) > 0
        },
        CreateNode: func(branch bool) fyne.CanvasObject {
            return widget.NewLabel("Collection Widgets")
        },
        UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
            t, ok := Tutorials[uid]
            if !ok {
                fyne.LogError("Missing tutorial panel: "+uid, nil)
                return
            }
            obj.(*widget.Label).SetText(t.Title)
        },
        OnSelected: func(uid string) {
            if t, ok := Tutorials[uid]; ok {
                a.Preferences().SetString(preferenceCurrentTutorial, uid)
                setTutorial(t)
            }
        },
    }

    if loadPrevious {
        currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
        tree.Select(currentPref)
    }

    //themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
        //widget.NewButton("Dark", func() {
            //a.Settings().SetTheme(theme.DarkTheme())
        //}),
        //widget.NewButton("Light", func() {
            //a.Settings().SetTheme(theme.LightTheme())
        //}),
    //)

    //return container.NewBorder(nil, themes, nil, nil, tree)

    return tree
}
