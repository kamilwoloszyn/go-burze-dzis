package app

import "github.com/kamilwoloszyn/burze-dzis/domain"

// SearchData is internal representation of data, can be modified by layers.
type SearchData struct {
	CityName     string
	CityKeyword  string
	Radius       int
	CountryCode  string
	CityLocation domain.CityLocation
}
