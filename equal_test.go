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
	"Equal1":      equal1,
	"Equal2":      equal2,
	"Equal3":      equal3,
	"Equal4":      equal4,
	"Equal5":      equal5,
	"Equal6":      equal6,
	"Equal7":      equal7,
	"Equal8":      equal8,
	"Equal9":      equal9,
	"Equal10":     equal10,
	"Equal11":     equal11,
	"Equal12":     equal12,
	"Equal13":     equal13,
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
