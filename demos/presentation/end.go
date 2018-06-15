package main

import (
	"fmt"

	"github.com/ngrok/tcell"
	"github.com/ngrok/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/rivo/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
