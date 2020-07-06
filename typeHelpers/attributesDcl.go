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

type Attributes struct {
	BaseDcl
}

type CurrentScope map[objects.ScopeIdentifier]objects.ScopeIdentifier

func (self *Attributes) UpdateIdlReferenceForDcl(currentScope CurrentScope, keyMap objects.LocalKeyMap, attributes []*objects.AttributeDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("No map data"))
	}
	for _, attribute := range attributes {
		err := self.ResolveType(currentScope, keyMap, attribute.AttributeType)
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil
}

func (self *Attributes) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	attributes []*objects.AttributeDcl,
	knownTypes objects.KnownTypes) error {
	for _, attribute := range attributes {
		err := self.ResolveFolderUsage(
			usage,
			fromWhere,
			attribute.AttributeType,
			knownTypes)
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil
}

func (self *Attributes) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *Attributes) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func NewAttributesDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *Attributes {
	return &Attributes{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenAttributes"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
