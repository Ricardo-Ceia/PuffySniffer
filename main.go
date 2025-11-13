package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Ricardo-Ceia/PuppySniffer/services"
)

func main() {

	services.ListAllDevs()
	myApp := app.New()
	myWindow := myApp.NewWindow("PuppySniffer")
	myWindow.Resize(fyne.NewSize(1280, 720))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the Protocol...(Ex:Ip)")

	content := services.ReadUserInput("Enter the Protocol...(Ex:Ip)")
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
