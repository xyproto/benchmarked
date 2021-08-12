#!/bin/sh
go test -bench=. "$@" | tee bench.out
head -4 bench.out > tmp.out
grep 'ns/op' bench.out | sort -r -n -k3 >> tmp.out
tail -2 bench.out >> tmp.out
mv -f tmp.out bench.out
