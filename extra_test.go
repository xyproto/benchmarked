// These tests are from Go, src/bytes/bytes_test.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package benchmarked

import (
	"fmt"
	"testing"
)

var bmbuf []byte
var isRaceBuilder = false

// For ease of reading, the test cases use strings that are converted to byte
// slices before invoking the functions.

func TestEqualExhaustive(t *testing.T) {
	var size = 128
	if testing.Short() {
		size = 32
	}
	a := make([]byte, size)
	b := make([]byte, size)
	bInit := make([]byte, size)
	// randomish but deterministic data
	for i := 0; i < size; i++ {
		a[i] = byte(17 * i)
		bInit[i] = byte(23*i + 100)
	}

	for len := 0; len <= size; len++ {
		for x := 0; x <= size-len; x++ {
			for y := 0; y <= size-len; y++ {
				copy(b, bInit)
				copy(b[y:y+len], a[x:x+len])
				if !Equal(a[x:x+len], b[y:y+len]) || !Equal(b[y:y+len], a[x:x+len]) {
					t.Errorf("Equal(%d, %d, %d) = false", len, x, y)
				}
			}
		}
	}
}

// make sure Equal returns false for minimally different strings. The data
// is all zeros except for a single one in one location.
func TestNotEqual(t *testing.T) {
	var size = 128
	if testing.Short() {
		size = 32
	}
	a := make([]byte, size)
	b := make([]byte, size)

	for len := 0; len <= size; len++ {
		for x := 0; x <= size-len; x++ {
			for y := 0; y <= size-len; y++ {
				for diffpos := x; diffpos < x+len; diffpos++ {
					a[diffpos] = 1
					if Equal(a[x:x+len], b[y:y+len]) || Equal(b[y:y+len], a[x:x+len]) {
						t.Errorf("NotEqual(%d, %d, %d, %d) = true", len, x, y, diffpos)
					}
					a[diffpos] = 0
				}
			}
		}
	}
}

func bmEqual(equal func([]byte, []byte) bool) func(b *testing.B, n int) {
	return func(b *testing.B, n int) {
		if len(bmbuf) < 2*n {
			bmbuf = make([]byte, 2*n)
		}
		buf1 := bmbuf[0:n]
		buf2 := bmbuf[n : 2*n]
		buf1[n-1] = 'x'
		buf2[n-1] = 'x'
		for i := 0; i < b.N; i++ {
			eq := equal(buf1, buf2)
			if !eq {
				b.Fatal("bad equal")
			}
		}
		buf1[n-1] = '\x00'
		buf2[n-1] = '\x00'
	}
}

func benchBytes(name string, b *testing.B, sizes []int, f func(b *testing.B, n int)) {
	for _, n := range sizes {
		if isRaceBuilder && n > 4<<10 {
			continue
		}
		b.Run(name+"_"+valName(n), func(b *testing.B) {
			if len(bmbuf) < n {
				bmbuf = make([]byte, n)
			}
			b.SetBytes(int64(n))
			f(b, n)
		})
	}
}

func valName(x int) string {
	if s := x >> 20; s<<20 == x {
		return fmt.Sprintf("%dM", s)
	}
	if s := x >> 10; s<<10 == x {
		return fmt.Sprintf("%dK", s)
	}
	return fmt.Sprint(x)
}
