package main

import (
	"io"
	"net/http"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
			SetWordWrap(true).
			SetChangedFunc(func() {
				app.Draw()
			})

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	textView.SetText(string(content))

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}