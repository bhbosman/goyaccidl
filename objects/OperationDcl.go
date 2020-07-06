package objects

import "encoding/json"

type OperationDcl struct {
	OperationName    string             `json:"operation_name"`
	OperationsParams []*OperationsParam `json:"operations_params"`
	OperationType    IDcl               `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
	Exceptions       IDclArray          `json:"-"` // json encoding resolved in MarshalJSON and UnmarshalJSON
}

func (self *OperationDcl) MarshalJSON() ([]byte, error) {
	OperationTypeDcl, err := NewScopeStream(self.OperationType)
	if err != nil {
		return nil, err
	}

	var ExceptionsDcl []ScopeStream
	for _, exception := range self.Exceptions {
		var dclStream ScopeStream
		dclStream, err = NewScopeStream(exception)
		ExceptionsDcl = append(ExceptionsDcl, dclStream)
	}

	type Alias OperationDcl
	return json.Marshal(
		&struct {
			*Alias
			OperationType ScopeStream   `json:"operation_type"`
			Exceptions    []ScopeStream `json:"exceptions"`
		}{
			Alias:         (*Alias)(self),
			OperationType: OperationTypeDcl,
			Exceptions:    ExceptionsDcl,
		})
}

func (self *OperationDcl) UnmarshalJSON(bytes []byte) error {
	type Alias OperationDcl
	data := &struct {
		*Alias
		OperationType ScopeStream   `json:"operation_type"`
		Exceptions    []ScopeStream `json:"exceptions"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.OperationType, err = data.OperationType.GetDcl()
	if err != nil {
		return err
	}

	for _, stream := range data.Exceptions {
		dcl, err := stream.GetDcl()
		if err != nil {
			return err
		}
		self.Exceptions = append(self.Exceptions, dcl)
	}

	return nil
}

func NewOperation(operationName string, operationType IDcl, operationsParams []*OperationsParam, exceptions IDclArray) *OperationDcl {
	return &OperationDcl{
		OperationName:    operationName,
		OperationType:    operationType,
		OperationsParams: operationsParams,
		Exceptions:       exceptions}
}
