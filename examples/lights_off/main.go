// Turns off the lights
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/lxsavage/veego"
)

const key = "GOVEE_KEY"

func main() {
	// Create a new controller to handle device interaction using the default HTTP
	// client
	controller, err := veego.NewController(key).
		WithClient(http.DefaultClient).
		Init()
	if errors.Is(err, veego.ErrInvalidKey) {
		log.Fatal("the key at the top of the file is invalid; check it and try again")
	} else if err != nil {
		log.Fatal(err)
	}

	err = controller.Devices().
		TypeIs(veego.DeviceLight).
		Off().
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}
