package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("pastrami")

	w.SetContent(widget.NewLabel("Welcome to PASTRAMI"))
	w.Resize(fyne.NewSize(300, 300))
	w.Show()

	a.Run()
}
