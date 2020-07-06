package objects

import (
	"fmt"
	"github.com/bhbosman/yaccidl"
	"strings"
)

type BaseDcl struct {
	PrimitiveType     yaccidl.PrimitiveType `json:"primitive_type"`
	Name              ScopeIdentifier       `json:"name"`
	DstSrcInformation *DstSrcInformation    `json:"dst_src_information"`
	OrderId           int                   `json:"order_id"`
	Forward           bool                  `json:"forward"`
	TypePrefix        string                `json:"type_prefix"`
	Imported          bool                  `json:"imported"`
	IsArray           bool                  `json:"is_array"`
	ArrayValue        []int64               `json:"array_value"`
}

func (self *BaseDcl) DclResolveFolderUsage() (id string, folderName string) {
	return self.GetLexemData().DclResolveFolderUsage()
}

func (d *BaseDcl) SetDestination(folderId, folder string) error {
	d.DstSrcInformation.Destination.Folder = folder
	d.DstSrcInformation.Destination.FolderId = folderId

	return nil
}

func (self *BaseDcl) GetTypePrefix() string {
	return self.TypePrefix
}

func (self *BaseDcl) GetIsArray() bool {
	return self.IsArray
}

func (self *BaseDcl) GetArrayValue() []int64 {
	return self.ArrayValue
}

func (self *BaseDcl) SetTypePrefix(s string) {
	self.TypePrefix = s
}

func (self *BaseDcl) GetImported() bool {
	return self.Imported
}

func (self *BaseDcl) BuildIdValue() (string, error) {
	ss := strings.Split(string(self.Name), "::")
	structName := strings.Join(ss, "/")
	version := "1.0"
	if self.TypePrefix == "" {
		return fmt.Sprintf("IDL:%v:%v", structName, version), nil
	}
	return fmt.Sprintf("IDL:%v/%v:%v", self.TypePrefix, structName, version), nil
}

func (self BaseDcl) GetForward() bool {
	return self.Forward
}

func (self *BaseDcl) GetOrderId() int {
	return self.OrderId
}

func (self BaseDcl) GetPrimitiveType() yaccidl.PrimitiveType {
	return self.PrimitiveType
}

func (self BaseDcl) GetName() ScopeIdentifier {
	return self.Name
}

func (self BaseDcl) GetLexemData() IDstSrcInformation {
	return self.DstSrcInformation
}

func (self *BaseDcl) UpdateIdlReference(s ScopeIdentifier) error {
	self.Name = s
	//self.DstSrcInformation.SourceLexemData = lexemData
	return nil
}

func NewBaseDcl(
	primitiveType yaccidl.PrimitiveType,
	name ScopeIdentifier,
	isArray bool,
	arrayValue []int64,
	orderId int,
	lexemData *yaccidl.LexemValue,
	forward bool,
	imported bool) *BaseDcl {
	return &BaseDcl{
		PrimitiveType: primitiveType,
		Name:          name,
		DstSrcInformation: &DstSrcInformation{
			SourceLexemData: lexemData,
			Destination:     Destination{},
		},
		OrderId:    orderId,
		Forward:    forward,
		TypePrefix: "",
		Imported:   imported,
		IsArray:    isArray,
		ArrayValue: arrayValue,
	}
}
