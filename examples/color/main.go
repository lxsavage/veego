// Cycles through colors
package main

import (
	"errors"
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
	if errors.Is(err, veego.ErrInvalidKey) {
		log.Fatal("the key at the top of the file is invalid; check it and try again")
	} else if err != nil {
		log.Fatal(err)
	}

	// Turn the lights on, cycle the lights red, green, then blue, then turn them off
	err = controller.Devices().
		TypeIs(veego.DeviceLight).
		On().
		Color(veego.RGB(255, 0, 0)).
		Sleep(2 * time.Second).
		Color(veego.RGB(0, 255, 0)).
		Sleep(2 * time.Second).
		Color(veego.RGB(0, 0, 255)).
		Off().
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}
