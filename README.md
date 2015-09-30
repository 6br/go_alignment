go-NW/SW/LCS/NWG
===

This package is an implementation of "Needleman-Bunsch" , "Smith-Waterman" , "Gotoh" and "LCS" Algorithm which is used in DNA analysing. 

## Description
DNA is composed by ATGC and when we want to know the degree of similarity between two sequences, we can use this tool to align those sequences and caclucate the score.

## Test-framework
This is tested by "Gospel", which is Behaviour Driven Development Testing Framework and named after a certain movie.
If you want to test this tool, you'll have to install it.

## Usage
./go\_alignment [flags] [int] [options]

flags(option): -c || --config -> write a config file including gap scores and a substitution matrix.

i.e. (config.txt)
               7 1 % d e => gap score is calculated by g(l) = -e * (l - 1) - d.
                1 -1 -1 -1
               -1  1 -1 -1 => substitution matrix. 
               -1 -1  1 -1
               -1 -1 -1  1

int : 0(default) -> NW, 1 -> LCS, 2 -> SW, 3-> Gotoh

options :

i.e. aagt aact -> alignment of "aagt" and "aact"

1.fasta 2.fasta -> alignment of fasta files. 

sequences.txt -> alignment of single fasta files.

## NOTION
Now, you can align only composed by ATGC(or atgc).

## Install
Please git clone. This repository includes garbage files. Please ignore.

