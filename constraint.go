package cps

// Rule is type of function that specifies constraints.
type Rule func(variables *Variables) bool

// Constraint defines a rule and associated variables
type Constraint struct {
	desc    string // The description of the constraint
	nameSet map[string]bool
	rule    Rule
}

// MakeContraint creates a new contraint
func MakeContraint(desc string) Constraint {
	r := Constraint{}
	r.desc = desc
	r.nameSet = make(map[string]bool)
	return r
}
