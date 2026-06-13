package veego

// Gets the power state of the device
func (dc deviceCapabilityStates) IsOn() bool {
	p, ok := dc.GetState("powerSwitch")
	if !ok {
		return false
	}

	// All numbers are parsed as float64 values with the default JSON unmarshaler
	return p.(float64) != 0
}

// Get the brightness of the device as a percentage between 0-100
func (dc deviceCapabilityStates) Brightness() (int, bool) {
	p, ok := dc.GetState("brightness")
	if !ok {
		return 0, false
	}

	// All numbers are parsed as float64 values with the default JSON unmarshaler
	return int(p.(float64)), true
}

// Turn the device off
func (d Device) Off(c *Controller) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    0,
	})
}

// Turn the device on
func (d Device) On(c *Controller) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    1,
	})
}

// Adjust the brightness of the device to a value between 0-100
func (d Device) Brightness(c *Controller, b percentage) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}

// Turn the devices off
func (fl cmdBuilderChain) Off() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    0,
	})
}

// Turn the devices on
func (fl cmdBuilderChain) On() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    1,
	})
}

// Adjust the brightness of the devices to a value between 0-100
func (fl cmdBuilderChain) Brightness(b percentage) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}
