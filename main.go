package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kamilwoloszyn/burze-dzis/adapter/burzedzis"
	"github.com/kamilwoloszyn/burze-dzis/app"
	"github.com/kamilwoloszyn/burze-dzis/config"
	"github.com/kamilwoloszyn/burze-dzis/domain/vxml"
)

func main() {
	config := config.NewConfig(
		"",
		"",
	)

	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancelFunc()

	var httpClient http.Client

	burzedzisService := burzedzis.NewClient(
		&httpClient,
		config.APIKey,
		config.APIHost,
	)

	service := app.NewService(burzedzisService)

	valid, err := service.IsValidKey(
		ctx,
		vxml.NewAPIKeyRequest(config.APIKey),
	)
	if err != nil {
		log.Printf("got an error: %v", err)
	}

	fmt.Printf("valid key: %v", valid)
}
