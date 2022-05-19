package main

import (
	"github.com/go-rod/rod"
)

func main() {
	rod.
		New().
		MustConnect().
		MustPage("https://www.wikipedia.org/").
		MustWaitLoad().
		MustScreenshot("screenshot.png")
}
