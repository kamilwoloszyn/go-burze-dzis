package app

import (
	"context"

	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/vxml"
)

type BurzeDzis interface {
	IsValidKey(context.Context, vxml.APIKeyRequest) (bool, error)
	CityLocation(context.Context, vxml.CityLocationRequest) (domain.CityLocation, error)
	Cities(context.Context, vxml.CitiesRequest) (domain.Cities, error)
	StormSearch(context.Context, vxml.StormSearchRequest) (domain.Storm, error)
	WeatherAlert(context.Context, vxml.WeatherAlertRequest) (domain.Alert, error)
}
