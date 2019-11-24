package main

import (
	"fmt"
	móduloDeClima "github.com/angelokurtis/golang-meetup/pkg/clima"
	"log"
)

func main() {
	medidor := móduloDeClima.NewOpenWeather()
	clima, err := medidor.AferirMinhaLocalização()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("o clima atual é de %.2f°C e %s\n", clima.Temperatura, clima.Descrição)
}
