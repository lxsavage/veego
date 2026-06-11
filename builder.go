package veego

import (
	"errors"
	"strconv"
)

type cmdBuilderChain struct {
	filters     []cmdBuilderFilter
	actions     []cmdBuilderAction
	memoDevices []Device
	controller  *Controller
	memoized    bool
}

type cmdBuilderCommand func(*cmdBuilderChain) *cmdBuilderChain
type cmdBuilderAction func(Controller, []Device) ([]Device, error)
type cmdBuilderFilter func(Controller, []Device) []Device

// Start a command chain with the list of all accessible devices
func (c *Controller) Devices() cmdBuilderChain {
	return cmdBuilderChain{
		controller: c,
		actions:    []cmdBuilderAction{},
	}
}

// See what devices would be affected by a command chain
func (f cmdBuilderChain) Query() []Device {
	var devices []Device
	if f.memoized {
		devices = f.memoDevices
	} else {
		var err error
		devices, err = f.controller.findDevices()
		if err != nil {
			return nil
		}
	}

	if len(devices) == 0 {
		return nil
	}

	for _, filter := range f.filters {
		devices = filter(*f.controller, devices)
	}

	if len(devices) == 0 {
		return nil
	}

	return devices
}

// Apply and cache the device/filter portion of the command chain
//
// Deprecated: avoid using command chains, and instead direct actions to individual devices instead
func (f cmdBuilderChain) Memoize() cmdBuilderChain {
	if !f.memoized {
		devices, err := f.controller.findDevices()
		if err != nil {
			f.memoDevices = []Device{}
			return f
		}
		f.memoDevices = devices
	}

	for _, filter := range f.filters {
		f.memoDevices = filter(*f.controller, f.memoDevices)
	}
	f.filters = []cmdBuilderFilter{}
	f.memoized = true
	return f
}

// Execute a command command chain, returning an error and stopping if there
// are any issues with the commands in the chain
//
// Deprecated: avoid using command chains, and instead direct actions to individual devices instead
func (f cmdBuilderChain) Exec() error {
	var devices []Device
	if f.memoized {
		devices = f.memoDevices
	} else {
		var err error
		devices, err = f.controller.findDevices()
		if err != nil {
			return err
		}
	}

	if len(devices) == 0 {
		return nil
	}

	for _, filter := range f.filters {
		devices = filter(*f.controller, devices)
	}

	if len(devices) == 0 {
		return nil
	}

	for i, cmd := range f.actions {
		var err error
		devices, err = cmd(*f.controller, devices)
		if err != nil {
			return errors.New(
				"failed to apply fluent query " +
					strconv.Itoa(i) +
					": " + err.Error(),
			)
		}
	}

	return nil
}
