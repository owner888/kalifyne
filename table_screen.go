package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/owner888/kaligo/database"
    sqlite "github.com/owner888/kaligo/database/driver/sqlite"
)

var (
    table string
    columns []database.Column
    colNames []string = []string{"Name", "Type", "Length", "Unsigned", "Allow Null", "Key", "Default", "Comment"}
) 

func tableScreen(_ fyne.Window) fyne.CanvasObject {
    a := fyne.CurrentApp()

    table = a.Preferences().String(preferenceCurrentTutorial)

    tabs := container.NewAppTabs(
		container.NewTabItem("Schema", widget.NewLabel("Schema")),
		container.NewTabItem("Data",   widget.NewLabel("Data")),
		container.NewTabItem("SQL",    widget.NewLabel("SQL")),
		container.NewTabItem("Log",    widget.NewLabel("Log")),
	)
    tabs.SetTabLocation(container.TabLocationTop)
    tabs.OnSelected = func(tabItem *container.TabItem) {
        //fmt.Printf("tabItem.Text = %v\n", tabItem.Text)
        if tabItem.Text == "Schema" {
            tabItem.Content = makeSchemaTable()
        } else if tabItem.Text == "Data" {
            tabItem.Content = makeDataTable()
        } else if tabItem.Text == "SQL" {
            tabItem.Content = widget.NewLabel("SQL")
        } else {
            tabItem.Content = widget.NewLabel("Log")
        }
    }

    //fmt.Printf("table = %v\n", table)

    return tabs
}

func makeSchemaTable() fyne.CanvasObject {
    db, err := database.Open(sqlite.Open("./test.db"))
    if err != nil {
        panic(err)
    }

    columns := db.Schema().ListColumns(table)
    fmt.Printf("jsonStr = %v\n", database.FormatJSON(columns))
    rowsNum := len(columns) + 1
    colsNum := 8
    
    var i = 0
    t := widget.NewTable(
        func() (int, int) { return rowsNum, colsNum },  // rows and columns
        func() fyne.CanvasObject {          // create
            i++
            fmt.Printf("cell i = %v\n", i)
            var c fyne.CanvasObject
            if i / 2 == 0 {
                c = widget.NewLabel("VARCHAR")
            } else {
                //c = widget.NewLabel("VARCHAR VARCHAR")
                c = widget.NewLabel("Cell 000, 000")
            }
            return c
        },
        func(id widget.TableCellID, cell fyne.CanvasObject) {   // update
            label := cell.(*widget.Label)
            switch id.Row { // 第一行
            case 0:
                label.SetText(colNames[id.Col])
            default:
                column := columns[id.Row - 1]

                switch id.Col {
                case 0:
                    //label.SetText(fmt.Sprintf("%d", id.Row+1))
                    label.SetText(column.DBName)
                case 1:
                    dataType := string(column.DataType)
                    if dataType == "int" {
                        if column.Size == 0 {
                            dataType = "INTEGER"
                        } else {
                            dataType = "INT"
                        }
                    } else if dataType == "string" {
                        if column.Size == 0 {
                            dataType = "TEXT"
                        } else {
                            dataType = "VARCHAR VARCHAR"
                        }
                    }
                    label.SetText(dataType)
                case 2:
                    label.SetText(database.ToString(column.Size))
                case 3:
                    if column.Unique {
                        label.SetText("YES")
                    } else {
                        label.SetText("NO")
                    }
                case 4:
                    if column.NotNull {
                        label.SetText("NO")
                    } else {
                        label.SetText("YES")
                    }
                case 5:
                    if column.PrimaryKey {
                        label.SetText("PRI")
                    } else {
                        label.SetText("")
                    }
                case 6:
                    label.SetText(column.DefaultValue)
                case 7:
                    label.SetText(column.Comment)
                default:
                    label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
                }
            }
        },
    )

    //t.SetColumnWidth(0, 34)
    //t.SetColumnWidth(1, 80)
    //t.SetColumnWidth(2, 80)
    //t.SetColumnWidth(3, 90)
    //t.SetColumnWidth(4, 100)
    //t.SetColumnWidth(5, 60)
	return t
}

func makeDataTable() fyne.CanvasObject {
	t := widget.NewTable(
        func() (int, int) { return 7, 8 },  // rows and columns
        func() fyne.CanvasObject {          // create
            return widget.NewLabel("Cell 000, 000")
        },
		func(id widget.TableCellID, cell fyne.CanvasObject) {   // update
			label := cell.(*widget.Label)
            switch id.Row { // 第一行
            case 0:
                label.SetText(colNames[id.Col])
            default:
                switch id.Col {
                case 0:
                    label.SetText(fmt.Sprintf("%d", id.Row+1))
                case 1:
                    label.SetText("A longer cell")
                default:
                    label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
                }
            }
		})
	//t.SetColumnWidth(0, 34)
	//t.SetColumnWidth(1, 102)
	return t
}

