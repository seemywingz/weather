package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/seemywingz/gotoolbox/epoch"
	"github.com/seemywingz/gotoolbox/geolocation"
)

// Flag Vars
var verbose bool
var location string
var units string

// Format
var unitFormat darksky.UnitMeasures
var unitDescription = mapToString(validUnits)

// Validation
var validUnits = map[string]string{
	"auto": "Determin Units Based on Location",
	"ca":   "same as si, except uses kilometers per hour",
	"uk":   "same as si, except uses kilometers, and miles per hour",
	"uk2":  "same as si, except uses miles, and miles per hour",
	"us":   "Imperial units",
	"si":   "International System of Units",
}
var validArgs = map[string]string{
	"daily":  "",
	"today":  "",
	"hourly": "",
	"now":    "",
}

func gatherData() (darksky.Data, geolocation.Data) {

	var locationData geolocation.Data

	if config.Location == "" {
		locationData, _ = geolocation.FromIP()
	} else {
		locationData, _ = geolocation.Locate(location)
	}

	weather, err := darksky.GetData(locationData.Latitude, locationData.Longitude, darkSkyAPIKey, config.Units)
	EoE("Error Getting DarkSky Data", err)
	unitFormat = darksky.UnitFormats[weather.Flags.Units]

	fmt.Println()
	fmt.Printf("      Location: %v, %v, %v\n", locationData.City, locationData.RegionCode, locationData.CountryCode)
	return weather, locationData
}

func displayCurrent(weather darksky.CurrentData) {
	icon := darksky.Icons[weather.Icon]

	fmt.Printf("          Time: %v\n", epoch.Format(weather.Time))
	fmt.Printf("       Weather: %v  %v %v\n", icon, weather.Summary, icon)
	fmt.Printf("          Temp: %v%v\n", weather.Temperature, unitFormat.Degrees)
	fmt.Printf("    Feels Like: %v%v\n", weather.ApparentTemperature, unitFormat.Degrees)
	if weather.PrecipProbability*100 > 1 {
		fmt.Printf("     Chance of: %v %.2f%%\n", weather.PrecipType, weather.PrecipProbability*100)
		fmt.Printf(" Precipitation: %v %v\n", weather.PrecipIntensity, unitFormat.Precipitation)
	}
	fmt.Printf("      Humidity: %.2f%%\n", weather.Humidity*100)
	fmt.Printf("   Cloud Cover: %.2f%%\n", weather.CloudCover*100)
	fmt.Printf("    Wind Speed: %v %v %v\n", weather.WindSpeed, unitFormat.Speed, darksky.GetBearings(weather.WindBearing))
	if weather.NearestStormDistance > 0 {
		fmt.Printf(" Nearest Storm: %v %v %v\n", weather.NearestStormDistance, unitFormat.Length, darksky.GetBearings(weather.NearestStormBearing))
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

func displayDaily(weather darksky.DailyData) {
	icon := darksky.Icons[weather.Icon]

	fmt.Printf("          Date: %v\n", epoch.FormatDate(weather.Time))
	fmt.Printf("       Weather: %v  %v %v\n", icon, weather.Summary, icon)
	fmt.Printf("       Sunrise: %v\n", epoch.FormatTime(weather.SunriseTime))
	fmt.Printf("        Sunset: %v\n", epoch.FormatTime(weather.SunsetTime))
	fmt.Printf("    Moon Phase: %v\n", darksky.MoonPhaseIcon(weather.MoonPhase))
	fmt.Printf("          High: %v%v\n", weather.TemperatureHigh, unitFormat.Degrees)
	fmt.Printf("           Low: %v%v\n", weather.TemperatureLow, unitFormat.Degrees)
	if weather.PrecipProbability*100 > 1 {
		fmt.Printf("     Chance of: %v %.2f%%\n", weather.PrecipType, weather.PrecipProbability*100)
		fmt.Printf(" Precipitation: %v %v\n", weather.PrecipIntensity, unitFormat.Precipitation)
	}
	fmt.Printf("      Humidity: %.2f%%\n", weather.Humidity*100)
	fmt.Printf("   Cloud Cover: %.2f%%\n", weather.CloudCover*100)
	fmt.Printf("    Wind Speed: %v %v %v\n", weather.WindSpeed, unitFormat.Speed, darksky.GetBearings(weather.WindBearing))
	if verbose {
		fmt.Printf("     Wind Gust: %v %v\n", weather.WindGust, unitFormat.Speed)
		fmt.Printf("     Dew Point: %v%v\n", weather.DewPoint, unitFormat.Degrees)
		fmt.Printf("      Pressure: %v hPa\n", weather.Pressure)
		fmt.Printf("         Ozone: %v DU\n", weather.Ozone)
		fmt.Printf("    Visibility: %v %v\n", weather.Visibility, unitFormat.Length)
		fmt.Printf("      UV Index: %v\n", weather.UvIndex)
	}

}

func displayAlerts(alerts []darksky.Alert) {
	for _, alert := range alerts {
		fmt.Printf("\n      ⚠️  %v ⚠️\n %v\n", alert.Title, alert.Description)
	}
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
	fmt.Println("")
	fmt.Println(mapToString(m))
	sel := GetInput()
	if _, found := m[sel]; found {
		return sel
	}
	fmt.Printf("%v is an Invalid Selection\n", sel)
	return SelectFromMap(m)
}

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "  %s: %s\n", key, value)
	}
	return b.String()
}
