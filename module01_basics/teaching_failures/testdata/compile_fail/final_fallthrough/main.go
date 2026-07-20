package finalfallthrough

func demonstrate(value int) {
	switch value {
	case 1:
		fallthrough
	}
}
