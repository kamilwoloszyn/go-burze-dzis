package vxml

import (
	"strings"

	"github.com/kamilwoloszyn/burze-dzis/domain"
)

const CitiesResponseSeparator = ","

type APIKeyResponse struct {
	Body struct {
		KeyAPIResponse struct {
			Return bool `xml:"return"`
		}
	}
}

type CityLocationResponse struct {
	Body struct {
		City struct {
			Return struct {
				CoordY float32 `xml:"y"`
				CoordX float32 `xml:"x"`
			} `xml:"return"`
		} `xml:"miejscowoscResponse"`
	}
}

func (c *CityLocationResponse) ToCityLocation() domain.CityLocation {
	return domain.CityLocation{
		CoordX: c.Body.City.Return.CoordX,
		CoordY: c.Body.City.Return.CoordY,
	}
}

type CitiesResponse struct {
	Body struct {
		List struct {
			Return string `xml:"return"`
		} `xml:"miejscowosci_listaResponse"`
	} `xml:"Body"`
}

func (c *CitiesResponse) ToCities() domain.Cities {
	response := c.Body.List.Return
	response = strings.ReplaceAll(response, "[", "")
	response = strings.ReplaceAll(response, "]", "")
	response = strings.ReplaceAll(response, "\"", "")

	cities := strings.Split(response, CitiesResponseSeparator)
	if len(cities) < 2 {
		return domain.Cities{}
	}
	// make from 1st position, because the first element is always a keyword you search for
	cities = cities[1:]
	return domain.Cities{
		Cities: cities,
	}
}

type StormSearchResponse struct {
	Body struct {
		StormResponse struct {
			Return struct {
				Quantity     int     `xml:"liczba"`
				Distance     float32 `xml:"odleglosc"`
				Direction    string  `xml:"kierunek"`
				PeriodOfTime int     `xml:"okres"`
			} `xml:"return"`
		} `xml:"szukaj_burzyResponse"`
	}
}

func (s *StormSearchResponse) ToStorm() domain.Storm {
	return domain.Storm{
		Quantity:     s.Body.StormResponse.Return.Quantity,
		Distance:     s.Body.StormResponse.Return.Distance,
		Direction:    s.Body.StormResponse.Return.Direction,
		PeriodOfTime: s.Body.StormResponse.Return.PeriodOfTime,
	}
}

type WeatherAlertResponse struct {
	Body struct {
		WeatherAlert struct {
			Return struct {
				AlertFromDate        string `xml:"od_dnia"`
				AlertToDate          string `xml:"do_dnia"`
				Fost                 int    `xml:"mroz"`
				FrostFromDate        string `xml:"mroz_od_dnia"`
				FrostToDate          string `xml:"mroz_do_dnia"`
				Heat                 int    `xml:"upal"`
				HeatFromDate         string `xml:"upal_od_dnia"`
				HeatToDate           string `xml:"upal_do_dnia"`
				Wind                 int    `xml:"wiatr"`
				WindFromDate         string `xml:"wiatr_od_dnia"`
				WindToDate           string `xml:"wiatr_do_dnia"`
				Rain                 int    `xml:"opad"`
				RainFromDate         string `xml:"opad_od_dnia"`
				RainToDate           string `xml:"opad_do_dnia"`
				ThunderStorm         int    `xml:"burza"`
				ThunderStormFromDate string `xml:"burza_od_dnia"`
				ThunderStormToDate   string `xml:"burza_do_dnia"`
				Cyclone              int    `xml:"traba"`
				CycloneFromDate      string `xml:"traba_od_dnia"`
				CycloneToDate        string `xml:"traba_do_dnia"`
			} `xml:"return"`
		} `xml:"ostrzezenia_pogodoweResponse"`
	}
}

func (w *WeatherAlertResponse) ToWeatherAlert() []domain.Alert {
	alerts := []domain.Alert{}
	if w.Body.WeatherAlert.Return.Cyclone != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameCyclone,
			FromDate:  w.Body.WeatherAlert.Return.CycloneFromDate,
			ToDate:    w.Body.WeatherAlert.Return.CycloneToDate,
		})
	}
	if w.Body.WeatherAlert.Return.Fost != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameFrost,
			FromDate:  w.Body.WeatherAlert.Return.FrostFromDate,
			ToDate:    w.Body.WeatherAlert.Return.FrostToDate,
		})
	}
	if w.Body.WeatherAlert.Return.Heat != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameHeat,
			FromDate:  w.Body.WeatherAlert.Return.HeatFromDate,
			ToDate:    w.Body.WeatherAlert.Return.HeatToDate,
		})
	}
	if w.Body.WeatherAlert.Return.Rain != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameRain,
			FromDate:  w.Body.WeatherAlert.Return.RainFromDate,
			ToDate:    w.Body.WeatherAlert.Return.RainToDate,
		})
	}
	if w.Body.WeatherAlert.Return.ThunderStorm != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameThunderStorm,
			FromDate:  w.Body.WeatherAlert.Return.ThunderStormFromDate,
			ToDate:    w.Body.WeatherAlert.Return.ThunderStormToDate,
		})
	}
	if w.Body.WeatherAlert.Return.Wind != 0 {
		alerts = append(alerts, domain.Alert{
			AlertName: domain.AlertNameWind,
			FromDate:  w.Body.WeatherAlert.Return.WindFromDate,
			ToDate:    w.Body.WeatherAlert.Return.WindToDate,
		})
	}
	return alerts
}
