package config

type Config struct {
	NBP_API_URL string
}

func GetConfig() Config {
	return Config{
		NBP_API_URL: "http://api.nbp.pl/api/exchangerates/tables",
	}
}
