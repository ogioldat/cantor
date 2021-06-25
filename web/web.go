package web

import (
	"fmt"

	"github.com/ogioldat/cantor/config"
)

func GetLatestCurrencies() {
	cfg, helpers := config.GetConfig()

	for endpoint := range cfg.Endpoints.Currencies {
		fmt.Println(endpoint)
	}

	fmt.Println(cfg.Endpoints)
	fmt.Println(cfg.NBP_BASE_URL)
	fmt.Println(helpers.ApplyURL("test"))
}
