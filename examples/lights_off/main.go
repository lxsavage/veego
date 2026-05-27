// Turns off the lights
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/lxsavage/veego"
)

const key = "GOVEE_KEY"

func main() {
	// Create a new controller to handle device interaction using the default HTTP
	// client
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	controller, err := veego.NewController(key, logger).
		WithClient(http.DefaultClient).
		Init()
	if err != nil {
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
