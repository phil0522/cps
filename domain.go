package cps

// Domain defines possible values for each variable.
type Domain string

const (
	// NotSet is a special to denote value of the variable is not set yet.
	NotSet = Domain("__not__")
)

// Domains is a list of Domain
type Domains []Domain

// makeNamesDomain create an array of domains with names
func makeNamesDomain(name ...string) []Domain {
	var r []Domain
	size := len(name)
	for i := 1; i <= size; i++ {
		r = append(r, Domain(name[i-1]))
	}
	return r
}
