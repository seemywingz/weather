# weather
Golang CLI for getting Weather Information provided by DarkSky API

## Installation
```shell
go install github.com/seemywingz/weather
```

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

### Sample Output
```shell
▶ weather

      Location: Pleasant Valley, NY, US
          Time: January 24, 10:22pm EST
       Weather: 🌙  Clear 🌙
          Temp: -0.33°C
    Feels Like: -2.44°C
      Humidity: 81.00%
   Cloud Cover: 19.00%
    Wind Speed: 3.8 mph ENE
 Nearest Storm: 107 miles NNW

▶ weather --units us

      Location: Pleasant Valley, NY, US
          Time: January 24, 10:23pm EST
       Weather: 🌙  Clear 🌙
          Temp: 31.38°F
    Feels Like: 27.59°F
      Humidity: 81.00%
   Cloud Cover: 19.00%
    Wind Speed: 3.8 mph ENE
 Nearest Storm: 107 miles NNW
```
```shell
▶ weather -l Seattle

      Location: Seattle, WA, US
          Time: January 24, 10:24pm EST
       Weather: 🌧  Light Rain 🌧
          Temp: 46.89°F
    Feels Like: 43.9°F
     Chance of: rain 100.00%
 Precipitation: 0.03 in/hr
      Humidity: 93.00%
   Cloud Cover: 47.00%
    Wind Speed: 6.12 mph S

      ⚠️  Special Weather Statement ⚠️
 ...INCREASED THREAT OF LANDSLIDES FROM RECENT HEAVY RAIN IN WESTERN WASHINGTON CONTINUES BUT WILL DIMINISH OVER THE NEXT COUPLE OF DAYS... The increased threat of landslides due to recent precipitation will slowly diminish over the next couple of days down to average levels for this time of year. At least a few landslides were reported earlier and an isolated landslide can not be ruled out. While additional rainfall is expected, it will not enough to keep the threat of landslides elevated or be a significant trigger for new landslides. Therefore the alert for enhance landslide threat will be allowed to expire this evening. For more information about current conditions, visit www.weather.gov/seattle, select Hydrology, and then scroll down for the links to the landslide information pages. For more information on landslides, visit the website for the Washington State Department of Natural Resources landslide geologic hazards at: http://bit.ly/2mtA3wn

```