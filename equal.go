package benchmarked

// Equal checks if two slices of bytes are equal
var Equal = equal14

func equal1(a, b []byte) bool {
	return string(a) == string(b)
}

func equal4(a, b []byte) bool {
	la := len(a)
	lb := len(b)
	for i := 0; i < la; i++ {
		if i >= lb {
			return false
		} else if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal5(a, b []byte) bool {
	la := len(a)
	lb := len(b)
	switch la {
	case 0:
		return lb == 0
	case 1:
		return lb == 1 && a[0] == b[0]
	case 2:
		return lb == 2 && a[0] == b[0] && a[1] == b[1]
	case 3:
		return lb == 3 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
	case 4:
		return lb == 4 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
	case lb:
		break
	default: // la != lb
		return false
	}
	// The length is 5 or above, start at index 4
	for i := 4; i < la; i++ {
		if i >= lb {
			return false
		} else if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal7(a, b []byte) bool {
	la := len(a)
	for i, v := range b {
		if i >= la || a[i] != v {
			return false
		}
	}
	return true
}

func equal8(a, b []byte) bool {
	la := len(a)
	lb := len(b)
	if la == 0 && lb == 0 {
		return true
	} else if la != lb {
		return false
	}
	for i := 0; i < lb; i++ {
		if i >= la || a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal14(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if i >= len(a) || a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal18(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
