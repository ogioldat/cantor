package main

import (
	"fmt"

	"github.com/ogioldat/cantor/web"
)

func main() {
	fmt.Println("Start")

	web.GetLatestCurrencies()
}
