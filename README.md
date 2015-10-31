go-NW/SW/LCS/NWG
===
[![Build Status](https://drone.io/github.com/6br/go_alignment/status.png)](https://drone.io/github.com/6br/go_alignment/latest)

This package is an implementation of "Needleman-Bunsch" , "Smith-Waterman" , "Gotoh" and "LCS" Algorithm which is used in DNA analysing.

## Description
DNA is composed by ATGC and when we want to know the degree of similarity between two sequences, we can use this tool to align those sequences and caclucate the score.

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

* -c || --config -> Write a config file including gap scores and a substitution matrix.
* -i -> Write an interval of alignment result sequences. (default: 50)

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
* 5 -> Linear-Memory Gotoh

options :

* aagt aact -> alignment of "aagt" and "aact"
* 1.fasta 2.fasta -> alignment of fasta files.
* sequences.txt -> alignment of single fasta files.

## NOTION
Now, you can align only composed by ATGC(or atgc).

This repository doesn't obey to best practice of golang. I'm working on.

## Install
Please git clone.
