package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	f := app.New()
	w := f.NewWindow("")
	label1 := widget.NewLabel("Label1")

	b1 := widget.NewButton("Button1", func() {})
	b2 := widget.NewButton("Button2", func() {})
	label2 := widget.NewLabel("Label3")

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), label1, layout.NewSpacer()),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), b1, b2, layout.NewSpacer()),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), label2, layout.NewSpacer()),
		),
	)

	w.Resize(fyne.Size{Height: 320, Width: 480})

	w.ShowAndRun()
}
