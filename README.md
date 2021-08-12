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

Currently, `equal14` is the function that is exported as `Equal` in the `benchmarked` package, but `equal9` is sometimes slight faster.


## Accuracy

I am aware that perfect benchmarking is a tricky.

Please let me know if you have improvements to how the functions are benchmarked, or how the benchmarks are interpreted!


## Go 1.17

After replacing the `bytes.Equal` function with the `benchmarked.Equal` function and running the benchmark again, these are the new results:

```
go version devel go1.17-c797850a69 Thu Aug 12 20:01:34 2021 +0200 linux/amd64
goos: linux
goarch: amd64
pkg: github.com/xyproto/benchmarked
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkEqual/equal6-12          	 1253846	       954.9 ns/op
BenchmarkEqual/equal12-12         	 1893031	       631.1 ns/op
BenchmarkEqual/equal1-12          	 1970276	       605.4 ns/op
BenchmarkEqual/equal3-12          	 2053842	       566.0 ns/op
BenchmarkEqual/equal16-12         	 2338264	       508.2 ns/op
BenchmarkEqual/equal19-12         	 2376043	       504.5 ns/op
BenchmarkEqual/equal5-12          	 2422234	       501.1 ns/op
BenchmarkEqual/equal4-12          	 2443252	       496.6 ns/op
BenchmarkEqual/equal18-12         	 2423817	       494.5 ns/op
BenchmarkEqual/equal8-12          	 2475429	       491.7 ns/op
BenchmarkEqual/equal7-12          	 2503693	       490.0 ns/op
BenchmarkEqual/equal11-12         	 2477918	       490.0 ns/op
BenchmarkEqual/equal14-12         	 2497464	       482.7 ns/op
BenchmarkEqual/equal10-12         	 2480304	       478.4 ns/op
BenchmarkEqual/equal2-12          	 2538212	       476.6 ns/op
BenchmarkEqual/equal15-12         	 2491797	       476.9 ns/op
BenchmarkEqual/equal17-12         	 2467926	       471.7 ns/op
BenchmarkEqual/equal13-12         	 2668387	       456.0 ns/op
BenchmarkEqual/equal20-12         	 2714343	       445.1 ns/op
BenchmarkEqual/equal9-12          	 2795047	       433.5 ns/op
BenchmarkEqual/bytes.Equal-12     	 2839177	       417.8 ns/op
PASS
ok  	github.com/xyproto/benchmarked	37.118s
```

The performance is pretty similar for `bytes.Equal` in Go 1.16 and Go 1.17.

With the updated `Equal` function, the performance is better.

## General info

* Version: 0.1.0
* License: BSD
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
