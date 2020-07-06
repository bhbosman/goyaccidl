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

type SequenceDcl struct {
	BaseDcl
}

func (self *SequenceDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType objects.ISequenceTypeDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	err := self.ResolveType(currentScope, keyMap, dclType)
	if err != nil {
		return self.logger.Error(err)
	}
	err = self.ResolveType(currentScope, keyMap, dclType.GetSequenceType())
	if err != nil {
		return self.logger.Error(err)
	}
	return nil
}

func (self *SequenceDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType objects.ISequenceTypeDcl,
	knownTypes objects.KnownTypes) error {
	if usage == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	err := self.ResolveFolderUsage(
		usage,
		fromWhere,
		self.LanguageTypeService.ReadAnyType,
		knownTypes)
	if err != nil {
		return err
	}
	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.WriteAnyType, knownTypes)
	if err != nil {
		return err
	}
	return nil
}

func (self *SequenceDcl) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *SequenceDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *SequenceDcl) SetDestination(dclType *objects.SequenceTypeDcl, types objects.KnownTypes) error {
	targetFileInfo := self.fileInformation.AddPath(self.appCtx.OutputFolder)
	err := dclType.SetDestination(targetFileInfo.FolderId, self.appCtx.OutputFolder)
	if err != nil {
		return err
	}
	return nil
}

func (self *SequenceDcl) RemoveScope(v *objects.SequenceTypeDcl, types objects.KnownTypes) error {
	var err error
	v.SequenceType, err = types.Find(v.SequenceType.GetName())
	if err != nil {
		return err
	}
	return nil
}

func NewSequenceDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *SequenceDcl {
	return &SequenceDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenSequenceDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
