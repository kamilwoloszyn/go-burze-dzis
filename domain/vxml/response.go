package vxml

import (
	"encoding/xml"

	"github.com/kamilwoloszyn/burze-dzis/domain"
)

type APIKeyResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName xml.Name
		IsValid bool `xml:"return"`
	} `xml:"Body"`
}

// <?xml version="1.0" encoding="UTF-8"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//     <SOAP-ENV:Body>
//         <ns1:KeyAPIResponse>
//             <return xsi:type="xsd:boolean">true</return>
//         </ns1:KeyAPIResponse>
//     </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>

type CityLocationResponse struct {
	XMLName xml.Name
	Body    struct {
		CityResponse struct {
			Return struct {
				CoordY float32 `xml:"y"`
				CoordX float32 `xml:"x"`
			} `xml:"return"`
		} `xml:"miejscowoscResponse"`
	} `xml:"Body"`
}

func (c *CityLocationResponse) ToCityLocation(l *domain.CityLocation) {
}

// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//     <SOAP-ENV:Body>
//         <ns1:miejscowoscResponse>
//             <return xsi:type="ns1:MyComplexTypeMiejscowosc">
//                 <y xsi:type="xsd:float">50.1</y>
//                 <x xsi:type="xsd:float">23.02</x>
//             </return>
//         </ns1:miejscowoscResponse>
//     </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>

type CitiesResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName xml.Name
		Cities  []string `xml:"return"`
	} `xml:"Body"`
}

func (c *CitiesResponse) ToCities(cities *domain.Cities) {

}

// <?xml version="1.0" encoding="UTF-8"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//     <SOAP-ENV:Body>
//         <ns1:miejscowosci_listaResponse>
//             <return xsi:type="xsd:string">["Ole",["Olecko", "Olędy (gm. Mordy)", "Olędy (gm. Zbuczyn)", "Olędzkie", "Oleksianka", "Oleksin (pow. Otwocki)", "Oleksin (pow. Siedlecki)", "Oleksin (woj. Podlaskie)", "Olendry (woj. Łódzkie)", "Olendry (woj. Podlaskie)", "Olendry (woj. Wielkopolskie)", "Olesin (pow. Krasnostawski)", "Olesin (pow. Puławski)", "Oleśnica (pow. Chodzieski)", "Oleśnica (pow. Poddębicki) ", "Oleśnica (woj. Dolnośląskie)", "Oleśnica (woj. Świętokrzyskie)", "Oleśniczka", "Oleśnik (woj. Mazowieckie)", "Oleśniki", "Olesno", "Olesno (woj. Małopolskie)", "Oleszka", "Oleszkowo", "Oleszna", "Oleszna Podgórska", "Oleszno (woj. Kujawsko-Pomorskie)", "Oleszno (woj. Świętokrzyskie)", "Oleszyce", "Olewin (woj. Łódzkie)", "Olewin (woj. Małopolskie)"]]</return>
//         </ns1:miejscowosci_listaResponse>
//     </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>

type StormSearchResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName xml.Name
		Return  struct {
			Quantity     int     `xml:"liczba"`
			Distance     float32 `xml:"odleglosc"`
			Direction    string  `xml:"kierunek"`
			PeriodOfTime int     `xml:"okres"`
		} `xml:"return"`
	} `xml:"Body"`
}

func (s *StormSearchResponse) ToStorm(storm *domain.Storm) {

}

// <?xml version="1.0" encoding="UTF-8"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//     <SOAP-ENV:Body>
//         <ns1:szukaj_burzyResponse>
//             <return xsi:type="ns1:MyComplexTypeBurza">
//                 <liczba xsi:type="xsd:int">0</liczba>
//                 <odleglosc xsi:type="xsd:float">0</odleglosc>
//                 <kierunek xsi:type="xsd:string"></kierunek>
//                 <okres xsi:type="xsd:int">15</okres>
//             </return>
//         </ns1:szukaj_burzyResponse>
//     </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>

type WeatherAlertResponse struct {
	XMLName xml.Name `xml:"SOAP-ENV"`
	Body    struct {
		XMLName xml.Name `xml:"ns1"`
		Return  struct {
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
		} `xml:"Return"`
	} `xml:"Body"`
}

func (w *WeatherAlertResponse) ToWeatherAlert(a *domain.Alert) {

}

// <?xml version="1.0" encoding="UTF-8"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//     <SOAP-ENV:Body>
//         <ns1:ostrzezenia_pogodoweResponse>
//             <return xsi:type="ns1:MyComplexTypeOstrzezenia">
//                 <od_dnia xsi:type="xsd:string"></od_dnia>
//                 <do_dnia xsi:type="xsd:string"></do_dnia>
//                 <mroz xsi:type="xsd:int">0</mroz>
//                 <mroz_od_dnia xsi:type="xsd:string">0</mroz_od_dnia>
//                 <mroz_do_dnia xsi:type="xsd:string">0</mroz_do_dnia>
//                 <upal xsi:type="xsd:int">0</upal>
//                 <upal_od_dnia xsi:type="xsd:string">0</upal_od_dnia>
//                 <upal_do_dnia xsi:type="xsd:string">0</upal_do_dnia>
//                 <wiatr xsi:type="xsd:int">0</wiatr>
//                 <wiatr_od_dnia xsi:type="xsd:string">0</wiatr_od_dnia>
//                 <wiatr_do_dnia xsi:type="xsd:string">0</wiatr_do_dnia>
//                 <opad xsi:type="xsd:int">0</opad>
//                 <opad_od_dnia xsi:type="xsd:string">0</opad_od_dnia>
//                 <opad_do_dnia xsi:type="xsd:string">0</opad_do_dnia>
//                 <burza xsi:type="xsd:int">0</burza>
//                 <burza_od_dnia xsi:type="xsd:string">0</burza_od_dnia>
//                 <burza_do_dnia xsi:type="xsd:string">0</burza_do_dnia>
//                 <traba xsi:type="xsd:int">0</traba>
//                 <traba_od_dnia xsi:type="xsd:string">0</traba_od_dnia>
//                 <traba_do_dnia xsi:type="xsd:string">0</traba_do_dnia>
//             </return>
//         </ns1:ostrzezenia_pogodoweResponse>
//     </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>
