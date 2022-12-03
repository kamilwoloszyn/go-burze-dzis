# burze-dzis
A golang client for burze.dzis.net

# Quick start

```
go get -u github.com/kamilwoloszyn/burze-dzis
```

Then, you can start using the client:

```
func main() {
	config := config.NewConfig(
		"<YOUR_API_KEY>",
		"https://burze.dzis.net",
	)

	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancelFunc()

	var httpClient http.Client

	burzedzisClient := burzedzis.NewClient(
		&httpClient,
		config.APIKey,
		config.APIHost,
	)

	valid, err := burzedzisClient.IsValidKey(
		ctx,
		"<YOUR_API_KEY>",
	)

	if err != nil {
		log.Printf("got an error during checking key validity: %v", err)
	}
	fmt.Printf("valid key: %v\n", valid)

	cityLocation, err := burzedzisClient.CityLocation(
		ctx,
		"Warszawa",
	)
	if err != nil {
		log.Printf("got an error during getting a city location: %v", err)
	}
	fmt.Printf("city location: %v\n", cityLocation)
	
	citiesSuggestion, err := burzedzisClient.Cities(
		ctx,
		"War",
		"PL",
	)
	if err != nil {
		log.Printf("got an error during getting a city suggestions: %v", err)
	}
	fmt.Printf("cities suggestion:%v\n", citiesSuggestion)
	
	storm, err := burzedzisClient.StormSearch(
		ctx,
		"Warszawa",
		15,
	)
	if err != nil {
		log.Printf("got an err during getting a storm: %v", err)
	}
	fmt.Printf("a storm result: %v", storm)

	alerts, err := burzedzisClient.WeatherAlert(
		ctx,
		"Warszawa",
	)
	if err != nil {
		log.Printf("got an err during getting alerts: %v", err)
	}
	fmt.Printf("weather alerts: %v", alerts)
}

```

# Docs

Available [here](https://pkg.go.dev/github.com/kamilwoloszyn/burze-dzis)

# Troubleshooting 

Through a label issue available on Github.