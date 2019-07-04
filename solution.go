package cps

import "fmt"

// A Solution denotes a feasbile solution of the problem
type Solution struct {
	assignment map[string]Domain
}

// OutputSolution outputs the solution.
func OutputSolution(s Solution) {
	reversed := make(map[Domain][]string)
	for name, domain := range s.assignment {
		reversed[domain] = append(reversed[domain], name)
	}

	for domain, names := range reversed {
		for _, name := range names {
			fmt.Printf("%s,", name)
		}
		fmt.Println(domain)
	}
}
