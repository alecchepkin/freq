#!/usr/bin/env bash

start=$(gdate +%s%N)

cat mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20

end=$(gdate +%s%N)

execution_time=$(echo "scale=6; ($end - $start) / 1000000" | bc)
echo "time execution:${execution_time}ms"