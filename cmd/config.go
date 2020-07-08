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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/seemywingz/gotoolbox/gtb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configName,
	configDir,
	configFile string
	config Config
)

// Config : User Defined Dfaults
type Config struct {
	Units    string `json:"units"`
	Location string `json:"location"`
	APIKey   string `json:"apiKey"`
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Custom Defaults",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		configure()
	},
}

func init() {
}

func initConfig() {
	homeDir, err := homedir.Dir()
	if err != nil {
		gtb.LoE("Unable to Find Home Directory, Not Using Config File", err)
		return
	}

	configName = "config"
	configDir = filepath.Join(homeDir, ".weather")
	configFile = filepath.Join(configDir, configName)

	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)

	viper.SetDefault("units", "auto")
	viper.SetDefault("location", "auto")
	viper.SetDefault("apiKey", "89dc0c059f63f8f283768862617b10f6")

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("It appears you don't have a config file... Using Current Settings")
			configure()
			saveConfig()
		}
	}

	parseConfig()

}

func parseConfig() {

	config = Config{
		viper.GetString("units"),
		viper.GetString("location"),
		viper.GetString("apiKey"),
	}
}

func printCurrentConfig() {
	parseConfig()
	fmt.Println("")
	// fmt.Println("   Weather Config:")
	// fmt.Println("      File:", configFile)
	fmt.Println("     Units:", config.Units)
	fmt.Println("  Location:", config.Location)
	// if config.APIKey != "" {
	// 	fmt.Println("   API Key:", config.APIKey)
	// }
}

func saveConfig() {
	fmt.Println("\nSaving Config As:", configFile)

	os.MkdirAll(configDir, 0744)
	_ = ioutil.WriteFile(configFile, make([]byte, 0, 0), 0644)

	gtb.EoE("Error Writing Config", viper.WriteConfig())

	printCurrentConfig()
}

func configAPIKey() {
	if config.APIKey == "" {
		fmt.Println("Enter Your Dark Sky API Key")
		fmt.Println("Don't Have One?")
		fmt.Println("Get One for Free At: https://darksky.net/dev")
		fmt.Printf(":")
		config.APIKey = gtb.GetInput()
	} else if gtb.Confirm("Want to Replace Your Dark Sky API Key?") {
		config.APIKey = ""
		configAPIKey()
	}
}

func confirmSave() {
	printCurrentConfig()
	if gtb.Confirm("\nWant to save these parameters?") {
		saveConfig()
		os.Exit(0)
	}
}

func configure() {

	fmt.Println("\nCurrent Parameters:")
	confirmSave()

	fmt.Println("\nOkay, let's make some choices:")
	viper.Set("units", gtb.SelectFromMap(darksky.ValidUnits))

	fmt.Println("\nEnter Your Default Location")
	fmt.Println("Examples:\n  12569, Beaverton, \"1600 Pennsylvania Ave\"")
	fmt.Println("  enter \"auto\" and your location will be determined from you public IP address!")
	fmt.Printf(":")
	viper.Set("location", gtb.GetInput())

	fmt.Println("\nOkay, Great!")
	configAPIKey()
	confirmSave()
}
