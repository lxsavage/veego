package veego

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

// Adjust the brightness to a value between 0-100
func (d Device) Brightness(c *Controller, b percentage) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}

// Turn the device(s) off
//
// Deprecated: use Device.Off() instead.
func (fl cmdBuilderChain) Off() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    0,
	})
}

// Turn the device(s) on
//
// Deprecated: use Device.On() instead.
func (fl cmdBuilderChain) On() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    1,
	})
}

// Adjust the brightness to a value between 0-100
//
// Deprecated: use Device.Brightness() instead.
func (fl cmdBuilderChain) Brightness(b percentage) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}
