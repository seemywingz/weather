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

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configName,
	configDir,
	configFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather CLI Tool",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		now()
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

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().StringVarP(&units, "units", "u", "auto", "System of units (e.g. auto, us, si, ca, uk2)")
	rootCmd.PersistentFlags().StringVarP(&locationArg, "location", "l", "", "Location to Report Weather Conditions Of (e.g 12569, Beaverton, \"1600 Pennsylvania Ave\")")

	viper.BindPFlag("units", rootCmd.PersistentFlags().Lookup("units"))
	viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location"))
	viper.SetDefault("units", "auto")
	viper.SetDefault("location", "")

	rootCmd.AddCommand(nowCmd)
	rootCmd.AddCommand(todayCmd)
	rootCmd.AddCommand(dailyCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(hourlyCmd)
}

func initConfig() {

	// Find home directory.
	homeDir, err := homedir.Dir()
	if err != nil {
		LoE("Error Getting Home Directory", err)
		return
	}

	configName = "config"
	configDir = homeDir + "/.weather"
	configFile = configDir + "/" + configName
	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		LoE("", err)
		if os.IsNotExist(err) {
			if _, err := os.Stat(configDir); os.IsNotExist(err) { // config dir not found
				os.MkdirAll(configDir, 0777) // create config dir
				config, _ := json.MarshalIndent(Config{}, "", "")
				err = ioutil.WriteFile(configFile, config, 0644)
			}
		}
	}

}
