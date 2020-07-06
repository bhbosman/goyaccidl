package ctx

import "strings"

type IdlDefinitions struct {
	definitions map[string]int
}

func (i IdlDefinitions) String() string {
	var result []string
	for k, _ := range i.definitions {
		result = append(result, k)
	}
	return strings.Join(result, ",")
}

func (i *IdlDefinitions) Set(v string) error {
	ss := strings.Split(v, ",")
	for _, s := range ss {
		i.definitions[s] = 1
	}
	return nil
}

func (i IdlDefinitions) AssignFlags() []string {
	var result []string
	for k, _ := range i.definitions {
		result = append(result, k)
	}
	return result
}

func NewIdlDefinitions() IdlDefinitions {
	return IdlDefinitions{
		definitions: make(map[string]int),
	}
}
