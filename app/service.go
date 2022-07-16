package app

import (
	"context"
	"fmt"

	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/domain/vxml"
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
	if stormReq.Body.StormSearch.CityName != "" {
		cityLocation, err := s.CityLocation(
			ctx,
			vxml.NewCityLocationRequest(
				stormReq.Body.StormSearch.CityName,
				stormReq.Body.StormSearch.APIKey,
			),
		)
		if err != nil {
			return domain.Storm{}, fmt.Errorf("StormSearch: couldn't obtain a city coords: %v", err)
		}
		stormReq.Body.StormSearch.CoordY = cityLocation.CoordY
		stormReq.Body.StormSearch.CoordX = cityLocation.CoordX
	}
	return s.burzeDzisClient.StormSearch(ctx, stormReq)
}

func (s *service) WeatherAlert(ctx context.Context, alertReq vxml.WeatherAlertRequest) ([]domain.Alert, error) {
	if alertReq.Body.WeatherAlert.CityName != "" {
		cityLocation, err := s.CityLocation(
			ctx,
			vxml.NewCityLocationRequest(
				alertReq.Body.WeatherAlert.CityName,
				alertReq.Body.WeatherAlert.APIKey,
			),
		)
		if err != nil {
			return nil, fmt.Errorf("StormSearch: couldn't obtain a city coords: %v", err)
		}
		alertReq.Body.WeatherAlert.CoordY = cityLocation.CoordY
		alertReq.Body.WeatherAlert.CoordX = cityLocation.CoordX
	}
	return s.burzeDzisClient.WeatherAlert(ctx, alertReq)
}
