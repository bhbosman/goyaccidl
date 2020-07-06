package objects

import "github.com/bhbosman/yaccidl"

type ModuleDcl struct {
	Dcl
	Members IDclArray `json:"members"`
}

func (m *ModuleDcl) GetList() IDclArray {
	return m.Members
}

func (m *ModuleDcl) Clear() {
	m.Members = nil
}

func NewModuleDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	members IDclArray) *ModuleDcl {
	return &ModuleDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlModule),
		Members: members,
	}
}
