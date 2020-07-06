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

type ConstantDcl struct {
	BaseDcl
}

func (self *ConstantDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.ConstantValue) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	return self.ResolveType(currentScope, keyMap, dclType.ConstantType)
}

func (self *ConstantDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.ConstantValue,
	knownTypes objects.KnownTypes) error {
	if usage == nil {
		return self.logger.Error(fmt.Errorf("No map data"))
	}
	return self.ResolveFolderUsage(
		usage,
		fromWhere,
		dclType,
		knownTypes)
}

func (self *ConstantDcl) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *ConstantDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *ConstantDcl) RemoveScope(dclType *objects.ConstantValue, types objects.KnownTypes) error {
	return nil
}

func (self *ConstantDcl) SetDestination(dclType *objects.ConstantValue, types objects.KnownTypes) error {
	targetFileInfo := self.fileInformation.AddPath(self.appCtx.OutputFolder)
	err := dclType.SetDestination(targetFileInfo.FolderId, self.appCtx.OutputFolder)
	if err != nil {
		return err
	}
	return nil
}

func NewConstantDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *ConstantDcl {
	return &ConstantDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenConstants"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
