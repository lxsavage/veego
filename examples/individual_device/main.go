// Performs operations on individual devices
package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
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

	// Get some devices
	devs := controller.Devices().
		NameMatches("Front Room Light \\d+|Dining Room Light \\d+").
		Query()

	// Skip processing if no devices were found
	if len(devs) == 0 {
		log.Println("0 devices found.")
		os.Exit(0)
	}

	// Run each update call in a goroutine to account for API latency
	var wg sync.WaitGroup
	for i, dev := range devs {
		wg.Go(func() {
			devState, err := dev.Stat(controller)
			if err != nil {
				log.Printf("failed to stat device %s: %v\n", dev.Name, err)
				return
			}

			// Get the current device brightness and increase it by 10%
			if brightness, ok := devState.Brightness(); ok {
				// Increase the brightness by 10%
				next := veego.Percentage(brightness + 10)
				if err := dev.Brightness(controller, next); err != nil {
					log.Printf("failed to set color for %s: %v\n", dev.Name, err)
					return
				}
			}

			// Wait depending on which index the device is
			time.Sleep(time.Duration(1+i) * time.Second)

			// Remove the red from the device color
			if _, g, b, ok := devState.Color(); ok {
				if err := dev.Color(controller, veego.RGB(0, g, b)); err != nil {
					log.Printf("failed to set color for %s: %v\n", dev.Name, err)
					return
				}
			}

			log.Printf("Succeeded with %s\n", dev.Name)
		})
	}

	wg.Wait()
}
