package typeHelpers

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
	"reflect"
)

type EnumDcl struct {
	BaseDcl
}

func (self *EnumDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	v *objects.EnumDcl,
	types objects.KnownTypes) error {

	err := self.ResolveFolderUsage(
		usage,
		fmt.Sprintf("%v/%v",
			fromWhere,
			reflect.TypeOf(self).String()),
		self.LanguageTypeService.ReadAnyType,
		types)
	if err != nil {
		return err
	}
	err = self.ResolveFolderUsage(
		usage,
		fmt.Sprintf("%v/%v",
			fromWhere,
			reflect.TypeOf(self).String()),
		self.LanguageTypeService.WriteAnyType,
		types)
	if err != nil {
		return err
	}
	return nil
}

func NewEnumDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *EnumDcl {
	return &EnumDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("typeHelpers.EnumDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
