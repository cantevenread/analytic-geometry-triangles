package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("†")
    w.Resize(fyne.Size{Width: 400, Height: 400})
    w.SetContent(widget.NewLabel("Hello Fyne!"))

	hello := widget.NewLabel("Hello Fyne!")
    w.SetContent(container.NewGridWithColumns(1,hello, widget.NewButton("Hi!", func() {
        hello.Hide()
        w.SetContent(widget.NewLabel("Hello Fyne!"))
    })))

	w.ShowAndRun()
}
