package forecast

type Checker interface {
	CheckByGeoCoordinates(latitude float64, longitude float64) (*Forecast, error)
}

type Forecast struct {
	Dt       int     `json:"dt"`
	Humidity float64 `json:"humidity"`
	Pressure float64 `json:"pressure"`
	Temp     struct {
		Average    float64 `json:"average"`
		AverageMax float64 `json:"average_max"`
		AverageMin float64 `json:"average_min"`
		RecordMax  float64 `json:"record_max"`
		RecordMin  float64 `json:"record_min"`
	} `json:"temp"`
	WindSpeed float64 `json:"wind_speed"`
}
