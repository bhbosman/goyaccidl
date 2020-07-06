package typeHelpers

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
	"log"
	"reflect"
)

type UnionDcl struct {
	BaseDcl
	typeDefDcl   *TypeDefDcl
	constantDcl  *ConstantDcl
	nativeDcl    *NativeDcl
	scopeDcl     *ScopeDcl
	primitiveDcl *PrimitiveDcl
	sequenceDcl  *SequenceDcl
	structDcl    *StructDcl
}

func (self *UnionDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.UnionDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}

	err := self.ResolveType(currentScope, keyMap, dclType.SwitchTypeDcl)
	if err != nil {
		return self.logger.Error(err)
	}

	for _, member := range dclType.UnionBodies {
		for _, caseValue := range member.CaseValues {
			err = self.ResolveType(currentScope, keyMap, caseValue.Value)
		}
		err = self.ResolveType(currentScope, keyMap, member.GetTypeSpec())
		if err != nil {
			return self.logger.Error(err)
		}
	}
	return nil
}

func (self *UnionDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.UnionDcl,
	knownTypes objects.KnownTypes) error {
	err := self.ResolveFolderUsage(usage, fromWhere, dclType, knownTypes)
	if err != nil {
		return self.logger.Error(err)
	}
	err = self.ResolveFolderUsage(usage, fromWhere, self.LanguageTypeService.IdlObjectType, knownTypes)
	if err != nil {
		return err
	}

	usage["__json__"] = objects.NewFindFoldersInUse("__json__", "encoding/json", "")

	if usage == nil {
		return self.logger.Error(fmt.Errorf("no usage data"))
	}
	for _, member := range dclType.UnionBodies {
		if member.Identifier == "profile" {
			fmt.Errorf("ddd")
		}
		for _, caseValue := range member.CaseValues {
			switch v := caseValue.Value.(type) {
			case *objects.ConstantValue:
				err = self.constantDcl.UpdateIdlFolderUsageForDcl(
					usage,
					fmt.Sprintf("%v/%v",
						fromWhere,
						reflect.TypeOf(self).String()),
					v,
					knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			default:
				s := fmt.Sprintf("33333333333333No handler found for %v(%v). LexemData: %v", reflect.TypeOf(v).String(), v.GetName(), v.GetLexemData())
				self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
			}
		}
	}
	dclType.SwitchTypeDcl, err = knownTypes.Find(dclType.SwitchTypeDcl.GetName())
	if err != nil {
		return err
	}
	return self.structDcl.UpdateIdlFolderUsageForDcl(
		usage,
		//true,
		fmt.Sprintf("%v/%v",
			fromWhere,
			reflect.TypeOf(self).String()),

		dclType,
		knownTypes)
	return nil
}

func (self *UnionDcl) Start(ctx context.Context) error {
	err := self.BaseDcl.Start(ctx)
	if err != nil {
		return err
	}
	b := true
	b = b && self.typeDefDcl != nil
	b = b && self.constantDcl != nil
	b = b && self.nativeDcl != nil
	b = b && self.scopeDcl != nil
	b = b && self.primitiveDcl != nil
	b = b && self.sequenceDcl != nil
	b = b && self.structDcl != nil
	if !b {
		return fmt.Errorf("UnionDcl not initialized")
	}
	return nil
}

func (self *UnionDcl) RemoveScope(dclType *objects.UnionDcl, knownTypes objects.KnownTypes) error {
	//self.structDcl.RemoveScope()
	for _, member := range dclType.UnionBodies {
		for _, caseValue := range member.CaseValues {
			switch v := caseValue.Value.(type) {
			case *objects.ScopeDcl:
				var err error
				caseValue.Value, err = knownTypes.Find(caseValue.Value.GetName())
				if err != nil {
					return err
				}
			//case *objects.ConstantValue:
			//	err := self.constantDcl.RemoveScope(v, knownTypes)
			//	if err != nil {
			//		return self.logger.Error(err)
			//	}
			default:
				s := fmt.Sprintf("33333333333333No handler found for %v(%v). LexemData: %v", reflect.TypeOf(v).String(), v.GetName(), v.GetLexemData())
				self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
			}
		}
		switch v := member.TypeSpec.(type) {
		case *objects.TypeDefDcl:
			err := self.typeDefDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		case *objects.ConstantValue:
			err := self.constantDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		case *objects.ScopeDcl:
			var err error
			member.TypeSpec, err = knownTypes.Find(member.TypeSpec.GetName())
			if err != nil {
				return self.logger.Error(err)
			}
			//err := self.scopeDcl.RemoveScope(v, knownTypes)
			//if err != nil {
			//	return self.logger.Error(err)
			//}
		case *objects.SequenceTypeDcl:
			err := self.sequenceDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		case *objects.NativeDcl:
			err := self.nativeDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		case *objects.StructDcl:
			err := self.structDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		case *objects.PrimitiveDcl:
			err := self.primitiveDcl.RemoveScope(v, knownTypes)
			if err != nil {
				return self.logger.Error(err)
			}
		default:
			s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(v).String(), v.GetName(), v.GetLexemData())
			self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
		}
	}
	return nil
}

func (self *UnionDcl) SetDestination(dclType *objects.UnionDcl, types objects.KnownTypes) error {
	targetFileInfo := self.fileInformation.AddPath(self.appCtx.OutputFolder)
	err := dclType.SetDestination(targetFileInfo.FolderId, self.appCtx.OutputFolder)
	if err != nil {
		return err
	}
	return nil
}

func (self *UnionDcl) Stop(ctx context.Context) error {
	self.typeDefDcl = nil
	self.constantDcl = nil
	self.nativeDcl = nil
	self.scopeDcl = nil
	self.primitiveDcl = nil
	self.sequenceDcl = nil
	self.structDcl = nil

	return self.BaseDcl.Stop(ctx)
}

func (self *UnionDcl) Init(
	interfaceDclHelper *InterfaceDcl,
	typeDefDclHelper *TypeDefDcl,
	constantDclHelper *ConstantDcl,
	nativeDclHelper *NativeDcl,
	scopeDclHelper *ScopeDcl,
	primitiveDclHelper *PrimitiveDcl,
	sequenceDcl *SequenceDcl,
	structDcl *StructDcl) error {
	self.typeDefDcl = typeDefDclHelper
	self.constantDcl = constantDclHelper
	self.nativeDcl = nativeDclHelper
	self.scopeDcl = scopeDclHelper
	self.primitiveDcl = primitiveDclHelper
	self.sequenceDcl = sequenceDcl
	self.structDcl = structDcl
	return nil
}

func NewUnionDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	idlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *UnionDcl {
	return &UnionDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenSequenceDcl"),
			idlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}
}
