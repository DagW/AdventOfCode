package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	Amount int
	Name   string
}

type Reaction struct {
	To   Ingredient
	From []Ingredient
}
type BaseCost struct {
	To   Ingredient
	From Ingredient
}

func RecurseRequirements(
	requirements map[string]int,
	reactions []Reaction,
	ingredientName string,
	multiplier int,
) (map[string]int, bool) {
	reaction, found := FindReaction(reactions, ingredientName)
	if !found {
		return requirements, false
	}
	rest := multiplier % reaction.To.Amount
	multiplier = multiplier / reaction.To.Amount
	if rest != 0 {
		multiplier += 1
	}

	for _, reactionIngredient := range reaction.From {
		newReq, success := RecurseRequirements(
			requirements, reactions,
			reactionIngredient.Name,
			reactionIngredient.Amount*multiplier,
		)
		if success {
			requirements = newReq
		} else {
			//If this produces more than we need? How do we tell?
			fmt.Println("Added",reactionIngredient.Amount * multiplier,reactionIngredient.Name)
			requirements[reactionIngredient.Name] += reactionIngredient.Amount * multiplier
		}
	}
	return requirements, true
}

func SolvePart1(str string) int {
	baseCosts, reactions := parseInput(str)
	fmt.Println("baseCosts", baseCosts)
	fmt.Println("reactions", reactions)
	// Collect the full requirements (All costs, ex 27xA, 2xB etc)
	requirements := map[string]int{}
	requirements, _ = RecurseRequirements(
		requirements, reactions,
		"FUEL", 1,
	)
	// Calculate the cost in ore
	//excess_resources := map[string]int{}
	fmt.Println(requirements)
	ore := 0
	for k, v := range requirements {
		baseCost, _ := FindBaseCost(baseCosts, k)
		times := v / baseCost.To.Amount
		rest := v % baseCost.To.Amount
		if rest > 0 {
			times++
		}
		ore += baseCost.From.Amount * times
		fmt.Println("Cost of", v, k, "->", baseCost.From.Amount*times, "ORE")
	}
	return ore
}

func FindReaction(reactions []Reaction, target string) (Reaction, bool) {
	for _, reaction := range reactions {
		if reaction.To.Name == target {
			return reaction, true
		}
	}
	return Reaction{}, false
}
func FindBaseCost(baseCosts []BaseCost, target string) (BaseCost, bool) {
	for _, baseCost := range baseCosts {
		if baseCost.To.Name == target {
			return baseCost, true
		}
	}
	return BaseCost{}, false
}

func parseInput(str string) ([]BaseCost, []Reaction) {
	var baseCosts []BaseCost
	var reactions []Reaction
	for _, row := range strings.Split(str, "\n") {
		if strings.TrimSpace(row) != "" {
			sides := strings.Split(row, " => ")
			if strings.Contains(sides[0], "ORE") {
				toAmount, _ := strconv.Atoi(strings.Split(sides[1], " ")[0])
				toName := strings.Split(sides[1], " ")[1]
				fromAmount, _ := strconv.Atoi(strings.Split(sides[0], " ")[0])
				fromName := strings.Split(sides[0], " ")[1]
				baseCosts = append(baseCosts, BaseCost{
					From: Ingredient{
						Amount: fromAmount,
						Name:   fromName,
					},
					To: Ingredient{
						Amount: toAmount,
						Name:   toName,
					},
				})
			} else {
				var from []Ingredient
				for _, val := range strings.Split(sides[0], ", ") {
					fromName := strings.Split(val, " ")[1]
					fromAmount, _ := strconv.Atoi(strings.Split(val, " ")[0])
					from = append(from, Ingredient{
						Amount: fromAmount,
						Name:   fromName,
					})
				}
				toAmount, _ := strconv.Atoi(strings.Split(sides[1], " ")[0])
				toName := strings.Split(sides[1], " ")[1]
				reactions = append(reactions, Reaction{
					From: from,
					To: Ingredient{
						Amount: toAmount,
						Name:   toName,
					},
				})
			}
		}
	}
	return baseCosts, reactions
}
func main() {
	file, err := os.Open("aoc_go/2019/day14/test4")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	fmt.Println("Part1 solution:", SolvePart1(input))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
