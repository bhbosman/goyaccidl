package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type InterfaceDcl struct {
	Dcl
	BaseInterfaces IDclArray       `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
	Local          bool            `json:"local"`
	Abstract       bool            `json:"abstract"`
	Custom         bool            `json:"custom"`
	ValueType      bool            `json:"value_type"`
	Operations     []*OperationDcl `json:"Operations"`
	Attributes     []*AttributeDcl `json:"attributes"`
}

func (self *InterfaceDcl) MarshalJSON() ([]byte, error) {
	var coll []ScopeStream
	for _, item := range self.BaseInterfaces {
		dclStream, err := NewScopeStream(item)
		if err != nil {
			return nil, err
		}
		coll = append(coll, dclStream)
	}

	type Alias InterfaceDcl
	return json.Marshal(
		&struct {
			*Alias
			BaseInterfaces []ScopeStream `json:"base_interfaces"`
		}{
			Alias:          (*Alias)(self),
			BaseInterfaces: coll,
		})
}

func (self *InterfaceDcl) UnmarshalJSON(bytes []byte) error {
	type Alias InterfaceDcl
	data := &struct {
		*Alias
		BaseInterfaces []ScopeStream `json:"base_interfaces"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	for _, item := range data.BaseInterfaces {
		dcl, err := item.GetDcl()
		if err != nil {
			return err
		}
		self.BaseInterfaces = append(self.BaseInterfaces, dcl)
	}
	return nil
}

func NewInterfaceDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	baseInterfaces IDclArray,
	local, abstract, custom, forward, valueType bool,
	operations []*OperationDcl,
	attributes []*AttributeDcl) *InterfaceDcl {
	return &InterfaceDcl{
		Dcl: *NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			forward,
			yaccidl.IdlInterface),
		BaseInterfaces: baseInterfaces,
		Local:          local,
		Abstract:       abstract,
		Custom:         custom,
		ValueType:      valueType,
		Operations:     operations,
		Attributes:     attributes,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*InterfaceDcl)(nil)),
		func() IDcl {
			return new(InterfaceDcl)
		})
}
