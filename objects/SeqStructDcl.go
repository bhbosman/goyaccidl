package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type SeqStructDcl struct {
	Dcl                IDcl            `json:"-"`
	StructBase         IDcl            `json:"-"`
	StructMembers      []*StructMember `json:"struct_members"`
	StructExceptionDcl bool            `json:"struct_exception_dcl"`
}

func (self *SeqStructDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.Dcl.DclResolveFolderUsage()
}

func (self *SeqStructDcl) SetDestination(folderId, folder string) error {
	return self.Dcl.SetDestination(folderId, folder)
}

func (self *SeqStructDcl) GetTypePrefix() string {
	return self.Dcl.GetTypePrefix()
}

func (self *SeqStructDcl) GetIsArray() bool {
	return self.Dcl.GetIsArray()
}

func (self *SeqStructDcl) GetArrayValue() []int64 {
	return self.Dcl.GetArrayValue()
}

func (self *SeqStructDcl) MarshalJSON() ([]byte, error) {
	dcl, err := NewIDclStream(self.Dcl)
	if err != nil {
		return nil, err
	}
	var structBase IDclStream
	structBase, err = NewIDclStream(self.StructBase)
	if err != nil {
		return nil, err
	}

	type Alias SeqStructDcl
	return json.Marshal(
		&struct {
			*Alias
			Dcl        IDclStream `json:"dcl"`
			StructBase IDclStream `json:"struct_base"`
		}{
			Alias:      (*Alias)(self),
			Dcl:        dcl,
			StructBase: structBase,
		})
}

func (self *SeqStructDcl) UnmarshalJSON(bytes []byte) error {
	type Alias SeqStructDcl
	data := &struct {
		*Alias
		Dcl        IDclStream `json:"dcl"`
		StructBase IDclStream `json:"struct_base"`
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
	self.StructBase, err = data.StructBase.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func (self *SeqStructDcl) SetTypePrefix(s string) {
	self.Dcl.SetTypePrefix(s)
}

func (self *SeqStructDcl) GetName() ScopeIdentifier {
	return self.Dcl.GetName()
}

func (self *SeqStructDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.Dcl.GetPrimitiveType()
}

func (self *SeqStructDcl) GetLexemData() IDstSrcInformation {
	return self.Dcl.GetLexemData()
}

func (self *SeqStructDcl) GetOrderId() int {
	return self.Dcl.GetOrderId()
}

func (self *SeqStructDcl) GetForward() bool {
	return self.Dcl.GetForward()
}

func (self *SeqStructDcl) GetImported() bool {
	return self.Dcl.GetImported()
}

func (self *SeqStructDcl) UpdateIdlReference(s ScopeIdentifier) error {
	return self.Dcl.UpdateIdlReference(s)
}

func (self *SeqStructDcl) BuildIdValue() (string, error) {
	return self.Dcl.BuildIdValue()
}

func (self *SeqStructDcl) GetStructExceptionDcl() (bool, error) {
	return self.StructExceptionDcl, nil
}

func (self *SeqStructDcl) GetStructBase() (ScopeIdentifier, error) {
	if self.StructBase != nil {
		return self.StructBase.GetName(), nil
	}
	return "", nil
}

func (self *SeqStructDcl) GetStructMembers() ([]*StructMember, error) {
	return self.StructMembers, nil
}

func (self *SeqStructDcl) IsIStructDcl() bool {
	return true
}

func (self *SeqStructDcl) GetList() IDclArray {
	var result IDclArray = nil
	for _, sm := range self.StructMembers {
		result = append(result, sm)
	}
	return result
}

func (self *SeqStructDcl) Clear() {
	// do nothing
}

func (self *SeqStructDcl) checkForIStructDcl() IStructDcl {
	return self
}

func (self *SeqStructDcl) toIDcl() IDcl {
	return self
}

func NewSeqStructDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	base IDcl,
	exceptionDcl bool,
	forward bool,
	members []*StructMember) *SeqStructDcl {
	return &SeqStructDcl{
		Dcl: NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			forward,
			yaccidl.IdlStruct),
		StructExceptionDcl: exceptionDcl,
		StructBase:         base,
		StructMembers:      members,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*SeqStructDcl)(nil)),
		func() IDcl {
			return new(SeqStructDcl)
		})
}
