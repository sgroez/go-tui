package main

import (
	"io"
	"net/http"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	url := ""
	app := tview.NewApplication()
	textView := tview.NewTextView().
			SetWordWrap(true).
			SetChangedFunc(func() {
				app.Draw()
			})
	
	urlInput := tview.NewInputField().
			SetLabel("URL").
			SetPlaceholder("enter a website url you want to fetch").
			SetChangedFunc(func(input string) {
				url = input
			}).
			SetDoneFunc(func(key tcell.Key) {
				if response, err := fetchUrl(url); err == nil {
					textView.SetText(response)
				}
			})

	grid := tview.NewGrid().
			SetRows(3, 0).
			SetColumns(0).
			SetBorders(true).
			AddItem(urlInput, 0, 0, 1, 1, 0, 0, true).
			AddItem(textView, 1, 0, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func fetchUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}