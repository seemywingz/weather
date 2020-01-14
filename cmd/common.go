package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

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
