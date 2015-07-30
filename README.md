MD5 cracker program in Go.  
=================

Abstract: 
---------

This program bruteforces MD5 hashes.

Description of the program:
----------------------------
```
Simple MD5 cracker.

Usage:
    md5cracker
    md5cracker <HashsFilePath> <WordListFilePath>

This program outputs some plaintext matching MD5 hashes that you give to the
program. There are two modes you can use with the cracker. You can crack one MD5
hash or multiple MD5 hashes. For the first mode, just run the program with no
arguments. Or provide two file paths for a set of hashes you wish to crack and
the word list file to crack the hashes with.
```

How to install on your system
-----------------------------

1. Setup Go, you can read how [here](https://golang.org/doc/install)
2. run `go get github.com/redragonx/md5-cracker-go`
3. cd into the src dir and run `go install`
4. If done properly, you can run the program anywhere.
