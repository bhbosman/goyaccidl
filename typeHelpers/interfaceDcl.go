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
)

type InterfaceDcl struct {
	BaseDcl
	operationsDcl          *Operations
	attributesDcl          *Attributes
	inheritedInterfacesDcl *InheritedInterfaces
	StructDcl              *StructDcl
}

func (self *InterfaceDcl) UpdateIdlReference(currentScope CurrentScope, keyMap objects.LocalKeyMap, dclType *objects.InterfaceDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}

	self.logger.LogWithLevel(10, func(logger *log.Logger) { logger.Printf("UpdateIdlReference types used in interface operations") })
	err := self.operationsDcl.UpdateIdlReferenceForDcl(currentScope, keyMap, dclType.Operations)
	if err != nil {
		return self.logger.Error(err)
	}

	self.logger.LogWithLevel(10, func(logger *log.Logger) {
		logger.Printf("UpdateIdlReference types used in interface attributes operation")
	})
	err = self.attributesDcl.UpdateIdlReferenceForDcl(currentScope, keyMap, dclType.Attributes)
	if err != nil {
		return self.logger.Error(err)
	}

	self.logger.LogWithLevel(10, func(logger *log.Logger) {
		logger.Printf("UpdateIdlReference types used in interface attributes operation")
	})
	err = self.inheritedInterfacesDcl.UpdateIdlReferenceForDcl(currentScope, keyMap, dclType.BaseInterfaces)
	if err != nil {
		return self.logger.Error(err)
	}

	return nil
}

func (self *InterfaceDcl) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	dclType *objects.InterfaceDcl,
	knownTypes objects.KnownTypes) error {

	if usage == nil {
		return self.logger.Error(fmt.Errorf("no usage data"))
	}

	err := self.ResolveFolderUsage(usage, "InterfaceDcl", self.LanguageTypeService.ReadAnyType, knownTypes)
	if err != nil {
		return self.logger.Error(err)
	}

	err = self.ResolveFolderUsage(usage, "InterfaceDcl", self.LanguageTypeService.WriteAnyType, knownTypes)
	if err != nil {
		return self.logger.Error(err)
	}

	//err = self.inheritedInterfacesDcl.UpdateIdlFolderUsageForDcl(
	//	usage,
	//	fmt.Sprintf("%v/%v",
	//		fromWhere,
	//		reflect.TypeOf(self).String()),
	//	dclType.BaseInterfaces,
	//	knownTypes)

	//dealWithDcl := func(dcl objects.IDcl) error {
	//	switch dcl.(type) {
	//	case *objects.SeqStructDcl:
	//		err := self.ResolveFolderUsage(usage, fromWhere, dcl, knownTypes)
	//		return err
	//
	//	case *objects.StructDcl:
	//		err := self.ResolveFolderUsage(usage, fromWhere, dcl, knownTypes)
	//		return err
	//	default:
	//		return nil
	//	}
	//}

	//for _, operation := range dclType.Operations {
	//	err := dealWithDcl(operation.OperationType)
	//	if err != nil {
	//		return err
	//	}
	//}

	if err != nil {
		return self.logger.Error(err)
	}

	return nil
}

func (self *InterfaceDcl) Start(ctx context.Context) error {
	err := self.BaseDcl.Start(ctx)
	if err != nil {
		return nil
	}
	b := true
	b = b && self.operationsDcl != nil
	b = b && self.attributesDcl != nil
	b = b && self.inheritedInterfacesDcl != nil
	if !b {
		return self.logger.Error(fmt.Errorf("InterfaceDcl not initialized"))
	}
	return nil
}

func (self *InterfaceDcl) Stop(ctx context.Context) error {
	self.operationsDcl = nil
	self.attributesDcl = nil
	self.inheritedInterfacesDcl = nil
	return self.BaseDcl.Stop(ctx)
}

func (self *InterfaceDcl) RemoveScope(dclType *objects.InterfaceDcl, types objects.KnownTypes) error {
	for i, dcl := range dclType.BaseInterfaces {
		var err error
		dclType.BaseInterfaces[i], err = types.Find(dcl.GetName())
		if err != nil {
			return err
		}
	}

	for _, dcl := range dclType.Attributes {
		var err error
		dcl.AttributeType, err = types.Find(dcl.AttributeType.GetName())
		if err != nil {
			return self.logger.Error(err)
		}
	}
	for _, dcl := range dclType.Operations {
		var err error
		dcl.OperationType, err = types.Find(dcl.OperationType.GetName())
		if err != nil {
			return self.logger.Error(err)
		}

		for exceptionIndex, exception := range dcl.Exceptions {
			var err error
			dcl.Exceptions[exceptionIndex], err = types.Find(exception.GetName())
			if err != nil {
				return self.logger.Error(err)
			}
		}

		for paramIndex, param := range dcl.OperationsParams {
			var err error
			dcl.OperationsParams[paramIndex].ParameterDcl, err = types.Find(param.ParameterDcl.GetName())
			if err != nil {
				return self.logger.Error(err)
			}
		}
	}

	return nil
}

func (self *InterfaceDcl) SetDestination(dclType *objects.InterfaceDcl, types objects.KnownTypes) error {
	targetFileInfo := self.fileInformation.AddPath(self.appCtx.OutputFolder)
	err := dclType.SetDestination(targetFileInfo.FolderId, self.appCtx.OutputFolder)
	if err != nil {
		return err
	}
	return nil
}

func NewInterfaceDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	flattenOperations *Operations,
	flattenAttributes *Attributes,
	flattenBaseInterfaces *InheritedInterfaces,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation,
	StructDcl *StructDcl) *InterfaceDcl {
	return &InterfaceDcl{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("InterfaceDclLogger"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
		operationsDcl:          flattenOperations,
		attributesDcl:          flattenAttributes,
		inheritedInterfacesDcl: flattenBaseInterfaces,
		StructDcl:              StructDcl,
	}
}
