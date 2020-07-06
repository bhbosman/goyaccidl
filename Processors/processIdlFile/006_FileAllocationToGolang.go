package processIdlFile

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"

	"github.com/bhbosman/goyaccidl/writers/gowriter"
)

type FileAllocationToGolang struct {
	WriteMembers           *gowriter.WriteMembers
	CreateKnownTypeService *CreateKnownTypeService
}

func (self FileAllocationToGolang) Name() string {
	return self.WriteMembers.Log.Name()
}

func (self FileAllocationToGolang) Run(incoming interface{}) (interface{}, error) {
	if flattenDeclsToFileAllocationOut, ok := incoming.(*FlattenDeclsToFileAllocationOut); ok {
		var err error
		err = self.BuildFiles(flattenDeclsToFileAllocationOut.Files, flattenDeclsToFileAllocationOut.ImportedDeclarationItems)
		if err != nil {
			return nil, self.WriteMembers.Log.Error(err)
		}
		result := NewFileAllocationToGolangOut(
			flattenDeclsToFileAllocationOut.Files,
			//flattenDeclsToFileAllocationOut.ScopeDeclarationItems,
			flattenDeclsToFileAllocationOut.ImportedDeclarationItems)
		return result, nil
	}
	return nil, self.WriteMembers.Log.Error(fmt.Errorf("type not supported objects.FileUsages"))
}

func (self FileAllocationToGolang) BuildFiles(
	files objects.FileUsages,
	ImportedDeclarationItems []*objects.ScopeItem) error {

	knownTypes, err := self.CreateKnownTypeService.CreateKnownTypes(files.ToDclArray(), ImportedDeclarationItems)
	if err != nil {
		return err
	}

	for _, fileUsage := range files {
		err := self.WriteMembers.Render(knownTypes, fileUsage)
		if err != nil {
			return self.WriteMembers.Log.Error(err)
		}
	}
	return nil
}

func NewFileAllocationToGolang(
	logFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	fileResolve *ctx2.ResolveFileName,
	GoWriterFactory *gowriter.Factory,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	CreateKnownTypeService *CreateKnownTypeService) (*FileAllocationToGolang, error) {
	writeMembers, err := gowriter.NewWriteMembers(
		ctx,
		logFactory.Create("RenderToGo"),
		GoWriterFactory,
		fileResolve,
		IdlDefaultTypes)
	if err != nil {
		return nil, err
	}
	return &FileAllocationToGolang{
		WriteMembers:           writeMembers,
		CreateKnownTypeService: CreateKnownTypeService,
	}, nil
}

type FileAllocationToGolangOut struct {
	Files                    objects.FileUsages   `json:"files"`
	ImportedDeclarationItems []*objects.ScopeItem `json:"imported_declaration_items"`
}

func NewFileAllocationToGolangOut(
	files objects.FileUsages,
	ImportedDeclarationItems []*objects.ScopeItem) *FileAllocationToGolangOut {
	return &FileAllocationToGolangOut{
		Files:                    files,
		ImportedDeclarationItems: ImportedDeclarationItems,
	}
}
