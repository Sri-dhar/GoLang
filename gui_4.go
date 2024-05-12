package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//create a simple gui window that takes input a string and displays it and has a button to take in put
	fmt.Println("Hello World")
	a := app.New()
	w := a.NewWindow("")
	w.Resize(fyne.NewSize(900, 800))
	stringToDisplay := "This is inside the GUI window"
	content := widget.NewLabel(stringToDisplay)
	content.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	w.SetContent(content)
	
	
	//create an input field
	input := widget.NewEntry()
	w.SetContent(widget.NewVBox(
		content,
		input,
		widget.NewButton("Submit", func() {
			content.SetText(input.Text)
		}),
	))

	
	
	w.ShowAndRun()

}	
