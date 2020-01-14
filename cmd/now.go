// Copyright © 2017 Kevin Jayne <seemywings@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Display Current Weather",
	Long: `
	
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		// if len(args) < 1 {
		// 	return errors.New("Must provide one argument")
		// }
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		now()
	},
}

func init() {
}

func display(weather WeatherData, location LocationData) {
	fmt.Println("\nCurrent Weather in Your Location:", weather.Summary)
	fmt.Println("        City:", location.City)
	fmt.Println("         Zip:", location.Zip)
	fmt.Println("        Temp:", weather.Temperature)
	fmt.Println("  Feels Like:", weather.ApparentTemperature)
}

func now() {
	location := getLocationData(zip)
	weather := getWeatherData(location.Latitude, location.Longitude)

	display(weather.Currently, location)
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

func getLocationData(zip string) LocationData {

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
