package main
/*
* Author: Stephen Chavez
* WWW: dicesoft.net
* Date: June 17, 2015
* License: Public domain / Do whatever you want.
*
* This program outputs some plaintext matching MD5 hashes that you give to the
* program. There are two modes you can use with the cracker. You can crack one
* MD5 hash or multiple MD5 hashes. For the first mode, just run the program with
* no arguments. Or provide two file paths for a set of hashes you wish to crack
* and the word list file to crack the hashes with.
*
*/

	import (
		"fmt"
		"os"
		"bufio"
		"crypto/md5"
		"encoding/hex"
		"time"
		"strings"
	)


	func main() {
		numOfArgs := len(os.Args)

		switch numOfArgs {
		case 1:
			crackSingleMD5Hash()
		case 3:
			crackLots()		
		default:
			printHelp()
		}
	}

	func crackLots() {
		var hashFileTextLine string
		var dictFileTextLine string
		foundHash := false

		hashFilePath := os.Args[1]
		dictFilePath := os.Args[2]

		hashFile, hashFileErr := os.Open(hashFilePath)
		dictFile, dictFileErr := os.Open(dictFilePath)
		

		if (hashFileErr != nil) {
			fmt.Printf("Can't read file: %s", hashFilePath)
		}
		
		if (dictFileErr != nil) {
			fmt.Printf("Can't read file: %s", dictFilePath)
		}

		defer hashFile.Close()
		defer dictFile.Close()

		hashFileScanner := bufio.NewScanner(hashFile)
		hashFileScanner.Split(bufio.ScanLines)
		
		dictFileScanner := bufio.NewScanner(dictFile)
		dictFileScanner.Split(bufio.ScanLines)
		

		fmt.Println("----------------------------------------")

		fmt.Printf("\t \t Scanning %d hashes \n", countLinesInFile(hashFilePath))

		for hashFileScanner.Scan() {
			hashFileTextLine = strings.ToLower(hashFileScanner.Text())
			fmt.Printf("Finding a word for %s \n\n", hashFileTextLine)

			for dictFileScanner.Scan() {
				dictFileTextLine = dictFileScanner.Text()
				dictFileHash := getMD5HashForString(dictFileTextLine)

				if (dictFileHash == hashFileTextLine ) {
					foundHash = true
					break
				}
			}

			if foundHash {
				fmt.Printf("    Hash %s matched %s \n", hashFileTextLine, dictFileTextLine)
			} else {
				fmt.Printf("No string in the dictionary file matched" +
						" the hash: %s", hashFileTextLine)
			}
			
			// We must reset the file for the wordlist.
			dictFile.Seek(0, 0)
			foundHash = false
		}
		fmt.Println("----------------------------------------")
	}

	func countLinesInFile(filePath string) int {
		countFile, countFileErr := os.Open(filePath)
		numberOfLines := 0

		if countFileErr != nil {
			fmt.Printf("Can't read file: %s", countFile)
		}

		countFileScanner := bufio.NewScanner(countFile)
		countFileScanner.Split(bufio.ScanLines)

		defer countFile.Close()

		for countFileScanner.Scan() {
			numberOfLines += 1
		}

		return numberOfLines
	}
	
	func printHelp() {

usage := `Simple MD5 cracker.

Usage:
    md5cracker
    md5cracker <HashsFilePath> <WordListFilePath>

This program outputs some plaintext matching MD5 hashes that you give to the
program. There are two modes you can use with the cracker. You can crack one MD5
hash or multiple MD5 hashes. For the first mode, just run the program with no
arguments. Or provide two file paths for a set of hashes you wish to crack and
the word list file to crack the hashes with.`

fmt.Println(usage)

	}

	func crackSingleMD5Hash() {
		fmt.Println("Enter a MD5 hash string...")
		userText := getUserInput()

		fmt.Println("Dict file path?")
		userDictFile := getUserInput()
		
		fmt.Println()
		fmt.Println()
		fmt.Println("----------------------------------------")
		
		tStart := time.Now()
		foundHash, findHashErr := findHash(userDictFile, userText)
		tEnd := time.Now()
		
		if (findHashErr != nil) {
			fmt.Print(findHashErr)
		} else {
			fmt.Printf("Found string: %s \t\t", foundHash)
			fmt.Printf("The search took %v seconds to run.\n", tEnd.Sub(tStart))
		}

		fmt.Println("----------------------------------------")
	}

	func getMD5HashForString(userString string) string {
		hasher := md5.New()
		hasher.Write([]byte(userString))
	
		return hex.EncodeToString(hasher.Sum(nil))
	}

	func getUserInput() string {
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')

		return strings.ToLower(strings.TrimSpace(input))
	}


	func findHash(path string, hashToCrack string) (string, error) {
		inFile, ioErr := os.Open(path)
		foundHash := false
		

		if (ioErr != nil) {
			return "", fmt.Errorf("Can't read file: %s", path)
		}

		defer inFile.Close()
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)
		
		var fileTextLine string


		for scanner.Scan() {
			fileTextLine = scanner.Text()
			fileHash := getMD5HashForString(fileTextLine)

			if (fileHash == hashToCrack ) {
				foundHash = true
				break
			}
		}
		
		if foundHash {
			return fileTextLine, nil
		} else {
			return "", fmt.Errorf("no string in the dict file matched" +
						" the hash: %s", hashToCrack)
		}
	}

