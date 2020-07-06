package typeHelpers

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
)

type PrimitiveDcl struct {
	BaseDcl
}

func (self *PrimitiveDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.PrimitiveDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	_, _, err := self.FindType(currentScope, keyMap, dclType.Name, dclType.GetLexemData())
	return err
}

func (self *PrimitiveDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.PrimitiveDcl,
	knownTypes objects.KnownTypes) error {
	if usage == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}

	err := self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.ReadAnyType, knownTypes)
	if err != nil {
		return err
	}
	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.WriteAnyType, knownTypes)
	if err != nil {
		return err
	}

	return self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)
}

func (self *PrimitiveDcl) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *PrimitiveDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *PrimitiveDcl) RemoveScope(dclType *objects.PrimitiveDcl, types objects.KnownTypes) error {
	return nil
}

func (self *PrimitiveDcl) SetDestination(dclType *objects.PrimitiveDcl, types objects.KnownTypes) error {
	return nil
}

func NewPrimitiveDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *PrimitiveDcl {
	return &PrimitiveDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("primitiveDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
