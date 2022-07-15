package vxml

import (
	"encoding/xml"
)

type APIKeyRequest struct {
	Envelope xml.Name `xml:"Envelope"`
	Body     struct {
		XMLName xml.Name `xml:"KeyAPI"`
		Key     string   `xml:"klucz"`
	}
}

func NewAPIKeyRequest(APIKey string) APIKeyRequest {
	return APIKeyRequest{
		Body: struct {
			XMLName xml.Name "xml:\"KeyAPI\""
			Key     string   "xml:\"klucz\""
		}{
			Key: APIKey,
		},
	}
}

type CityLocationRequest struct {
	XMLName xml.Name
	Body    struct {
		XMLName xml.Name
		Name    string `xml:"nazwa"`
		APIKey  string `xml:"klucz"`
	}
}

type CitiesRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName     xml.Name
		Name        string `xml:"nazwa"`
		CountryCode string `xml:"kraj"`
		APIKey      string `xml:"klucz"`
	} `xml:"Body"`
}

// <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
//     <Body>
//         <miejscowosci_lista xmlns="https://burze.dzis.net/soap.php">
//             <nazwa>[string]</nazwa>
//             <kraj>[string]</kraj>
//             <klucz>[string]</klucz>
//         </miejscowosci_lista>
//     </Body>
// </Envelope>

type StormSearchRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"szukaj_burzy"`
		CoordY  float32  `xml:"y"`
		CoordX  float32  `xml:"x"`
		Radius  int      `xml:"promien"`
		APIKey  string   `xml:"klucz"`
	} `xml:"Body"`
}

type WeatherAlertRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"ostrzezenia_pogodowe"`
		CoordY  float32  `xml:"y"`
		CoordX  float32  `xml:"x"`
		APIKey  string   `xml:"klucz"`
	} `xml:"Body"`
}

// <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
//     <Body>
//         <ostrzezenia_pogodowe xmlns="https://burze.dzis.net/soap.php">
//             <y>[float]</y>
//             <x>[float]</x>
//             <klucz>[string]</klucz>
//         </ostrzezenia_pogodowe>
//     </Body>
// </Envelope>
