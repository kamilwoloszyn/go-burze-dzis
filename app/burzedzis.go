package app

import (
	"context"

	"github.com/kamilwoloszyn/burze-dzis/domain"
)

type BurzeDzisService interface {
	IsValidKey(ctx context.Context, key string) (bool, error)
	CityLocation(ctx context.Context, searchData SearchData) (domain.CityLocation, error)
	Cities(ctx context.Context, searchData SearchData) (domain.Cities, error)
	StormSearch(ctx context.Context, searchData SearchData) (domain.Storm, error)
	WeatherAlert(ctx context.Context, searchData SearchData) ([]domain.Alert, error)
}
