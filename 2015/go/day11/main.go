package main

import "fmt"

func main() {

	input := "cqjxjnds"

	iteration := 0
	for {
		incrementPassword(&input, len(input)-1)
		if isValidPassword(input) {
			iteration++
			fmt.Println(iteration, input)
			if iteration == 2 {
				break
			}
		}
	}
}

func isValidPassword(s string) bool {
	return hasIncreasingStraight(s) && !hasForbiddenChars(s) && hasTwoPairs(s)
}

func hasIncreasingStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			return true
		}
	}
	return false
}

func hasForbiddenChars(s string) bool {
	for _, c := range s {
		if c == 'i' || c == 'o' || c == 'l' {
			return true
		}
	}
	return false
}

func hasTwoPairs(s string) bool {
	pairs := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func replaceAtIndex(in *string, r rune, i int) {
	out := []rune(*in)
	out[i] = r
	*in = string(out)
}

func incrementCharAtIndex(s *string, idx int) {
	out := []rune(*s)
	out[idx]++
	*s = string(out)
}

func incrementPassword(s *string, idx int) {
	if idx < 0 {
		*s = "a" + *s
		return
	}

	if (*s)[idx] == 'z' {
		replaceAtIndex(s, 'a', idx)
		incrementPassword(s, idx-1)
	} else {
		incrementCharAtIndex(s, idx)
	}
}
