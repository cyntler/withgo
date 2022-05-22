package main

import (
	"github.com/go-rod/rod"
)

func main() {
	wikipedia := rod.
		New().
		MustConnect().
		MustPage("https://www.wikipedia.org/").
		MustWaitLoad()

	wikipedia.MustScreenshot("screenshot1.png")
	wikipedia.MustScreenshot("screenshot2.png")
}
