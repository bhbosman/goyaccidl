package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
)

type StructMember struct {
	MemberName string
	IsArray    bool
	ArrayValue []int64
	MemberType IDcl
}

func (self *StructMember) DclResolveFolderUsage() (id string, folderName string) {
	return self.MemberType.DclResolveFolderUsage()
}

func (self *StructMember) SetDestination(folderId, folder string) error {
	return self.MemberType.SetDestination(folderId, folder)
}

func (self *StructMember) GetTypePrefix() string {
	return self.MemberType.GetTypePrefix()
}

func (self *StructMember) GetIsArray() bool {
	return self.MemberType.GetIsArray()
}

func (self *StructMember) GetArrayValue() []int64 {
	return self.MemberType.GetArrayValue()
}

func (self *StructMember) SetTypePrefix(s string) {
	self.MemberType.SetTypePrefix(s)
}

func (self *StructMember) GetImported() bool {
	return self.MemberType.GetImported()
}

type structMemberJson struct {
	MemberName string      `json:"member_name"`
	MemberType ScopeStream `json:"member_type"`
}

func (self *StructMember) UnmarshalJSON(bytes []byte) error {
	jsonData := &structMemberJson{}
	err := json.Unmarshal(bytes, jsonData)
	if err != nil {
		return err
	}
	self.MemberName = jsonData.MemberName
	self.MemberType, err = jsonData.MemberType.GetDcl()
	if err != nil {
		return err
	}
	return nil
}

func (self *StructMember) MarshalJSON() ([]byte, error) {
	raw, err := NewScopeStream(self.MemberType)
	if err != nil {
		return nil, err
	}
	return json.Marshal(
		&structMemberJson{
			MemberName: self.MemberName,
			MemberType: raw,
		})
}

func (self *StructMember) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.MemberType.GetPrimitiveType()
}

func (self *StructMember) BuildIdValue() (string, error) {
	return self.MemberType.BuildIdValue()
}

func (self *StructMember) GetName() ScopeIdentifier {
	return self.MemberType.GetName()
}

func (self *StructMember) GetLexemData() IDstSrcInformation {
	return self.MemberType.GetLexemData()
}

func (self *StructMember) GetForward() bool {
	return self.MemberType.GetForward()
}

func (self *StructMember) UpdateIdlReference(ss ScopeIdentifier) error {
	return self.MemberType.UpdateIdlReference(ss)
}

func (self *StructMember) GetOrderId() int {

	return self.MemberType.GetOrderId()
}

func NewStructMember(
	name string,
	IsArray bool,
	ArrayValue []int64,
	typeDcl IDcl) *StructMember {
	return &StructMember{
		MemberName: name,
		IsArray:    IsArray,
		ArrayValue: ArrayValue,
		MemberType: typeDcl,
	}
}
