package cmd

// GeoLocationData : Data returned from opendatasoft
type GeoLocationData struct {
	City        string  `json:"city"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"region_code"`
	PostalCode  string  `json:"postal_code"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
	IP          string  `json:"ip"`
}

// LocationResponse : Data returned from opendatasoft
// type LocationResponse struct {
// 	Nhits int `json:"nhits"`

// 	Parameters struct {
// 		Dataset  string `json:"dataset"`
// 		TimeZone string `json:"timesone"`
// 		Q        string `json:"q"`
// 		Rows     int    `json:"rows"`
// 		Format   string `json:"format"`
// 	} `json:"parameters"`

// 	Records []struct {
// 		Fields LocationData `json:"fields"`
// 	} `json:"records"`
// }

// // LocationData : Data returned from opendatasoft
// type LocationData struct {
// 	City      string  `json:"city"`
// 	Latitude  float32 `json:"latitude"`
// 	Longitude float32 `json:"longitude"`
// 	State     string  `json:"state"`
// 	Zip       string  `json:"zip"`
// 	Timezone  int     `json:"timezone"`
// }
