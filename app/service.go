package app

import (
	"context"

	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/vxml"
)

type service struct {
	burzeDzisClient BurzeDzis
}

func NewService(burzeDzisClient BurzeDzis) *service {
	return &service{
		burzeDzisClient: burzeDzisClient,
	}
}

func (s *service) IsValidKey(ctx context.Context, keyReq vxml.APIKeyRequest) (bool, error) {
	return s.burzeDzisClient.IsValidKey(ctx, keyReq)
}

func (s *service) CityLocation(ctx context.Context, locationReq vxml.CityLocationRequest) (domain.CityLocation, error) {
	return s.burzeDzisClient.CityLocation(ctx, locationReq)
}

func (s *service) Cities(ctx context.Context, citiesReq vxml.CitiesRequest) (domain.Cities, error) {
	return s.burzeDzisClient.Cities(ctx, citiesReq)
}

func (s *service) StormSearch(ctx context.Context, stormReq vxml.StormSearchRequest) (domain.Storm, error) {
	return s.burzeDzisClient.StormSearch(ctx, stormReq)
}

func (s *service) WeatherAlert(ctx context.Context, alertReq vxml.WeatherAlertRequest) (domain.Alert, error) {
	return s.burzeDzisClient.WeatherAlert(ctx, alertReq)
}
