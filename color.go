package veego

// Set the color of the device
func (d Device) Color(c *Controller, color colorRGB) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorRgb",
		Value:    int(color),
	})
}

// Set the color temperature of the device
func (d Device) Temperature(c *Controller, t colorTemp) error {
	return d.SendPayload(*c, DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorTemperatureK",
		Value:    int(t),
	})
}

// Get the current color of the device, the last return is true if found, otherwise false
func (dc deviceCapabilityStates) Color() (uint8, uint8, uint8, bool) {
	o, ok := dc.GetState("colorRgb")
	if !ok {
		return 0, 0, 0, false
	}

	r, g, b := colorRGB(o.(float64)).Components()
	return r, g, b, true
}

// Get the current color temperature of the device, the last return is true if found, otherwise false
func (dc deviceCapabilityStates) Temperature() (int, bool) {
	k, ok := dc.GetState("colorTemperatureK")
	if !ok {
		return 0, false
	}

	return int(k.(float64)), true
}

// Set the color of the devices
func (fl cmdBuilderChain) Color(c colorRGB) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorRgb",
		Value:    int(c),
	})
}

// Set the color temperature of the devices
func (fl cmdBuilderChain) Temperature(t colorTemp) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorTemperatureK",
		Value:    int(t),
	})
}
