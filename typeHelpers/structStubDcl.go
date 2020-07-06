package typeHelpers

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	yaccToken "github.com/bhbosman/yaccidl"
)

type StructStubDcl struct {
	BaseDcl
}

func (self *StructStubDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.StructDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	var err error
	if dclType.StructBase != nil {
		err = self.ResolveType(currentScope, keyMap, dclType.StructBase)
		if err != nil {
			return self.logger.Error(err)
		}
	}

	for _, member := range dclType.StructMembers {
		err = self.ResolveType(currentScope, keyMap, member.MemberType)
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil
}


func (self *StructStubDcl) Start(ctx context.Context) error {
	err := self.BaseDcl.Start(ctx)
	if err != nil {
		return err
	}
	b := true
	if !b {
		return fmt.Errorf("StructDcl not initialized")
	}
	return nil
}

func (self *StructStubDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *StructStubDcl) Init(
	fileInformation *yaccToken.FileInformation) error {
	self.fileInformation = fileInformation

	return nil
}

func (self *StructStubDcl) RemoveScope(dclType *objects.StructDcl, types objects.KnownTypes) error {
	if dclType.StructBase != nil && dclType.StructBase.GetPrimitiveType() == yaccidl.IdlScope {
		var err error
		dclType.StructBase, err = types.Find(dclType.StructBase.GetName())
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *StructStubDcl) SetDestination(dclType objects.IStructDcl, types objects.KnownTypes) error {
	return nil
}

func (self *StructStubDcl) UpdateIdlFolderUsageForDclAsStub(usage objects.FileUsage, fromWhere string, dclType *objects.StructDcl, knownTypes objects.KnownTypes) error {
	err := self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)
	if err != nil {
		return self.logger.Error(err)
	}

	return nil
}

//func (self *StructStubDcl) UpdateIdlFolderUsageForDclAsSkeleton(usage objects.FileUsage, fromWhere string, dclType *objects.StructDcl, knownTypes objects.KnownTypes) error {
//	err := self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)
//	if err != nil {
//		return self.logger.Error(err)
//	}
//	base, err := dclType.GetStructBase()
//	if err != nil {
//		return err
//	}
//	StructBase, _ := knownTypes[base]
//
//	if StructBase == nil {
//		err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.IdlObjectType, knownTypes)
//		if err != nil {
//			return err
//		}
//	} else {
//		if st, ok := StructBase.(objects.IStructDcl); ok {
//			err := self.ResolveFolderUsage(usage, fromWhere, st, knownTypes)
//			if err != nil {
//				return self.logger.Error(err)
//			}
//		}
//		err = self.ResolveFolderUsage(usage, fromWhere, StructBase, knownTypes)
//		if err != nil {
//			return err
//		}
//	}
//	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.ReadAnyType, knownTypes)
//	if err != nil {
//		return err
//	}
//	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.WriteAnyType, knownTypes)
//	if err != nil {
//		return err
//	}
//	usage["__json__"] = objects.NewFindFoldersInUse("__json__", "encoding/json", "")
//	return nil
//}

func NewStructStubDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *StructStubDcl {
	return &StructStubDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("typeHelpers.StructStubDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
