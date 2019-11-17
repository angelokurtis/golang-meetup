package location

import (
	"encoding/json"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"github.com/angelokurtis/golang-meetup/pkg/location"
	"net/http"
	"os"
)

type IPStack struct {
	log *log.Logger
}

func (i *IPStack) WhereAmI() (*location.Location, error) {
	(*i.log).Info("searching for my current location")

	ak := os.Getenv("IPSTACK_API_ACCESS_KEY")
	r, err := http.Get("http://api.ipstack.com/check?access_key=" + ak)
	if err != nil {
		return nil, err
	}
	b := r.Body
	defer b.Close()
	l := &location.Location{}
	err = json.NewDecoder(b).Decode(l)
	if err != nil {
		return nil, err
	}
	(*i.log).Info("location found. You are in " + l.City + "/" + l.RegionCode)
	return l, nil
}
