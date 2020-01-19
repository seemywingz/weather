package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var verbose bool
var location string
var units string
var unitFormat UnitMeasures
var validUnits = map[string]string{
	"auto": "Determin Units Based on Location",
	"ca":   "same as si, except uses kilometers per hour",
	"uk":   "same as si, except uses kilometers, and miles per hour",
	"uk2":  "same as si, except uses miles, and miles per hour",
	"us":   "Imperial units",
	"si":   "International System of Units",
}
var unitDescription = mapToString(validUnits)

func gatherData() (WeatherResponse, GeoLocationData) {

	var locationData GeoLocationData

	if config.Location == "" {
		locationData = getLocationDataFromIP()
	} else {
		locationData = geoLocate(location)
	}

	weather := getWeatherData(locationData.Latitude, locationData.Longitude)
	unitFormat = UnitFormats[weather.Flags.Units]

	fmt.Println()
	fmt.Printf("      Location: %v, %v, %v\n", locationData.City, locationData.RegionCode, locationData.CountryCode)
	return weather, locationData
}

func display(weather WeatherData) {
	icon := Icons[weather.Icon]

	fmt.Printf("          Time: %v\n", epochFormat(weather.Time))
	fmt.Printf("       Weather: %v  %v %v\n", icon, weather.Summary, icon)
	fmt.Printf("          Temp: %v%v\n", weather.Temperature, unitFormat.Degrees)
	fmt.Printf("    Feels Like: %v%v\n", weather.ApparentTemperature, unitFormat.Degrees)
	if weather.PrecipProbability*100 > 1 {
		fmt.Printf("     Chance of: %v %.2f%%\n", weather.PrecipType, weather.PrecipProbability*100)
		fmt.Printf(" Precipitation: %v %v\n", weather.PrecipIntensity, unitFormat.Precipitation)
	}
	fmt.Printf("      Humidity: %.2f%%\n", weather.Humidity*100)
	fmt.Printf("   Cloud Cover: %.2f%%\n", weather.CloudCover*100)
	fmt.Printf("    Wind Speed: %v %v %v\n", weather.WindSpeed, unitFormat.Speed, getBearings(weather.WindBearing))
	if weather.NearestStormDistance > 0 {
		fmt.Printf(" Nearest Storm: %v %v %v\n", weather.NearestStormDistance, unitFormat.Length, getBearings(weather.NearestStormBearing))
	}
	if verbose {
		fmt.Printf("     Wind Gust: %v %v\n", weather.WindGust, unitFormat.Speed)
		fmt.Printf("     Dew Point: %v%v\n", weather.DewPoint, unitFormat.Degrees)
		fmt.Printf("      Pressure: %v hPa\n", weather.Pressure)
		fmt.Printf("         Ozone: %v DU\n", weather.Ozone)
		fmt.Printf("    Visibility: %v %v\n", weather.Visibility, unitFormat.Length)
		fmt.Printf("      UV Index: %v\n", weather.UvIndex)
	}
}

func displayDaily(weather DailyWeatherData) {
	icon := Icons[weather.Icon]

	fmt.Printf("          Date: %v\n", epochFormatDate(weather.Time))
	fmt.Printf("       Weather: %v  %v %v\n", icon, weather.Summary, icon)
	fmt.Printf("       Sunrise: %v\n", epochFormatTime(weather.SunriseTime))
	fmt.Printf("        Sunset: %v\n", epochFormatTime(weather.SunsetTime))
	fmt.Printf("    Moon Phase: %v\n", getMoonPhase(weather.MoonPhase))
	fmt.Printf("          High: %v%v\n", weather.TemperatureHigh, unitFormat.Degrees)
	fmt.Printf("           Low: %v%v\n", weather.TemperatureLow, unitFormat.Degrees)
	if weather.PrecipProbability*100 > 1 {
		fmt.Printf("     Chance of: %v %.2f%%\n", weather.PrecipType, weather.PrecipProbability*100)
		fmt.Printf(" Precipitation: %v %v\n", weather.PrecipIntensity, unitFormat.Precipitation)
	}
	fmt.Printf("      Humidity: %.2f%%\n", weather.Humidity*100)
	fmt.Printf("   Cloud Cover: %.2f%%\n", weather.CloudCover*100)
	fmt.Printf("    Wind Speed: %v %v %v\n", weather.WindSpeed, unitFormat.Speed, getBearings(weather.WindBearing))
	if verbose {
		fmt.Printf("     Wind Gust: %v %v\n", weather.WindGust, unitFormat.Speed)
		fmt.Printf("     Dew Point: %v%v\n", weather.DewPoint, unitFormat.Degrees)
		fmt.Printf("      Pressure: %v hPa\n", weather.Pressure)
		fmt.Printf("         Ozone: %v DU\n", weather.Ozone)
		fmt.Printf("    Visibility: %v %v\n", weather.Visibility, unitFormat.Length)
		fmt.Printf("      UV Index: %v\n", weather.UvIndex)
	}

}

