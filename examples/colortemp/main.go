// Adjusts the color temperature
//
// Created by Logan Savage
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"logansavage.dev/veego"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	key := os.Getenv("GOVEE_KEY")
	if len(key) == 0 {
		log.Fatal("GOVEE_KEY environment variable not defined!")
	}

	// Create a new controller to handle device interaction using the default
	// HTTP client
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	controller, err := veego.NewController(key, logger).
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
