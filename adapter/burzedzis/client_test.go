package burzedzis_test

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/kamilwoloszyn/burze-dzis/adapter/burzedzis"
	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/domain/vxml"
	"github.com/kamilwoloszyn/burze-dzis/generics"
	"github.com/kamilwoloszyn/burze-dzis/mock"
)

const (
	apiKey = "testAPIKey12356"
	host   = "burze.dzis.net"
)

func TestIsValidKey(t *testing.T) {
	testCases := []struct {
		name                string
		request             vxml.APIKeyRequest
		mockedHTTPResponse  *http.Response
		expectedRequestBody *string
		expectedResult      bool
		expectedErr         bool
	}{
		{
			name:    "valid key",
			request: vxml.NewAPIKeyRequest(apiKey),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:KeyAPIResponse>
									<return xsi:type="xsd:boolean">true</return>
								</ns1:KeyAPIResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: generics.Ptr(""),
			expectedResult:      true,
		},
		{
			name: "invalid key",
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:KeyAPIResponse>
									<return xsi:type="xsd:boolean">false</return>
								</ns1:KeyAPIResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: generics.Ptr(""),
			expectedResult:      false,
		},
		{
			name: "bad request",
			mockedHTTPResponse: &http.Response{
				Status:     "500 Internal Server Error",
				StatusCode: 500,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
							<SOAP-ENV:Body>
								<SOAP-ENV:Fault>
									<faultcode>SOAP-ENV:Client</faultcode>
									<faultstring>Bad Request</faultstring>
								</SOAP-ENV:Fault>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockedHTTP := mock.MockedHTTP{
				MockedDoFunc: func(r *http.Request) (*http.Response, error) {
					// TODO: checking a request
					return tt.mockedHTTPResponse, nil
				},
			}
			client := burzedzis.NewClient(mockedHTTP, apiKey, host)
			valid, err := client.IsValidKey(
				ctx,
				tt.request,
			)
			if err != nil && !tt.expectedErr {
				t.Errorf("unexpected err %v in %s", err, tt.name)
			}
			if valid != tt.expectedResult {
				t.Errorf("wrong response, %v != %v", valid, tt.expectedResult)
			}

		})
	}
}

func TestCityLocation(t *testing.T) {
	testCases := []struct {
		name                string
		request             vxml.CityLocationRequest
		mockedHTTPResponse  *http.Response
		expectedRequestBody *string
		expectedResult      domain.CityLocation
		expectedErr         bool
	}{
		{
			name: "a valid response",
			request: vxml.NewCityLocationRequest(
				"Zadupie",
				generics.Ptr(apiKey),
			),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:miejscowoscResponse>
									<return xsi:type="ns1:MyComplexTypeMiejscowosc">
										<y xsi:type="xsd:float">10.1</y>
										<x xsi:type="xsd:float">33.02</x>
									</return>
								</ns1:miejscowoscResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>						
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedResult: domain.CityLocation{
				CoordX: 33.02,
				CoordY: 10.1,
			},
		},
		{
			name: "a wrong response",
			request: vxml.NewCityLocationRequest(
				"Zadupie Wielkie",
				nil,
			),
			mockedHTTPResponse: &http.Response{
				Status:     "500 Internal Server Error",
				StatusCode: 500,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
							<SOAP-ENV:Body>
								<SOAP-ENV:Fault>
									<faultcode>SOAP-ENV:Server</faultcode>
									<faultstring>W celu skorzystania z funkcji, wymagane jest uwierzytelnienie lub przekroczono limit zapytań / To use the function, authentication is required or the query limit has been exceeded</faultstring>
								</SOAP-ENV:Fault>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>												
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedResult:      domain.CityLocation{},
			expectedErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockedHTTP := mock.MockedHTTP{
				MockedDoFunc: func(r *http.Request) (*http.Response, error) {
					// TODO: checking a request
					return tt.mockedHTTPResponse, nil
				},
			}
			client := burzedzis.NewClient(mockedHTTP, apiKey, host)
			result, err := client.CityLocation(ctx, tt.request)
			if err != nil && !tt.expectedErr {
				t.Errorf("unexpected err %v in %s", err, tt.name)
			}
			if equal := reflect.DeepEqual(result, tt.expectedResult); !equal {
				t.Errorf("result diff %v != %v", result, tt.expectedResult)
			}
		})
	}
}

