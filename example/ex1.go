// Ex1: https://www.brainzilla.com/logic/logic-grid/thanksgiving-dinner/
package main

import (
	"fmt"

	"github.com/phil0522/cps"
)

func ageOf(vs *cps.Variables, name string) int {
	return vs.IntPropertyOf(vs.Of(name), PropertyAge)
}

const (
	// PropertyAge is property for age
	PropertyAge = "age"

	// AllAges is name that includes all name of age variables.
	AllAges = "age*"
)

func ex1() {

	const (
		// Domains
		Larry   = "Larry"
		Nicolas = "Nicholas"
		Philip  = "Philip"
		Thomas  = "Thomas"

		// Variables
		// Food
		PropertyFood = "food"
		HAM          = "ham"
		POTATO       = "potato"
		PUMPKIN      = "pumpkin"
		TURKEY       = "Turkey"
	)

	p := cps.NewProblemWithNames(Larry, Nicolas, Philip, Thomas)

	p.AddIntegerVariables(PropertyAge, 8, 9, 10, 11)
	p.AddStringVariables(PropertyFood, HAM, POTATO, PUMPKIN, TURKEY)

	// rule1: Larry is look forward to eating turkey
	p.AddConstraintRule("rule1", func(vs *cps.Variables) bool {
		return Larry == vs.Of(TURKEY)
	}, TURKEY)

	// rule2: The body who likes pumpkin pie is one year younger than Philip
	p.AddConstraintRule("rule2", func(vs *cps.Variables) bool {
		return ageOf(vs, PUMPKIN) == ageOf(vs, Philip)-1
	}, PUMPKIN, AllAges)

	// rule3: Thomas is younger than the boy that loves turkey.
	p.AddConstraintRule("rule3", func(vs *cps.Variables) bool {
		return ageOf(vs, Thomas) < ageOf(vs, TURKEY)
	}, TURKEY, AllAges)

	// rule4: The boy who likes ham is two years older than Philp
	p.AddConstraintRule("rule4", func(vs *cps.Variables) bool {
		return ageOf(vs, HAM) == ageOf(vs, Philip)+2
	}, HAM, AllAges)

	p.Solve()
	fmt.Println("Results:")
	p.PrintSolutions()
}
