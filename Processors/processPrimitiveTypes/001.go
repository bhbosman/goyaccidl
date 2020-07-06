package processPrimitiveTypes

import (
	"github.com/bhbosman/goyaccidl/Service"
)

type ProcessorStep01 struct {
	defaultTypes *Service.IdlDefaultTypes
}

func NewProcessorStep01(defaultTypes *Service.IdlDefaultTypes) *ProcessorStep01 {
	return &ProcessorStep01{
		defaultTypes: defaultTypes,
	}
}

func (self ProcessorStep01) Name() string {
	return "get primitive types"
}

func (self ProcessorStep01) Run(_ interface{}) (interface{}, error) {
	return self.defaultTypes.IdlPrimitivesInScopedFormat(), nil
}
