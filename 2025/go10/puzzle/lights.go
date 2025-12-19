package puzzle

func CountButtonsForLights(m *Machine) []int {
	for count := 1; count < 99; count++ {
		r := tryButtonsForLights(m, count)
		if r != nil {
			return r
		}
	}
	return nil
}

func tryButtonsForLights(m *Machine, count int) []int {

	for b, but := range m.Buttons {
		pushButtonForLights(m, but)

		if count > 1 {
			r := tryButtonsForLights(m, count-1)
			if r != nil {
				r = append(r, b)
				return r
			}
		} else {
			if allLightsOff(m) {
				r := make([]int, 1)
				r[0] = b
				return r
			}
		}
		pushButtonForLights(m, but)
	}
	return nil
}

func allLightsOff(m *Machine) bool {
	for _, l := range m.Lights {
		if l {
			return false
		}
	}
	return true
}

func pushButtonForLights(m *Machine, b *Button) {

	lights := b.Numbers
	for _, l := range lights {
		m.Lights[l] = !m.Lights[l]
	}
}
