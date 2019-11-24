package localização

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type IPStack struct{}

func NewIPStack() *IPStack {
	return &IPStack{}
}

func (i *IPStack) OndeEstou() (*Local, error) {
	fmt.Println("buscando minha localização atual")

	ak := os.Getenv("IPSTACK_API_ACCESS_KEY")
	r, err := http.Get("http://api.ipstack.com/check?access_key=" + ak)
	if err != nil {
		return nil, err
	}
	b := r.Body
	defer b.Close()
	l := &Local{}
	err = json.NewDecoder(b).Decode(l)
	if err != nil {
		return nil, err
	}
	fmt.Println("localização encontrada")
	fmt.Println("estamos em " + l.Cidade + "/" + l.CódigoRegional)
	return l, nil
}
