package weather

type Checker interface {
	CheckByCurrentLocation() (*Weather, error)
	CheckByCoord(lat float64, long float64) (*Weather, error)
}
