package processIdlFile

import (
	"context"
	"fmt"
	"github.com/bhbosman/goyaccidl/Service"

	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/typeHelpers"
	"github.com/bhbosman/yaccidl"
	"log"
	"path/filepath"
	"reflect"
)

type FlattenDeclsToFileAllocation struct {
	logger                 *log2.SubSystemLogger
	StructDclHelper        *typeHelpers.StructDcl
	StructStubDclHelper    *typeHelpers.StructStubDcl
	UnionDclHelper         *typeHelpers.UnionDcl
	InterfaceDclHelper     *typeHelpers.InterfaceDcl
	TypeDefDclHelper       *typeHelpers.TypeDefDcl
	ConstantDclHelper      *typeHelpers.ConstantDcl
	NativeDclHelper        *typeHelpers.NativeDcl
	ScopeDclHelper         *typeHelpers.ScopeDcl
	PrimitiveDclHelper     *typeHelpers.PrimitiveDcl
	IdlDefaultTypes        *Service.IdlDefaultTypes
	ctx                    *ctx2.GoYaccAppCtx
	IdlToGoTranslation     *Service.IdlToGoTranslation
	fileInformation        *yaccidl.FileInformation
	DclHelpers             *Service.DclHelpers
	CreateKnownTypeService *CreateKnownTypeService
	SequenceDcl            *typeHelpers.SequenceDcl
	EnumDcl                *typeHelpers.EnumDcl
	RemoveTypedef          *Service.RemoveTypedef
	SeqStructDcl           *typeHelpers.StructSequenceDcl
	ExpandInterfaces       *Service.ExpandInterfaces
}

func (self *FlattenDeclsToFileAllocation) Name() string {
	return self.logger.Name()
}

func (self *FlattenDeclsToFileAllocation) FindFile(files map[string]*objects.FileData, fileName string) (*objects.FileData, bool) {
	fd, ok := files[fileName]
	return fd, ok
}

func (self *FlattenDeclsToFileAllocation) Run(input interface{}) (interface{}, error) {
	if flattenDeclaredTypesOut, ok := input.(*objects.FlattenDeclaredTypesOut); ok {
		self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Printf("%v", len(flattenDeclaredTypesOut.InputStreamDcl.Members)) })

		obs := flattenDeclaredTypesOut.InputStreamDcl.Members.ToObs().
			Map(
				func(ctx context.Context, i interface{}) (interface{}, error) {
					if array, ok := i.(objects.IDclArray); ok {
						knownTypes, err := self.CreateKnownTypeService.CreateKnownTypes(
							array,
							flattenDeclaredTypesOut.ImportedDeclarationItems)
						if err != nil {
							return nil, err
						}
						dcl, err := self.RemoveTypedef.Remove(array, knownTypes)
						if err != nil {
							return nil, err
						}
						return dcl, nil
					}
					return nil, fmt.Errorf("can not typecast to objects.IDclArray")
				}).
			Map(
				func(ctx context.Context, i interface{}) (interface{}, error) {
					if array, ok := i.(objects.IDclArray); ok {
						knownTypes, err := self.CreateKnownTypeService.CreateKnownTypes(
							array,
							flattenDeclaredTypesOut.ImportedDeclarationItems)
						if err != nil {
							return nil, err
						}
						return self.ExpandInterfaces.Expand(array, knownTypes)
					}
					return nil, fmt.Errorf("can not typecast to objects.IDclArray")

				}).
			Map(
				func(ctx context.Context, i interface{}) (interface{}, error) {
					if array, ok := i.(objects.IDclArray); ok {
						files, err := self.CreateFileStructure(
							array,
							flattenDeclaredTypesOut.ImportedDeclarationItems)
						if err != nil {
							return nil, self.logger.Error(err)
						}
						result := NewFlattenDeclsToFileAllocationOut(
							files,
							flattenDeclaredTypesOut.ImportedDeclarationItems)
						return result, nil
					}
					return nil, fmt.Errorf("can not typecast to objects.IDclArray")
				})
		obs.Run()
		result := obs.FirstOrDefault(nil)
		get, err := result.Get()
		if err != nil {
			return nil, err
		}
		if get.E != nil {
			return nil, get.E
		}
		if get.V != nil {
			if result, ok := get.V.(*FlattenDeclsToFileAllocationOut); ok {
				return result, nil
			}
		}

	}
	return nil, self.logger.Error(fmt.Errorf("input type wrong"))
}

