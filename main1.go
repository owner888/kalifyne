package main

import (
    //"fmt"
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
	"image/color"
    //"os"
)

//func init() {
    ////设置中文字体
    //dir, err := os.UserHomeDir()
    //if err != nil {
        //panic(err.Error())
    //}
    //var fontfile = fmt.Sprintf("%s/%s", dir, "Documents/YaHeiMonacoHybrid.ttf")
    //if FileExists(fontfile) {
        //os.Setenv("FYNE_FONT", fontfile)
    //}
//}


//// FileExists 判断所给路径文件/文件夹是否存在
//func FileExists(file string) bool {
    //_, err := os.Stat(file) //os.Stat获取文件信息
    //if err == nil {
        //return true
    //}
    //if os.IsExist(err) {
        //return true
    //}
    //return false
//}
func main1() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("Border Layout", color.Black)
    top.Resize(fyne.NewSize(200, 300))
    // NewGridLayoutWithColumns：返回一个 gridLayout 结构体，可以指定列数。如果需要垂直布局，可以替换成 NewGridLayoutWithRows
    // NewContainerWithLayout 返回一个 Container 实例，使布局生效
    left := container.NewWithoutLayout(canvas.NewText("Left ", color.Black), canvas.NewText("Right ", color.Black))
    //left.Resize(fyne.NewSize(325, 700))
	//left := canvas.NewText("left", color.Black)
	//middle := canvas.NewText("content", color.Black)

    img := canvas.NewImageFromResource(theme.FyneLogo())
    text4 := canvas.NewText("centered", color.Black)
    middle := container.NewMax(img, text4)

	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(1125, 700))
	myWindow.ShowAndRun()
}
