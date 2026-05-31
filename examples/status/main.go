// Checks and prints the status of the first device it sees
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

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

	// Get the power status of the light
	state := capabilities.IsOn()
	fmt.Printf("Is the device switched on? %t\n", state)

	// Get the color of the light
	if r, g, b, ok := capabilities.Color(); ok {
		fmt.Printf("Color code of the light's current state is: RGB(%d, %d, %d)\n", r, g, b)
	}

	// Get the state of a specific capability on the device not implemented by the
	// library directly yet
	if k, ok := capabilities.GetState("colorTemperatureK"); ok {
		// All numerical values should be cast to a float64 before operating on them
		// due to how the JSON standard defines a number as a floating point value
		colorTempK := k.(float64)
		fmt.Printf("Color temperature: %.2f K\n", colorTempK)
	}
}
