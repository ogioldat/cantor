package web

import (
	"fmt"

	"github.com/ogioldat/cantor/config"
)

func GetLatestCurrencies() {
	url := config.GetConfig().NBP_API_URL

	fmt.Println(url)
}
