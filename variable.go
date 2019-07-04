package cps

import (
	"fmt"
	"strconv"
	"strings"
)

// Variable defines a variable which can associate constraints later.
type Variable struct {
	name  string
	value Domain
}

// Variables defines a list of variable, providing utility methods to access it.
type Variables struct {
	items   map[string]*Variable
	domains map[string]Domain
}

// Init initializes an empty Variable
func (vs *Variables) Init() {
	vs.items = make(map[string]*Variable)
	vs.domains = make(map[string]Domain)
}

// Add adds a new variable of given name
func (vs *Variables) Add(name string) *Variables {
	_, exists := vs.items[name]
	if exists {
		panic(fmt.Sprintf("Value [%s] already exists", name))
	}

	v := Variable{name, NotSet}
	vs.items[name] = &v
	return vs
}

// Set sets value for the variable
func (vs *Variables) Set(name string, domain Domain) *Variables {
	vs.items[name].value = domain
	return vs
}

// Of returns the variable of given name
func (vs *Variables) Of(name string) Domain {
	domain, isDomain := vs.domains[name]
	if isDomain {
		return domain
	}
	v, ok := vs.items[name]

	if ok {
		return v.value
	}
	panic(fmt.Sprintf("can not find variable with name [%s]", name))
}

// IntPropertyOf returns integer value of property given the variable name.
func (vs *Variables) IntPropertyOf(d Domain, propName string) int {
	for k, v := range vs.items {
		if strings.HasPrefix(k, propName+"*") && v.value == d {
			intValue, err := strconv.Atoi(strings.Replace(k, propName+"*", "", -1))
			if err != nil {
				panic(fmt.Sprintf("unexpected variable name [%s], can not convert to int value.", k))
			}
			return intValue
		}
	}
	panic(fmt.Sprintf("Can not find a varialbe [%s] with assigned value of [%s]\n This may also happens when constraints were added before variables", d, propName))
}
