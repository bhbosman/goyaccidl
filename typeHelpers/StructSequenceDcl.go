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
	"reflect"
)

type StructSequenceDcl struct {
	BaseDcl
}

func (self *StructSequenceDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.StructDcl) error {
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

func (self *StructSequenceDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	primary bool,
	fromWhere string,
	dclType objects.IBaseStructDcl,
	knownTypes objects.KnownTypes) error {

	err := self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)
	if err != nil {
		return self.logger.Error(err)
	}

	base, err := dclType.GetStructBase()
	if err != nil {
		return err
	}
	StructBase, _ := knownTypes[base]

	if StructBase == nil {
		err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.IdlObjectType, knownTypes)
		if err != nil {
			return err
		}
	} else {
		if st, ok := StructBase.(objects.IStructDcl); ok {
			err = self.UpdateIdlFolderUsageForDcl(
				usage,
				false,
				fmt.Sprintf("%v/%v",
					fromWhere,
					reflect.TypeOf(self).String()),
				st,
				knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		}
		err = self.ResolveFolderUsage(usage, fromWhere, StructBase, knownTypes)
		if err != nil {
			return err
		}
	}
	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.ReadAnyType, knownTypes)
	if err != nil {
		return err
	}
	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.WriteAnyType, knownTypes)
	if err != nil {
		return err
	}

	if primary {
		usage["__json__"] = objects.NewFindFoldersInUse("__json__", "encoding/json", "")
		usage["__reflect__"] = objects.NewFindFoldersInUse("__reflect__", "reflect", "")
	}
	StructExceptionDcl, err := dclType.GetStructExceptionDcl()
	if err != nil {
		return err
	}
	if StructExceptionDcl {
		usage["__fmt__"] = objects.NewFindFoldersInUse("__fmt__", "fmt", "")
	}

	if usage == nil {
		return self.logger.Error(fmt.Errorf("no usage data"))
	}
	StructMembers, err := dclType.GetStructMembers()
	if err != nil {
		return err
	}

	for _, member := range StructMembers {
		find, b := self.LanguageTypeService.Find(member.MemberType.GetPrimitiveType())
		if b {
			if find.GetLexemData().GetSourceLexemData() == self.LanguageTypeService.NoLexem {
			} else {
				err = self.ResolveFolderUsage(usage, fromWhere+"ddd", find, knownTypes)
				if err != nil {
					return err
				}
			}
		} else {
			err = self.ResolveFolderUsage(usage, fromWhere+"ddd", member.MemberType, knownTypes)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *StructSequenceDcl) Start(ctx context.Context) error {
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

func (self *StructSequenceDcl) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func (self *StructSequenceDcl) Init(
	fileInformation *yaccToken.FileInformation) error {

	self.fileInformation = fileInformation

	return nil
}

func (self *StructSequenceDcl) RemoveScope(dclType *objects.SeqStructDcl, types objects.KnownTypes) error {
	if dclType.StructBase != nil && dclType.StructBase.GetPrimitiveType() == yaccidl.IdlScope {
		var err error
		dclType.StructBase, err = types.Find(dclType.StructBase.GetName())
		if err != nil {
			return err
		}
	}
	for _, member := range dclType.StructMembers {
		if member.MemberType != nil {
			var err error
			member.MemberType, err = types.Find(member.MemberType.GetName())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *StructSequenceDcl) SetDestination(dclType objects.IStructDcl, types objects.KnownTypes) error {
	return nil
}

func NewSeqStructDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *StructSequenceDcl {
	return &StructSequenceDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("typeHelpers.StructDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
