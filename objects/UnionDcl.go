package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type UnionDcl struct {
	Dcl
	UnionBodies   []*UnionBody `json:"union_bodies"`
	SwitchTypeDcl IDcl         `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
}

func (d *UnionDcl) GetStructExceptionDcl() (bool, error) {
	return false, nil
}

func (d *UnionDcl) SetDestination(folderId, folder string) error {
	err := d.Dcl.SetDestination(folderId, folder)
	return err
}

func (self *UnionDcl) GetStructBase() (ScopeIdentifier, error) {
	return "", nil
}

func (self *UnionDcl) GetStructMembers() ([]*StructMember, error) {
	var structMembers []*StructMember
	structMembers = append(structMembers, NewStructMember(
		"Switch",
		false,
		nil,
		self.SwitchTypeDcl))
	for _, item := range self.UnionBodies {
		structMembers = append(
			structMembers,
			NewStructMember(
				item.Identifier,
				item.Array,
				item.Int64s,
				item.TypeSpec))
	}
	return structMembers, nil
}

func (self *UnionDcl) IsIUnionDcl() bool {
	return true
}

func (self *UnionDcl) toIUnionDcl() IUnionDcl {
	return self
}

func (self *UnionDcl) MarshalJSON() ([]byte, error) {
	TypeSpec, err := NewIDclStream(self.SwitchTypeDcl)
	if err != nil {
		return nil, err
	}

	type Alias UnionDcl
	return json.Marshal(
		&struct {
			*Alias
			TypeSpec IDclStream `json:"type_spec"`
		}{
			Alias:    (*Alias)(self),
			TypeSpec: TypeSpec,
		})
}

func (self *UnionDcl) UnmarshalJSON(bytes []byte) error {
	type Alias UnionDcl
	data := &struct {
		*Alias
		TypeSpec IDclStream `json:"type_spec"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.SwitchTypeDcl, err = data.TypeSpec.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func NewUnionDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	switchTypeDcl IDcl,
	unionBodies []*UnionBody) *UnionDcl {
	return &UnionDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlUnion),
		UnionBodies:   unionBodies,
		SwitchTypeDcl: switchTypeDcl,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*UnionDcl)(nil)),
		func() IDcl {
			return new(UnionDcl)
		})
}
