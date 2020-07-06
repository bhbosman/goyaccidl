package objects

import (
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type Dcl struct {
	BaseDcl
}

func NewDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	forward bool,
	primitiveType yaccidl.PrimitiveType) *Dcl {

	return &Dcl{
		BaseDcl: *NewBaseDcl(
			primitiveType,
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			forward,
			false),
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*Dcl)(nil)),
		func() IDcl {
			return new(Dcl)
		})
}
