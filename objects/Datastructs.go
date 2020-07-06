package objects

import (
	"go.uber.org/multierr"
)

func ErrorToString(err error) []string {
	var result []string
	for _, e := range multierr.Errors(err) {
		result = append(result, e.Error())
	}
	return result
}
