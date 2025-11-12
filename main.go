package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PuppySniffer")
	myWindow.Resize(fyne.NewSize(1280, 720))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the Protocol...(Ex:Ip)")

	content := readUserInput("Enter the Protocol...(Ex:Ip)")

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func readUserInput(placeholder string) *fyne.Container {
	styledInputContainer, input := StyledProtocolInput(placeholder)

	content := container.NewVBox(styledInputContainer, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		input.SetText("")
	}))
	return content
}

func StyledProtocolInput(placeholder string) (*fyne.Container, *widget.Entry) {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)

	rect := canvas.NewRectangle(color.RGBA{168, 187, 163, 255})
	rect.CornerRadius = 5

	content := container.NewPadded(input)

	return container.NewStack(rect, content), input
}
