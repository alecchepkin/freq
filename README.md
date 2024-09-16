# Finding word frequency in text

## Aims
### The aim of this exercise is to test your ability to write a working program in Go.
- We are looking for you to demonstrate that you can write good solid code, and debug it to get a good working solution. We consider it preferable to have a good solution that is finished and works rather than half-written code that tries to be too clever.
- Comments in your code will help us understand what you are trying to do in your program. But you can assume we are not stupid.
- We are looking for a pragmatic solution with the right trade-offs between performance and reliability given the time constraints.
- We are interested in allowing you to demonstrate your ability to design and implement data structures and algorithms to solve problems, so use of any language standard container libraries is not permitted (Go maps are not allowed, Go slices are allowed). The use of I/O streams is permitted, the use of C++/Java/Go strings is discouraged.
- You will be asked to explain the design of your solution and talk us through the code.
- Most candidates take between 1 and 5 hours. Please do not spend longer than 6 hours on this task. In any case, please indicate approximately how long you spent writing your solution.

### Languages and Tools
- The reference platform for the task is a Linux 64bit system. If you do not have access to such a platform you should implement it on a Linux or Windows environment of your choosing, however, you should be prepared to explain how your solution may need to be modified to run on the reference platform.
- The solution is to be written entirely in C, C++, Java or Go.
- Your program should also handle binary files (e.g. /boot/vmlinuz) without crashing.

## The Problem
Given the attached text file as an argument, your program will read the file, and output the 20 most frequently used words in the file in order, along with their frequency. The output should be the same to that of the following bash program:
``` bash
#!/usr/bin/env bash
cat $1 | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20
```


```bash
start=$(gdate +%s%N); cat ./testdata/mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20; end=$(gdate +%s%N); execution_time=$(echo "scale=6; ($end - $start) / 1000000" | bc); echo "time execution:${execution_time}ms"
4284 the
2192 and
2185 of
1861 a
1685 to
1366 in
1056 i
1024 that
889 his
821 it
783 he
616 but
603 was
595 with
577 s
564 is
551 for
542 all
541 as
458 at
time execution:60.620000ms

```

## Solution

We are using to store words a trie. [https://en.wikipedia.org/wiki/Trie]

Trie.list is a slice with most frequent words.
During inserting we will update the list.
There is a text.converter ror parsing strings to words.


```bash
go run cmd/freq/main.go testdata/mobydick.txt 
4244 the
2173 of
2169 and
1810 a
1667 to
1358 in
989 i
971 that
886 his
749 it
726 he
600 but
598 was
591 with
558 is
536 for
534 as
527 all
456 at
438 this
time execution:420.754167ms

```

