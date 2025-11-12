package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the Protocol...(Ex:Ip)")

	content := readUserInput(input)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func readUserInput(input *widget.Entry) *fyne.Container {
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		input.SetText("")
	}))

	return content
}
