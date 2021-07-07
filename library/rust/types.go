package rust

type RustResponse struct {
	Version string `json:"version"`
	Latest  int    `json:"latest"`
	Updates []struct {
		Timestamp  int `json:"timestamp"`
		BuildID    int `json:"buildId"`
		Forecast   int `json:"forecast"`
		MarginLow  int `json:"marginLow"`
		MarginHigh int `json:"marginHigh"`
	} `json:"updates"`
	ServerUpdates []struct {
		Timestamp int `json:"timestamp"`
		BuildID   int `json:"buildId"`
	} `json:"serverUpdates"`
	Estimate       int `json:"estimate"`
	MarginLow      int `json:"marginLow"`
	MarginHigh     int `json:"marginHigh"`
	StagingUpdates []struct {
		Timestamp int `json:"timestamp"`
		BuildID   int `json:"buildId"`
	} `json:"stagingUpdates"`
}
