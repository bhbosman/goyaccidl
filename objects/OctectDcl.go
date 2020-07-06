package objects

import "github.com/bhbosman/yaccidl"

type IOctetDcl interface {
	IDcl
	IsIOctetDcl() bool
}

type OctetDcl struct {
	BaseDcl
}

func (o OctetDcl) IsIOctetDcl() bool {
	return true
}

func NewOctetDcl(
	PrimitiveType yaccidl.PrimitiveType,
	Name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	LexemData *yaccidl.LexemValue,
	forward bool,
	TypeDefScopeName ScopeIdentifier,
	TypeDefIdlType yaccidl.PrimitiveType) *OctetDcl {
	return &OctetDcl{
		BaseDcl: *NewBaseDcl(
			PrimitiveType,
			Name,
			isArray,
			arrayValue,
			orderId,
			LexemData,
			forward,
			true),
	}
}

func testOctetDcl(v *OctetDcl) IOctetDcl {
	return v
}
