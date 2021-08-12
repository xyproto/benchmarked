package benchmarked

// Equal checks if two slices of bytes are equal
var Equal = equal7 // previously equal14

func equal1(a, b []byte) bool {
	return string(a) == string(b)
}

func equal2(a, b []byte) bool {
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

func equal3(a, b []byte) bool {
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

func equal4(a, b []byte) bool {
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
	case 5:
		return lb == 5 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4]
	case lb:
		break
	default: // la != lb
		return false
	}
	// The length is 6 or above, so start at index 5
	// First check the exponential locations, from 5
	for x := 5; x < la; x*=2 {
		if x >= lb || a[x] != b[x] {
			return false
		}
	}
	// Index 6 is now the first unchecked position
	for i := 6; i < la; i++ {
		if i >= lb || a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal5(a, b []byte) bool {
	la := len(a)
	for i, v := range b {
		if i >= la || a[i] != v {
			return false
		}
	}
	return true
}

func equal6(a, b []byte) bool {
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

func equal7(a, b []byte) bool {
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

func equal8(a, b []byte) bool {
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
