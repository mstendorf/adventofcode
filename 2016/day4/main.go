package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	roomName string
	sectorID int
	checksum string
}

func (r *Room) isValid() bool {
	name := strings.ReplaceAll(r.roomName, "-", "")
	counts := countChars(name)
	for i, char := range r.checksum {
		// fmt.Println("Char:", string(char), "Counts:", string(counts[i]))
		if char != counts[i] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`([\w-]+)-(\d+)\[(\w+)\]`)
	scanner := bufio.NewScanner(file)
	sum := 0
	validRooms := make([]Room, 0)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		sectorID, _ := strconv.Atoi(matches[2])
		room := Room{roomName: matches[1], sectorID: sectorID, checksum: matches[3]}

		if room.isValid() {
			sum += sectorID
			validRooms = append(validRooms, room)
		}
	}
	fmt.Println("Part1:", sum)

	for _, room := range validRooms {
		decrypted := decryptRoom(room)
		if strings.Contains(decrypted, "north") {
			fmt.Println("Part2:", room.sectorID)
			break
		}
	}
}

func decryptRoom(room Room) string {
	decrypted := ""
	for _, char := range room.roomName {
		if char == '-' {
			decrypted += " "
		} else {
			decrypted += string(rotateChar(char, room.sectorID))
		}
	}
	return decrypted
}

func rotateChar(char rune, times int) rune {
	// if char == ' ' {
	// 	return ' '
	// }
	rotated := char
	for i := 0; i < times; i++ {
		if rotated == 'z' {
			rotated = 'a'
		} else {
			rotated++
		}
	}
	return rotated
}

func countChars(roomCode string) []rune {
	counts := make(map[rune]int)
	for _, char := range roomCode {
		counts[char]++
	}
	keys := make([]rune, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if counts[keys[i]] == counts[keys[j]] {
			return keys[i] < keys[j]
		}
		return counts[keys[i]] > counts[keys[j]]
	})
	return keys
}
