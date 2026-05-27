package veego

import "regexp"

// Filter devices down by name
func (f cmdBuilderChain) NameMatches(regex string) cmdBuilderChain {
	filter := func(c Controller, devs []Device) []Device {
		re, err := regexp.Compile(regex)
		if err != nil {
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

// Filters to devices of the specified type(s)
func (f cmdBuilderChain) TypeIs(types ...DeviceType) cmdBuilderChain {
	allowedTypes := map[DeviceType]bool{}
	for _, t := range types {
		allowedTypes[t] = true
	}

	filter := func(c Controller, devs []Device) []Device {
		newDevs := []Device{}
		for _, d := range devs {
			if _, ok := allowedTypes[d.Type]; ok {
				newDevs = append(newDevs, d)
			}
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}

// Filters to the device that has the specified MAC address
func (f cmdBuilderChain) MacIs(addrs ...string) cmdBuilderChain {
	allowedAddrs := map[string]bool{}
	for _, a := range addrs {
		allowedAddrs[a] = true
	}
	filter := func(c Controller, devs []Device) []Device {
		newDevs := []Device{}
		for _, d := range devs {
			if _, ok := allowedAddrs[d.MACAddress]; ok {
				newDevs = append(newDevs, d)
			}
		}

		return newDevs
	}

	f.filters = append(f.filters, filter)
	return f
}

// Take the first count devices and discard the rest
func (f cmdBuilderChain) Take(count int) cmdBuilderChain {
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
