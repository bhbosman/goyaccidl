package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
)

type ScopeItem struct {
	Dcl IDcl `json:"-"`
}

func (self *ScopeItem) MarshalJSON() ([]byte, error) {
	TypeSpec, err := NewIDclStream(self.Dcl)
	if err != nil {
		return nil, err
	}

	type Alias ScopeItem
	return json.Marshal(
		&struct {
			*Alias
			Dcl IDclStream `json:"dcl"`
		}{
			Alias: (*Alias)(self),
			Dcl:   TypeSpec,
		})
}

func (self *ScopeItem) UnmarshalJSON(bytes []byte) error {
	type Alias ScopeItem
	data := &struct {
		*Alias
		Dcl IDclStream `json:"dcl"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.Dcl, err = data.Dcl.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func NewScopeItem(
	Dcl IDcl,
	idlType yaccidl.PrimitiveType,
	orderId int,
	scopeName ScopeIdentifier,
	lexemData *yaccidl.LexemValue) *ScopeItem {
	return &ScopeItem{
		Dcl: Dcl,
	}
}
