package veego

// Set the color of the device(s)
func (fl fluentChain) Color(c colorRGB) fluentChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorRgb",
		Value:    int(c),
	})
}

// Set the color temperature of the device(s)
func (fl fluentChain) Temperature(t colorTemp) fluentChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeColorSetting,
		Instance: "colorTemperatureK",
		Value:    int(t),
	})
}
