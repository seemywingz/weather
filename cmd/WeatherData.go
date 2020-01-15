package cmd

// WeatherResponse : basic response from Dark Sky API
type WeatherResponse struct {
	Currently WeatherData `json:"currently"`
	Minutely  struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                 int64   `json:"time"`
			PrecipIntensity      float64 `json:"precipIntensity"`
			PrecipIntensityError float64 `json:"precipIntensityError"`
			PrecipProbability    float64 `json:"precipProbability"`
			PrecipType           string  `json:"precipType"`
		} `json:"minutely"`
	} `json:"minutely"`
	Hourly struct {
		Summary string        `json:"summary"`
		Icon    string        `json:"icon"`
		Data    []WeatherData `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string        `json:"summary"`
		Icon    string        `json:"icon"`
		Data    []WeatherData `json:"data"`
	} `json:"daily"`
	Alerts []WeatherAlert `json:"alerts"`
	Flags  struct {
		Units          string `json:"units"`
		Unavalable     string `json:"darksky-unavailable"`
		NearestStation int64  `json:"nearest-station"`
		Sources        string `json:"sources"`
	} `json:"flags"`
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

// WeatherAlert : format for dark sky weather alert
type WeatherAlert struct {
	Title       string `json:"title"`
	Time        int    `json:"time"`
	Expires     int    `json:"expires"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}
