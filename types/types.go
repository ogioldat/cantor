package types

type CurrenciesResponse struct {
	Table         string `json:"table"`
	No            string `json:"no"`
	EffectiveDate string `json:"effectiveDate"`
	Rates         []struct {
		Currency string  `json:"currency"`
		Code     string  `json:"code"`
		Mid      float32 `json:"mid"`
	} `json:"rates"`
}
