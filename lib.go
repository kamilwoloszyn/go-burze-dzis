package burzedzis

import (
	"context"

	"github.com/kamilwoloszyn/burze-dzis/adapter/burzedzisnet"
	"github.com/kamilwoloszyn/burze-dzis/app"
	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/domain/vxml"
)

type BurzeDzisClient struct {
	s app.Service
}

func NewClient(client app.HTTPDoer, apiKey string, host string) BurzeDzisClient {
	service := app.NewService(
		burzedzisnet.NewClient(client, apiKey, host),
	)
	return BurzeDzisClient{
		s: *service,
	}
}

func (b *BurzeDzisClient) IsValidKey(ctx context.Context, keyReq vxml.APIKeyRequest) (bool, error) {
	return b.s.IsValidKey(ctx, keyReq)
}

func (b *BurzeDzisClient) CityLocation(ctx context.Context, locationReq vxml.CityLocationRequest) (domain.CityLocation, error) {
	return b.s.CityLocation(ctx, locationReq)
}

func (b *BurzeDzisClient) Cities(ctx context.Context, citiesReq vxml.CitiesRequest) (domain.Cities, error) {
	return b.s.Cities(ctx, citiesReq)
}

func (b *BurzeDzisClient) StormSearch(ctx context.Context, stormReq vxml.StormSearchRequest) (domain.Storm, error) {
	return b.StormSearch(ctx, stormReq)
}

func (b *BurzeDzisClient) WeatherAlert(ctx context.Context, alertReq vxml.WeatherAlertRequest) ([]domain.Alert, error) {
	return b.WeatherAlert(ctx, alertReq)
}
