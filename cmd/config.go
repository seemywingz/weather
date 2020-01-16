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
	"github.com/spf13/cobra"
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
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Custom Defaults",
	Long: `
	
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		configure()
	},
}

func configure() {
	confirmConfigDefaults()
}

func initConfig() {

	homeDir, err := homedir.Dir()
	if err != nil {
		LoE("Unable to Find Home Directory, Not Using Config File", err)
		return
	}

	configName = "config"
	configDir = filepath.Join(homeDir, ".weather")
	configFile = filepath.Join(configDir, configName)

	if _, err := os.Stat(configFile); err == nil {
		readConfig()
		if location != "" {
			config.Location = location
		}
		if units != "" {
			config.Units = units
		}
	} else if os.IsNotExist(err) {
		fmt.Println("It appears you don't have a config file")
		confirmConfigDefaults()
	} else {
		LoE("Error Accessing Config Directory...", err)
	}

}

func readConfig() {
	file, _ := ioutil.ReadFile(configFile)
	_ = json.Unmarshal([]byte(file), &config)
}

func confirmConfigDefaults() {
	yes := Confirm("Want to use your current parameters?")
	if yes {
		createConfig()
	} else {
		fmt.Println("Skipping Config Creation")
	}
}

func createConfig() {
	os.MkdirAll(configDir, 0744)
	fmt.Println("")
	fmt.Println("   Updating Weather Config:")
	fmt.Println("      File:", configFile)
	fmt.Println("     Units:", units)
	fmt.Println("  Location:", location)
	jsonData, _ := json.MarshalIndent(Config{
		units,
		location,
	}, "", "")
	_ = ioutil.WriteFile(configFile, jsonData, 0644)
}
