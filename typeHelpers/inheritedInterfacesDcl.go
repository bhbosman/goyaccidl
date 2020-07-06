package typeHelpers

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
)

type InheritedInterfaces struct {
	BaseDcl
}

func (self *InheritedInterfaces) UpdateIdlReferenceForDcl(
	currentScope CurrentScope,
	keyMap objects.LocalKeyMap,
	interfaces objects.IDclArray) error {
	for _, s := range interfaces {
		err := self.ResolveType(currentScope, keyMap, s)
		if err != nil {
			return self.logger.Error(err)
		}

	}
	return nil
}

func (self *InheritedInterfaces) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	interfaces objects.IDclArray,
	knownTypes objects.KnownTypes) error {

	for _, s := range interfaces {
		err := self.ResolveFolderUsage(usage, fromWhere, s, knownTypes)
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil

}

func (self *InheritedInterfaces) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *InheritedInterfaces) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func NewInheritedInterfacesDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *InheritedInterfaces {
	return &InheritedInterfaces{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("InheritedInterfaces"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
