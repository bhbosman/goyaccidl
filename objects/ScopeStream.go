package objects

import (
	"github.com/bhbosman/yaccidl"
)

type ScopeStream struct {
	Assigned   bool                `json:"assigned"`
	Name       ScopeIdentifier     `json:"name"`
	IsArray    bool                `json:"is_array"`
	ArrayValue []int64             `json:"array_value"`
	OrderId    int                 `json:"order_id"`
	LexemData  *yaccidl.LexemValue `json:"lexem_data"`
}

func NewScopeStream(dcl IDcl) (ScopeStream, error) {
	if dcl == nil {
		return ScopeStream{
			Assigned: false,
		}, nil
	}

	return ScopeStream{
		Assigned:   true,
		Name:       dcl.GetName(),
		IsArray:    dcl.GetIsArray(),
		ArrayValue: dcl.GetArrayValue(),
		OrderId:    dcl.GetOrderId(),
		LexemData:  dcl.GetLexemData().GetSourceLexemData(),
	}, nil
}

func (self *ScopeStream) GetDcl() (IDcl, error) {
	if !self.Assigned {
		return nil, nil
	}
	return NewScopeDcl(self.Name, self.IsArray, self.ArrayValue, self.OrderId, self.LexemData), nil

}
