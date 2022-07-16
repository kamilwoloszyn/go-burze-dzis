package vxml

type APIKeyRequest struct {
	Body struct {
		KeyAPI struct {
			Key string `xml:"klucz"`
		}
	}
}

func NewAPIKeyRequest(APIKeyParam string) APIKeyRequest {
	return APIKeyRequest{
		Body: struct {
			KeyAPI struct {
				Key string "xml:\"klucz\""
			}
		}{
			KeyAPI: struct {
				Key string "xml:\"klucz\""
			}{
				Key: APIKeyParam,
			},
		},
	}
}

type CityLocationRequest struct {
	Body struct {
		City struct {
			Name   string  `xml:"nazwa"`
			APIKey *string `xml:"klucz,omitempty"`
		} `xml:"miejscowosc"`
	}
}

func NewCityLocationRequest(cityName string, APIKey *string) CityLocationRequest {
	cityLocation := CityLocationRequest{
		Body: struct {
			City struct {
				Name   string  "xml:\"nazwa\""
				APIKey *string "xml:\"klucz,omitempty\""
			} "xml:\"miejscowosc\""
		}{
			City: struct {
				Name   string  "xml:\"nazwa\""
				APIKey *string "xml:\"klucz,omitempty\""
			}{
				Name: cityName,
			},
		},
	}
	if APIKey != nil {
		cityLocation.Body.City.APIKey = APIKey
	}
	return cityLocation
}

type CitiesRequest struct {
	Body struct {
		CitiesList struct {
			Name        string  `xml:"nazwa"`
			CountryCode string  `xml:"kraj"`
			APIKey      *string `xml:"klucz,omitempty"`
		} `xml:"miejscowosci_lista"`
	} `xml:"Body"`
}

func NewCitiesRequest(cityKeyword, countryCode string, APIKey *string) CitiesRequest {
	citiesRequest := CitiesRequest{
		Body: struct {
			CitiesList struct {
				Name        string  "xml:\"nazwa\""
				CountryCode string  "xml:\"kraj\""
				APIKey      *string "xml:\"klucz,omitempty\""
			} "xml:\"miejscowosci_lista\""
		}{
			CitiesList: struct {
				Name        string  "xml:\"nazwa\""
				CountryCode string  "xml:\"kraj\""
				APIKey      *string "xml:\"klucz,omitempty\""
			}{
				Name:        cityKeyword,
				CountryCode: countryCode,
			},
		},
	}
	if APIKey != nil {
		citiesRequest.Body.CitiesList.APIKey = APIKey
	}
	return citiesRequest
}

type StormSearchRequest struct {
	Body struct {
		StormSearch struct {
			CityName string  `xml:"-"`
			CoordY   float32 `xml:"y"`
			CoordX   float32 `xml:"x"`
			Radius   int     `xml:"promien"`
			APIKey   *string `xml:"klucz,omitempty"`
		} `xml:"szukaj_burzy"`
	} `xml:"Body"`
}

func NewStormSearchRequest(cityName string, radius int, APIKey *string) StormSearchRequest {
	stormSearchRequest := StormSearchRequest{
		Body: struct {
			StormSearch struct {
				CityName string  "xml:\"-\""
				CoordY   float32 "xml:\"y\""
				CoordX   float32 "xml:\"x\""
				Radius   int     "xml:\"promien\""
				APIKey   *string "xml:\"klucz,omitempty\""
			} "xml:\"szukaj_burzy\""
		}{
			StormSearch: struct {
				CityName string  "xml:\"-\""
				CoordY   float32 "xml:\"y\""
				CoordX   float32 "xml:\"x\""
				Radius   int     "xml:\"promien\""
				APIKey   *string "xml:\"klucz,omitempty\""
			}{
				CityName: cityName,
				Radius:   radius,
			},
		},
	}
	if APIKey != nil {
		stormSearchRequest.Body.StormSearch.APIKey = APIKey
	}
	return stormSearchRequest
}

type WeatherAlertRequest struct {
	Body struct {
		WeatherAlert struct {
			CityName string  `xml:"-"`
			CoordY   float32 `xml:"y"`
			CoordX   float32 `xml:"x"`
			APIKey   *string `xml:"klucz,omitempty"`
		} `xml:"ostrzezenia_pogodowe"`
	}
}

func NewWeatherAlertRequest(cityName string, APIKey *string) WeatherAlertRequest {
	weatherAlertReq := WeatherAlertRequest{
		Body: struct {
			WeatherAlert struct {
				CityName string  "xml:\"-\""
				CoordY   float32 "xml:\"y\""
				CoordX   float32 "xml:\"x\""
				APIKey   *string "xml:\"klucz,omitempty\""
			} "xml:\"ostrzezenia_pogodowe\""
		}{
			WeatherAlert: struct {
				CityName string  "xml:\"-\""
				CoordY   float32 "xml:\"y\""
				CoordX   float32 "xml:\"x\""
				APIKey   *string "xml:\"klucz,omitempty\""
			}{
				CityName: cityName,
			},
		},
	}
	if APIKey != nil {
		weatherAlertReq.Body.WeatherAlert.APIKey = APIKey
	}
	return weatherAlertReq
}
