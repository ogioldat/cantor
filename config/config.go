package config

import (
	"github.com/ogioldat/cantor/utils"
)

type Config struct {
	NBP_BASE_URL string
	Endpoints    Endpoints
}

type Helpers struct {
	ApplyURL func(string) string
}

type Endpoints struct {
	Currencies []string
}

var DEFAULT_PARAMS = map[string]string{"format": "json"}

func ApplyURL(endpoint string) string {
	cfg, _ := GetConfig()
	return cfg.NBP_BASE_URL + endpoint + "/" + utils.ParseURLParams(DEFAULT_PARAMS)
}

func GetConfig() (Config, Helpers) {
	return Config{
			NBP_BASE_URL: "http://api.nbp.pl/api/exchangerates/",
			Endpoints: Endpoints{
				Currencies: []string{"A", "B", "C"},
			},
		},
		Helpers{
			ApplyURL,
		}
}
