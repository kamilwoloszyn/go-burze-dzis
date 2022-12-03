package app

import (
	"context"
	"fmt"
	"reflect"

	"github.com/kamilwoloszyn/burze-dzis/domain"
)

// Service contains all components necessary to run app properly.
type Service struct {
	burzeDzisClient BurzeDzisService
}

// NewService creates a new service
func NewService(burzeDzisClient BurzeDzisService) *Service {
	return &Service{
		burzeDzisClient: burzeDzisClient,
	}
}

// IsValidKey checks validity provided key.
// If an error will occur, then returning value will be set to false
// with error specified. If value is set false without any error, then
// a provided key is wrong.
func (s *Service) IsValidKey(ctx context.Context, key string) (bool, error) {
	return s.burzeDzisClient.IsValidKey(ctx, key)
}

// CityLocation returns a coordinate of a city. If the city doesn't exist, then empty response without error will be returned.
func (s *Service) CityLocation(ctx context.Context, cityName string) (domain.CityLocation, error) {
	return s.burzeDzisClient.CityLocation(ctx, SearchData{CityName: cityName})
}

// Cities returns a list of suggestion of cities. If a provided keyword won't match any city, then empty list will be returned.
func (s *Service) Cities(ctx context.Context, keyword string, countryCode string) (domain.Cities, error) {
	return s.burzeDzisClient.Cities(ctx, SearchData{CityKeyword: keyword, CountryCode: countryCode})
}

// StormSearch returns some data about thunderstorm in / arround the provided city.
// If the city doesn't exist, expect an error.
func (s *Service) StormSearch(ctx context.Context, cityName string, radius int) (domain.Storm, error) {
	if cityName == "" {
		return domain.Storm{}, fmt.Errorf("StormSearch: No city provided")
	}
	cityLocation, err := s.CityLocation(
		ctx,
		cityName,
	)
	if err != nil {
		return domain.Storm{}, fmt.Errorf("StormSearch: couldn't obtain a city coords: %v", err)
	}
	if equal := reflect.DeepEqual(cityLocation, domain.CityLocation{}); equal {
		return domain.Storm{}, fmt.Errorf("StormSearch: wrong coords received. Is a correct city provided ? ")
	}

	return s.burzeDzisClient.StormSearch(ctx, SearchData{
		CityName: cityName,
		Radius:   radius,
		CityLocation: domain.CityLocation{
			CoordX: cityLocation.CoordX,
			CoordY: cityLocation.CoordY,
		},
	})
}

// WeatherAlert returns weather alerts based on a provided city. If the city does't exist
// expect an error.
func (s *Service) WeatherAlert(ctx context.Context, cityName string) ([]domain.Alert, error) {
	if cityName == "" {
		return nil, fmt.Errorf("WeatherAlert: no city provided")
	}
	cityLocation, err := s.CityLocation(
		ctx,
		cityName,
	)
	if err != nil {
		return nil, fmt.Errorf("WeatherAlert: couldn't obtain a city coords: %v", err)
	}
	if equal := reflect.DeepEqual(cityLocation, domain.CityLocation{}); equal {
		return nil, fmt.Errorf("WeatherAlert: wrong coords received. Is a correct city provided ? ")
	}
	return s.burzeDzisClient.WeatherAlert(
		ctx,
		SearchData{
			CityName: cityName,
			CityLocation: domain.CityLocation{
				CoordX: cityLocation.CoordX,
				CoordY: cityLocation.CoordY,
			},
		})
}
