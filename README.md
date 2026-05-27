# VeeGo: An Unofficial GO Library for Govee Devices

A builder-style library for interacting with Govee devices

## Usage

Turning off the lights is as simple as:

```go
import (
	"log"
	"log/slog"
	"net/http"

	"logansavage.dev/veego"
)

const key = "<GOVEE API KEY>"

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

	// Get all light devices whose names start with "Living Room" and turn them off
	err = controller.Devices().
		TypeIs(veego.DeviceLight).
		NameMatches("^Living Room").
		Off().
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}
```

Additional examples can be found under the example programs in `examples/`.

### Structure

This library is designed around a builder syntax, with two sections:

1. Filters, which sequentially filter down the devices shown
2. Actions, which are a sequence of actions to perform on the devices that pass
   through every filter

With this, the structure of a command is:

```go
err := controller.Devices().
  <filterFunctions>.
  <actionFunctions>.
  Exec()
```

Applying this to the lights off example:

```go
err = controller.Devices().
	TypeIs(veego.DeviceLight).   // Filter function
	NameMatches("^Living Room"). // Filter function
	Off().                       // Action function
	Exec()
````

### Looping and Memoization

If doing operations in a loop, the library provides a `.Memoize()` action that
will pull down the device list and cache it after running all filters. This will
prevent recalculating the device list every iteration of the loop.

```go
allLights := controller.Devices().
	TypeIs(veego.DeviceLight).
	Memoize()

err = allLights.On().Exec()
if err != nil {
	log.Fatal(err)
}

// Turn up the brightness from 0% to 100% in increments of 5%
for i := 0; i <= 100; i += 5 {
	err := allLights.
		Brightness(veego.Percentage(i)).
		Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lamps at %d%%\n", i)
	time.Sleep(1 * time.Second)
}
```

Unlike other actions, additional filters could be applied to the query after
memoizing it.
