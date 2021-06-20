package utils

import (
	"net/url"
)

func ParseURLParams(params map[string]string) string {
	query := url.Values{}

	for param, val := range params {
		query.Add(param, val)
	}

	return query.Encode()
}
