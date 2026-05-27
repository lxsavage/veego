package veego

// Turn the device(s) off
func (fl cmdBuilderChain) Off() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    0,
	})
}

// Turn the device(s) on
func (fl cmdBuilderChain) On() cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    1,
	})
}

// Adjust the brightness to a value between 0-100
func (fl cmdBuilderChain) Brightness(b percentage) cmdBuilderChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}
