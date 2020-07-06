package objects

import "encoding/json"

type AttributeDcl struct {
	AttributeName      string `json:"attribute_name"`
	AttributeType      IDcl   `json:"-"` // fixed
	AttributeWriteable bool   `json:"attribute_writeable"`
	AttributeReadable  bool   `json:"attribute_readable"`
}

func (self *AttributeDcl) MarshalJSON() ([]byte, error) {
	attributeTypeDcl, err := NewScopeStream(self.AttributeType)
	if err != nil {
		return nil, err
	}

	type Alias AttributeDcl
	return json.Marshal(
		&struct {
			AttributeTypeDcl ScopeStream `json:"attribute_type_dcl"`
			*Alias
		}{
			AttributeTypeDcl: attributeTypeDcl,
			Alias:            (*Alias)(self),
		})
}

func (self *AttributeDcl) UnmarshalJSON(bytes []byte) error {
	type Alias AttributeDcl
	data := &struct {
		AttributeTypeDcl ScopeStream `json:"attribute_type_dcl"`
		*Alias
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	self.AttributeType, err = data.AttributeTypeDcl.GetDcl()
	if err != nil {
		return err
	}
	return nil
}

func NewAttribute(attributeName string, attributeType IDcl, attributeWriteable bool, attributeReadable bool) *AttributeDcl {
	return &AttributeDcl{
		AttributeName:      attributeName,
		AttributeType:      attributeType,
		AttributeWriteable: attributeWriteable,
		AttributeReadable:  attributeReadable,
	}
}
