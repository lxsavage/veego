// Adjusts the color temperature
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lxsavage/veego"
)

const key = "GOVEE_KEY"

func main() {
	// Create a new controller to handle device interaction using the default HTTP
	// client
	controller, err := veego.NewController(key).
		WithClient(http.DefaultClient).
		Init()
	if err != nil {
		log.Fatal(err)
	}

	// Turn the lights on, adjust the color temperature a bit, then turn them off
	err = controller.Devices().
		TypeIs(veego.DeviceLight).
		On().
		Temperature(veego.K(6000)).
		Sleep(2 * time.Second).
		Temperature(veego.K(2000)).
		Temperature(veego.K(4000)).
		Sleep(2 * time.Second).
		Off().
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}
