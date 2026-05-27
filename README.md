# VeeGo: An Unofficial GO Library for Govee Devices

A builder-style library for interacting with Govee devices

## Usage

Turning off the lights is as simple as:

```go
import (
	"log"
	"log/slog"
	"net/http"

	"github.com/lxsavage/veego"
)

const key = "<GOVEE API KEY>"

func main() {
	// Create a new controller to handle device interaction using the default HTTP
	// client
	controller, err := veego.NewController(key).
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

For additional usage information, see the [wiki](https://github.com/lxsavage/veego/wiki). Additional examples can be found under the example programs in `examples/`.
