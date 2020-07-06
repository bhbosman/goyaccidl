package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type SequenceTypeDcl struct {
	Dcl          IDcl
	SequenceType IDcl
}

func (self *SequenceTypeDcl) SetSequenceType(v IDcl) error {
	self.SequenceType = v
	return nil
}

func (self *SequenceTypeDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.Dcl.DclResolveFolderUsage()
}

func (self *SequenceTypeDcl) SetDestination(folderId, folder string) error {
	return self.Dcl.SetDestination(folderId, folder)
}

func (self *SequenceTypeDcl) GetTypePrefix() string {
	return self.Dcl.GetTypePrefix()
}

func (self *SequenceTypeDcl) SetTypePrefix(s string) {
	self.Dcl.SetTypePrefix(s)
}

func (self *SequenceTypeDcl) MarshalJSON() ([]byte, error) {
	var Dcl IDclStream
	var sequenceType IDclStream
	var err error
	Dcl, err = NewIDclStream(self.Dcl)
	if err != nil {
		return nil, err
	}
	sequenceType, err = NewIDclStream(self.SequenceType)
	if err != nil {
		return nil, err
	}
	return json.Marshal(
		&struct {
			Dcl          IDclStream `json:"dcl"`
			SequenceType IDclStream `json:"sequence_type"`
		}{
			Dcl:          Dcl,
			SequenceType: sequenceType,
		})
}

func (self *SequenceTypeDcl) UnmarshalJSON(bytes []byte) error {
	data := &struct {
		Dcl          IDclStream `json:"dcl"`
		SequenceType IDclStream `json:"sequence_type"`
	}{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	self.Dcl, err = data.Dcl.GetDcl()
	if err != nil {
		return err
	}
	self.SequenceType, err = data.SequenceType.GetDcl()
	if err != nil {
		return err
	}
	return nil
}

func (self *SequenceTypeDcl) GetName() ScopeIdentifier {
	return self.Dcl.GetName()
}

func (self *SequenceTypeDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.Dcl.GetPrimitiveType()
}

func (self *SequenceTypeDcl) GetLexemData() IDstSrcInformation {
	return self.Dcl.GetLexemData()
}

func (self *SequenceTypeDcl) GetOrderId() int {
	return self.Dcl.GetOrderId()
}

func (self *SequenceTypeDcl) GetForward() bool {
	return self.Dcl.GetForward()
}

func (self *SequenceTypeDcl) GetImported() bool {
	return self.Dcl.GetImported()
}

func (self *SequenceTypeDcl) UpdateIdlReference(s ScopeIdentifier) error {
	return self.Dcl.UpdateIdlReference(s)
}

func (self *SequenceTypeDcl) BuildIdValue() (string, error) {
	return self.Dcl.BuildIdValue()
}

func (self *SequenceTypeDcl) GetSequenceType() IDcl {
	return self.SequenceType
}

func (self *SequenceTypeDcl) IsISequenceTypeDcl() bool {
	return true
}

func (self *SequenceTypeDcl) toIDcl() IDcl {
	return self
}

func (self *SequenceTypeDcl) GetIsArray() bool {
	return self.Dcl.GetIsArray()
}

func (self *SequenceTypeDcl) GetArrayValue() []int64 {
	return self.Dcl.GetArrayValue()
}

func NewSequenceTypeDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	sequenceType IDcl) ISequenceTypeDcl {
	return &SequenceTypeDcl{
		Dcl: NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false, yaccidl.IdlSequence),
		SequenceType: sequenceType,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*SequenceTypeDcl)(nil)),
		func() IDcl {
			return new(SequenceTypeDcl)
		})
}
