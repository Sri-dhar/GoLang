package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Hello World")
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello Fyne!:SetContent"))
	myWindow.Resize(fyne.NewSize(900, 800))
	myWindow.ShowAndRun()
}

