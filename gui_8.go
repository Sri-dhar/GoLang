package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("A new Window for checkboxes")
	w.Resize(fyne.NewSize(900, 900))

	checkbox1 := widget.NewCheck("Checkbox 1", func(value bool) {
		fmt.Println("Checkbox 1 is now", value)
	}) // Added closing parenthesis here

	w.SetContent(checkbox1)
	w.ShowAndRun()
}