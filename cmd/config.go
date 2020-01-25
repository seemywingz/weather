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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/seemywingz/gotoolbox/gtb"
	"github.com/spf13/cobra"
)

var (
	configName,
	configDir,
	configFile string
	config Config
)

var auto bool

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

	configCmd.Flags().BoolVarP(&auto, "auto", "a", false, "set all config values to auto")
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

	if _, err := os.Stat(configFile); err == nil {
		readConfig()
		configOverride()
	} else if os.IsNotExist(err) {
		fmt.Println("It appears you don't have a saved config file")
		configOverride()
		configure()
	} else {
		gtb.LoE("Error Accessing Config Directory...", err)
	}

}

func readConfig() {
	file, _ := ioutil.ReadFile(configFile)
	_ = json.Unmarshal([]byte(file), &config)
}

func configOverride() {
	if location != "" {
		config.Location = location
	}
	if units != "" {
		config.Units = units
	}
	if config.Units == "" {
		config.Units = "auto"
	}
	if config.Location == "" {
		config.Location = "auto"
	}
}

func printCurrentConfig() {
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
	jsonData, _ := json.MarshalIndent(config, "", "")
	_ = ioutil.WriteFile(configFile, jsonData, 0644)
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
		configAPIKey()
		saveConfig()
		os.Exit(0)
	}
}

func configure() {

	// if auto {
	// 	config.Location = "auto"
	// 	config.Units = "auto"
	// 	configAPIKey()
	// 	return
	// }

	fmt.Println("\nCurrent Parameters:")
	confirmSave()

	fmt.Println("\nOkay, let's make some choices:")
	config.Units = gtb.SelectFromMap(darksky.ValidUnits)
	fmt.Println("\nEnter Your Default Location")
	fmt.Println("Examples:\n  12569, Beaverton, \"1600 Pennsylvania Ave\"")
	fmt.Println("  enter \"auto\" and your location will be determined from you public IP address!")
	fmt.Printf(":")
	config.Location = gtb.GetInput()
	fmt.Println("\nOkay, Great!")
	configAPIKey()
	confirmSave()
}
