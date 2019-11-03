package places

import (
	"context"

	"github.com/alexsniffin/go-places/internal/api/clients"

	"go.uber.org/zap"
	"googlemaps.github.io/maps"
)

// Places struct
type Places struct {
	log        *zap.SugaredLogger
	mapsClient clients.Client
}

// NewPlaces creates a places struct
func NewPlaces(log *zap.SugaredLogger, mapsClient clients.Client) Places {
	return Places{
		log:        log,
		mapsClient: mapsClient,
	}
}

// GetPlaces get places in an area
func (p *Places) GetPlaces(input string, radius uint, location maps.LatLng) maps.AutocompleteResponse {
	r := &maps.PlaceAutocompleteRequest{
		Input:    input,
		Radius:   radius,
		Location: &location,
	}

	res, err := p.mapsClient.C.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		p.log.Fatalf("fatal error: %s", err)
	}

	return res
}