func TestCities(t *testing.T) {
	testCases := []struct {
		name                string
		request             vxml.CitiesRequest
		mockedHTTPResponse  *http.Response
		expectedRequestBody *string
		expectedResult      domain.Cities
		expectedErr         bool
	}{
		{
			name: "a valid response",
			request: vxml.NewCitiesRequest(
				"Ole",
				"PL",
				generics.Ptr(apiKey),
			),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:miejscowosci_listaResponse>
									<return xsi:type="xsd:string">["Ole",["Olecko", "Olędy (gm. Mordy)", "Olędy (gm. Zbuczyn)", "Olędzkie", "Oleksianka", "Oleksin (pow. Otwocki)", "Oleksin (pow. Siedlecki)"]]</return>
								</ns1:miejscowosci_listaResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>												
						`,
					),
				),
				ContentLength: -1,
			},
			expectedResult: domain.Cities{
				Cities: []string{
					"Olecko",
					"Olędy (gm. Mordy)",
					"Olędy (gm. Zbuczyn)",
					"Olędzkie",
					"Oleksianka",
					"Oleksin (pow. Otwocki)",
					"Oleksin (pow. Siedlecki)",
				},
			},
			expectedRequestBody: nil,
		},
		{
			name: "an empty response",
			request: vxml.NewCitiesRequest(
				"Oleee",
				"PL",
				generics.Ptr(apiKey),
			),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:miejscowosci_listaResponse>
									<return xsi:type="xsd:string">["Oleee",[]]</return>
								</ns1:miejscowosci_listaResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>																	
						`,
					),
				),
				ContentLength: -1,
			},
			expectedResult:      domain.Cities{},
			expectedRequestBody: nil,
		},
		{
			name: "invalid response",
			mockedHTTPResponse: &http.Response{
				Status:     "500 Internal Server Error",
				StatusCode: 500,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
							<SOAP-ENV:Body>
								<SOAP-ENV:Fault>
									<faultcode>SOAP-ENV:Server</faultcode>
									<faultstring>W celu skorzystania z funkcji, wymagane jest uwierzytelnienie lub przekroczono limit zapytań / To use the function, authentication is required or the query limit has been exceeded</faultstring>
								</SOAP-ENV:Fault>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>												
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedResult:      domain.Cities{},
			expectedErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockedHTTP := mock.MockedHTTP{
				MockedDoFunc: func(r *http.Request) (*http.Response, error) {
					// TODO: checking a request
					return tt.mockedHTTPResponse, nil
				},
			}
			client := burzedzis.NewClient(mockedHTTP, apiKey, host)
			result, err := client.Cities(ctx, tt.request)
			if err != nil && !tt.expectedErr {
				t.Errorf("unexpected err %v in %s", err, tt.name)
			}
			if equal := reflect.DeepEqual(result, tt.expectedResult); !equal {
				t.Errorf("result diff %v != %v", result, tt.expectedResult)
			}
		})
	}
}

