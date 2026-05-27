// Checks and prints the status of the first device it sees
package main

import (
	"fmt"
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

	// Get the list of all devices
	devices := controller.Devices().Query()
	if len(devices) == 0 {
		fmt.Println("No devices found.")
		os.Exit(0)
	}

	// Get the status of the first device returned
	capabilities, err := devices[0].Stat(controller)
	if err != nil {
		log.Fatal(err)
	}

	// List all capabilities available
	for _, s := range capabilities {
		fmt.Println(s.String())
	}

	// Get the state of a specific capability on the device
	if p, ok := capabilities.GetState("powerSwitch"); ok {
		// Convert the JSON number to a boolean; it must be done this way due to how
		// JSON is deserialized internally
		isOn := p.(float64) != 0
		fmt.Printf("The state of the power is: %t\n", isOn)
	}
}
