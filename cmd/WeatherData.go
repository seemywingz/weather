package cmd

// WeatherResponse : basic response from Dark Sky API
type WeatherResponse struct {
	Currently WeatherData `json:"currently"`
	Minutely  struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
	} `json:"minutely"`
	Alerts []WeatherAlert
}

// WeatherAlert : format for dark sky weather alert
type WeatherAlert struct {
	Title       string `json:""`
	Time        int    `json:"time"`
	Expires     int    `json:"expires"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}

// WeatherData : Struct containg json data from DarkSky API
type WeatherData struct {
	Time                 int64   `json:"time"`
	Summary              string  `json:"summary"`
	Icon                 string  `json:"icon"`
	NearestStormDistance int     `json:"nearestStormDistance"`
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipIntensityError float64 `json:"precipIntensityError"`
	PrecipProbability    float64 `json:"precipProbability"`
	PrecipType           string  `json:"precipType"`
	Temperature          float64 `json:"temperature"`
	ApparentTemperature  float64 `json:"apparentTemperature"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	Pressure             float64 `json:"pressure"`
	WindSpeed            float64 `json:"windSpeed"`
	WindGust             float64 `json:"windGust"`
	WindBearing          float64 `json:"windBearing"`
	CloudCover           float64 `json:"cloudCover"`
	UvIndex              float64 `json:"uvIndex"`
	Visibility           float64 `json:"visibility"`
	Ozone                float64 `json:"ozone"`
}
