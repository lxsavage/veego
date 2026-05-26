// Checks the status of the first device it sees
//
// Created by Logan Savage
package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

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

	// Get the list of all devices
	devices := controller.Devices().Query()
	if len(devices) == 0 {
		log.Fatalln("No devices found.")
	}

	// Get the status of the first device returned
	states, err := devices[0].Stat(controller)
	if err != nil {
		log.Fatal(err)
	}

	// Print out the status
	for _, s := range states {
		fmt.Printf("%v\n", s)
	}
}
