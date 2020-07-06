package objects

import "encoding/json"

type UnionBody struct {
	DefaultCase bool              `json:"default_case"`
	CaseValues  []*UnionCaseValue `json:"case_values"`
	TypeSpec    IDcl              `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
	Identifier  string            `json:"identifier"`
	Array       bool              `json:"array"`
	Int64s      []int64           `json:"int_64_s"`
}

func (self *UnionBody) MarshalJSON() ([]byte, error) {
	TypeSpec, err := NewIDclStream(self.TypeSpec)
	if err != nil {
		return nil, err
	}

	type Alias UnionBody
	return json.Marshal(
		&struct {
			*Alias
			TypeSpec IDclStream `json:"type_spec"`
		}{
			Alias:    (*Alias)(self),
			TypeSpec: TypeSpec,
		})
}

func (self *UnionBody) UnmarshalJSON(bytes []byte) error {
	type Alias UnionBody
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

	self.TypeSpec, err = data.TypeSpec.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func (u *UnionBody) GetTypeSpec() IDcl {
	return u.TypeSpec
}

func (u UnionBody) GetDefaultCase() bool {
	return u.DefaultCase
}

func (u UnionBody) GetCaseValues() []*UnionCaseValue {
	return u.CaseValues
}

func NewUnionBody(
	defaultCase bool,
	caseValues []*UnionCaseValue,
	Identifier string,
	array bool,
	int64s []int64,
	typeSpec IDcl) *UnionBody {
	return &UnionBody{
		DefaultCase: defaultCase,
		CaseValues:  caseValues,
		TypeSpec:    typeSpec,
		Identifier:  Identifier,
		Array:       array,
		Int64s:      int64s,
	}
}
