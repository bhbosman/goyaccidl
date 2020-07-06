package objects

import (
	"encoding/json"
	"github.com/bhbosman/yaccidl"
	"reflect"
)

type IBaseStructDcl interface {
	IDcl
	GetStructBase() (ScopeIdentifier, error)
	GetStructMembers() ([]*StructMember, error)
	GetStructExceptionDcl() (bool, error)
}

type IStructDcl interface {
	IBaseStructDcl
	IsIStructDcl() bool
}

type StructType uint8

const StructTypeAsDefined StructType = 0
const StructTypeAsIdlStub StructType = 1
const StructTypeAsIdlSkeleton StructType = 2
const StructTypeAsMethodParamIn StructType = 3
const StructTypeAsMethodParamOut StructType = 4

type StructDcl struct {
	Dcl                 IDcl            `json:"-"`
	StructBase          IDcl            `json:"-"`
	StructMembers       []*StructMember `json:"struct_members"`
	StructExceptionDcl  bool            `json:"struct_exception_dcl"`
	StructType          StructType      `json:"struct_type"`
	ImplementInterfaces IDclArray       `json:"-"`
}

func (self *StructDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.Dcl.DclResolveFolderUsage()
}

func (self *StructDcl) SetDestination(folderId, folder string) error {
	return self.Dcl.SetDestination(folderId, folder)
}

func (self *StructDcl) GetTypePrefix() string {
	return self.Dcl.GetTypePrefix()
}

func (self *StructDcl) GetIsArray() bool {
	return self.Dcl.GetIsArray()
}

func (self *StructDcl) GetArrayValue() []int64 {
	return self.Dcl.GetArrayValue()
}

func (self *StructDcl) MarshalJSON() ([]byte, error) {
	dcl, err := NewIDclStream(self.Dcl)
	if err != nil {
		return nil, err
	}
	var structBase IDclStream
	structBase, err = NewIDclStream(self.StructBase)
	if err != nil {
		return nil, err
	}

	type Alias StructDcl
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

func (self *StructDcl) UnmarshalJSON(bytes []byte) error {
	type Alias StructDcl
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

func (self *StructDcl) SetTypePrefix(s string) {
	self.Dcl.SetTypePrefix(s)
}

func (self *StructDcl) GetName() ScopeIdentifier {
	return self.Dcl.GetName()
}

func (self *StructDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.Dcl.GetPrimitiveType()
}

func (self *StructDcl) GetLexemData() IDstSrcInformation {
	return self.Dcl.GetLexemData()
}

func (self *StructDcl) GetOrderId() int {
	return self.Dcl.GetOrderId()
}

func (self *StructDcl) GetForward() bool {
	return self.Dcl.GetForward()
}

func (self *StructDcl) GetImported() bool {
	return self.Dcl.GetImported()
}

func (self *StructDcl) UpdateIdlReference(s ScopeIdentifier) error {
	return self.Dcl.UpdateIdlReference(s)
}

func (self *StructDcl) BuildIdValue() (string, error) {
	return self.Dcl.BuildIdValue()
}

func (self *StructDcl) GetStructExceptionDcl() (bool, error) {
	return self.StructExceptionDcl, nil
}

func (self *StructDcl) GetStructBase() (ScopeIdentifier, error) {
	if self.StructBase != nil {
		return self.StructBase.GetName(), nil
	}
	return "", nil
}

func (self *StructDcl) GetStructMembers() ([]*StructMember, error) {
	return self.StructMembers, nil
}

func (self *StructDcl) IsIStructDcl() bool {
	return true
}

func (self *StructDcl) GetList() IDclArray {
	var result IDclArray = nil
	for _, sm := range self.StructMembers {
		result = append(result, sm)
	}
	return result
}

func (self *StructDcl) Clear() {
	// do nothing
}

func (self *StructDcl) checkForIStructDcl() IStructDcl {
	return self
}

func (self *StructDcl) toIDcl() IDcl {
	return self
}

func NewStructDcl(
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	base IDcl,
	exceptionDcl bool,
	forward bool,
	StructType StructType,
	members []*StructMember,
	implementInterfaces IDclArray) *StructDcl {
	return &StructDcl{
		Dcl: NewDcl(
			name,
			isArray,
			arrayValue,
			orderId,
			lexemData,
			forward,
			yaccidl.IdlStruct),
		StructBase:          base,
		StructMembers:       members,
		StructExceptionDcl:  exceptionDcl,
		StructType:          StructType,
		ImplementInterfaces: implementInterfaces,
	}
}

func init() {
	RegisterIDclCodec(
		reflect.TypeOf((*StructDcl)(nil)),
		func() IDcl {
			return new(StructDcl)
		})
}
