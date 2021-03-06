package benchmarked

import (
	"bytes"
	"math/rand"
	"testing"
)

// max length of byte slices that are created for the benchmarks
const maxLen = 32

type bytesEqualFunctionType func([]byte, []byte) bool

var (
	result bool

	functions = map[string]bytesEqualFunctionType{
		"bytes.Equal": bytes.Equal,
		//"equal1":      equal1,
		//"equal2": equal2,
		//"equal3":      equal3,
		//"equal4":      equal4,
		//"equal6":      equal6,
		//"equal7":      equal7,
		//"equal8":      equal8,
		//"equal9":      equal9,
		//"equal10": equal10,
		"equal11": equal11,
		//"equal12": equal12,
		//"equal13": equal13,
		//"equal14":     equal14,
		//"equal15":     equal15,
		//"equal16":     equal16,
		//"equal17":     equal17,
		//"equal18":     equal18,
		//"equal19":     equal19,
		//"equal20":     equal20,
		//"equal21": equal21,
		//"equal22":     equal22,
		//"equal23":     equal23,
		//"equal24":     equal24,
		//"equal25":     equal25,
		//"equal26":     equal26,
		//"equal27":     equal27,
		//"equal28":     equal28,
		//"equal29":     equal29,
		//"equal30":     equal30,
		//"equal31":     equal31,
		//"equal32":     equal32,
		"equal33": equal33,
		//"equal34": equal34,
		//"equal35": equal35,
		//"equal36": equal36,
	}
)

func randomBytes(maxIndex int) []byte {
	if maxIndex == 0 {
		return []byte{}
	}
	b := make([]byte, rand.Intn(maxIndex))
	for i := range b {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func BenchmarkEqual(b *testing.B) {
	for k, v := range functions {
		b.Run(k, func(b *testing.B) {
			r := false
			for x := 0; x < maxLen; x++ {
				for y := 0; y < maxLen; y++ {
					b.StopTimer()
					bx := randomBytes(x)
					by := randomBytes(y)
					b.StartTimer()
					for i := 0; i < b.N; i++ {
						r = v(bx, by)
					}
				}
			}

			result = r
		})
		// From go/src/bytes/bytes_test.go
		b.Run(k+"_0", func(b *testing.B) {
			var buf [4]byte
			buf1 := buf[0:0]
			buf2 := buf[1:1]
			for i := 0; i < b.N; i++ {
				eq := v(buf1, buf2)
				if !eq {
					b.Fatal("bad equal")
				}
			}
		})
		sizes := []int{1, 6, 9, 15, 16, 20, 32, 4 << 10, 4 << 20, 64 << 20, 128 << 20}
		benchBytes(k, b, sizes, bmEqual(v))
	}
}

func TestEqual(t *testing.T) {
	if !Equal(nil, nil) {
		t.Fatal("nil and nil should be equal")
	}
}
