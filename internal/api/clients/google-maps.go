package clients

import (
	"go.uber.org/zap"
	"googlemaps.github.io/maps"
)

// Client struct
type Client struct {
	log *zap.SugaredLogger
	C   *maps.Client
}

// NewClient create a client struct
func NewClient(log *zap.SugaredLogger, apiKey string) Client {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return Client{
		log: log,
		C:   c,
	}
}
