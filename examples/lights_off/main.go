// Turns off the lights
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lxsavage/veego"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	key := os.Getenv("GOVEE_KEY")
	if len(key) == 0 {
		log.Fatal("GOVEE_KEY environment variable not defined!")
	}

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
