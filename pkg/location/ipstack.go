package location

import (
	"encoding/json"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"net/http"
	"os"
)

type IPStack struct {
	log log.Logger
}

func NewIPStack(log log.Logger) *IPStack {
	return &IPStack{log: log}
}

func (i *IPStack) WhereAmI() (*Location, error) {
	i.log.Info("buscando minha localização atual")

	ak := os.Getenv("IPSTACK_API_ACCESS_KEY")
	r, err := http.Get("http://api.ipstack.com/check?access_key=" + ak)
	if err != nil {
		return nil, err
	}
	b := r.Body
	defer b.Close()
	l := &Location{}
	err = json.NewDecoder(b).Decode(l)
	if err != nil {
		return nil, err
	}
	i.log.Info("localização encontrada. Estamos em " + l.City + "/" + l.RegionCode)
	return l, nil
}
