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
	"os"

	"github.com/spf13/cobra"
)

var verbose,
	imperialUnits bool

var units string = "ca"
var zip string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather CLI Tool",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		now()

		fmt.Println("Public IP", getPubIP())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(nowCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().BoolVarP(&imperialUnits, "imperial", "i", false, "Use Imperial Units")
	rootCmd.PersistentFlags().BoolVarP(&imperialUnits, "fahrenheit", "f", false, "Use Imperial Units")
	rootCmd.PersistentFlags().StringVarP(&zip, "zip", "z", "", "Zipcode to gather weather info for")
}

func display(weather WeatherData, location Location) {
	fmt.Println()
	fmt.Printf("    Location: %v, %v, %v\n", location.City, location.RegionCode, location.PostalCode)
	fmt.Println("     Weather:", weather.Summary)
	fmt.Printf("        Temp: %v°\n", weather.Temperature)
	fmt.Printf("  Feels Like: %v°\n", weather.ApparentTemperature)
}

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

func getLocationData(zip string) LocationData {

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
