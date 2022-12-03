package burzedzis

import (
	"context"

	"github.com/kamilwoloszyn/burze-dzis/adapter/burzedzisnet"
	"github.com/kamilwoloszyn/burze-dzis/app"
	"github.com/kamilwoloszyn/burze-dzis/domain"
)

// Service contains all components necessary to run app properly.
type BurzeDzisClient struct {
	s app.Service
}

// NewService creates a new burzedzis client
func NewClient(client app.HTTPDoer, apiKey string, host string) BurzeDzisClient {
	service := app.NewService(
		burzedzisnet.NewClient(client, apiKey, host),
	)
	return BurzeDzisClient{
		s: *service,
	}
}

// IsValidKey checks validity provided key.
// If an error will occur, then returning value will be set to false
// with error specified. If value is set false without any error, then
// a provided key is wrong.
func (b *BurzeDzisClient) IsValidKey(ctx context.Context, key string) (bool, error) {
	return b.s.IsValidKey(ctx, key)
}

// CityLocation returns a coordinate of a city. If the city doesn't exist, then empty response without error will be returned.
func (b *BurzeDzisClient) CityLocation(ctx context.Context, cityName string) (domain.CityLocation, error) {
	return b.s.CityLocation(ctx, cityName)
}

// Cities returns a list of suggestion of cities. If a provided keyword won't match any city, then empty list will be returned.
func (b *BurzeDzisClient) Cities(ctx context.Context, keyword string, countryCode string) (domain.Cities, error) {
	return b.s.Cities(ctx, keyword, countryCode)
}

// StormSearch returns some data about thunderstorm in / arround the provided city.
// If the city doesn't exist, expect an error.
func (b *BurzeDzisClient) StormSearch(ctx context.Context, cityName string, radius int) (domain.Storm, error) {
	return b.s.StormSearch(ctx, cityName, radius)
}

// WeatherAlert returns weather alerts based on a provided city. If the city does't exist
// expect an error.
func (b *BurzeDzisClient) WeatherAlert(ctx context.Context, cityName string) ([]domain.Alert, error) {
	return b.s.WeatherAlert(ctx, cityName)
}
