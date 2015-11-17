go-NW/SW/LCS/NWG/MEA
===
[![Build Status](https://drone.io/github.com/6br/go_alignment/status.png)](https://drone.io/github.com/6br/go_alignment/latest)

This package is an implementation of "Needleman-Bunsch" , "Smith-Waterman" , "Gotoh" and "LCS" Algorithm which is used in DNA analysing.

## Description
DNA is composed by ATGC and when we want to know the degree of similarity between two sequences, we can use this tool to align those sequences and calculate the score.

## Test-framework
This is tested by "Gospel", which is Behaviour Driven Development Testing Framework and named after a certain movie.
If you want to test this tool, you'll have to install it.

```bash
go get github.com/r7kamura/gospel
cd src
go test -v
```

## Usage
```bash
./go_alignment [flags] [int] [options]
go run main.go [flags] [int] [options]
```

flags(option):

* -c || --config -> Write a configure file including gap scores and a substitution matrix.
* -i -> Write an interval of alignment result sequences. (default: 50)
* -d -> Set if you want to examine memory usage.(Access http://localhost:6060/debug/pprof/heap?debug=1)

i.e. (config.txt)
```text
7 1 % d e   #=> gap score is calculated by g(l) = -e * (l - 1) - d.
 1 -1 -1 -1
-1  1 -1 -1 #=> substitution matrix.
-1 -1  1 -1
-1 -1 -1  1
```

int :

* 0(default) -> NW
* 1 -> LCS
* 2 -> SW
* 3 -> Gotoh
* 4 -> MEA
* 5 -> Pseudo Linear-Memory Gotoh(For debug)
* 6 -> Real Linear-Memory Gotoh

options :

* aagt aact -> alignment of "aagt" and "aact"
* 1.fasta 2.fasta -> alignment of fasta files.
* sequences.txt -> alignment of single fasta files.

## NOTION
Now, you can align only composed by ATGC(or atgc).

I have recognized it has a lot of issue such as below.
* This repository doesn't obey to best practice of golang.
* Test's coverage is not 100%.

## Install
Please git clone.
