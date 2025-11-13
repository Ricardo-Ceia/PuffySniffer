package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Ricardo-Ceia/PuppySniffer/components"
	"log"
)

func ReadUserInput(placeholder string) *fyne.Container {
	styledInputContainer, input := components.StyledProtocolInput(placeholder)

	content := container.NewVBox(styledInputContainer, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		input.SetText("")
	}))
	return content
}