func displayAlerts(alerts []WeatherAlert) {
	for _, alert := range alerts {
		fmt.Printf("\n      ⚠️  %v ⚠️\n %v\n", alert.Title, alert.Description)
	}
}

func getWeatherData(lat, long float32) WeatherResponse {

	url := fmt.Sprintf("https://api.darksky.net/forecast/b0e78d287f75fb03eba6022344d3b944/%v,%v?units=%v", lat, long, config.Units)
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	resData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	weatherResponse := WeatherResponse{}
	json.Unmarshal(resData, &weatherResponse)

	return weatherResponse

}

func getBearings(degrees float64) string {
	index := int(math.Mod((degrees+11.25)/22.5, 16))
	return Directions[index]
}

func getMoonPhase(phase float64) string {
	var icon string

	switch {
	case phase == 0:
		icon = "🌑"
	case phase > 0 && phase < 0.25:
		icon = "🌒"
	case phase == 0.25:
		icon = "🌓"
	case phase > 0.25 && phase < 0.5:
		icon = "🌔"
	case phase == 0.5:
		icon = "🌕"
	case phase >= 0.5 && phase < 0.75:
		icon = "🌖"
	case phase == 0.75:
		icon = "🌗"
	case phase > 0.75:
		icon = "🌘"
	}

	return icon
}

func getLocationDataFromIP() GeoLocationData {

	url := "https://telize.j3ss.co/geoip"
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	responseData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	locationData := GeoLocationData{}
	json.Unmarshal(responseData, &locationData)

	return locationData

}

func geoLocate(location string) GeoLocationData {

	url := "https://geocode.jessfraz.com/geocode"

	reqBody, _ := json.Marshal(map[string]string{
		"Location": location,
	})

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	EoE("Error Getting GeoLocation Response", err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	locationData := GeoLocationData{}
	json.Unmarshal(body, &locationData)

	return locationData

}

func epochFormat(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2, 3:04pm MST")
}

func epochFormatDate(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 (Monday)")
}

func epochFormatTime(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("3:04pm MST")
}

func epochFormatHour(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	s := epochTime.Format("3pm")
	s = s[:len(s)-1]
	if len(s) == 2 {
		s += " "
	}
	return s
}

// LoE : if error, log to console
func LoE(msg string, err error) {
	if err != nil {
		log.Printf("\n❌  %s\n   %v\n", msg, err)
	}
}

// EoE : exit with error code 1 and print, if err is not nil
func EoE(msg string, err error) {
	if err != nil {
		fmt.Printf("\n❌  %s\n   %v\n", msg, err)
		os.Exit(1)
		panic(err)
	}
}

// Confirm : return confirmation based on user input
func Confirm(q string) bool {
	print(q + " (Y/n) ")
	a := GetInput()
	var res bool
	switch strings.ToLower(a) {
	case "":
		fallthrough
	case "y":
		fallthrough
	case "yes":
		res = true
	case "n":
		fallthrough
	case "no":
		res = false
	default:
		return Confirm(q)
	}
	return res
}

// GetInput : return string of user input
func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// SelectFromArray : select an element in the provided array
func SelectFromArray(a []string) string {
	fmt.Println("Choices:")
	for i := range a {
		fmt.Println("[", i, "]: "+a[i])
	}
	fmt.Println("Enter Number of Selection: ")
	sel, err := strconv.Atoi(GetInput())
	EoE("Error Getting Integer Input from User", err)
	if sel <= len(a)-1 {
		return a[sel]
	}
	return SelectFromArray(a)
}

// SelectFromMap : select an element in the provided map
func SelectFromMap(m map[string]string) string {
	fmt.Println(mapToString(m))
	sel := GetInput()
	if _, found := m[sel]; found {
		return sel
	}
	fmt.Printf("%v is an Invalid Selection", sel)
	return SelectFromMap(m)
}

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "  %s: %s\n", key, value)
	}
	return b.String()
}
