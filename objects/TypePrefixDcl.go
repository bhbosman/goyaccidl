package objects

import "github.com/bhbosman/yaccidl"

type TypePrefixDcl struct {
	Dcl
	Value string
}

func NewTypePrefixDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	value string) *TypePrefixDcl {
	return &TypePrefixDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlTypePrefix),
		Value: value,
	}
}
