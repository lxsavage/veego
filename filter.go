package veego

import "regexp"

// Filter devices down by name
func (f fluentChain) NameMatches(regex string) fluentChain {
	filter := func(c Controller, devs []Device) []Device {
		re, err := regexp.Compile(regex)
		if err != nil {
			f.controller.logger.Error("Invalid regular expression: " + regex)
			return []Device{}
		}

		newDevs := []Device{}
		for _, d := range devs {
			if re.MatchString(d.Name) {
				newDevs = append(newDevs, d)
			}
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}

// Filters to devices of the specified type
func (f fluentChain) TypeIs(t DeviceType) fluentChain {
	filter := func(c Controller, devs []Device) []Device {
		newDevs := []Device{}
		for _, d := range devs {
			if d.Type == t {
				newDevs = append(newDevs, d)
			}
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}

// Filters to the device that has the specified MAC address
func (f fluentChain) MacIs(s string) fluentChain {
	filter := func(c Controller, devs []Device) []Device {
		newDevs := []Device{}
		for _, d := range devs {
			if d.MACAddress == s {
				newDevs = append(newDevs, d)
			}
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}

// Take the first count devices and discard the rest
func (f fluentChain) Take(count int) fluentChain {
	filter := func(c Controller, devs []Device) []Device {
		newDevs := []Device{}
		if count == 0 {
			return newDevs
		}

		takeCount := min(len(devs), count)
		for i := 0; i < takeCount; i++ {
			newDevs = append(newDevs, devs[i])
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}
