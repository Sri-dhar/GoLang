package main

import(
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main(){

	fmt.Println("Enter a string to display in the GUI window")
	var stringToDisplay string
	fmt.Scanln(&stringToDisplay)

	a := App.new()
	w := a.NewWindow("Box Layout")
	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	text4 := canvas.NewText("centered", color.White)
	text5 := canvas.NewText(stringToDisplay, color.White)

	content := container
}