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
	card := widget.NewCard("Hi tehre", "substring", w.Canvas().Content())
	w.SetContent(card)
	w.MainMenu()
	w.Show()

	a.Run()
}
