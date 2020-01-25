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

	"github.com/seemywingz/gotoolbox/darksky"
	"github.com/spf13/cobra"
)

var numHours = 12

var hourlyCmd = &cobra.Command{
	Use:   "hourly",
	Short: "Display Hourly Weather",
	Long: `
	
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		// if len(args) < 1 {
		// 	return errors.New("Must provide one argument")
		// }
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		hourly()
	},
}

func init() {
	hourlyCmd.Flags().IntVarP(&numHours, "num-hours", "n", 12, "Number of Hours to Display, Max: 48")
}

func hourly() {

	if numHours > 48 {
		EoE("Max Hours is 48", fmt.Errorf("Error: %v is too many hours", numHours))
	}

	weather, _ := gatherData()
	icon := darksky.Icons[weather.Hourly.Icon]
	fmt.Printf("       Weather: %v  %v %v\n\n", icon, weather.Hourly.Summary, icon)

	for i := 0; i < numHours; i++ {
		displayCurrent(weather.Hourly.Data[i])
		fmt.Println()
	}
	displayAlerts(weather.Alerts)
}
