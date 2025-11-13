package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func StyledProtocolInput(placeholder string) (*fyne.Container, *widget.Entry) {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)

	rect := canvas.NewRectangle(color.RGBA{168, 187, 163, 255})
	rect.CornerRadius = 5

	content := container.NewPadded(input)

	return container.NewStack(rect, content), input
}
