package benchmarked

import (
	"sync"
)

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
	for x := 5; x < la; x *= 2 {
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

func equal9(a, b []byte) bool {
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
	return string(a) == string(b)
}

func equal10(a, b []byte) bool {
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
	return string(a) == string(b)
}

func equal11(a, b []byte) bool {
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
	case 6:
		return lb == 6 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5]
	case lb:
		break
	default: // la != lb
		return false
	}
	return string(a) == string(b)
}

func equal12(a, b []byte) bool {
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
	case 6:
		return lb == 6 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5]
	case 7:
		return lb == 7 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5] && a[6] == b[6]
	case lb:
		break
	default: // la != lb
		return false
	}
	return string(a) == string(b)
}

func examineCenter(start, stop int, a, b *[]byte, wg *sync.WaitGroup, differ *bool) {
	if start == stop {
		wg.Done()
		return
	}
	m := start + (stop-start)/2
	//fmt.Printf("range %d to %d, center %d\n", start, stop, m)
	if (*a)[m] != (*b)[m] {
		*differ = true
		wg.Done()
		return
	}
	wg.Add(2)
	go examineCenter(start, m, a, b, wg, differ)
	go examineCenter(m, stop, a, b, wg, differ)
	wg.Done()
}

func equal13(a, b []byte) bool {
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
	case 6:
		return lb == 6 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5]
	case 7:
		return lb == 7 && a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4] && a[5] == b[5] && a[6] == b[6]
	case lb:
		break
	default: // la != lb
		return false
	}

	differ := false

	var wg sync.WaitGroup
	wg.Add(1)
	go examineCenter(0, la-1, &a, &b, &wg, &differ)
	wg.Wait()

	return differ
}

func equal14(a, b []byte) bool {
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
		return !(string(a) != string(b))
	default: // la != lb
		return false
	}
}

func equal15(a, b []byte) bool {
	switch len(a) {
	case 0:
		return len(b) == 0
	case 1:
		return len(b) == 1 && a[0] == b[0]
	case 2:
		return len(b) == 2 && a[0] == b[0] && a[1] == b[1]
	}
	return string(a) == string(b)
}
