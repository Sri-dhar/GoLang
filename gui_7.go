package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func printHello() {
	fmt.Println("Hello World")
}

func main() {
	printHello()
	fmt.Println("Enter a number")
	var input int
	fmt.Scanln(&input)

	a := app.New()
	w := a.NewWindow("Box Layout")
	w.Resize(fyne.NewSize(900, 900))

	button := widget.NewButton("Click me", func() {
		fmt.Println("Button clicked")
	})
	w.SetContent(button)
	w.ShowAndRun()
}