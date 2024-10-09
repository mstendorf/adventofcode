package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	secret := "yzbqklnj"

	found := false
	for i := 0; i < 10000000; i++ {
		inp := fmt.Sprintf("%s%d", secret, i)
		hash := GetMD5Hash(inp)
		// fmt.Println(inp)
		// fmt.Println(hash)

		if !found && hash[:5] == "00000" {
			println("Five 0's", i)
			found = true
			// break
		}
		if hash[:6] == "000000" {
			println("Six 0's", i)
			break
		}
	}

}
