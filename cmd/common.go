package cmd

import (
	"fmt"

	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/seemywingz/gotoolbox/epoch"
	"github.com/seemywingz/gotoolbox/geolocation"
	"github.com/seemywingz/gotoolbox/gtb"
)

// Flag Vars
var verbose bool
var location string
var units string

// Format
var unitFormat darksky.UnitMeasures
var unitDescription = gtb.MapToString(darksky.ValidUnits)

var validArgs = map[string]string{
	"daily":  "",
	"today":  "",
	"hourly": "",
	"now":    "",
}

func gatherData() (darksky.Data, geolocation.Data) {

	var locationData geolocation.Data
	var err error

	if config.Location == "" || config.Location == "auto" {
		locationData, _ = geolocation.FromIP()
	} else {
		locationData, err = geolocation.Locate(config.Location)
	}

	weather, err := darksky.GetData(locationData.Latitude, locationData.Longitude, config.APIKey, config.Units)
	gtb.EoE("Error Getting DarkSky Data", err)
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
		descs := gtb.SplitMulti(alert.Description, ".*")
		fmt.Printf("\n      ⚠️  %v ⚠️\n", alert.Title)
		for _, desc := range descs {
			fmt.Println(desc)
		}
	}
}