func TestStormSearch(t *testing.T) {
	testCases := []struct {
		name                string
		request             vxml.StormSearchRequest
		mockedHTTPResponse  *http.Response
		expectedRequestBody *string
		expectedResult      domain.Storm
		expectedErr         bool
	}{
		{
			name:    "valid response",
			request: vxml.NewStormSearchRequest("Zadupie Wielkie", 10, generics.Ptr(apiKey)),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:szukaj_burzyResponse>
									<return xsi:type="ns1:MyComplexTypeBurza">
										<liczba xsi:type="xsd:int">11</liczba>
										<odleglosc xsi:type="xsd:float">0.84</odleglosc>
										<kierunek xsi:type="xsd:string">NW</kierunek>
										<okres xsi:type="xsd:int">15</okres>
									</return>
								</ns1:szukaj_burzyResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>																								
						`,
					),
				),
				ContentLength: -1,
			},
			expectedResult: domain.Storm{
				Quantity:     11,
				Distance:     0.84,
				Direction:    "NW",
				PeriodOfTime: 15,
			},
		},
		{
			name: "invalid response",
			mockedHTTPResponse: &http.Response{
				Status:     "500 Internal Server Error",
				StatusCode: 500,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
							<SOAP-ENV:Body>
								<SOAP-ENV:Fault>
									<faultcode>SOAP-ENV:Server</faultcode>
									<faultstring>W celu skorzystania z funkcji, wymagane jest uwierzytelnienie lub przekroczono limit zapytań / To use the function, authentication is required or the query limit has been exceeded</faultstring>
								</SOAP-ENV:Fault>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>												
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedResult:      domain.Storm{},
			expectedErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockedHTTP := mock.MockedHTTP{
				MockedDoFunc: func(r *http.Request) (*http.Response, error) {
					// TODO: checking a request
					return tt.mockedHTTPResponse, nil
				},
			}
			client := burzedzis.NewClient(mockedHTTP, apiKey, host)
			result, err := client.StormSearch(ctx, tt.request)
			if err != nil && !tt.expectedErr {
				t.Errorf("unexpected err %v in %s", err, tt.name)
			}
			if equal := reflect.DeepEqual(result, tt.expectedResult); !equal {
				t.Errorf("result diff %v != %v", result, tt.expectedResult)
			}
		})
	}
}

func WeatherAlert(t *testing.T) {
	testCases := []struct {
		name                string
		request             vxml.WeatherAlertRequest
		mockedHTTPResponse  *http.Response
		expectedRequestBody *string
		expectedResult      domain.WeatherAlert
		expectedErr         bool
	}{
		{
			name:    "valid response",
			request: vxml.NewWeatherAlertRequest("Zadupie Wielkie", generics.Ptr(apiKey)),
			mockedHTTPResponse: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://burze.dzis.net/soap.php" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
							<SOAP-ENV:Body>
								<ns1:ostrzezenia_pogodoweResponse>
									<return xsi:type="ns1:MyComplexTypeOstrzezenia">
										<od_dnia xsi:type="xsd:string"></od_dnia>
										<do_dnia xsi:type="xsd:string"></do_dnia>
										<mroz xsi:type="xsd:int">0</mroz>
										<mroz_od_dnia xsi:type="xsd:string">0</mroz_od_dnia>
										<mroz_do_dnia xsi:type="xsd:string">0</mroz_do_dnia>
										<upal xsi:type="xsd:int">1</upal>
										<upal_od_dnia xsi:type="xsd:string">2022-07-23 07:30</upal_od_dnia>
										<upal_do_dnia xsi:type="xsd:string">2022-07-23 18:00</upal_do_dnia>
										<wiatr xsi:type="xsd:int">0</wiatr>
										<wiatr_od_dnia xsi:type="xsd:string">0</wiatr_od_dnia>
										<wiatr_do_dnia xsi:type="xsd:string">0</wiatr_do_dnia>
										<opad xsi:type="xsd:int">0</opad>
										<opad_od_dnia xsi:type="xsd:string">0</opad_od_dnia>
										<opad_do_dnia xsi:type="xsd:string">0</opad_do_dnia>
										<burza xsi:type="xsd:int">0</burza>
										<burza_od_dnia xsi:type="xsd:string">0</burza_od_dnia>
										<burza_do_dnia xsi:type="xsd:string">0</burza_do_dnia>
										<traba xsi:type="xsd:int">0</traba>
										<traba_od_dnia xsi:type="xsd:string">0</traba_od_dnia>
										<traba_do_dnia xsi:type="xsd:string">0</traba_do_dnia>
									</return>
								</ns1:ostrzezenia_pogodoweResponse>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>																														
						`,
					),
				),
				ContentLength: -1,
			},
			expectedResult: domain.WeatherAlert{
				Alerts: &[]domain.Alert{
					{
						AlertName: domain.AlertNameHeat,
						FromDate:  "2022-07-23 07:30",
						ToDate:    "2022-07-23 18:00",
					},
				},
			},
		},
		{
			name: "invalid response",
			mockedHTTPResponse: &http.Response{
				Status:     "500 Internal Server Error",
				StatusCode: 500,
				Body: io.NopCloser(
					strings.NewReader(
						`<?xml version="1.0" encoding="UTF-8"?>
						<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
							<SOAP-ENV:Body>
								<SOAP-ENV:Fault>
									<faultcode>SOAP-ENV:Server</faultcode>
									<faultstring>W celu skorzystania z funkcji, wymagane jest uwierzytelnienie lub przekroczono limit zapytań / To use the function, authentication is required or the query limit has been exceeded</faultstring>
								</SOAP-ENV:Fault>
							</SOAP-ENV:Body>
						</SOAP-ENV:Envelope>												
						`,
					),
				),
				ContentLength: -1,
			},
			expectedRequestBody: nil,
			expectedResult:      domain.WeatherAlert{},
			expectedErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockedHTTP := mock.MockedHTTP{
				MockedDoFunc: func(r *http.Request) (*http.Response, error) {
					// TODO: checking a request
					return tt.mockedHTTPResponse, nil
				},
			}
			client := burzedzis.NewClient(mockedHTTP, apiKey, host)
			result, err := client.WeatherAlert(ctx, tt.request)
			if err != nil && !tt.expectedErr {
				t.Errorf("unexpected err %v in %s", err, tt.name)
			}
			if equal := reflect.DeepEqual(result, tt.expectedResult); !equal {
				t.Errorf("result diff %v != %v", result, tt.expectedResult)
			}
		})
	}
}
