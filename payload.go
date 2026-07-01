package veego

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Send a raw device payload to the API; should only be used if the needed
// functionality is not yet implemented.
//
// Note that this payload is sent as the "capability" section of the following
// API method:
//
// POST /router/api/v1/device/control
//
// See https://developer.govee.com/reference/control-you-devices for more
// details.
func (d Device) SendPayload(c Controller, pl DevicePayload) error {
	body, err := json.Marshal(map[string]any{
		"requestId": uuid.NewString(),
		"payload": map[string]any{
			"sku":        d.SKU,
			"device":     d.MACAddress,
			"capability": pl,
		},
	})
	if err != nil {
		return err
	}

	_, err = Request[map[string]any](c,
		"POST",
		"/device/control",
		body,
	)
	if err != nil {
		return err
	}

	return nil
}

// Send a raw device payload to the API for each device in the chain; should
// only be used if the needed functionality is not yet implemented.
//
// Note that this payload is sent as the "capability" section of the following
// API method:
//
// POST /router/api/v1/device/control
//
// See https://developer.govee.com/reference/control-you-devices for more
// details.
func (fl cmdBuilderChain) SendPayload(pl DevicePayload) cmdBuilderChain {
	cmd := func(c Controller, devs []Device) ([]Device, error) {
		for i, d := range devs {
			if err := d.SendPayload(c, pl); err != nil {
				return devs, fmt.Errorf("%d, %v", i, err)
			}
		}
		return devs, nil
	}

	fl.actions = append(fl.actions, cmd)
	return fl
}
