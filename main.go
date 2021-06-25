package main

import (
	"github.com/ogioldat/cantor/ui/menu"
	"github.com/ogioldat/cantor/ui/welcomescreen"
)

func main() {
	// currencies := web.GetLatestCurrencies()
	welcomescreen.IntroScreen()
	menu.Showmenu()

}
