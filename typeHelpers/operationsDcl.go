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

type Operations struct {
	BaseDcl
}

func (self *Operations) UpdateIdlReferenceForDcl(currentScope CurrentScope, keyMap objects.LocalKeyMap, operations []*objects.OperationDcl) error {
	if keyMap == nil {
		return self.logger.Error(fmt.Errorf("no map data"))
	}

	for _, operation := range operations {
		err := self.ResolveType(currentScope, keyMap, operation.OperationType)
		if err != nil {
			return self.logger.Error(err)
		}

		err = self.UpdateIdlReferenceForDclForOperation(currentScope, keyMap, operation)
		if err != nil {
			return self.logger.Error(err)
		}

		for _, exception := range operation.Exceptions {
			err := self.ResolveType(currentScope, keyMap, exception)
			if err != nil {
				return self.logger.Error(err)
			}

		}
	}

	return nil
}

func (self *Operations) UpdateIdlReferenceForDclForOperation(currentScope CurrentScope, keyMap objects.LocalKeyMap, operations *objects.OperationDcl) error {
	for _, param := range operations.OperationsParams {
		err := self.ResolveType(currentScope, keyMap, param.ParameterDcl)
		if err != nil {
			return self.logger.Error(err)
		}

	}
	return nil
}

func (self *Operations) UpdateIdlFolderUsageForDcl(
	usage objects.FileUsage,
	fromWhere string,
	operations []*objects.OperationDcl,
	knownTypes objects.KnownTypes) error {

	//if usage == nil {
	//	return self.logger.Error(fmt.Errorf("no usage data"))
	//}
	//for _, operation := range operations {
	//	ans, err := self.DclHelpers.FindPrimitiveTypeForUpdateIdlFolderUsage(operation.OperationType, knownTypes)
	//	if err != nil {
	//		return err
	//	}
	//	err = self.ResolveFolderUsage(usage, "Operations", ans.Dcl, knownTypes)
	//	if err != nil {
	//		return self.logger.Error(err)
	//	}
	//	for _, param := range operation.OperationsParams {
	//		ans, err = self.DclHelpers.FindPrimitiveTypeForUpdateIdlFolderUsage(param.ParameterDcl, knownTypes)
	//		if err != nil {
	//			return err
	//		}
	//		err := self.ResolveFolderUsage(usage, "Operations", ans.Dcl, knownTypes)
	//		if err != nil {
	//			return self.logger.Error(err)
	//		}
	//	}
	//	for _, exception := range operation.Exceptions {
	//		exceptionName := exception.GetName()
	//		fmt.Sprintf("%v", exceptionName)
	//		ans, err = self.DclHelpers.FindPrimitiveTypeForUpdateIdlFolderUsage(exception, knownTypes)
	//		if err != nil {
	//			return err
	//		}
	//		err := self.ResolveFolderUsage(usage, "Operations", ans.Dcl, knownTypes)
	//		if err != nil {
	//			return self.logger.Error(err)
	//		}
	//	}
	//}
	return nil
}

func (self *Operations) Start(ctx context.Context) error {
	return self.BaseDcl.Start(ctx)
}

func (self *Operations) Stop(ctx context.Context) error {
	return self.BaseDcl.Stop(ctx)
}

func NewOperationsDcl(
	appCtx *ctx2.GoYaccAppCtx,
	factory *log2.LogFactory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) *Operations {
	return &Operations{
		BaseDcl: NewBaseDcl(
			appCtx,
			factory.Create("FlattenOperations"),
			IdlDefaultTypes,
			LanguageTypeService,
			DclHelpers,
			fileInformation),
	}

}
