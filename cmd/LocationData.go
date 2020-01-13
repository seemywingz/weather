package cmd

// LocationData : Data returned from opendatasoft
type LocationData struct {
	Nhits int `json:"nhits"`

	Parameters struct {
		Dataset  string `json:"dataset"`
		TimeZone string `json:"timesone"`
		Q        string `json:"q"`
		Rows     int    `json:"rows"`
		Format   string `json:"format"`
	} `json:"parameters"`

	Records []struct {
		Fields struct {
			City      string  `json:"city"`
			Latitude  float32 `json:"latitude"`
			Longitude float32 `json:"longitude"`
			State     string  `json:"state"`
			Zip       string  `json:"zip"`
			Timezone  int     `json:"timezone"`
		} `json:"fields"`
	} `json:"records"`
}