package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"path/filepath"
)

//noinspection GoSnakeCaseUsage
type WriteStructDcl struct {
	writeBaseType
	helper *WriteStructHelper
}

func (self *WriteStructDcl) Write(
	w *FileGenerationContext,
	v objects.IStructDcl,
	knownTypes objects.KnownTypes) error {
	var err error = nil
	var structExceptionDcl bool
	structExceptionDcl, err = v.GetStructExceptionDcl()
	if err != nil {
		return err
	}

	if structExceptionDcl {
		err = self.WriteHeader(w, "Exception", "WriteStructDcl", v)
		if err != nil {
			return err
		}
	} else {
		err = self.WriteHeader(w, "Struct", "WriteStructDcl", v)
		if err != nil {
			return err
		}
	}
	err = self.helper.WriteStruct(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteUniqueId01(w, v)
	if err != nil {
		return nil
	}

	err = self.helper.WriteStringFunction(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteErrorFunction(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.helper.WriteGoStringFunction(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteExtractValue(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteExtract(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteInsert(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteHelperStruct(w, v, knownTypes)
	if err != nil {
		return err
	}

	//err = self.WriteHelperStructCreateTypeCode(w, v, knownTypes)
	//if err != nil {
	//	return err
	//}

	//err = self.WriteHelperStructAddMember(w, v, knownTypes)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (self *WriteStructDcl) WriteErrorFunction(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
	var structExceptionDcl bool
	var err error
	structExceptionDcl, err = v.GetStructExceptionDcl()
	if err != nil {
		return err
	}
	if structExceptionDcl {
		buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
		_, _ = fmt.Fprintf(w.writer, "func (self *%v) Error() string{\n", buildTypeName)
		buildTypeName = self.idlToGoTranslation.BuildTypeName(v.GetName())
		_, _ = fmt.Fprintf(w.writer, "\treturn \t__fmt__.Sprintf(\"Error of type %v(%%v)\", self.String())\n", buildTypeName)
		_, _ = fmt.Fprintf(w.writer, "}\n")
	}
	return nil
}

func (self *WriteStructDcl) WriteGlobalVariables(w FileGenerationContext, v objects.IStructDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())

	_, _ = fmt.Fprintf(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self *WriteStructDcl) WritePackageInit(w *FileGenerationContext, v objects.IStructDcl) error {
	//
	registrationName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.AddRegistrationFunc)
	newRegistrationInformationName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.NewRegistrationInformationFunc)
	sourceFile := v.GetLexemData().GetSourceLexemData().SourceFileName
	targetFile := filepath.Base(w.TargetFile)
	constName := self.BuildConstName(
		w,
		func() objects.ScopeIdentifier {
			return v.GetName()
		})
	CorbaIIdlObjectName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.CorbaIIdlObject)
	typeName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		v)
	c := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.StructConst)

	_, _ = fmt.Fprintf(w.writer, "\t%v(\n", registrationName)
	_, _ = fmt.Fprintf(w.writer, "\t\t%v(\n", newRegistrationInformationName)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t%v,\n", constName)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t%v,\n", c)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\"%v\",\n", sourceFile)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\"%v\",\n", targetFile)
	_, _ = fmt.Fprintf(w.writer, "\t\t\tfunc() %v {\n", CorbaIIdlObjectName)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\treturn &%v{}\n", typeName)
	_, _ = fmt.Fprintf(w.writer, "\t\t\t},\n")
	members, err := v.GetStructMembers()
	if err != nil {
		return err
	}
	if len(members) > 0 {
		typeNameRandomDataGenerator := self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			self.LanguageTypeService.RandomDataGenerator)
		_, _ = fmt.Fprintf(w.writer, "\t\t\tfunc(generator %v) %v {\n", typeNameRandomDataGenerator, CorbaIIdlObjectName)
		_, _ = fmt.Fprintf(w.writer, "\t\t\t\treturn &%v{}\n", typeName)
		_, _ = fmt.Fprintf(w.writer, "\t\t\t},\n")
	} else {
		typeNameRandomDataGenerator := self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			self.LanguageTypeService.RandomDataGenerator)
		_, _ = fmt.Fprintf(w.writer, "\t\t\tfunc(generator %v) %v {\n", typeNameRandomDataGenerator, CorbaIIdlObjectName)
		_, _ = fmt.Fprintf(w.writer, "\t\t\t\treturn &%v{}\n", typeName)
		_, _ = fmt.Fprintf(w.writer, "\t\t\t},\n")
	}
	_, _ = fmt.Fprintf(w.writer, "\t\t\t__reflect__.TypeOf((*%v)(nil))))\n", typeName)

	return nil
}

func (self *WriteStructDcl) WriteExtractValue(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
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

func (self *WriteStructDcl) WriteExtract(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
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

func (self *WriteStructDcl) WriteInsert(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
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

func (self *WriteStructDcl) Start(ctx context.Context) error {
	err := self.writeBaseType.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (self *WriteStructDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}


func NewWriteStructDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	helper *WriteStructHelper,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteStructDcl {
	return &WriteStructDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		helper: helper,
	}
}
