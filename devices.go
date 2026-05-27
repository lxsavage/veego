package veego

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type rangeParam struct {
	Min       int
	Max       int
	Precision int
}

type optionParam struct {
	Name  string
	Value int
}

type deviceCapability struct {
	Type       DeviceCapabilityType
	Instance   string
	Parameters struct {
		DataType string
		Range    *rangeParam
		Options  []optionParam
	}
}

type deviceCapabilityStates []deviceCapabilityState
type deviceCapabilityState struct {
	Type     DeviceCapabilityType
	Instance string
	State    struct{ Value any }
}

type DevicePayload struct {
	// The type of value that the property is
	Type DeviceCapabilityType `json:"type"`

	// The name of the property to change
	Instance string `json:"instance"`

	// Type is device type dependent
	Value any `json:"value"`
}

// An IoT device that can be used in context of other Govee commands
type Device struct {
	SKU          string `json:"sku"`
	MACAddress   string `json:"device"`
	Name         string `json:"deviceName"`
	Type         DeviceType
	Capabilities []deviceCapability
}

// Get a list of all devices accessible. This is a fallback function and should
// only be used if the fluent version does not cover the needed use.
func (c Controller) FindDevices() ([]Device, error) {
	res, err := Request[struct {
		Data []Device
	}](c, "GET", "/user/devices", nil)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (d Device) String() string {
	caps := []string{}
	for _, c := range d.Capabilities {
		caps = append(caps, c.Instance)
	}
	return fmt.Sprintf("<SKU: %s, MACAddress: %s, Name: %s ,Type: %s, Capabilities: [%s]>",
		d.SKU, d.MACAddress, d.Name, d.Type,
		strings.Join(caps, ", "),
	)
}

// Get the status of the device as an array of capabilities and their states
func (d Device) Stat(c *Controller) (deviceCapabilityStates, error) {
	body, err := json.Marshal(map[string]any{
		"requestId": uuid.NewString(),
		"payload": map[string]string{
			"sku":    d.SKU,
			"device": d.MACAddress,
		},
	})
	if err != nil {
		return nil, errors.New("failed to marshal JSON body for " + d.MACAddress + ": " + err.Error())
	}

	status, err := Request[struct {
		Payload struct {
			Capabilities []deviceCapabilityState
		}
	}](*c, "POST", "/device/state", body)
	if err != nil {
		return nil, errors.New("failed to get device stat for " + d.MACAddress + ": " + err.Error())
	}

	return status.Payload.Capabilities, nil
}

func (dcs deviceCapabilityState) String() string {
	return fmt.Sprintf("<Name: %s, Type: %s, State: %v>", dcs.Instance, dcs.Type, dcs.State)
}

// Returns the state of a specific capability by its name; the second value is true if found, otherwise false
func (dc deviceCapabilityStates) GetState(key string) (any, bool) {
	for _, cap := range dc {
		if cap.Instance == key {
			return cap.State.Value, true
		}
	}
	return nil, false
}
