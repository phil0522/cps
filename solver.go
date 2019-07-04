package cps

import (
	"fmt"
	"sort"
)

// ByAffectedVariable implements sort.Interface for []Constraint based on
// number of affect affected varaibles.
type ByAffectedVariable []Constraint

func (a ByAffectedVariable) Len() int      { return len(a) }
func (a ByAffectedVariable) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares two item at given indexes
func (a ByAffectedVariable) Less(i, j int) bool {
	lenA := len(a[i].nameSet)
	lenB := len(a[j].nameSet)
	if lenA != lenB {
		return lenA < lenB
	}
	return a[i].desc < a[j].desc
}

type orderByRankSorter struct {
	names []string
	ranks map[string]int
}

func (a orderByRankSorter) Len() int      { return len(a.names) }
func (a orderByRankSorter) Swap(i, j int) { a.names[i], a.names[j] = a.names[j], a.names[i] }
func (a orderByRankSorter) Less(i, j int) bool {
	rank1 := a.ranks[a.names[i]]
	rank2 := a.ranks[a.names[j]]
	if rank1 == rank2 {
		return a.names[i] < a.names[j]
	}
	return rank1 < rank2
}

func (a orderByRankSorter) Sort(names []string) {
	a.names = names
	sort.Sort(a)
}

func orderByRank(ranks map[string]int) *orderByRankSorter {
	return &orderByRankSorter{
		ranks: ranks,
	}
}

type varContext struct {
	name             string       // variable name
	accumulatedNames []string     // All assigned variable when this variable is assigned.
	newConstraints   []Constraint // The newly added constraint for this variable
}
type context struct {
	varContexts []varContext
}

func exists(s map[string]bool, name string) bool {
	_, ok := s[name]
	return ok
}

// Sorts all varialbes and constraints for better performance
func (p *Problem) reorderVariables() {
	sort.Sort(ByAffectedVariable(p.constraints))
}

func cloneStringArray(array []string) []string {
	newArray := []string{}
	newArray = append(newArray, array...)
	return newArray
}

func inStringList(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func isConstraintSolvable(nameSet map[string]bool, accumulatedNames []string) bool {
	for name := range nameSet {
		if !inStringList(accumulatedNames, name) {
			return false
		}
	}
	return true
}

// Pre-process, fill in context
func (p *Problem) preprocess() {
	p.reorderVariables()

	variableRank := make(map[string]int)
	for rank, c := range p.constraints {
		for name := range c.nameSet {
			_, exists := variableRank[name]
			if exists {
				continue
			}
			variableRank[name] = rank
		}
	}

	allVariables := []string{}
	for _, v := range p.variables.items {
		allVariables = append(allVariables, v.name)
	}
	orderByRank(variableRank).Sort(allVariables)

	accumulatedNames := []string{}
	for _, name := range allVariables {
		varContext := varContext{}
		varContext.name = name
		accumulatedNames = append(accumulatedNames, name)
		varContext.accumulatedNames = cloneStringArray(accumulatedNames)

		// Add constraints that are complete on the accumulated names.
		for _, c := range p.constraints {
			if exists(c.nameSet, name) && isConstraintSolvable(c.nameSet, accumulatedNames) {
				varContext.newConstraints = append(varContext.newConstraints, c)
			}
		}
		p.context = append(p.context, varContext)
	}

	p.displayOrderedVariablesAndConstraints()
}

func (p *Problem) displayOrderedVariablesAndConstraints() {
	fmt.Println("reordered variables and constraints")
	for _, ctx := range p.context {
		fmt.Print(ctx.name)
		fmt.Print(" ")
		fmt.Print(ctx.accumulatedNames)
		fmt.Print(" [")
		for _, c := range ctx.newConstraints {
			fmt.Print(c.desc)
			fmt.Print(", ")
		}
		fmt.Print("]")
		fmt.Println()
	}

	fmt.Println()

	for _, c := range p.constraints {
		fmt.Printf("constraint %s %v\n", c.desc, c.nameSet)
	}
}

// Solve solves to solutions
func (p *Problem) Solve() {
	p.preprocess()
	p.recursiveBackTrace(0)
}

func (p *Problem) recursiveBackTrace(depth int) {
	if depth >= len(p.variables.items) {
		// We get a solution, as we satisfy all constraints and use up all variables.
		solution := Solution{}
		solution.assignment = make(map[string]Domain)
		for _, ctx := range p.context {
			solution.assignment[ctx.name] = p.variables.Of(ctx.name)
		}
		p.solutions = append(p.solutions, solution)
		return
	}

	variable := p.variables.items[p.context[depth].name]
	for _, d := range p.domains {
		variable.value = d
		for i := 0; i < depth; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("set variable [%s] to [%s]", variable.name, variable.value)
		// Check all constraints
		pass := true
		for _, c := range p.context[depth].newConstraints {
			if !c.rule(&p.variables) {
				fmt.Printf(" failed at [%s]\n", c.desc)
				pass = false
				break
			}
		}
		if pass {
			fmt.Printf(" pass\n")
			p.recursiveBackTrace(depth + 1)
		}
	}
}
