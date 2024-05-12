package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

func main() {
    fmt.Println("Hello World")
    a := app.New()
    w := a.NewWindow("")


    stringToDisplay := "This is inside the GUI window"

    content := widget.NewLabel(stringToDisplay)
    w.SetContent(content)
    //how to change the size of the content
    w.Resize(fyne.NewSize(900, 800))
    //how to increase the font size of the content
    content.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
}