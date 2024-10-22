package main

import "fmt"

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

func main() {
	player_health := 100
	boss_health := 109
	boss_damage := 8
	boss_armor := 2

	weapons := []Item{
		Item{"Dagger", 8, 4, 0},
		Item{"Shortsword", 10, 5, 0},
		Item{"Warhammer", 25, 6, 0},
		Item{"Longsword", 40, 7, 0},
		Item{"Greataxe", 74, 8, 0},
	}

	armor := []Item{
		Item{"Leather", 13, 0, 1},
		Item{"Chainmail", 31, 0, 2},
		Item{"Splintmail", 53, 0, 3},
		Item{"Bandedmail", 75, 0, 4},
		Item{"Platemail", 102, 0, 5},
	}

	rings := []Item{
		Item{"Damage +1", 25, 1, 0},
		Item{"Damage +2", 50, 2, 0},
		Item{"Damage +3", 100, 3, 0},
		Item{"Defense +1", 20, 0, 1},
		Item{"Defense +2", 40, 0, 2},
		Item{"Defense +3", 80, 0, 3},
	}
	part1(player_health, boss_health, boss_damage, boss_armor, weapons, armor, rings)
	part2(player_health, boss_health, boss_damage, boss_armor, weapons, armor, rings)
}

// please dont look at this, i am ashamed
func part1(player_health, boss_health, boss_damage, boss_armor int, weapons, armor, rings []Item) {
	min_cost := 1000
	for _, weapon := range weapons {
		if simulate(player_health, boss_health, boss_damage, boss_armor, weapon.damage, 0) && weapon.cost < min_cost {
			min_cost = weapon.cost
		}
		for _, a := range armor {
			if simulate(player_health, boss_health, boss_damage, boss_armor, weapon.damage, a.armor) && weapon.cost+a.cost < min_cost {
				min_cost = weapon.cost + a.cost
			}
			for _, r1 := range rings {
				cost := weapon.cost + a.cost + r1.cost
				damage := weapon.damage + r1.damage
				armor := a.armor + r1.armor
				if simulate(player_health, boss_health, boss_damage, boss_armor, damage, armor) && cost < min_cost {
					min_cost = cost
				}
				for _, r2 := range rings {
					if r1 == r2 {
						continue
					}
					cost := weapon.cost + a.cost + r1.cost + r2.cost
					damage := weapon.damage + r1.damage + r2.damage
					armor := a.armor + r1.armor + r2.armor
					if simulate(player_health, boss_health, boss_damage, boss_armor, damage, armor) && cost < min_cost {
						min_cost = cost
					}
				}
			}
		}
	}
	fmt.Println(min_cost)
}

func part2(player_health, boss_health, boss_damage, boss_armor int, weapons, armor, rings []Item) {
	max_cost := 0

	for _, weapon := range weapons {
		for _, ring1 := range rings {
			for _, ring2 := range rings {
				if ring1 == ring2 {
					continue
				}
				cost := ring1.cost + ring2.cost + weapon.cost
				damage := ring1.damage + ring2.damage + weapon.damage
				defense := ring1.armor + ring2.armor
				if !simulate(player_health, boss_health, boss_damage, boss_armor, damage, defense) && cost > max_cost {
					max_cost = cost
				}
				for _, armor := range armor {
					cost := ring1.cost + ring2.cost + armor.cost + weapon.cost
					damage := ring1.damage + ring2.damage + armor.damage + weapon.damage
					defense := ring1.armor + ring2.armor + armor.armor
					if !simulate(player_health, boss_health, boss_damage, boss_armor, damage, defense) && cost > max_cost {
						max_cost = cost
					}
				}
			}
		}
	}

	fmt.Println(max_cost)
}

func simulate(player_health, boss_health, boss_damage, boss_armor, damage, armor int) bool {
	player_turn := true
	for player_health > 0 && boss_health > 0 {
		if player_turn {
			boss_health -= max(1, damage-boss_armor)
		} else {
			player_health -= max(1, boss_damage-armor)
		}
		player_turn = !player_turn
	}
	return player_health > 0
}
