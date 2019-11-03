package main

import (
	"github.com/alexsniffin/go-places/internal/api/clients"
	"github.com/alexsniffin/go-places/internal/api/config"
	"github.com/alexsniffin/go-places/internal/api/places"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("Running google places sdk example")

	config := config.NewConfig(sugar, "config")

	mapsClient := clients.NewClient(sugar, config.Cfg.MapsAPI.APIKey)
	places := places.NewPlaces(sugar, mapsClient)

	result := places.GetPlaces("pizza", 1000, maps.LatLng{
		Lat: 33.911701,
		Lng: -84.357011,
	})

	pretty.Println(result)
}
