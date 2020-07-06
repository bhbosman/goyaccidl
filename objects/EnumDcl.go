package objects

import (
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type EnumDcl struct {
	Dcl
	Members []string `json:"members"`
}

func (d *EnumDcl) AddMember(identifier string) {
	d.Members = append(d.Members, identifier)
}

func NewEnumDcl(
	identifier ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue) *EnumDcl {
	return &EnumDcl{
		Dcl: *NewDcl(
			identifier,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlEnum),
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*EnumDcl)(nil)),
		func() IDcl {
			return new(EnumDcl)
		})
}
