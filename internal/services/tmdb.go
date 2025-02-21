package services

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/config"
	"github.com/go-resty/resty/v2"
)

func FetchFromTMDB(url string) (*resty.Response, error) {
	client := resty.New()
	response, err := client.R().
		SetHeader("accept", "application/json").
		SetAuthToken(config.LoadConfig().TMOBAPIKey).
		Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
