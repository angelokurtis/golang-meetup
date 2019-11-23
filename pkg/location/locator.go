package location

type Locator interface {
	WhereAmI() (*Location, error)
}
