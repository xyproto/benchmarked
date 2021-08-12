#!/bin/sh
go test -bench=. "$@" | tee bench.out
head -4 bench.out > sorted.out
grep 'ns/op' bench.out | sort -r -n -k3 >> sorted.out
tail -2 bench.out >> sorted.out
cat sorted.out
