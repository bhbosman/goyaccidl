package objects

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type IImportedDcl interface {
	IDcl
	IsImportedDcl() bool
}

type ImportedDcl struct {
	BaseDcl IDcl `json:"-"`
}

func (self *ImportedDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.BaseDcl.DclResolveFolderUsage()
}

func (self *ImportedDcl) SetDestination(folderId, folder string) error {
	return nil
}

func (self *ImportedDcl) GetTypePrefix() string {
	return self.BaseDcl.GetTypePrefix()
}

func (self *ImportedDcl) GetIsArray() bool {
	return self.BaseDcl.GetIsArray()
}

func (self *ImportedDcl) GetArrayValue() []int64 {
	return self.BaseDcl.GetArrayValue()
}

func (self *ImportedDcl) GetDcl() IDcl {
	return self.BaseDcl
}

func (self *ImportedDcl) MarshalJSON() ([]byte, error) {
	TypeSpec, err := NewIDclStream(self.BaseDcl)
	if err != nil {
		return nil, err
	}

	type Alias ImportedDcl
	return json.Marshal(
		&struct {
			*Alias
			BaseDcl IDclStream `json:"base_dcl"`
		}{
			Alias:   (*Alias)(self),
			BaseDcl: TypeSpec,
		})
}

func (self *ImportedDcl) UnmarshalJSON(bytes []byte) error {
	type Alias ImportedDcl
	data := &struct {
		*Alias
		BaseDcl IDclStream `json:"base_dcl"`
	}{
		Alias: (*Alias)(self),
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	self.BaseDcl, err = data.BaseDcl.GetDcl()
	if err != nil {
		return err
	}

	return nil
}

func (self *ImportedDcl) SetTypePrefix(s string) {
	self.BaseDcl.SetTypePrefix(s)
}

func (self *ImportedDcl) GetImported() bool {
	return self.BaseDcl.GetImported()
}

func (self *ImportedDcl) GetName() ScopeIdentifier {
	return self.BaseDcl.GetName()
}

func (self *ImportedDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.BaseDcl.GetPrimitiveType()
}

func (self *ImportedDcl) GetLexemData() IDstSrcInformation {
	return self.BaseDcl.GetLexemData()
}

func (self *ImportedDcl) GetOrderId() int {
	return self.BaseDcl.GetOrderId()
}

func (self *ImportedDcl) GetForward() bool {
	return self.BaseDcl.GetForward()
}

func (self *ImportedDcl) UpdateIdlReference(s ScopeIdentifier) error {
	return self.BaseDcl.UpdateIdlReference(s)
}

func (self *ImportedDcl) BuildIdValue() (string, error) {
	return self.BaseDcl.BuildIdValue()
}

func (self *ImportedDcl) GetStructExceptionDcl() (bool, error) {
	if self.IsIStructDcl() {
		if structDcl, ok := self.BaseDcl.(IStructDcl); ok {
			return structDcl.GetStructExceptionDcl()
		}
		return false, fmt.Errorf("marked as struct, but it is not")
	}
	return false, fmt.Errorf("imported type is not a struct")
}

func (self *ImportedDcl) GetStructBase() (ScopeIdentifier, error) {
	if self.IsIStructDcl() {
		if structDcl, ok := self.BaseDcl.(IStructDcl); ok {
			return structDcl.GetStructBase()
		}
		return "", fmt.Errorf("marked as struct, but it is not")
	}
	return "", fmt.Errorf("imported type is not a struct")
}

func (self *ImportedDcl) GetStructMembers() ([]*StructMember, error) {
	if self.IsIStructDcl() {
		if structDcl, ok := self.BaseDcl.(IStructDcl); ok {
			return structDcl.GetStructMembers()
		}
		return nil, fmt.Errorf("marked as struct, but it is not")
	}
	return nil, fmt.Errorf("imported type is not a struct")
}

func (self *ImportedDcl) IsIStructDcl() bool {
	return self.BaseDcl.GetPrimitiveType() == yaccidl.IdlStruct
}

func (self *ImportedDcl) TypeDeclPrimitiveType() (yaccidl.PrimitiveType, error) {
	if self.IsTypeDefDcl() {
		if t, ok := self.BaseDcl.(ITypeDefDcl); ok {
			return t.TypeDeclPrimitiveType()
		}
		return yaccidl.IdlInvalid, fmt.Errorf("marked as typedef, but it is not")
	}
	return yaccidl.IdlInvalid, fmt.Errorf("imported type is not a typdef")
}

func (self *ImportedDcl) TypeDeclName() (ScopeIdentifier, error) {
	if self.IsTypeDefDcl() {
		if t, ok := self.BaseDcl.(ITypeDefDcl); ok {
			return t.TypeDeclName()
		}
		return "", fmt.Errorf("marked as ITypeDefDcl, but it is not")

	}
	return "", fmt.Errorf("imported type is not a typdef")
}

func (self *ImportedDcl) IsTypeDefDcl() bool {
	return self.BaseDcl.GetPrimitiveType() == yaccidl.IdlTypedef
}

func (self *ImportedDcl) checkForIStructDcl() IStructDcl {
	return self
}

func (self *ImportedDcl) checkForITypeDefDcl() ITypeDefDcl {
	return self
}

func (self *ImportedDcl) IsImportedDcl() bool {
	return true
}

func NewImportedDcl(
	BaseDcl IDcl) *ImportedDcl {
	return &ImportedDcl{
		BaseDcl: BaseDcl,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*ImportedDcl)(nil)),
		func() IDcl {
			return new(ImportedDcl)
		})
}
