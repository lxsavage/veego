package veego

import "strconv"

type fluentChain struct {
	filters     []fluentFilter
	actions     []fluentAction
	memoDevices []Device
	controller  *Controller
	memoized    bool
}

type fluentCommand func(*fluentChain) *fluentChain
type fluentAction func(Controller, []Device) ([]Device, error)
type fluentFilter func(Controller, []Device) []Device

// Start a fluent command chain with the list of all accessible devices
func (c *Controller) Devices() fluentChain {
	return fluentChain{
		controller: c,
		actions:    []fluentAction{},
	}
}

// See what devices would be affected by a fluent chain
func (f fluentChain) Query() []Device {
	var devices []Device
	if f.memoized {
		devices = f.memoDevices
	} else {
		var err error
		devices, err = f.controller.FindDevices()
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

// Apply and cache the device/filter portion of the fluent chain
func (f fluentChain) Memoize() fluentChain {
	if !f.memoized {
		devices, err := f.controller.FindDevices()
		if err != nil {
			f.memoDevices = []Device{}
			return f
		}
		f.memoDevices = devices
	}

	for _, filter := range f.filters {
		f.memoDevices = filter(*f.controller, f.memoDevices)
	}
	f.filters = []fluentFilter{}
	f.memoized = true
	return f
}

// Execute a fluent command chain, returning an error and stopping if there
// are any issues with the commands in the chain
func (f fluentChain) Exec() error {
	var devices []Device
	if f.memoized {
		devices = f.memoDevices
	} else {
		var err error
		devices, err = f.controller.FindDevices()
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
			f.controller.logger.Error(
				"failed to apply fluent query " +
					strconv.Itoa(i) +
					": " + err.Error(),
			)
			return err
		}
	}

	return nil
}
