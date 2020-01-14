package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var verbose,
	imperialUnits bool

var units string = "si"
var zip string

func getWeatherData(lat, long float32) WeatherResponse {

	if imperialUnits {
		units = "us"
	}

	url := fmt.Sprintf("https://api.darksky.net/forecast/b0e78d287f75fb03eba6022344d3b944/%v,%v?units=%v", lat, long, units)
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	resData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	weatherResponse := WeatherResponse{}
	json.Unmarshal(resData, &weatherResponse)

	return weatherResponse

}

func display(weather WeatherData, location Location) {
	fmt.Println()
	fmt.Printf("    Location: %v, %v, %v\n", location.City, location.RegionCode, location.PostalCode)
	fmt.Println("     Weather:", weather.Summary)
	fmt.Printf("        Temp: %v°\n", weather.Temperature)
	fmt.Printf("  Feels Like: %v°\n", weather.ApparentTemperature)
}

func epochFormat(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 at 3:04pm MST")
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

// GetIP : get local ip address
func getPubIP() string {
	// we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	// https://ifconfig.co
	// https://ifconfig.me
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	EoE("Error Getting IP Address", err)
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	EoE("Error Reading IP Address", err)
	return string(ip)
}

func getLocationDataFromIP() Location {

	url := "https://telize.j3ss.co/geoip"
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	responseData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	locationData := GeoLocationData{}
	json.Unmarshal(responseData, &locationData)

	return Location{
		locationData.City,
		locationData.Region,
		locationData.RegionCode,
		locationData.PostalCode,
		locationData.Country,
		locationData.CountryCode,
		locationData.Timezone,
		locationData.Latitude,
		locationData.Longitude,
	}

}

func searchLocationData(zip string) LocationData {

	if zip == "" {
		zip = "12569"
		fmt.Println("Using Default Zipcode:", zip)
	}

	url := "https://public.opendatasoft.com/api/records/1.0/search/?dataset=us-zip-code-latitude-and-longitude&q=" + zip
	res, err := http.Get(url)
	EoE("Error Getting Location Data", err)

	responseData, err := ioutil.ReadAll(res.Body)
	EoE("Error Reading Location Data", err)

	locationResponse := LocationResponse{}
	json.Unmarshal(responseData, &locationResponse)

	if locationResponse.Nhits < 1 {
		EoE("Sorry, Could Not Find Weather Data Fror ZIP: "+zip, errors.New(""))
	}

	return locationResponse.Records[0].Fields

}
