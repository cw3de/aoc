package puzzle

func CountButtonsForLights(m *Machine) []int {
	for count := 1; count < 99; count++ {
		r := TryButtonsForLights(m, count)
		if r != nil {
			return r
		}
	}
	return nil
}

func TryButtonsForLights(m *Machine, count int) []int {

	for b := 0; b < len(m.Buttons); b++ {
		PushButtonForLights(m, b)

		if count > 1 {
			r := TryButtonsForLights(m, count-1)
			if r != nil {
				r = append(r, b)
				return r
			}
		} else {
			if AllLightsOff(m) {
				r := make([]int, 1)
				r[0] = b
				return r
			}
		}
		PushButtonForLights(m, b)
	}
	return nil
}

func AllLightsOff(m *Machine) bool {
	for _, l := range m.Lights {
		if l {
			return false
		}
	}
	return true
}

func PushButtonForLights(m *Machine, b int) {

	lights := m.Buttons[b].Numbers
	for _, l := range lights {
		m.Lights[l] = !m.Lights[l]
	}
}
