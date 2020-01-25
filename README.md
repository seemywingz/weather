# weather
Golang CLI for getting Weather Information provided by DarkSky API

## Usage
```shell
▶ weather --help


Usage:
  weather [flags]
  weather [command]

Available Commands:
  config      Configure Custom Defaults
  daily       Display 7 Day Weather Forecast
  help        Help about any command
  hourly      Display Hourly Weather
  now         Display Current Weather
  today       Display Today's Weather

Flags:
  -h, --help              help for weather
  -l, --location string   Location to Report Weather Conditions Of 
                          Examples: 12569, Beaverton, "1600 Pennsylvania Ave"
                          If empty, your location will be determined from you public IP address
  -u, --units string      System of Units:
                            auto: Determin Units Based on Location
                            ca: same as si, except uses kilometers per hour
                            uk: same as si, except uses kilometers, and miles per hour
                            uk2: same as si, except uses miles, and miles per hour
                            us: Imperial units
                            si: International System of Units
                          
  -v, --verbose           Verbose Output, print less common weather data

Use "weather [command] --help" for more information about a command.

```