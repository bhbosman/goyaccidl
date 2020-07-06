package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"sort"
)

type WriteInterface struct {
	writeBaseType
	InterfaceHelper *Service.InterfaceHelper
}

func (self *WriteInterface) Write(ctx *FileGenerationContext, v *objects.InterfaceDcl, types objects.KnownTypes) error {
	err := self.WriteHeader(ctx, "Interface", "WriteInterface", v)
	if err != nil {
		return nil
	}


	err = self.WriteInterface(ctx, v, types)
	if err != nil {
		return nil
	}

	err = self.WriteHelperStruct(ctx, v, types)
	if err != nil {
		return err
	}
	err = self.WriteUniqueId01(ctx, v)
	if err != nil {
		return err
	}
	err = self.WriteUniqueId02(ctx, v)
	if err != nil {
		return err
	}
	err = self.WriteExtract(ctx, v, types)
	if err != nil {
		return err
	}
	err = self.WriteInsert(ctx, v, types)
	if err != nil {
		return err
	}

	return nil
}

type declaredOperations struct {
	operations      map[string]*objects.OperationDcl
	duplicatesFound bool
	seq             []string
}

func (o *declaredOperations) AddOperation(operation *objects.OperationDcl) {
	o.operations[operation.OperationName] = operation
	o.seq = append(o.seq, operation.OperationName)
}

func (o *declaredOperations) getOperations() []*objects.OperationDcl {
	var result []*objects.OperationDcl = nil
	sort.Strings(o.seq)
	for _, s := range o.seq {
		result = append(result, o.operations[s])
	}
	return result
}


func (self WriteInterface) WriteOperation(ctx *FileGenerationContext, v *objects.InterfaceDcl, operation *objects.OperationDcl, knownTypes objects.KnownTypes) error {
	if len(operation.Exceptions) > 0 {
		_, _ = fmt.Fprintf(ctx.writer, "\t//Exceptions for : %v\n", self.idlToGoTranslation.ExportMemberName(operation.OperationName))
		for _, exception := range operation.Exceptions {
			typeName := self.idlToGoTranslation.AbsTypeName(
				v.GetLexemData(),
				exception)
			_, _ = fmt.Fprintf(ctx.writer, "\t//\t%v\n", typeName)
		}
	}
	_, _ = fmt.Fprintf(ctx.writer, "\t// Original name: \"%v\"\n", operation.OperationName)
	operationName := self.idlToGoTranslation.ExportMemberName(operation.OperationName)
	inType, err := knownTypes.Find(v.GetName().Append(operation.OperationName).Append("In"))
	if err != nil {
		return err
	}
	inTypeName := self.idlToGoTranslation.AbsTypeName(v.DstSrcInformation, inType)
	outType, err := knownTypes.Find(v.GetName().Append(operation.OperationName).Append("Out"))
	if err != nil {
		return err
	}
	outTypeName := self.idlToGoTranslation.AbsTypeName(v.DstSrcInformation, outType)
	_, _ = fmt.Fprintf(ctx.writer, "\t%v(params %v) (%v, error)", operationName, inTypeName, outTypeName)
	_, _ = fmt.Fprintf(ctx.writer, "\n")
	return nil
}

func (self WriteInterface) WriteAttribute(
	ctx *FileGenerationContext,
	v *objects.InterfaceDcl,
	attribute *objects.AttributeDcl,
	knownTypes objects.KnownTypes) error {
	_, _ = fmt.Fprintf(ctx.writer, "\t// Property %v\n", self.idlToGoTranslation.ExportMemberName(attribute.AttributeName))
	if attribute.AttributeReadable {
		_, _ = fmt.Fprintf(ctx.writer, "\t// Get Property %v\n", self.idlToGoTranslation.ExportMemberName(attribute.AttributeName))
		ans, err := self.idlToGoTranslation.CreateTypeReference(
			v.GetLexemData(),
			nil,
			attribute.AttributeType,
			knownTypes)
		if err != nil {
			return err
		}

		switch ans.PrimitiveType {
		default:
		}
		_, _ = fmt.Fprintf(ctx.writer, "\tGet%v() (%v, error)\n",
			self.idlToGoTranslation.ExportMemberName(attribute.AttributeName),
			ans.Reference)
	}
	if attribute.AttributeWriteable {
		_, _ = fmt.Fprintf(ctx.writer, "\t// Set Property %v\n", self.idlToGoTranslation.ExportMemberName(attribute.AttributeName))

		ans, err := self.idlToGoTranslation.CreateTypeReference(
			v.GetLexemData(),
			nil,
			attribute.AttributeType,
			knownTypes)
		if err != nil {
			return err
		}

		_, _ = fmt.Fprintf(ctx.writer, "\tSet%v(value %v) error\n",
			self.idlToGoTranslation.ExportMemberName(attribute.AttributeName),
			ans.Reference)
	}
	return nil
}

func (self WriteInterface) FindFolderInUser(v *objects.InterfaceDcl, result objects.FileUsage) {

}

func (self WriteInterface) WriteGlobalVariables(w FileGenerationContext, v *objects.InterfaceDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprint(w.writer, "\n")
	_, _ = fmt.Fprint(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self WriteInterface) WritePackageInit(w *FileGenerationContext, v *objects.InterfaceDcl) error {
	return nil
}

func (self WriteInterface) WriteExtract(w *FileGenerationContext, v *objects.InterfaceDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	readAnyTypeValue := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Read(stream %v) (%v, error) {\n", buildTypeName, readAnyTypeValue, buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "\treturn nil, nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteInterface) WriteInsert(w *FileGenerationContext, v *objects.InterfaceDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	writeAnyTypeValue := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.WriteAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Write(stream %v, v %v) error {\n", buildTypeName, writeAnyTypeValue, buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteInterface) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteInterface) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func (self *WriteInterface) WriteInterface(ctx *FileGenerationContext, v *objects.InterfaceDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.Name)
	_, _ = fmt.Fprintf(ctx.writer, "type %v interface {\n", buildTypeName)
	var operations []string
	for item := range self.InterfaceHelper.GetOperationsSorted(v, types).Observe() {
		if operation, ok := item.V.(*objects.OperationDcl); ok {
			err := self.WriteOperation(ctx, v, operation, types)
			if err != nil {
				return err
			}
			operations = append(operations, operation.OperationName)
		}
	}

	for _, attribute := range v.Attributes {
		err := self.WriteAttribute(ctx, v, attribute, types)
		if err != nil {
			return err
		}
	}
	_, _ = fmt.Fprintf(ctx.writer, "}\n")
	_, _ = fmt.Fprintf(ctx.writer, "\n")

	for _, s := range operations{
		_, _ = fmt.Fprintf(ctx.writer, "const %v%vOperation = \"%v\"\n", buildTypeName, self.idlToGoTranslation.BuildTypeName(objects.ScopeIdentifier(s)), s)
	}
	return nil


}

func NewWriteInterface(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService,
	InterfaceHelper *Service.InterfaceHelper) *WriteInterface {
	return &WriteInterface{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		InterfaceHelper: InterfaceHelper,
	}
}
