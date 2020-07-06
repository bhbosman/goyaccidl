package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"reflect"
)

//noinspection GoSnakeCaseUsage
type WriteStructSkeletonDcl struct {
	writeBaseType
	helper *WriteStructHelper
}

func (self *WriteStructSkeletonDcl) Write(
	w *FileGenerationContext,
	v objects.IStructDcl,
	knownTypes objects.KnownTypes) error {
	var err error = nil

	err = self.WriteHeader(w, "Struct", reflect.TypeOf(self).String(), v)
	if err != nil {
		return err
	}

	err = self.WriteStruct(w, v, knownTypes)
	if err != nil {
		return err
	}


	return nil
}




func (self *WriteStructSkeletonDcl) WriteExtractValue(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
	var structBase objects.ScopeIdentifier
	var err error
	structBase, err = v.GetStructBase()
	if err != nil {
		return err
	}
	structBaseDcl, _ := knownTypes[structBase]

	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	readAnyTypeValue := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) ReadValue(stream %v) error {\n", buildTypeName, readAnyTypeValue)
	base, _ := self.helper.findBaseRepresentation(
		w,
		v.GetLexemData(),
		structBaseDcl,
		knownTypes)
	_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")
	_, _ = fmt.Fprintf(w.writer, "\terr = self.%v.ReadValue(stream)\n", base.Name())
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	var StructMembers []*objects.StructMember
	StructMembers, err = v.GetStructMembers()
	if err != nil {
		return err
	}
	if len(StructMembers) > 0 {
		for _, member := range StructMembers {
			err := self.helper.WriteStructMemberExtractValue(w, v, member, 0, knownTypes)
			if err != nil {
				return err
			}
		}
	}
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSkeletonDcl) WriteExtract(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	readAnyTypeValue := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) Read(stream %v) error {\n", buildTypeName, readAnyTypeValue)
	_, _ = fmt.Fprintf(w.writer, "\terr := self.ReadValue(stream)\n")
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSkeletonDcl) WriteInsert(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	writeAnyTypeValue := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.WriteAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) Write(stream %v) error {\n", buildTypeName, writeAnyTypeValue)

	var StructMembers []*objects.StructMember
	var err error
	StructMembers, err = v.GetStructMembers()
	if err != nil {
		return err
	}

	var structBase objects.ScopeIdentifier
	structBase, err = v.GetStructBase()
	if err != nil {
		return err
	}

	structBaseDcl, _ := knownTypes[structBase]
	if structBaseDcl != nil {
	}
	_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")
	base, _ := self.helper.findBaseRepresentation(
		w,
		v.GetLexemData(),
		structBaseDcl,
		knownTypes)
	_, _ = fmt.Fprintf(w.writer, "\terr = self.%v.Write(stream)\n", base.Name())
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	for _, member := range StructMembers {
		err := self.helper.WriteStructMemberInsert(w, v, member, 0, knownTypes)
		if err != nil {
			return err
		}
	}
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSkeletonDcl) Start(ctx context.Context) error {
	err := self.writeBaseType.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (self *WriteStructSkeletonDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func (self *WriteStructSkeletonDcl) WriteStruct(w *FileGenerationContext, v objects.IBaseStructDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v struct {\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}


func NewWriteStructSkeletonDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	helper *WriteStructHelper,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteStructSkeletonDcl {
	return &WriteStructSkeletonDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		helper: helper,
	}
}
