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

type ScopeDcl struct {
	BaseDcl
}

func (self *ScopeDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.ScopeDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	return self.ResolveType(currentScope, keyMap, dclType)
}

func (self *ScopeDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.ScopeDcl,
	knownTypes objects.KnownTypes) error {
	if usage == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	return self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)

}

func (self *ScopeDcl) Start(context.Context) error {
	return nil
}

func (self *ScopeDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *ScopeDcl) RemoveScope(dclType *objects.ScopeDcl, types objects.KnownTypes) error {
	return nil
}

func (self *ScopeDcl) SetDestination(dclType *objects.ScopeDcl, types objects.KnownTypes) error {
	return nil
}

func NewScopeDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *ScopeDcl {
	return &ScopeDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("ScopeDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
