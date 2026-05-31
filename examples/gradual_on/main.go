// Gradually brightens all light devices seen from 0-100% over the course of
// 20 seconds
package main

import (
	"errors"
	"fmt"
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

	// Get a cached reference to every light device accessible. Memoize should be
	// used if the devices list is reused in its current state to prevent
	// refetching devices and reapplying filters every time an Exec() call is made
	// on them.
	allLights := controller.Devices().
		TypeIs(veego.DeviceLight).
		Memoize()

	err = allLights.On().Exec()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= 100; i += 5 {
		// Gradually increase the brightness of these lights from 0-100% by
		// increments of 5%
		err := allLights.
			Brightness(veego.Percentage(i)).
			Exec()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Lamps at %d%%\n", i)
		time.Sleep(1 * time.Second)
	}

	// Turn off the lights after the demo
	_ = allLights.Off().Exec()
}
