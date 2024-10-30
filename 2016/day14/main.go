package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type PotentialKey struct {
	hash string
	char byte
	idx  int
}

type Hash struct {
	hash string
	idx  int
}

func main() {
	salt := "ngcjuoqr"
	// salt := "abc"

	part1(salt)
	part2(salt)
}

func part1(salt string) {
	hashes := make([]string, 0)
	for i := 0; i < 20000; i++ {
		hash := GetMD5Hash(salt + strconv.Itoa(i))
		hashes = append(hashes, hash)
	}

	fmt.Println("Part 1", findKeys(hashes))
}

func part2(salt string) {
	hashes := make([]string, 0)
	for i := 0; i < 22000; i++ {
		hash := GetMD5Hash(salt + strconv.Itoa(i))
		for j := 0; j < 2016; j++ {
			hash = GetMD5Hash(hash)
		}
		hashes = append(hashes, hash)
	}

	fmt.Println("Part 2", findKeys(hashes))
}

func findKeys(hashes []string) int {

	keyCount := 0
	for i := 0; i < len(hashes); i++ {
		hasTriple, char := hasTriple(hashes[i])

		isKey := false
		if hasTriple {
			for j := i + 1; j <= i+1000; j++ {
				if hasFive(hashes[j], char) {
					isKey = true
					break
				}
			}
		}
		if isKey {
			// fmt.Println("Key", hashes[i], i)
			keyCount++
		}
		if keyCount == 64 {
			// fmt.Println("Part 1", i)
			return i
		}
	}
	return 0
}

func hasTriple(hex string) (bool, byte) {
	for i := 0; i < len(hex)-2; i++ {
		if hex[i] == hex[i+1] && hex[i] == hex[i+2] {
			return true, hex[i]
		}
	}
	return false, '0'
}

func hasFive(hex string, char byte) bool {
	for i := 0; i < len(hex)-4; i++ {
		if hex[i] == char && hex[i] == hex[i+1] && hex[i] == hex[i+2] && hex[i] == hex[i+3] && hex[i] == hex[i+4] {
			return true
		}
	}
	return false
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
