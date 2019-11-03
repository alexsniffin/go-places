package models

//Configuration todo
type Configuration struct {
	MapsAPI MapsAPI
}

// MapsAPI maps api config
type MapsAPI struct {
	APIKey string
}
