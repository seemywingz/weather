package cmd

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/user"
	"time"
)

var verbose bool
var locationArg string
var units string
var unitFormat UnitMeasures

func gatherData() (WeatherResponse, GeoLocationData) {

	var location GeoLocationData

	if locationArg == "" {
		location = getLocationDataFromIP()
	} else {
		location = geoLocate(locationArg)
	}

	weather := getWeatherData(location.Latitude, location.Longitude)
	unitFormat = UnitFormats[weather.Flags.Units]

	fmt.Println()
	fmt.Printf("      Location: %v, %v, %v\n", location.City, location.RegionCode, location.CountryCode)
	return weather, location
}

func getWeatherData(lat, long float32) WeatherResponse {

	url := fmt.Sprintf("https://api.darksky.net/forecast/b0e78d287f75fb03eba6022344d3b944/%v,%v?units=%v", lat, long, units)
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	resData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	weatherResponse := WeatherResponse{}
	json.Unmarshal(resData, &weatherResponse)

	return weatherResponse

}

func display(weather WeatherData, alerts []WeatherAlert) {
	icon := Icons[weather.Icon]

	fmt.Printf("          Time: %v\n", epochFormat(weather.Time))
	fmt.Printf("       Weather: %v  %v %v\n", icon, weather.Summary, icon)
	fmt.Printf("          Temp: %v%v\n", weather.Temperature, unitFormat.Degrees)
	fmt.Printf("    Feels Like: %v%v\n", weather.ApparentTemperature, unitFormat.Degrees)
	fmt.Printf("      Humidity: %.2f%%\n", weather.Humidity*100)
	fmt.Printf("   Cloud Cover: %.2f%%\n", weather.CloudCover*100)
	fmt.Printf("          Wind: %v%v %v\n", weather.WindSpeed, unitFormat.Speed, getBearings(weather.WindBearing))
	if weather.PrecipProbability*100 > 1 || verbose {
		fmt.Printf("   Precip Prob: %.2f%% chance of %v\n", weather.PrecipProbability*100, weather.PrecipType)
		fmt.Printf("        Precip: %v%v\n", weather.PrecipIntensity, unitFormat.Precipitation)
	}

	for _, alert := range alerts {
		fmt.Printf("\n      ⚠️  %v ⚠️\n %v\n", alert.Title, alert.Description)
	}
}

func getBearings(degrees float64) string {
	index := int(math.Mod((degrees+11.25)/22.5, 16))
	return Directions[index]
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

// LoE : log with error code 1 and print if err is notnull
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

// WriteVar : write gob to local folesystem
func WriteVar(file string, data interface{}) error {
	gobFile, err := os.Create(file)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(gobFile)
	encoder.Encode(data)
	gobFile.Close()
	return nil
}

// ReadVar : read gob from loacal filesystem
func ReadVar(file string, object interface{}) error {
	gobFile, err := os.Open(file)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(gobFile)
	err = decoder.Decode(object)
	gobFile.Close()
	return nil
}

// GetHomeDir : returns a full path to user's home dorectory
func GetHomeDir() string {
	usr, err := user.Current()
	EoE("Failed to get Current User", err)
	if usr.HomeDir != "" {
		return usr.HomeDir
	}
	return os.Getenv("HOME")
}
