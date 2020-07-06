package objects

import (
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type ScopeDcl struct {
	Dcl
}

func NewScopeDcl2(
	node yaccidl.IYaccNode,
	orderId int,
	lexemData *yaccidl.LexemValue) *ScopeDcl {
	return &ScopeDcl{
		Dcl: *NewDcl(
			ScopeIdentifier(node.Identifier()),
			node.IsArray(),
			node.Array(),
			orderId,
			lexemData,
			false,
			yaccidl.IdlScope),
	}
}

func NewScopeDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue) *ScopeDcl {
	return &ScopeDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlScope),
	}
}

func NewScopeDclWithType(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	primitiveType yaccidl.PrimitiveType) *ScopeDcl {
	return &ScopeDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			primitiveType),
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*ScopeDcl)(nil)),
		func() IDcl {
			return new(ScopeDcl)
		},
	)
}