func (self *FlattenDeclsToFileAllocation) FindCreate(
	files map[string]*objects.FileData,
	dcl objects.IDcl,
	targetFolder string) (*objects.FileData, error) {
	var key string
	var targetFileName string

	targetFileNameFunc := func(sourceFileName, v2 string) string {
		sourceFileName = filepath.Base(sourceFileName)
		//v2 = self.IdlToGoTranslation.BuildTypeName(dcl)
		return fmt.Sprintf("xdl_%v", self.IdlToGoTranslation.BuildTypeName(dcl.GetName()))
	}

	switch dcl.GetPrimitiveType() {
	case yaccidl.IdlInterface:
		key = filepath.Join(
			dcl.GetLexemData().GetSourceFolderName(),
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
		targetFileName = targetFileNameFunc(
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
	case yaccidl.IdlStruct:
		key = filepath.Join(
			dcl.GetLexemData().GetSourceFolderName(),
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
		targetFileName = targetFileNameFunc(
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
	case yaccidl.IdlEnum:
		key = filepath.Join(
			dcl.GetLexemData().GetSourceFolderName(),
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
		targetFileName = targetFileNameFunc(
			dcl.GetLexemData().GetSourceFileName(),
			dcl.GetName().Combined())
	//case yaccidl.IdlTypedef:
	//	key = filepath.Join(
	//		dcl.GetLexemData().GetSourceFolderName(),
	//		dcl.GetLexemData().GetSourceFileName(),
	//		dcl.GetName().Name())
	//	targetFileName = targetFileNameFunc(
	//		dcl.GetLexemData().GetSourceFileName(),
	//		dcl.GetName().Name())
	default:
		key = filepath.Join(
			dcl.GetLexemData().GetSourceFolderName(),
			dcl.GetLexemData().GetSourceFileName())
		targetFileName = dcl.GetLexemData().GetSourceFileName()

	}

	targetFileInfo := self.fileInformation.AddPath(targetFolder)
	err := dcl.SetDestination(targetFileInfo.FolderId, targetFolder)
	if err != nil {
		return nil, err
	}

	if fd, ok := self.FindFile(files, key); ok {
		return fd, nil
	}
	fd := objects.NewFileData(
		dcl.GetLexemData().GetSourceFolderId(),
		targetFileInfo.FolderId,
		targetFileName)

	files[key] = fd
	return fd, nil
}

func (self *FlattenDeclsToFileAllocation) CreateFileStructure(
	dcls objects.IDclArray,

	ImportedDeclarationItems []*objects.ScopeItem) (objects.FileUsages, error) {
	targetFolder := self.ctx.OutputFolder
	files, err := self.CreateFiles(dcls, targetFolder)
	if err != nil {
		return nil, err
	}
	knownTypes, err := self.CreateKnownTypeService.CreateKnownTypes(
		dcls,
		ImportedDeclarationItems)
	if err != nil {
		return nil, err
	}
	err = self.RemoveScope(files, knownTypes, ImportedDeclarationItems)
	if err != nil {
		return nil, err
	}

	err = self.SetDestination(files, targetFolder, knownTypes)
	if err != nil {
		return nil, err
	}

	err = self.UpdateIdlFolderUsageForDcl(files, knownTypes)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (self *FlattenDeclsToFileAllocation) UpdateIdlFolderUsageForDcl(
	files objects.FileUsages,
	knownTypes objects.KnownTypes) error {
	for _, file := range files {
		for _, fileMember := range file.Members {
			switch dclType := fileMember.(type) {
			case *objects.InterfaceDcl:
				err := self.InterfaceDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.TypeDefDcl:
				err := self.TypeDefDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.StructDcl:
				switch dclType.StructType {
				case objects.StructTypeAsDefined, objects.StructTypeAsMethodParamIn, objects.StructTypeAsMethodParamOut:
					err := self.StructDclHelper.UpdateIdlFolderUsageForDcl(
						file.FolderUsage,
						fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
						dclType,
						knownTypes)
					if err != nil {
						return err
					}

				case objects.StructTypeAsIdlStub:
					err := self.StructStubDclHelper.UpdateIdlFolderUsageForDclAsStub(
						file.FolderUsage,
						fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
						dclType,
						knownTypes)
					if err != nil {
						return err
					}
				case objects.StructTypeAsIdlSkeleton:
					err := self.StructDclHelper.UpdateIdlFolderUsageForDclAsSkeleton(
						file.FolderUsage,
						fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
						dclType,
						knownTypes)
					if err != nil {
						return err
					}
				}

			case *objects.SeqStructDcl:
				err := self.StructDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}

			case *objects.UnionDcl:
				err := self.UnionDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.ConstantValue:
				err := self.ConstantDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.NativeDcl:
				err := self.NativeDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.EnumDcl:
				err := self.EnumDcl.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
			case *objects.ModuleDcl:
				if len(dclType.Members) > 0 {
					err := self.logger.Error(fmt.Errorf("error: At this point module members must have been flattened"))
					if err != nil {
						return err
					}
				}
			case *objects.SequenceTypeDcl:
				self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Printf("Implement %v", reflect.TypeOf(dclType).String()) })
				break
			case *objects.ScopeDcl:
				err := self.ScopeDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
				break
			case *objects.PrimitiveDcl:
				err := self.PrimitiveDclHelper.UpdateIdlFolderUsageForDcl(
					file.FolderUsage,
					fmt.Sprintf("%v(%v)", reflect.TypeOf(self).String(), fileMember.GetName()),
					dclType,
					knownTypes)
				if err != nil {
					return err
				}
				break
			default:
				s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(fileMember).String(), fileMember.GetName(), fileMember.GetLexemData())
				self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
				err := fmt.Errorf(s)
				return err

			}
		}
	}
	return nil
}

func (self *FlattenDeclsToFileAllocation) CreateFiles(dcl objects.IDclArray, targetFolder string) (objects.FileUsages, error) {
	files := make(objects.FileUsages)
	for _, baseDcl := range dcl {
		if dcl, ok01 := baseDcl.(objects.IDcl); ok01 {
			if _, ok02 := dcl.(*objects.PrimitiveDcl); !ok02 {
				fd, err := self.FindCreate(files, dcl, targetFolder)
				if err != nil {
					return nil, err
				}
				fd.Add(dcl)
			}
		}
	}
	return files, nil
}

func (self *FlattenDeclsToFileAllocation) RemoveScope(
	files objects.FileUsages,
	knownTypes objects.KnownTypes,
	ImportedDeclarationItems []*objects.ScopeItem) error {
	for _, item := range knownTypes {
		err := self.InternalRemoveScope(item, knownTypes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *FlattenDeclsToFileAllocation) InternalRemoveScope(fileMember objects.IDcl, knownTypes objects.KnownTypes) error {
	switch dclType := fileMember.(type) {
	case *objects.InterfaceDcl:
		return self.InterfaceDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.TypeDefDcl:
		return self.TypeDefDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.SeqStructDcl:
		return self.SeqStructDcl.RemoveScope(dclType, knownTypes)
	case *objects.StructDcl:
		return self.StructDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.UnionDcl:
		return self.UnionDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.ImportedDcl:
		return nil

	case *objects.ConstantValue:
		return self.ConstantDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.NativeDcl:
		return self.NativeDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.EnumDcl:
		return nil
	case *objects.ModuleDcl:
		if len(dclType.Members) > 0 {
			return self.logger.Error(fmt.Errorf("Error: At this point module members must have been flattened"))
		}
		return nil
	//case *objects.SequenceTypeDcl:
	//	self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Printf("Implement %v", reflect.TypeOf(dclType).String()) })
	//	return nil
	//	break
	case *objects.ScopeDcl:
		return self.ScopeDclHelper.RemoveScope(dclType, knownTypes)
	case *objects.PrimitiveDcl:
		return self.PrimitiveDclHelper.RemoveScope(dclType, knownTypes)
	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(fileMember).String(), fileMember.GetName(), fileMember.GetLexemData())
		self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
		err := fmt.Errorf(s)
		return err
	}
}

func (self *FlattenDeclsToFileAllocation) SetDestination(files objects.FileUsages, targetFile string, knownTypes objects.KnownTypes) error {
	for _, file := range files {
		for _, fileMember := range file.Members {
			switch dclType := fileMember.(type) {
			case *objects.InterfaceDcl:
				err := self.InterfaceDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.TypeDefDcl:
				err := self.TypeDefDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.StructDcl:
				err := self.StructDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.SeqStructDcl:
				err := self.StructDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.UnionDcl:
				err := self.UnionDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.ConstantValue:
				err := self.ConstantDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.NativeDcl:
				err := self.NativeDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
			case *objects.EnumDcl:
				break
			case *objects.ModuleDcl:
				if len(dclType.Members) > 0 {
					err := self.logger.Error(fmt.Errorf("error: At this point module members must have been flattened"))
					if err != nil {
						return err
					}
				}
			case *objects.SequenceTypeDcl:
				err := self.SequenceDcl.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
				break
			case *objects.ScopeDcl:
				err := self.ScopeDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
				break
			case *objects.PrimitiveDcl:
				err := self.PrimitiveDclHelper.SetDestination(dclType, knownTypes)
				if err != nil {
					return self.logger.Error(err)
				}
				break
			default:
				s := fmt.Sprintf("SetDestination: No handler found for %v(%v). LexemData: %v", reflect.TypeOf(fileMember).String(), fileMember.GetName(), fileMember.GetLexemData())
				self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
				err := fmt.Errorf(s)
				return err
			}
		}
	}
	return nil
}

func NewFlattenDeclsToFileAllocation(
	logFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	StructDclHelper *typeHelpers.StructDcl,
	InterfaceDclHelper *typeHelpers.InterfaceDcl,
	TypeDefDclHelper *typeHelpers.TypeDefDcl,
	ConstantDclHelper *typeHelpers.ConstantDcl,
	NativeDclHelper *typeHelpers.NativeDcl,
	ScopeDclHelper *typeHelpers.ScopeDcl,
	PrimitiveDclHelper *typeHelpers.PrimitiveDcl,
	UnionDclHelper *typeHelpers.UnionDcl,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	IdlToGoTranslation *Service.IdlToGoTranslation,
	fileInformation *yaccidl.FileInformation,
	DclHelpers *Service.DclHelpers,
	CreateKnownTypeService *CreateKnownTypeService,
	SequenceDcl *typeHelpers.SequenceDcl,
	EnumDcl *typeHelpers.EnumDcl,
	RemoveTypedef *Service.RemoveTypedef,
	SeqStructDcl *typeHelpers.StructSequenceDcl,
	ExpandInterfaces *Service.ExpandInterfaces,
	StructStubDclHelper    *typeHelpers.StructStubDcl) *FlattenDeclsToFileAllocation {
	return &FlattenDeclsToFileAllocation{
		logger:                 logFactory.Create("FlattenDeclsToFileAllocation"),
		StructDclHelper:        StructDclHelper,
		StructStubDclHelper:    StructStubDclHelper,
		UnionDclHelper:         UnionDclHelper,
		InterfaceDclHelper:     InterfaceDclHelper,
		TypeDefDclHelper:       TypeDefDclHelper,
		ConstantDclHelper:      ConstantDclHelper,
		NativeDclHelper:        NativeDclHelper,
		ScopeDclHelper:         ScopeDclHelper,
		PrimitiveDclHelper:     PrimitiveDclHelper,
		IdlDefaultTypes:        IdlDefaultTypes,
		ctx:                    ctx,
		IdlToGoTranslation:     IdlToGoTranslation,
		fileInformation:        fileInformation,
		DclHelpers:             DclHelpers,
		CreateKnownTypeService: CreateKnownTypeService,
		SequenceDcl:            SequenceDcl,
		EnumDcl:                EnumDcl,
		RemoveTypedef:          RemoveTypedef,
		SeqStructDcl:           SeqStructDcl,
		ExpandInterfaces:       ExpandInterfaces,
	}
}

type FlattenDeclsToFileAllocationOut struct {
	Files objects.FileUsages `json:"files"`
	//ScopeDeclarationItems    objects.ScopeDeclarationItems `json:"scope_declaration_items"`
	ImportedDeclarationItems []*objects.ScopeItem `json:"imported_declaration_items"`
}

func NewFlattenDeclsToFileAllocationOut(
	files objects.FileUsages,
	//ScopeDeclarationItems objects.ScopeDeclarationItems,
	importedDeclarationItems []*objects.ScopeItem) *FlattenDeclsToFileAllocationOut {
	return &FlattenDeclsToFileAllocationOut{
		Files: files,
		//ScopeDeclarationItems:    ScopeDeclarationItems,
		ImportedDeclarationItems: importedDeclarationItems,
	}
}
