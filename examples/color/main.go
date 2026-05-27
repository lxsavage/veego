// Cycles through colors
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

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
