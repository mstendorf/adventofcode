package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Cookie struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	ingredients := builIngredientList(file)
	part1(ingredients)
	part2(ingredients)
}

func part1(ingredients []Ingredient) {
	spoons := make(map[Ingredient]int)
	for _, ingredient := range ingredients {
		spoons[ingredient] = 0
	}
	maxScore := 0
	for i := 1; i < 101; i++ {
		spoons[ingredients[0]] = i
		for j := 1; i+j < 101; j++ {
			spoons[ingredients[1]] = j
			for k := 1; i+j+k < 101; k++ {
				spoons[ingredients[2]] = k
				remainder := 100 - i - j - k
				spoons[ingredients[3]] = remainder
				if i+j+k+remainder == 100 {
					total, _ := getTotal(spoons)
					if total > maxScore {
						maxScore = total
					}
				}
			}
		}
	}
	fmt.Println(maxScore)
}

func part2(ingredients []Ingredient) {
	spoons := make(map[Ingredient]int)
	for _, ingredient := range ingredients {
		spoons[ingredient] = 0
	}
	maxScore := 0
	for i := 1; i < 101; i++ {
		spoons[ingredients[0]] = i
		for j := 1; i+j < 101; j++ {
			spoons[ingredients[1]] = j
			for k := 1; i+j+k < 101; k++ {
				spoons[ingredients[2]] = k
				remainder := 100 - i - j - k
				spoons[ingredients[3]] = remainder
				if i+j+k+remainder == 100 {
					total, calories := getTotal(spoons)

					if calories == 500 && total > maxScore {
						maxScore = total
					}
				}
			}
		}
	}
	fmt.Println(maxScore)
}

func getTotal(spoons map[Ingredient]int) (int, int) {
	cookie := Cookie{0, 0, 0, 0, 0}
	for ingredient, amount := range spoons {
		cookie.capacity += ingredient.capacity * amount
		cookie.durability += ingredient.durability * amount
		cookie.flavor += ingredient.flavor * amount
		cookie.texture += ingredient.texture * amount
		cookie.calories += ingredient.calories * amount
	}
	if cookie.capacity < 0 {
		cookie.capacity = 0
	}
	if cookie.durability < 0 {
		cookie.durability = 0
	}
	if cookie.flavor < 0 {
		cookie.flavor = 0
	}
	if cookie.texture < 0 {
		cookie.texture = 0
	}
	return cookie.capacity * cookie.durability * cookie.flavor * cookie.texture, cookie.calories
}

func builIngredientList(file *os.File) []Ingredient {
	file.Seek(0, 0)
	ingredients := make([]Ingredient, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		capacity, _ := strconv.Atoi(strings.Trim(parts[2], ","))
		durability, _ := strconv.Atoi(strings.Trim(parts[4], ","))
		flavor, _ := strconv.Atoi(strings.Trim(parts[6], ","))
		texture, _ := strconv.Atoi(strings.Trim(parts[8], ","))
		calories, _ := strconv.Atoi(strings.Trim(parts[10], ","))
		ingredients = append(ingredients, Ingredient{name, capacity, durability, flavor, texture, calories})
	}
	return ingredients
}
