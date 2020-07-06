package typeHelpers

import (
	"context"
	"errors"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	yaccToken "github.com/bhbosman/yaccidl"
	"reflect"
)

type TypeDefDcl struct {
	BaseDcl
	sequenceDcl  *SequenceDcl
	primitiveDcl *PrimitiveDcl
	scopeDcl     *ScopeDcl
	structDcl    *StructDcl
	EnumDcl      *EnumDcl
}

func (self *TypeDefDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dcl *objects.TypeDefDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}
	switch dclType := dcl.TypeDecl.(type) {
	case *objects.SequenceTypeDcl:
		err := self.sequenceDcl.UpdateIdlReference(currentScope, keyMap, dclType)
		if err != nil {
			return self.logger.Error(err)
		}

	case *objects.PrimitiveDcl:
		switch dclType.PrimitiveType {
		case yaccidl.IdlObjectKind:
		default:
			err := self.primitiveDcl.UpdateIdlReference(currentScope, keyMap, dclType)
			if err != nil {
				return self.logger.Error(err)
			}
		}

	case *objects.ScopeDcl:
		err := self.scopeDcl.UpdateIdlReference(currentScope, keyMap, dclType)
		if err != nil {
			return self.logger.Error(err)
		}

	default:
		err := self.logger.Error(fmt.Errorf("needs an implementation for %v", reflect.TypeOf(dclType).String()))
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil
}

func (self *TypeDefDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dcl *objects.TypeDefDcl,
	knownTypes objects.KnownTypes) error {
	if usage == nil {
		return self.logger.Error(fmt.Errorf("no usage data"))
	}
	switch dclType := dcl.TypeDecl.(type) {
	case *objects.StructDcl, *objects.SeqStructDcl:
		panic("fix remove typedef")
	case *objects.EnumDcl:
		return self.EnumDcl.UpdateIdlFolderUsageForDcl(
			usage,
			fmt.Sprintf("%v/%v(%v %v)",
				fromWhere,
				reflect.TypeOf(self).String(),
				dclType.GetName(),
				dclType.GetPrimitiveType()),
			dclType,
			knownTypes)
	case *objects.TypeDefDcl:
		return self.UpdateIdlFolderUsageForDcl(
			usage,
			fmt.Sprintf("%v/%v(%v %v)",
				fromWhere,
				reflect.TypeOf(self).String(),
				dclType.GetName(),
				dclType.GetPrimitiveType()),
			dclType,
			knownTypes)
	case *objects.PrimitiveDcl:
		switch dclType.PrimitiveType {
		case yaccidl.IdlObjectKind:
			self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.IdlObjectType, knownTypes)
			return nil
		default:
			return self.primitiveDcl.UpdateIdlFolderUsageForDcl(
				usage,
				fmt.Sprintf("%v/%v(%v %v)",
					fromWhere,
					reflect.TypeOf(self).String(),
					dclType.GetName(),
					dclType.GetPrimitiveType()),
				dclType,
				knownTypes)
		}
	default:
		return self.logger.Error(fmt.Errorf("needs an implementation for %v", reflect.TypeOf(dclType).String()))
	}
}

func (self *TypeDefDcl) Start(ctx context.Context) error {
	err := self.BaseDcl.Start(ctx)
	if err != nil {
		return err
	}
	b := true
	b = b && self.sequenceDcl != nil
	b = b && self.primitiveDcl != nil
	b = b && self.scopeDcl != nil
	if !b {
		return fmt.Errorf("TypeDefDcl not initialized")
	}
	return nil
}

func (self *TypeDefDcl) Stop(ctx context.Context) error {
	self.sequenceDcl = nil
	self.primitiveDcl = nil
	self.scopeDcl = nil

	return self.BaseDcl.Stop(ctx)
}

var TypeDefDclRemoveScope = errors.New("typeHelpers::TypeDefDcl::RemoveScope error")

func (self *TypeDefDcl) RemoveScope(dclType *objects.TypeDefDcl, types objects.KnownTypes) error {
	switch dclType.TypeDecl.GetPrimitiveType() {
	case yaccToken.IdlSequence:
		if seqType, ok := dclType.TypeDecl.(*objects.SequenceTypeDcl); ok {
			return self.sequenceDcl.RemoveScope(seqType, types)
		}
		return TypeDefDclRemoveScope
	case yaccidl.IdlScope:
		var err error
		dclType.TypeDecl, err = types.Find(dclType.TypeDecl.GetName())
		if err != nil {
			return err
		}
		//return TypeDefDclRemoveScope
	default:
		return nil
	}
	return nil
}

func (self *TypeDefDcl) SetDestination(dclType *objects.TypeDefDcl, types objects.KnownTypes) error {
	//targetFileInfo := self.fileInformation.AddPath(self.appCtx.OutputFolder)
	//err := dclType.SetDestination(targetFileInfo.FolderId, self.appCtx.OutputFolder)
	//if err != nil {
	//	return  err
	//}
	return nil
}

func NewTypeDefDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	flattenSequenceDcl *SequenceDcl,
	flattenPrimitiveDcl *PrimitiveDcl,
	flattenScopeDcl *ScopeDcl,
	DclHelpers *Service.DclHelpers,
	structDcl *StructDcl,
	LanguageTypeService *Service.LanguageTypeService,
	EnumDcl *EnumDcl,
	fileInformation *yaccToken.FileInformation) *TypeDefDcl {
	return &TypeDefDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("typeHelpers.TypeDefDcl"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
		sequenceDcl:  flattenSequenceDcl,
		primitiveDcl: flattenPrimitiveDcl,
		scopeDcl:     flattenScopeDcl,
		structDcl:    structDcl,
		EnumDcl:      EnumDcl,
	}
}
