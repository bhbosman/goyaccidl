package objects

import (
	"fmt"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

//type IPrimitiveDcl interface {
//	IDcl
//	//IsIPrimitiveDcl() bool
//	//IsNative() bool
//}
type PrimitiveDcl struct {
	Dcl
	ReferenceType bool
	isNative      bool
}

func (self *PrimitiveDcl) IsNative() bool {
	return self.isNative
}

func (self *PrimitiveDcl) IsIPrimitiveDcl() bool {
	return true
}

func (self *PrimitiveDcl) BuildIdValue() (string, error) {
	return fmt.Sprintf(""), nil
}

func (self *PrimitiveDcl) UpdateIdlReference(_ ScopeIdentifier) error {
	// override, but do nothing
	return nil
}

func NewPrimitiveDcl(
	node yaccidl.IPrimitiveTypeDcl,
	ReferenceType bool) *PrimitiveDcl {
	return NewPrimitiveDclFull(
		ScopeIdentifier(node.Identifier()),
		false,
		nil,
		-1,
		node.PrimitiveType(),
		node.LexemData(),
		ReferenceType)
}

func NewPrimitiveDclFull(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	PrimitiveType yaccidl.PrimitiveType,
	lexemData *yaccidl.LexemValue,
	ReferenceType bool) *PrimitiveDcl {
	return &PrimitiveDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false, PrimitiveType),
		ReferenceType: ReferenceType,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*PrimitiveDcl)(nil)),
		func() IDcl {
			return new(PrimitiveDcl)
		},
	)
}
