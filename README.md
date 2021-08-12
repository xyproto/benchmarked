# Benchmarked

The quest to find a faster `bytes.Equal` function.

So far, the best performing function is 30% faster than `bytes.Equal`.


## Code comparison

```go
func Equal(a, b []byte) bool {
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
```

If `if i >= len(a) ||` is removed, the performance is slightly worse. This is a bit counter-intuitive.

For comparison, `bytes.Equal` looks like [this](https://cs.opensource.google/go/go/+/refs/tags/go1.16.7:src/bytes/bytes.go;l=18):

```go
func Equal(a, b []byte) bool {
    return string(a) == string(b)
}
```


## Benchmark results

Tested on Arch Linux, using `go version go1.16.7 linux/amd64`.

```
goos: linux
goarch: amd64
pkg: github.com/xyproto/benchmarked
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkEqual/equal6-12         	 1000000	      1066.0 ns/op
BenchmarkEqual/equal3-12         	 1452951	       817.6 ns/op
BenchmarkEqual/equal1-12         	 1553632	       803.4 ns/op
BenchmarkEqual/bytes.Equal-12    	 1640098	       768.0 ns/op
BenchmarkEqual/equal10-12        	 1699735	       753.1 ns/op
BenchmarkEqual/equal8-12         	 1636986	       740.7 ns/op
BenchmarkEqual/equal12-12        	 1777388	       685.6 ns/op
BenchmarkEqual/equal15-12        	 1789273	       674.1 ns/op
BenchmarkEqual/equal11-12        	 1798375	       667.2 ns/op
BenchmarkEqual/equal16-12        	 1767123	       663.7 ns/op
BenchmarkEqual/equal5-12         	 1892774	       653.7 ns/op
BenchmarkEqual/equal4-12         	 1916486	       617.6 ns/op
BenchmarkEqual/equal7-12         	 1945572	       610.1 ns/op
BenchmarkEqual/equal2-12         	 1943868	       610.6 ns/op
BenchmarkEqual/equal13-12        	 2085339	       575.5 ns/op
BenchmarkEqual/equal9-12         	 2122830	       574.2 ns/op
BenchmarkEqual/equal14-12        	 2230186	       534.8 ns/op
PASS
ok  	github.com/xyproto/benchmarked	31.966s
```

Currently, `equal14` is the function that is exported as `Equal` in the `benchmarked` package.


## Accuracy

I am aware that perfect benchmarking is a tricky.

Please let me know if you have improvements to how the functions are benchmarked, or how the benchmarks are interpreted!


## General info

* Version: 0.1.0
* License: BSD
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
