package benchmarked

import (
	"bytes"
	"math/rand"
	"testing"
)

// max length of byte slices that are created for the benchmarks
const maxLen = 16

type bytesEqualFunctionType func([]byte, []byte) bool

var (
	result bool

	functions = map[string]bytesEqualFunctionType{
		"bytes.Equal": bytes.Equal,
		"equal1":      equal1,
		"equal2":      equal2,
		"equal3":      equal3,
		"equal4":      equal4,
		"equal5":      equal5,
		"equal6":      equal6,
		"equal7":      equal7,
		"equal8":      equal8,
		"equal9":      equal9,
		"equal10":     equal10,
		"equal11":     equal11,
		"equal12":     equal12,
		"equal13":     equal13,
		"equal14":     equal14,
		"equal15":     equal15,
		"equal16":     equal16,
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
			/*
				b.StopTimer()
				b1s := make([][]byte, b.N)
				b2s := make([][]byte, b.N)
				for i := 0; i < b.N; i++ {
					b1s[i] = RandomBytes(numbytes)
					b2s[i] = RandomBytes(numbytes)
				}
				b.StartTimer()
			*/
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
	}
}
