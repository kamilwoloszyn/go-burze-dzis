package domain

type Storm struct {
	Quantity     int     `xml:"liczba"`
	Distance     float32 `xml:"odleglosc"`
	Direction    string  `xml:"kierunek"`
	PeriodOfTime int     `xml:"okres"`
}
