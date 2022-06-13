package domain

import "time"

type AlertType string

const (
	AlertNameFrost        = "Frost"
	AlertNameHeat         = "Heat"
	AlertNameWind         = "Wind"
	AlertNameThunderStorm = "Thunerstorm"
	AlertNameCyclone      = "Cyclone"
)

type WeatherAlert struct {
	FromDate string `xml:"od_dnia"`
	ToDate   string `xml:"do_dnia"`
	Alerts   *[]Alert
}

func (a *WeatherAlert) AddAlert(alert Alert) error {
	alerts := []Alert{}
	if err := alert.Validate(); err != nil {
		return err
	}
	alerts = append(alerts, *a.Alerts...)
	alerts = append(alerts, alert)
	a.Alerts = &alerts
	return nil
}

type Alert struct {
	AlertName AlertType
	FromDate  string
	ToDate    string
}

func (a *Alert) Validate() error {
	if found := alertEnum(string(a.AlertName)); !found {
		return ErrWrongAlertType
	}
	return nil
}

func NewAlert(
	name string,
	fromDate string,
	toDate string,
) (Alert, error) {
	var alert Alert
	if found := alertEnum(name); !found {
		return Alert{}, ErrWrongAlertType
	}
	alert.AlertName = AlertType(name)
	if fromDate == "" {
		alert.FromDate = time.Now().String()
	}
	if toDate != "" {
		alert.ToDate = toDate
	}
	return alert, nil
}

func alertEnum(alertName string) bool {
	switch alertName {
	case string(AlertNameFrost),
		string(AlertNameHeat),
		string(AlertNameWind),
		string(AlertNameThunderStorm),
		string(AlertNameCyclone):
		return true
	default:
		return false
	}
}
