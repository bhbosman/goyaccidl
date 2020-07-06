package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
)

type OperationsParam struct {
	ParameterName string                 `json:"parameter_name"`
	Direction     yaccidl.ParamDirection `json:"direction"`
	ParameterDcl  IDcl                   `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
}

func (self *OperationsParam) MarshalJSON() ([]byte, error) {
	ParameterDcl, err := NewScopeStream(self.ParameterDcl)
	if err != nil {
		return nil, err
	}

	type Alias OperationsParam
	return json.Marshal(
		&struct {
			*Alias
			ParameterDcl ScopeStream `json:"parameter_dcl"`
		}{
			Alias:        (*Alias)(self),
			ParameterDcl: ParameterDcl,
		})
}

func (self *OperationsParam) UnmarshalJSON(bytes []byte) error {
	type Alias OperationsParam
	data := &struct {
		*Alias
		ParameterDcl ScopeStream `json:"parameter_dcl"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.ParameterDcl, err = data.ParameterDcl.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func NewOperationsParam(parameterName string, parameterType IDcl, direction yaccidl.ParamDirection) *OperationsParam {
	return &OperationsParam{
		ParameterName: parameterName,
		ParameterDcl:  parameterType,
		Direction:     direction}
}
