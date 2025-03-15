package main

import (
	"fmt"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/net/html"
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

	plaintext := ""

	domDoc := html.NewTokenizer(resp.Body)
    previousStartToken := domDoc.Token()
	loopDom:
		for {
			current := domDoc.Next() // Gets type of next token
			switch current {
			case html.ErrorToken:
				break loopDom // End of the document,  done
			case html.StartTagToken:
				previousStartToken = domDoc.Token() // Sets previousStartToken to html tag of token
			case html.TextToken:
				if previousStartToken.Data == "script" || previousStartToken.Data == "style" {
					continue // Ignores text inside script or style tags
				}
				plaintext += fmt.Sprintf("%s\n", html.UnescapeString(string(domDoc.Text()))) // Appends text from current token to plaintext return value
			}
		}

	return string(plaintext), nil
}