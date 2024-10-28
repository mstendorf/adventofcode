package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"slices"
	"strings"
)

func main() {
	door := "ojvtpuvg"
	// door := "abc"

	part1(door)
	part2(door)
}

func part1(door string) {

	password := make([]string, 0)
	for i := 0; true; i++ {
		value := fmt.Sprintf("%s%d", door, i)

		md5 := GetMD5Hash(value)
		if md5[:5] == "00000" {
			password = append(password, string(md5[5]))
			out := "\rPart 1: " + strings.Join(password, "")
			fmt.Printf(out)
		}
		if len(password) == 8 {
			fmt.Println()
			break
		}
	}
	// fmt.Println("Part 1:", strings.Join(password, ""))
}

func part2(door string) {

	password := make([]string, 8)
	for i := 0; true; i++ {
		value := fmt.Sprintf("%s%d", door, i)

		md5 := GetMD5Hash(value)
		if md5[:5] == "00000" {
			position := md5[5]
			if position >= '0' && position <= '7' {
				index := position - '0'
				if password[index] == "" {
					password[index] = string(md5[6])
					current := ""
					for _, p := range password {
						if p == "" {
							current += "."
						} else {
							current += p
						}
					}
					out := "\rPart 2: " + current
					fmt.Printf(out)
				}
			}
		}
		if !slices.Contains(password, "") {
			fmt.Println()
			break
		}
	}
	// fmt.Println("Part 2:", strings.Join(password, ""))
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
