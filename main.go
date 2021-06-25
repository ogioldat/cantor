package main

import (
	"fmt"

	"github.com/ogioldat/cantor/web"
)

func main() {
	fmt.Println("Start")
	// welcomescreen.IntroScreen()

	web.GetLatestCurrencies()
}
