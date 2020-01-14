package cmd

// WeatherResponse : basic response from Dark Sky API
type WeatherResponse struct {
	Currently WeatherData `json:"currently"`
	Minutely  struct {
		Summary string `json:"summary"`
	} `json:"minutely"`
}

// WeatherData : Struct containg json data from DarkSky API
type WeatherData struct {
	Time                 int     `json:"time"`
	Summary              string  `json:"summary"`
	Icon                 string  `json:"icon"`
	NearestStormDistance int     `json:"nearestStormDistance"`
	PrecipIntensity      float32 `json:"precipIntensity"`
	PrecipIntensityError float32 `json:"precipIntensityError"`
	PrecipProbability    float32 `json:"precipProbability"`
	PrecipType           string  `json:"precipType"`
	Temperature          float32 `json:"temperature"`
	ApparentTemperature  float32 `json:"apparentTemperature"`
	DewPoint             float32 `json:"dewPoint"`
	Humidity             float32 `json:"humidity"`
	Pressure             float32 `json:"pressure"`
	WindSpeed            float32 `json:"windSpeed"`
	WindGust             float32 `json:"windGust"`
	WindBearing          float32 `json:"windBearing"`
	CloudCover           float32 `json:"cloudCover"`
	UvIndex              float32 `json:"uvIndex"`
	Visibility           float32 `json:"visibility"`
	Ozone                float32 `json:"ozone"`
}
