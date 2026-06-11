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

// Set the color of the device(s)
//
// Deprecated: avoid using command chains, and instead direct actions to individual devices instead
func (fl cmdBuilderChain) Color(c colorRGB) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorRgb",
		Value:    int(c),
	})
}

// Set the color temperature of the device(s)
//
// Deprecated: avoid using command chains, and instead direct actions to individual devices instead
func (fl cmdBuilderChain) Temperature(t colorTemp) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorTemperatureK",
		Value:    int(t),
	})
}
