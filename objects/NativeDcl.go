package objects

import (
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type NativeDcl struct {
	Dcl
}

func (self *NativeDcl) toIDcl() IDcl {
	return self
}

func NewNativeDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue) *NativeDcl {
	return &NativeDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlNative),
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*NativeDcl)(nil)),
		func() IDcl {
			return new(NativeDcl)
		})
}
