package veego

// Turn the device(s) off
func (fl fluentChain) Off() fluentChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    0,
	})
}

// Turn the device(s) on
func (fl fluentChain) On() fluentChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeOnOff,
		Instance: "powerSwitch",
		Value:    1,
	})
}

// Adjust the brightness to a value between 0-100
func (fl fluentChain) Brightness(b percentage) fluentChain {
	return fl.SendPayload(DevicePayload{
		Type:     TypeRange,
		Instance: "brightness",
		Value:    int(b),
	})
}
