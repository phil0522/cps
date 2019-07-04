/// Einstein's Puzzle, for description and answer, check
/// https://web.stanford.edu/~laurik/fsmbook/examples/Einstein%27sPuzzle.html

package main

import (
	"github.com/phil0522/cps"
)

const (
	kindPos   = "pos"
	kindColor = "color"
	kindDrink = "drink"
	kindCigar = "cigar"
	kindPet   = "pet"

	allPos = "pos*"
)

// Domains
const (
	English   = cps.Domain("English")
	Swede     = cps.Domain("Swede")
	Dane      = cps.Domain("Dane")
	Norwegian = cps.Domain("Norwegian")
	German    = cps.Domain("German")
)

const (
	red    string = "red"
	white  string = "white"
	green  string = "green"
	yellow string = "yellow"
	blue   string = "blue"

	tea    string = "tea"
	coffee string = "coffee"
	milk   string = "milk"
	bier   string = "bier"
	water  string = "water"

	// Cigar
	pallMall   string = "PallMall"
	dunhills   string = "Dunhills"
	blueMaster string = "BlueMaster"
	prince     string = "Prince"
	blend      string = "Blend"

	// Dog and others
	dog   string = "dog"
	bird  string = "bird"
	cat   string = "cat"
	horse string = "horse"
	fish  string = "fish"
)

func posOf(vs *cps.Variables, name string) int {
	return vs.IntPropertyOf(vs.Of(name), kindPos)
}

func posOfDomain(vs *cps.Variables, domain cps.Domain) int {
	return vs.IntPropertyOf(domain, kindPos)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// The fish puzzle
func ex2() {
	problem := cps.NewProblemWithNames(English, Swede, Dane, Norwegian, German)
	problem.AddIntegerVariables(kindPos, 1, 2, 3, 4, 5)
	problem.AddStringVariables(kindColor, red, white, green, yellow, blue)
	problem.AddStringVariables(kindDrink, tea, coffee, milk, bier, water)
	problem.AddStringVariables(kindCigar, pallMall, dunhills, blueMaster, prince, blend)
	problem.AddStringVariables(kindPet, dog, bird, cat, horse, fish)

	problem.AddConstraintRule("rule 1", func(vs *cps.Variables) bool {
		return vs.Of(red) == English
	}, red)

	problem.AddConstraintRule("rule 2", func(vs *cps.Variables) bool {
		return vs.Of(dog) == Swede
	}, dog)

	problem.AddConstraintRule("rule 3", func(vs *cps.Variables) bool {
		return vs.Of(tea) == Dane
	}, tea)

	problem.AddConstraintRule("rule 4", func(vs *cps.Variables) bool {
		return posOf(vs, green) == posOf(vs, white)-1
	}, green, white, allPos)

	problem.AddConstraintRule("rule 5", func(vs *cps.Variables) bool {
		return vs.Of(green) == vs.Of(coffee)
	}, green, coffee)

	problem.AddConstraintRule("rule 6", func(vs *cps.Variables) bool {
		return vs.Of(pallMall) == vs.Of(bird)
	}, pallMall, bird)

	problem.AddConstraintRule("rule 7", func(vs *cps.Variables) bool {
		return vs.Of(yellow) == vs.Of(dunhills)
	}, yellow, dunhills)

	problem.AddConstraintRule("rule 8", func(vs *cps.Variables) bool {
		return posOf(vs, milk) == 3
	}, milk, allPos)

	problem.AddConstraintRule("rule 9", func(vs *cps.Variables) bool {
		return vs.Of("pos*1") == Norwegian
	}, "pos*1")

	problem.AddConstraintRule("rule 10", func(vs *cps.Variables) bool {
		catPos := posOf(vs, cat)
		blendPos := posOf(vs, blend)
		return abs(catPos-blendPos) == 1
	}, blend, cat, allPos)

	problem.AddConstraintRule("rule 11", func(vs *cps.Variables) bool {
		return vs.Of(blueMaster) == vs.Of(bier)
	}, blueMaster, bier)

	problem.AddConstraintRule("rule 12", func(vs *cps.Variables) bool {
		horsePos := posOf(vs, horse)
		dunhillsPos := posOf(vs, dunhills)
		return abs(horsePos)-abs(dunhillsPos) == 1
	}, horse, dunhills, allPos)

	problem.AddConstraintRule("rule 13", func(vs *cps.Variables) bool {
		return vs.Of(prince) == German
	}, prince)

	problem.AddConstraintRule("rule 14", func(vs *cps.Variables) bool {
		norwegianPos := posOfDomain(vs, Norwegian)
		bluePos := posOf(vs, blue)
		return abs(norwegianPos-bluePos) == 1
	}, blue, allPos)

	problem.AddConstraintRule("rule 15", func(vs *cps.Variables) bool {
		blendPos := posOf(vs, blend)
		waterPos := posOf(vs, water)
		return abs(blendPos-waterPos) == 1
	}, blend, water, allPos)

	problem.Solve()
	problem.PrintSolutions()
}
