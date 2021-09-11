# Benchmarked

The quest to find a faster `bytes.Equal` function, for short slices of bytes that are likely to be unequal.

So far, the best performing function is 30% faster than `bytes.Equal`, for random bytes (or very short slices).

`bytes.Equal` is much faster for non-random slices, or for slices that are more likely to be equal.

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

Tested on Arch Linux, using `go version go1.17.1 linux/amd64`.

```
go version go1.17.1 linux/amd64
goos: linux
goarch: amd64
pkg: github.com/xyproto/benchmarked
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkEqual/equal7-12         	 2334206	       521.3 ns/op
BenchmarkEqual/equal7_0-12       	558665079	         2.161 ns/op
BenchmarkEqual/equal7_1-12       	425021043	         2.914 ns/op	 343.22 MB/s
BenchmarkEqual/equal7_6-12       	259927176	         4.502 ns/op	1332.66 MB/s
BenchmarkEqual/equal7_9-12       	213153759	         5.477 ns/op	1643.18 MB/s
BenchmarkEqual/equal7_15-12      	153948649	         7.690 ns/op	1950.60 MB/s
BenchmarkEqual/equal7_16-12      	148206308	         8.126 ns/op	1969.05 MB/s
BenchmarkEqual/equal7_20-12      	137458686	         8.959 ns/op	2232.49 MB/s
BenchmarkEqual/equal7_32-12      	95213845	        12.27 ns/op	2608.44 MB/s
BenchmarkEqual/equal7_4K-12      	  900271	      1267 ns/op	3232.81 MB/s
BenchmarkEqual/equal7_4M-12      	     942	   1305363 ns/op	3213.13 MB/s
BenchmarkEqual/equal7_64M-12     	      50	  20388883 ns/op	3291.44 MB/s
BenchmarkEqual/equal8-12         	 2281830	       522.7 ns/op
BenchmarkEqual/equal8_0-12       	558422462	         2.172 ns/op
BenchmarkEqual/equal8_1-12       	416260861	         2.897 ns/op	 345.24 MB/s
BenchmarkEqual/equal8_6-12       	268363809	         4.477 ns/op	1340.10 MB/s
BenchmarkEqual/equal8_9-12       	223445011	         5.570 ns/op	1615.70 MB/s
BenchmarkEqual/equal8_15-12      	153536294	         7.836 ns/op	1914.12 MB/s
BenchmarkEqual/equal8_16-12      	146779784	         8.082 ns/op	1979.74 MB/s
BenchmarkEqual/equal8_20-12      	135378856	         8.841 ns/op	2262.29 MB/s
BenchmarkEqual/equal8_32-12      	96787196	        12.10 ns/op	2644.44 MB/s
BenchmarkEqual/equal8_4K-12      	  964761	      1258 ns/op	3254.79 MB/s
BenchmarkEqual/equal8_4M-12      	     934	   1273029 ns/op	3294.74 MB/s
BenchmarkEqual/equal8_64M-12     	      56	  20785395 ns/op	3228.65 MB/s
BenchmarkEqual/equal9-12         	 1907461	       622.1 ns/op
BenchmarkEqual/equal9_0-12       	535467026	         2.269 ns/op
BenchmarkEqual/equal9_1-12       	463374744	         2.615 ns/op	 382.35 MB/s
BenchmarkEqual/equal9_6-12       	222498549	         5.368 ns/op	1117.75 MB/s
BenchmarkEqual/equal9_9-12       	216767678	         5.557 ns/op	1619.56 MB/s
BenchmarkEqual/equal9_15-12      	215504280	         5.544 ns/op	2705.73 MB/s
BenchmarkEqual/equal9_16-12      	216054140	         5.528 ns/op	2894.25 MB/s
BenchmarkEqual/equal9_20-12      	199416088	         6.078 ns/op	3290.29 MB/s
BenchmarkEqual/equal9_32-12      	173851706	         6.818 ns/op	4693.73 MB/s
BenchmarkEqual/equal9_4K-12      	17150572	        73.87 ns/op	55451.38 MB/s
BenchmarkEqual/equal9_4M-12      	   15087	     79908 ns/op	52489.18 MB/s
BenchmarkEqual/equal9_64M-12     	     938	   1428522 ns/op	46977.83 MB/s
BenchmarkEqual/equal11-12        	 1920991	       632.3 ns/op
BenchmarkEqual/equal11_0-12      	542482678	         2.250 ns/op
BenchmarkEqual/equal11_1-12      	453833128	         2.547 ns/op	 392.61 MB/s
BenchmarkEqual/equal11_6-12      	332427776	         3.681 ns/op	1630.05 MB/s
BenchmarkEqual/equal11_9-12      	217858291	         5.488 ns/op	1640.07 MB/s
BenchmarkEqual/equal11_15-12     	216759604	         5.522 ns/op	2716.48 MB/s
BenchmarkEqual/equal11_16-12     	220136926	         5.517 ns/op	2899.97 MB/s
BenchmarkEqual/equal11_20-12     	201039819	         6.042 ns/op	3310.27 MB/s
BenchmarkEqual/equal11_32-12     	173880058	         6.805 ns/op	4702.19 MB/s
BenchmarkEqual/equal11_4K-12     	17682602	        72.72 ns/op	56328.52 MB/s
BenchmarkEqual/equal11_4M-12     	   12907	     93974 ns/op	44632.61 MB/s
BenchmarkEqual/equal11_64M-12    	     816	   1351236 ns/op	49664.80 MB/s
BenchmarkEqual/bytes.Equal-12    	 2030947	       582.5 ns/op
BenchmarkEqual/bytes.Equal_0-12  	241109206	         4.950 ns/op
BenchmarkEqual/bytes.Equal_1-12  	262490774	         4.622 ns/op	 216.35 MB/s
BenchmarkEqual/bytes.Equal_6-12  	259128402	         4.565 ns/op	1314.48 MB/s
BenchmarkEqual/bytes.Equal_9-12  	247350720	         4.896 ns/op	1838.22 MB/s
BenchmarkEqual/bytes.Equal_15-12 	251240676	         4.823 ns/op	3110.13 MB/s
BenchmarkEqual/bytes.Equal_16-12 	245152966	         4.938 ns/op	3240.36 MB/s
BenchmarkEqual/bytes.Equal_20-12 	217075081	         5.572 ns/op	3589.68 MB/s
BenchmarkEqual/bytes.Equal_32-12 	185060070	         6.526 ns/op	4903.41 MB/s
BenchmarkEqual/bytes.Equal_4K-12 	17811034	        71.99 ns/op	56896.26 MB/s
BenchmarkEqual/bytes.Equal_4M-12 	   13807	     87231 ns/op	48082.81 MB/s
BenchmarkEqual/bytes.Equal_64M-12         	     768	   1450212 ns/op	46275.21 MB/s
BenchmarkEqual/equal2-12                  	 2247710	       522.6 ns/op
BenchmarkEqual/equal2_0-12                	619726996	         1.962 ns/op
BenchmarkEqual/equal2_1-12                	426232502	         2.840 ns/op	 352.06 MB/s
BenchmarkEqual/equal2_6-12                	264056676	         4.537 ns/op	1322.42 MB/s
BenchmarkEqual/equal2_9-12                	215042818	         5.670 ns/op	1587.21 MB/s
BenchmarkEqual/equal2_15-12               	146994534	         8.159 ns/op	1838.51 MB/s
BenchmarkEqual/equal2_16-12               	100000000	        10.70 ns/op	1495.68 MB/s
BenchmarkEqual/equal2_20-12               	100000000	        10.10 ns/op	1980.03 MB/s
BenchmarkEqual/equal2_32-12               	80992852	        14.94 ns/op	2142.02 MB/s
BenchmarkEqual/equal2_4K-12               	  776301	      1548 ns/op	2646.17 MB/s
BenchmarkEqual/equal2_4M-12               	     805	   1511476 ns/op	2774.97 MB/s
BenchmarkEqual/equal2_64M-12              	      48	  24528166 ns/op	2735.99 MB/s
BenchmarkEqual/equal3-12                  	 2265630	       543.8 ns/op
BenchmarkEqual/equal3_0-12                	631188696	         1.934 ns/op
BenchmarkEqual/equal3_1-12                	484716854	         2.530 ns/op	 395.30 MB/s
BenchmarkEqual/equal3_6-12                	359801810	         3.378 ns/op	1776.21 MB/s
BenchmarkEqual/equal3_9-12                	260861240	         4.602 ns/op	1955.75 MB/s
BenchmarkEqual/equal3_15-12               	190626165	         6.361 ns/op	2358.08 MB/s
BenchmarkEqual/equal3_16-12               	180260571	         6.691 ns/op	2391.15 MB/s
BenchmarkEqual/equal3_20-12               	145482079	         8.207 ns/op	2437.06 MB/s
BenchmarkEqual/equal3_32-12               	92331108	        12.58 ns/op	2543.57 MB/s
BenchmarkEqual/equal3_4K-12               	  796726	      1497 ns/op	2737.01 MB/s
BenchmarkEqual/equal3_4M-12               	     798	   1522812 ns/op	2754.32 MB/s
BenchmarkEqual/equal3_64M-12              	      49	  24303025 ns/op	2761.34 MB/s
BenchmarkEqual/equal5-12                  	 2447942	       503.4 ns/op
BenchmarkEqual/equal5_0-12                	564325531	         2.153 ns/op
BenchmarkEqual/equal5_1-12                	411798536	         2.948 ns/op	 339.23 MB/s
BenchmarkEqual/equal5_6-12                	201277873	         6.028 ns/op	 995.28 MB/s
BenchmarkEqual/equal5_9-12                	163258918	         7.704 ns/op	1168.25 MB/s
BenchmarkEqual/equal5_15-12               	100000000	        10.33 ns/op	1451.70 MB/s
BenchmarkEqual/equal5_16-12               	100000000	        10.73 ns/op	1491.82 MB/s
BenchmarkEqual/equal5_20-12               	89721282	        12.60 ns/op	1587.52 MB/s
BenchmarkEqual/equal5_32-12               	59720516	        18.25 ns/op	1753.59 MB/s
BenchmarkEqual/equal5_4K-12               	  602617	      1940 ns/op	2111.87 MB/s
BenchmarkEqual/equal5_4M-12               	     612	   1976828 ns/op	2121.73 MB/s
BenchmarkEqual/equal5_64M-12              	      36	  31586882 ns/op	2124.58 MB/s
BenchmarkEqual/equal6-12                  	 2448990	       496.6 ns/op
BenchmarkEqual/equal6_0-12                	718796142	         1.714 ns/op
BenchmarkEqual/equal6_1-12                	316314543	         3.840 ns/op	 260.44 MB/s
BenchmarkEqual/equal6_6-12                	207891271	         5.716 ns/op	1049.70 MB/s
BenchmarkEqual/equal6_9-12                	205740954	         5.893 ns/op	1527.30 MB/s
BenchmarkEqual/equal6_15-12               	163198696	         7.361 ns/op	2037.87 MB/s
BenchmarkEqual/equal6_16-12               	158763963	         7.536 ns/op	2123.08 MB/s
BenchmarkEqual/equal6_20-12               	139646366	         8.605 ns/op	2324.28 MB/s
BenchmarkEqual/equal6_32-12               	97905802	        11.94 ns/op	2679.23 MB/s
BenchmarkEqual/equal6_4K-12               	  923911	      1229 ns/op	3331.83 MB/s
BenchmarkEqual/equal6_4M-12               	     958	   1252183 ns/op	3349.59 MB/s
BenchmarkEqual/equal6_64M-12              	      57	  19893001 ns/op	3373.49 MB/s
BenchmarkEqual/equal10-12                 	 1985851	       613.8 ns/op
BenchmarkEqual/equal10_0-12               	539994705	         2.195 ns/op
BenchmarkEqual/equal10_1-12               	479189469	         2.553 ns/op	 391.76 MB/s
BenchmarkEqual/equal10_6-12               	227510841	         5.162 ns/op	1162.37 MB/s
BenchmarkEqual/equal10_9-12               	222745418	         5.406 ns/op	1664.86 MB/s
BenchmarkEqual/equal10_15-12              	221719364	         5.437 ns/op	2758.74 MB/s
BenchmarkEqual/equal10_16-12              	219725887	         5.347 ns/op	2992.55 MB/s
BenchmarkEqual/equal10_20-12              	201677560	         5.948 ns/op	3362.35 MB/s
BenchmarkEqual/equal10_32-12              	177628647	         6.835 ns/op	4681.87 MB/s
BenchmarkEqual/equal10_4K-12              	17965774	        71.39 ns/op	57378.82 MB/s
BenchmarkEqual/equal10_4M-12              	   13089	     91528 ns/op	45825.57 MB/s
BenchmarkEqual/equal10_64M-12             	     896	   1389237 ns/op	48306.26 MB/s
BenchmarkEqual/equal14-12                 	 1820478	       647.0 ns/op
BenchmarkEqual/equal14_0-12               	545735204	         2.166 ns/op
BenchmarkEqual/equal14_1-12               	474083642	         2.571 ns/op	 388.91 MB/s
BenchmarkEqual/equal14_6-12               	194126050	         6.136 ns/op	 977.83 MB/s
BenchmarkEqual/equal14_9-12               	172116336	         6.980 ns/op	1289.34 MB/s
BenchmarkEqual/equal14_15-12              	171179192	         7.027 ns/op	2134.53 MB/s
BenchmarkEqual/equal14_16-12              	171069920	         6.960 ns/op	2299.01 MB/s
BenchmarkEqual/equal14_20-12              	158673866	         7.558 ns/op	2646.07 MB/s
BenchmarkEqual/equal14_32-12              	154551900	         7.891 ns/op	4055.06 MB/s
BenchmarkEqual/equal14_4K-12              	17288450	        72.89 ns/op	56193.65 MB/s
BenchmarkEqual/equal14_4M-12              	   10000	    103520 ns/op	40516.72 MB/s
BenchmarkEqual/equal14_64M-12             	     897	   1380290 ns/op	48619.39 MB/s
BenchmarkEqual/equal12-12                 	 1923289	       621.5 ns/op
BenchmarkEqual/equal12_0-12               	552239088	         2.211 ns/op
BenchmarkEqual/equal12_1-12               	472473561	         2.572 ns/op	 388.85 MB/s
BenchmarkEqual/equal12_6-12               	360257554	         3.373 ns/op	1778.57 MB/s
BenchmarkEqual/equal12_9-12               	215363626	         5.593 ns/op	1609.22 MB/s
BenchmarkEqual/equal12_15-12              	213708351	         5.649 ns/op	2655.51 MB/s
BenchmarkEqual/equal12_16-12              	211506676	         5.641 ns/op	2836.57 MB/s
BenchmarkEqual/equal12_20-12              	194994757	         6.215 ns/op	3218.01 MB/s
BenchmarkEqual/equal12_32-12              	169220380	         7.024 ns/op	4556.10 MB/s
BenchmarkEqual/equal12_4K-12              	17733861	        71.63 ns/op	57181.98 MB/s
BenchmarkEqual/equal12_4M-12              	   13513	     89113 ns/op	47066.99 MB/s
BenchmarkEqual/equal12_64M-12             	     855	   1384961 ns/op	48455.41 MB/s
BenchmarkEqual/equal15-12                 	 1882747	       635.0 ns/op
BenchmarkEqual/equal15_0-12               	559555354	         2.214 ns/op
BenchmarkEqual/equal15_1-12               	463593916	         2.534 ns/op	 394.68 MB/s
BenchmarkEqual/equal15_6-12               	243634449	         5.005 ns/op	1198.74 MB/s
BenchmarkEqual/equal15_9-12               	214986054	         5.662 ns/op	1589.47 MB/s
BenchmarkEqual/equal15_15-12              	214355602	         5.649 ns/op	2655.43 MB/s
BenchmarkEqual/equal15_16-12              	211211875	         5.670 ns/op	2822.07 MB/s
BenchmarkEqual/equal15_20-12              	192436266	         6.174 ns/op	3239.40 MB/s
BenchmarkEqual/equal15_32-12              	169811239	         7.013 ns/op	4563.07 MB/s
BenchmarkEqual/equal15_4K-12              	17921160	        70.51 ns/op	58092.05 MB/s
BenchmarkEqual/equal15_4M-12              	   13153	    101313 ns/op	41399.57 MB/s
BenchmarkEqual/equal15_64M-12             	     915	   1339161 ns/op	50112.63 MB/s
BenchmarkEqual/equal1-12                  	 2042114	       580.4 ns/op
BenchmarkEqual/equal1_0-12                	277536656	         4.414 ns/op
BenchmarkEqual/equal1_1-12                	240858116	         5.063 ns/op	 197.51 MB/s
BenchmarkEqual/equal1_6-12                	237483036	         4.976 ns/op	1205.79 MB/s
BenchmarkEqual/equal1_9-12                	230119875	         5.268 ns/op	1708.52 MB/s
BenchmarkEqual/equal1_15-12               	227345976	         5.242 ns/op	2861.63 MB/s
BenchmarkEqual/equal1_16-12               	231247798	         5.162 ns/op	3099.86 MB/s
BenchmarkEqual/equal1_20-12               	207778452	         5.753 ns/op	3476.39 MB/s
BenchmarkEqual/equal1_32-12               	177478395	         6.697 ns/op	4778.26 MB/s
BenchmarkEqual/equal1_4K-12               	17718289	        72.11 ns/op	56801.56 MB/s
BenchmarkEqual/equal1_4M-12               	   13119	     89140 ns/op	47052.85 MB/s
BenchmarkEqual/equal1_64M-12              	     854	   1387512 ns/op	48366.31 MB/s
BenchmarkEqual/equal4-12                  	 2050443	       583.5 ns/op
BenchmarkEqual/equal4_0-12                	456952554	         2.666 ns/op
BenchmarkEqual/equal4_1-12                	462868502	         2.668 ns/op	 374.76 MB/s
BenchmarkEqual/equal4_6-12                	338981548	         3.593 ns/op	1669.88 MB/s
BenchmarkEqual/equal4_9-12                	221671224	         5.478 ns/op	1642.94 MB/s
BenchmarkEqual/equal4_15-12               	176277367	         6.882 ns/op	2179.65 MB/s
BenchmarkEqual/equal4_16-12               	169787241	         7.097 ns/op	2254.55 MB/s
BenchmarkEqual/equal4_20-12               	146820382	         8.260 ns/op	2421.38 MB/s
BenchmarkEqual/equal4_32-12               	95709738	        12.15 ns/op	2634.14 MB/s
BenchmarkEqual/equal4_4K-12               	  967862	      1224 ns/op	3346.26 MB/s
BenchmarkEqual/equal4_4M-12               	     951	   1245648 ns/op	3367.17 MB/s
BenchmarkEqual/equal4_64M-12              	      56	  19965131 ns/op	3361.30 MB/s
PASS
ok  	github.com/xyproto/benchmarked	291.584s
```

Currently, `equal14` is the function that is exported as `Equal` in the `benchmarked` package, but `equal9` is sometimes slight faster.


## Accuracy

I am aware that perfect benchmarking is a tricky.

Please let me know if you have improvements to how the functions are benchmarked, or how the benchmarks are interpreted!


## General info

* Version: 0.1.0
* License: BSD
