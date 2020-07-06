package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type ITypeDefDcl interface {
	IDcl
	IsTypeDefDcl() bool
	TypeDeclName() (ScopeIdentifier, error)
	TypeDeclPrimitiveType() (yaccidl.PrimitiveType, error)
	GetDcl() IDcl
}

type TypeDefDcl struct {
	Dcl      IDcl
	TypeDecl IDcl
}

func (self *TypeDefDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.Dcl.DclResolveFolderUsage()
}

func (self *TypeDefDcl) SetDestination(folderId, folder string) error {
	err := self.TypeDecl.SetDestination(folderId, folder)
	if err != nil {
		return err
	}

	return self.Dcl.SetDestination(folderId, folder)
}

func (self *TypeDefDcl) GetTypePrefix() string {
	return self.Dcl.GetTypePrefix()
}

func (self *TypeDefDcl) GetIsArray() bool {
	return self.Dcl.GetIsArray()
}

func (self *TypeDefDcl) GetArrayValue() []int64 {
	return self.Dcl.GetArrayValue()
}

func (self *TypeDefDcl) GetDcl() IDcl {
	return self.Dcl
}

func (self *TypeDefDcl) SetTypePrefix(s string) {
	self.Dcl.SetTypePrefix(s)
}

func (self *TypeDefDcl) MarshalJSON() ([]byte, error) {
	var Dcl IDclStream
	var TypeDecl IDclStream
	var err error
	Dcl, err = NewIDclStream(self.Dcl)
	if err != nil {
		return nil, err
	}
	TypeDecl, err = NewIDclStream(self.TypeDecl)
	if err != nil {
		return nil, err
	}
	return json.Marshal(
		&struct {
			Dcl      IDclStream `json:"dcl"`
			TypeDecl IDclStream `json:"type_decl"`
		}{
			Dcl:      Dcl,
			TypeDecl: TypeDecl,
		})
}

func (self *TypeDefDcl) UnmarshalJSON(bytes []byte) error {
	data := &struct {
		Dcl      IDclStream `json:"dcl"`
		TypeDecl IDclStream `json:"type_decl"`
	}{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	self.Dcl, err = data.Dcl.GetDcl()
	if err != nil {
		return err
	}
	self.TypeDecl, err = data.TypeDecl.GetDcl()
	if err != nil {
		return err
	}
	return nil
}

func (self *TypeDefDcl) GetName() ScopeIdentifier {
	return self.Dcl.GetName()
}

func (self *TypeDefDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.Dcl.GetPrimitiveType()
}

func (self *TypeDefDcl) GetLexemData() IDstSrcInformation {
	return self.Dcl.GetLexemData()
}

func (self *TypeDefDcl) GetOrderId() int {
	return self.Dcl.GetOrderId()
}

func (self *TypeDefDcl) GetForward() bool {
	return self.Dcl.GetForward()
}

func (self *TypeDefDcl) GetImported() bool {
	return self.Dcl.GetImported()
}

func (self *TypeDefDcl) UpdateIdlReference(s ScopeIdentifier) error {
	return self.Dcl.UpdateIdlReference(s)
}

func (self *TypeDefDcl) BuildIdValue() (string, error) {
	return self.Dcl.BuildIdValue()
}

func (self *TypeDefDcl) TypeDeclPrimitiveType() (yaccidl.PrimitiveType, error) {
	return self.TypeDecl.GetPrimitiveType(), nil
}

func (self *TypeDefDcl) TypeDeclName() (ScopeIdentifier, error) {
	return self.TypeDecl.GetName(), nil
}

func (self *TypeDefDcl) IsTypeDefDcl() bool {
	return true
}

func (self *TypeDefDcl) asITypeDefDcl() ITypeDefDcl {
	return self
}

func NewTypeDefDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	typeDecl IDcl) *TypeDefDcl {
	return &TypeDefDcl{
		Dcl: NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			false,
			yaccidl.IdlTypedef),
		TypeDecl: typeDecl,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*TypeDefDcl)(nil)),
		func() IDcl {
			return new(TypeDefDcl)
		})
}
