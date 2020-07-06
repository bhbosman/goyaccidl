package objects

import (
	"github.com/bhbosman/yaccidl"
)

type InterfaceDclContainer struct {
	Dcl
	BaseInterfaces IDclArray
	Local          bool
	Abstract       bool
	Custom         bool
	ValueType      bool
	Operations     []*OperationDcl
	Attributes     []*AttributeDcl
	Members        IDclArray
}

func (i *InterfaceDclContainer) GetList() IDclArray {
	return i.Members
}

func (i *InterfaceDclContainer) Clear() {
	i.Members = nil
}

func NewInterfaceDclContainer(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	baseInterfaces IDclArray,
	local, abstract, custom, forward, valueType bool,
	members IDclArray,
	operations []*OperationDcl,
	attributes []*AttributeDcl) *InterfaceDclContainer {
	return &InterfaceDclContainer{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			forward,
			yaccidl.IdlInterfaceContainer),
		Members:        members,
		BaseInterfaces: baseInterfaces,
		Local:          local,
		Abstract:       abstract,
		Custom:         custom,
		ValueType:      valueType,
		Operations:     operations,
		Attributes:     attributes,
	}
}
