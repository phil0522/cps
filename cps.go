package cps

// NewProblemWithNames creates a new Problem.
func NewProblemWithNames(names ...Domain) Problem {
	p := Problem{}
	p.variables.Init()

	if len(names) < 2 {
		panic("must provide at least 2 domain values")
	}
	for _, name := range names {
		p.domains = append(p.domains, name)
	}
	for _, d := range p.domains {
		p.variables.domains[string(d)] = d
	}
	return p
}
