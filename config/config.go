package config

import (
	"github.com/ogioldat/cantor/utils"
)

type MenuOption struct {
	Number int
	Name   string
}

type Config struct {
	NBP_BASE_URL string
	Endpoints    Endpoints
	MenuOptions  map[string]MenuOption
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
	return cfg.NBP_BASE_URL + endpoint + "?" + utils.ParseURLParams(DEFAULT_PARAMS)
}

func GetConfig() (Config, Helpers) {
	options := make(map[string]MenuOption)

	options["displayCurrentCurrencies"] = MenuOption{Number: 1, Name: "Wyświetl aktualny kursów walut obcych"}
	options["displayCurrenciesOnDay"] = MenuOption{Number: 2, Name: "Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika"}
	options["displayCurrentCurrenciesBuySell"] = MenuOption{Number: 3, Name: "Wyświetl aktuale kursy kupna i sprzedaży walut obcych"}
	options["displayCurrentCurrenciesBuySellOnDay"] = MenuOption{Number: 4, Name: "Wyświetl kursy kupna i sprzedaży walut obcych z wybranego dnia"}
	options["currencyConversion"] = MenuOption{Number: 5, Name: "Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia"}
	options["settings"] = MenuOption{Number: 6, Name: "Ustawienia"}
	options["exit"] = MenuOption{Number: 7, Name: "Zakończ program"}

	return Config{
			NBP_BASE_URL: "http://api.nbp.pl/api/exchangerates/tables/",
			Endpoints: Endpoints{
				Currencies: []string{"A", "B", "C"},
			},
			MenuOptions: options,
		},
		Helpers{
			ApplyURL,
		}
}
