package location

type Finder interface {
	WhereAmI() (*Location, error)
}

type Location struct {
	CountryCode string  `json:"country_code"`
	RegionCode  string  `json:"region_code"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
