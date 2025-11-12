package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	output := canvas.NewText(time.Now().Format(time.TimeOnly), color.NRGBA{G: 0xff, A: 0xff})
	output.TextStyle.Monospace = true
	output.TextSize = 32
	myWindow.SetContent(output)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fyne.Do(func() {
				output.Text = time.Now().Format(time.TimeOnly)
				output.Refresh()
			})
		}
	}()
	myWindow.ShowAndRun()
}
