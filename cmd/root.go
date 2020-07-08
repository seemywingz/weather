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
	"os"

	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather CLI Tool",
	Long: `
	
	`,
	ValidArgs: []string{"today"},
	Args: func(cmd *cobra.Command, args []string) error {
		if darksky.ValidUnits[config.Units] == "" {
			return fmt.Errorf("❌  Invalid Unit Type: %s", units)
		}

		for _, arg := range args {
			if _, ok := validArgs[arg]; !ok {
				return fmt.Errorf("❌  Invalid Arg: %s", arg)
			}
		}

		return nil
	},
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

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output, print less common weather data")

	rootCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Location to Report Weather Conditions Of \nExamples: 12569, Beaverton, \"1600 Pennsylvania Ave\"\nIf empty, your location will be determined from you public IP address")
	viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location"))

	rootCmd.PersistentFlags().StringVarP(&units, "units", "u", "", "System of Units:\n"+unitDescription)
	viper.BindPFlag("units", rootCmd.PersistentFlags().Lookup("units"))

	rootCmd.AddCommand(nowCmd)
	rootCmd.AddCommand(todayCmd)
	rootCmd.AddCommand(dailyCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(hourlyCmd)
}
