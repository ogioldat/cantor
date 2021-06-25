package main

import (
	"fmt"

	"github.com/ogioldat/cantor/web"
	"github.com/ogioldat/cantor/welcomescreen"
)

func main() {
	fmt.Println("Start")
	welcomescreen.IntroScreen()
	web.GetLatestCurrencies()
}
