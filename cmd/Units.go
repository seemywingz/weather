package cmd

// UnitMeasures are the location specific terms for weather data.
type UnitMeasures struct {
	Degrees       string
	Speed         string
	Length        string
	Precipitation string
}

// UnitFormats describe each regions UnitMeasures.
var UnitFormats = map[string]UnitMeasures{
	"us": {
		Degrees:       "°F",
		Speed:         "mph",
		Length:        "miles",
		Precipitation: "in/hr",
	},
	"si": {
		Degrees:       "°C",
		Speed:         "m/s",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	"ca": {
		Degrees:       "°C",
		Speed:         "km/h",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	// deprecated, use "uk2" in stead
	"uk": {
		Degrees:       "°C",
		Speed:         "mph",
		Length:        "kilometers",
		Precipitation: "mm/h",
	},
	"uk2": {
		Degrees:       "°C",
		Speed:         "mph",
		Length:        "miles",
		Precipitation: "mm/h",
	},
}
