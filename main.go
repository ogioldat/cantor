package main

import (
	"github.com/ogioldat/cantor/ui/menu"
	"github.com/ogioldat/cantor/ui/welcomescreen"
	"github.com/ogioldat/cantor/web"
)

func main() {
	web.GetLatestCurrencies()
	welcomescreen.IntroScreen()
	menu.Showmenu()

}
