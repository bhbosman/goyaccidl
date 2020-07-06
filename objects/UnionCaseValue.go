package objects

import "encoding/json"

type UnionCaseValue struct {
	Value IDcl `json:"-"`
}

func (self *UnionCaseValue) MarshalJSON() ([]byte, error) {
	value, err := NewIDclStream(self.Value)
	if err != nil {
		return nil, err
	}

	type Alias UnionCaseValue
	return json.Marshal(
		&struct {
			*Alias
			Value IDclStream `json:"value"`
		}{
			Alias: (*Alias)(self),
			Value: value,
		})
}

func (self *UnionCaseValue) UnmarshalJSON(bytes []byte) error {
	type Alias UnionCaseValue
	data := &struct {
		*Alias
		Value IDclStream `json:"value"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.Value, err = data.Value.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func NewUnionCaseValue(value IDcl) *UnionCaseValue {
	return &UnionCaseValue{
		Value: value,
	}
}
