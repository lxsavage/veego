package veego

import "time"

// Pause for a moment
func (fl cmdBuilderChain) Sleep(t time.Duration) cmdBuilderChain {
	cmd := func(_ Controller, d []Device) ([]Device, error) {
		time.Sleep(t)
		return d, nil
	}

	fl.actions = append(fl.actions, cmd)
	return fl
}
