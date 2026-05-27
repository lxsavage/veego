package veego

// Set the color of the device(s)
func (fl cmdBuilderChain) Color(c colorRGB) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorRgb",
		Value:    int(c),
	})
}

// Set the color temperature of the device(s)
func (fl cmdBuilderChain) Temperature(t colorTemp) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorTemperatureK",
		Value:    int(t),
	})
}
