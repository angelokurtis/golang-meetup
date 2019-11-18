package forecast

type Checker interface {
	CheckByGeoCoordinates(latitude float64, longitude float64) (*Weather, error)
}

type Weather struct {
	Description string
	Temp        float64
}
