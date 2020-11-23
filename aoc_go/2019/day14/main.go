package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Ingredient struct {
	Amount int
	Name   string
}

type Reaction struct {
	To   Ingredient
	From []Ingredient
}
type Reactions struct {
	reactions []Reaction
}

func (reactions *Reactions) FindReaction(target string) (Reaction, bool) {
	for _, reaction := range reactions.reactions {
		if reaction.To.Name == target {
			return reaction, true
		}
	}
	return Reaction{}, false
}
func (reactions *Reactions) add(reaction Reaction) {
	reactions.reactions = append(reactions.reactions, reaction)
}

type Stack struct {
	items []Ingredient
}

func (stack *Stack) push(ingredient Ingredient) {
	stack.items = append(stack.items, ingredient)
}
func (stack *Stack) pop() Ingredient {
	ingredient := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return ingredient
}
func (stack *Stack) empty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) Clear() {
	stack.items = []Ingredient{}
}

const TRILLION int = 1000000000000

func SolvePart2(str string, solverFunction func(reactions Reactions, numFuel int) int) int {
	reactions := parseInput(str)
	fuel := 1.0
	runs := 0
	stepsize := 1000000
	for stepsize > 0 {
		runs++
		ore := solverFunction(reactions, int(fuel))
		if float64(ore)/float64(TRILLION) > 1 {
			fuel -= float64(stepsize)
			stepsize /= 2
			fuel += float64(stepsize)
		} else {
			fuel += float64(stepsize)
		}
	}

	return int(fuel)
}

func SolvePart1(str string, solverFunction func(reactions Reactions, numFuel int) int) int {
	reactions := parseInput(str)
	return solverFunction(reactions, 1)
}

func findOreForReactionsWithRecursion(reactions Reactions, numFuel int) int {
	supplies := map[string]int{}
	return recurse(Ingredient{Name: "FUEL", Amount: numFuel}, &reactions, &supplies)
}

func recurse(ingredient Ingredient, reactions *Reactions, supplies *map[string]int) int {
	if ingredient.Name == "ORE" {
		return ingredient.Amount
	}
	if amountInSupply := int(math.Min(float64((*supplies)[ingredient.Name]), float64(ingredient.Amount))); amountInSupply > 0 {
		ingredient.Amount -= amountInSupply
		(*supplies)[ingredient.Name] -= amountInSupply
	}
	if ingredient.Amount == 0 {
		return 0
	}
	ore := 0
	if reaction, found := reactions.FindReaction(ingredient.Name); found {
		multiplier := int(math.Ceil(float64(ingredient.Amount) / float64(reaction.To.Amount)))
		for _, subingredient := range reaction.From {
			subingredient.Amount *= multiplier
			ore += recurse(subingredient, reactions, supplies)
		}
		if extraSupplies := (multiplier * reaction.To.Amount) - ingredient.Amount; extraSupplies > 0 {
			(*supplies)[reaction.To.Name] += extraSupplies
		}
	}
	return ore
}

func findOreForReactionsWithQueue(reactions Reactions, numFuel int) int {
	supplies := map[string]int{}
	var stack Stack
	stack.push(Ingredient{
		Name:   "FUEL",
		Amount: numFuel,
	})
	ore := 0
	for !stack.empty() {
		ingredient := stack.pop()
		if ingredient.Name == "ORE" {
			ore += ingredient.Amount
		} else {
			//Check if we have some first
			if amountInSupply := int(math.Min(float64(supplies[ingredient.Name]), float64(ingredient.Amount))); amountInSupply > 0 {
				ingredient.Amount -= amountInSupply
				supplies[ingredient.Name] -= amountInSupply
			}
			// fmt.Println("Searching for cost of", ingredient.Amount, ingredient.Name)
			if reaction, found := reactions.FindReaction(ingredient.Name); found {
				// fmt.Println("  found that it needs ", reaction.From, "to produce", reaction.To.Amount, reaction.To.Name)
				multiplier := int(math.Ceil(float64(ingredient.Amount) / float64(reaction.To.Amount)))
				// fmt.Println("  So we need to run the reaction", multiplier, "times to produce", reaction.To.Amount*multiplier, reaction.To.Name)
				for _, subingredient := range reaction.From {
					subingredient.Amount = subingredient.Amount * multiplier
					stack.push(subingredient)
				}
				if extraSupplies := (multiplier * reaction.To.Amount) - ingredient.Amount; extraSupplies > 0 {
					supplies[reaction.To.Name] += extraSupplies
					// fmt.Println("We received some extra from this reaction", supplies)
				}
			}
		}
	}

	return ore
}

func parseInput(str string) Reactions {
	var reactions Reactions
	for _, row := range strings.Split(str, "\n") {
		if strings.TrimSpace(row) != "" {
			sides := strings.Split(row, " => ")
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
			reactions.add(Reaction{
				From: from,
				To: Ingredient{
					Amount: toAmount,
					Name:   toName,
				},
			})
		}
	}
	return reactions
}

func readInput(path string) string {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	contents := ""
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}
	return contents
}

func main() {
	/*for _, filename := range []string{""}{
		fmt.Println("Part1", filename, SolvePart1(readInput("aoc_go/2019/day14/test0")))
	}*/
	fmt.Println("findOreForReactionsWithRecursion")
	fmt.Println("Part1 solution:", SolvePart1(readInput("aoc_go/2019/day14/input"), findOreForReactionsWithRecursion), "114125")
	fmt.Println("Part2 solution:", SolvePart2(readInput("aoc_go/2019/day14/input"), findOreForReactionsWithRecursion), "12039407")

	fmt.Println("findOreForReactionsWithQueue")
	fmt.Println("Part1 solution:", SolvePart1(readInput("aoc_go/2019/day14/input"), findOreForReactionsWithQueue), "114125")
	fmt.Println("Part2 solution:", SolvePart2(readInput("aoc_go/2019/day14/input"), findOreForReactionsWithQueue), "12039407")

}
