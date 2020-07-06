package objects

import (
	"fmt"
	"strings"
)

type ScopeIdentifier string

func (self ScopeIdentifier) Scope() ScopeIdentifier {
	ss := strings.Split(string(self), "::")
	if len(ss) <= 1 {
		return ScopeIdentifier("")
	}
	return ScopeIdentifier(strings.Join(ss[:len(ss)-1], "::"))
}

func (self ScopeIdentifier) Name() string {
	ss := strings.Split(string(self), "::")
	if len(ss) <= 1 {
		return string(self)
	}
	return ss[len(ss)-1]
}

func (self ScopeIdentifier) First() string {
	ss := strings.Split(string(self), "::")
	if len(ss) <= 1 {
		return string(self)
	}
	return ss[0]

}
func (self ScopeIdentifier) Append(s string) ScopeIdentifier {
	result := fmt.Sprintf("%v_%v", self, s)
	return ScopeIdentifier(result)
}

func (self ScopeIdentifier) Combined() string {
	ss := strings.Split(string(self), "::")
	return strings.Join(ss, "")
}
