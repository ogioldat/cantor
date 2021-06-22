package main

import (
	"github.com/ogioldat/cantor/web"
	"github.com/ogioldat/cantor/welcomescreen"
)

func main() {
	welcomescreen.IntroScreen()

	web.GetLatestCurrencies()
}
