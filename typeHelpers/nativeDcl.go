package typeHelpers

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
)

type NativeDcl struct {
	BaseDcl
}

func (self *NativeDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.NativeDcl) error {
	return nil
}

func (self *NativeDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.NativeDcl,
	knownTypes objects.KnownTypes) error {
	return nil
}

func (self *NativeDcl) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *NativeDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *NativeDcl) RemoveScope(dclType *objects.NativeDcl, types objects.KnownTypes) error {
	return nil
}

func (self *NativeDcl) SetDestination(dclType *objects.NativeDcl, types objects.KnownTypes) error {
	return nil
}

func NewNative(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *NativeDcl {
	return &NativeDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenNative"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
