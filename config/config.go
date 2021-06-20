package config

import (
	"github.com/ogioldat/cantor/utils"
)

type Config struct {
	NBP_API_URL string
	endpoint    Endpoints
}

type Endpoints struct {
	currencies string
}

var ConfigMap = Config{
	NBP_API_URL: "http://api.nbp.pl/api/exchangerates/",
	endpoint: Endpoints{
		currencies: "/tables",
	},
}

var DEFAULT_PARAMS = map[string]string{"format": "json"}

func GetEndpointFn(params map[string]string) string {

	for key, value := range DEFAULT_PARAMS {
		params[key] = value
	}

	return utils.ParseURLParams(params)
}

func GetConfig() Config {
	return ConfigMap
}
