#!/bin/sh

N=${1:-2}

for f in 128M 64M 4M 4K 20 16 15 9 6 1 0; do
  # top N
  grep _$f- bench.out | cut -f1,3 | sort -k2 -n | head -n $N | cut -d- -f1 | cut -d/ -f2 | cut -d_ -f1 > best_$f
  # bottom N
  grep _$f- bench.out | cut -f1,3 | sort -r -k2 -n | head -n $N | cut -d- -f1 | cut -d/ -f2 | cut -d_ -f1 > worst_$f

  echo -e "\n--- Best for the $f test:"
  cat best_$f
  echo -e "\n--- Worst for the $f test:"
  cat worst_$f
done

echo -e "\n--- Overall worst"
cat worst_* | sort | uniq -c | sort -r -n

echo -e "\n--- Overall best"
cat best_* | sort | uniq -c | sort -r -n
