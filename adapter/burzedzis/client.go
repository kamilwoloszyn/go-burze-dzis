package burzedzis

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/kamilwoloszyn/burze-dzis/domain"
	"github.com/kamilwoloszyn/burze-dzis/domain/vxml"
)

const (
	pathAPI = "/soap.php"
)

type BurzeDzisClient struct {
	apiKey     string
	host       string
	httpClient *http.Client
}

func NewClient(client *http.Client, apiKey string, host string) *BurzeDzisClient {
	return &BurzeDzisClient{
		apiKey:     apiKey,
		host:       host,
		httpClient: client,
	}
}

func (c *BurzeDzisClient) IsValidKey(ctx context.Context, keyReq vxml.APIKeyRequest) (bool, error) {
	data, err := xml.Marshal(keyReq)
	if err != nil {
		return false, err
	}
	response, err := c.makeRequest(ctx, data)
	if err != nil {
		return false, fmt.Errorf("IsValidKey: make a request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("IsValidKey: got unexpected err code: %d", response.StatusCode)
	}
	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("IsValidKey: reading a response: %v", err)
	}
	var responseData vxml.APIKeyResponse
	if err := xml.Unmarshal(rawResponse, &responseData); err != nil {
		return false, fmt.Errorf("IsValidKey: unmarshall a response: %v", err)
	}
	return responseData.Body.IsValid, nil

}

func (c *BurzeDzisClient) CityLocation(ctx context.Context, locationReq vxml.CityLocationRequest) (domain.CityLocation, error) {
	data, err := xml.Marshal(locationReq)
	if err != nil {
		return domain.CityLocation{}, err
	}
	response, err := c.makeRequest(ctx, data)
	if err != nil {
		return domain.CityLocation{}, fmt.Errorf("CityLocation: make a request: %v", err)
	}
	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return domain.CityLocation{}, fmt.Errorf("CityLocation: reading a response: %v", err)
	}
	var responseData vxml.CityLocationResponse
	if err := xml.Unmarshal(rawResponse, &responseData); err != nil {
		return domain.CityLocation{}, fmt.Errorf("CityLocation: unmarshall a response: %v", err)
	}
	var cityLocation domain.CityLocation
	responseData.ToCityLocation(&cityLocation)
	return cityLocation, nil
}

func (c *BurzeDzisClient) Cities(ctx context.Context, citiesReq vxml.CitiesRequest) (domain.Cities, error) {
	data, err := xml.Marshal(citiesReq)
	if err != nil {
		return domain.Cities{}, err
	}
	response, err := c.makeRequest(ctx, data)
	if err != nil {
		return domain.Cities{}, fmt.Errorf("Cities: make a request: %v", err)
	}
	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return domain.Cities{}, fmt.Errorf("Cities: reading a response: %v", err)
	}
	var responseData vxml.CitiesResponse
	if err := xml.Unmarshal(rawResponse, &responseData); err != nil {
		return domain.Cities{}, fmt.Errorf("Cities: unmarshall a response: %v", err)
	}
	var cities domain.Cities
	responseData.ToCities(&cities)
	return cities, nil
}

func (c *BurzeDzisClient) StormSearch(ctx context.Context, stormReq vxml.StormSearchRequest) (domain.Storm, error) {
	data, err := xml.Marshal(stormReq)
	if err != nil {
		return domain.Storm{}, err
	}
	response, err := c.makeRequest(ctx, data)
	if err != nil {
		return domain.Storm{}, fmt.Errorf("StormSearch: make a request: %v", err)
	}
	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return domain.Storm{}, fmt.Errorf("StormSearch: reading a response: %v", err)
	}
	var responseData vxml.StormSearchResponse
	if err := xml.Unmarshal(rawResponse, &responseData); err != nil {
		return domain.Storm{}, fmt.Errorf("StormSearch: unmarshall a response: %v", err)
	}
	var storm domain.Storm
	responseData.ToStorm(&storm)
	return storm, nil

}

func (c *BurzeDzisClient) WeatherAlert(ctx context.Context, alertReq vxml.WeatherAlertRequest) (domain.Alert, error) {
	data, err := xml.Marshal(alertReq)
	if err != nil {
		return domain.Alert{}, err
	}
	response, err := c.makeRequest(ctx, data)
	if err != nil {
		return domain.Alert{}, fmt.Errorf("WeatherAlert: make a request: %v", err)
	}
	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return domain.Alert{}, fmt.Errorf("WeatherAlert: reading a response: %v", err)
	}
	var responseData vxml.WeatherAlertResponse
	if err := xml.Unmarshal(rawResponse, &responseData); err != nil {
		return domain.Alert{}, fmt.Errorf("WeatherAlert: unmarshall a response: %v", err)
	}
	var alert domain.Alert
	responseData.ToWeatherAlert(&alert)
	return alert, nil

}

func (c *BurzeDzisClient) makeRequest(ctx context.Context, data []byte) (*http.Response, error) {
	reader := bytes.NewReader(data)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.host+pathAPI,
		reader,
	)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: %v", err)
	}
	return c.httpClient.Do(req)
}
