package objects

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type ConstantValue struct {
	Dcl
	ConstantType  IDcl        `json:"-"`
	ConstantValue interface{} `json:"-"`
}

func (self *ConstantValue) MarshalJSON() ([]byte, error) {
	TypeSpec, err := NewIDclStream(self.ConstantType)
	if err != nil {
		return nil, err
	}

	type Alias ConstantValue
	return json.Marshal(
		&struct {
			*Alias
			TypeSpec      IDclStream `json:"type_spec"`
			ConstantValue string     `json:"constant_value"`
		}{
			Alias:         (*Alias)(self),
			TypeSpec:      TypeSpec,
			ConstantValue: fmt.Sprintf("%v", self.ConstantValue),
		})
}

func (self *ConstantValue) UnmarshalJSON(bytes []byte) error {
	type Alias ConstantValue
	data := &struct {
		*Alias
		TypeSpec      IDclStream `json:"type_spec"`
		ConstantValue string     `json:"constant_value"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.ConstantType, err = data.TypeSpec.GetDcl()
	if err != nil {
		return err
	}

	self.ConstantValue = data.ConstantValue

	return nil
}

func (c ConstantValue) GetValue() interface{} {
	return c.ConstantValue
}

func (c ConstantValue) GetConstantType() IDcl {
	return c.ConstantType
}

func (c ConstantValue) IsConstantValue() bool {
	return true
}

func NewConstantValue(
	constantName ScopeIdentifier,
	orderId int,
	lexemData *yaccidl.LexemValue,
	constantType IDcl,
	constantValue interface{}) *ConstantValue {

	return &ConstantValue{
		Dcl: *NewDcl(
			constantName,
			false,
			nil,
			orderId,
			lexemData,
			false,
			yaccidl.IdlConst),
		ConstantType:  constantType,
		ConstantValue: constantValue,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*ConstantValue)(nil)),
		func() IDcl {
			return new(ConstantValue)
		})
}
