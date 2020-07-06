package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	"path/filepath"
)

//noinspection GoSnakeCaseUsage
type WriteStructSequenceDcl struct {
	writeBaseType
}

func (self *WriteStructSequenceDcl) Write(
	w *FileGenerationContext,
	v objects.IStructDcl,
	knownTypes objects.KnownTypes) error {
	var err error = nil

	err = self.WriteHeader(w, "Struct", "WriteStructSequenceDcl", v)
	if err != nil {
		return err
	}

	err = self.WriteStruct(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteUniqueId01(w, v)
	if err != nil {
		return nil
	}

	err = self.WriteStringFunction(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteErrorFunction(w, v, knownTypes)
	if err != nil {
		return err
	}

	err = self.WriteGoStringFunction(w, v, knownTypes)
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

	return nil
}

func (self *WriteStructSequenceDcl) WriteErrorFunction(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
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

func (self *WriteStructSequenceDcl) WriteGlobalVariables(w FileGenerationContext, v objects.IStructDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())

	_, _ = fmt.Fprintf(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self *WriteStructSequenceDcl) WritePackageInit(w *FileGenerationContext, v objects.IStructDcl) error {
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
		self.LanguageTypeService.SequenceConst)
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

func (self *WriteStructSequenceDcl) WriteExtract(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	readAnyTypeValue := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), self.LanguageTypeService.ReadAnyType)
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

func (self *WriteStructSequenceDcl) WriteExtractValue(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	readAnyTypeValue := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), self.LanguageTypeService.ReadAnyType)
	idlObjectValue := self.LanguageTypeService.IdlObjectType.GetName()
	UnsignedLongValue := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), self.LanguageTypeService.UInt32TypeHelper)

	structMembers, err := v.GetStructMembers()
	if err != nil {
		return err
	}
	member := structMembers[0]
	memberType, err := self.DclHelpers.GetActualType(member.MemberType, knownTypes)
	if err != nil {
		return err
	}

	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlInterface:
		_, _ = fmt.Fprintf(w.writer, "func (self *%v) ReadValue(stream %v) error {\n", buildTypeName, readAnyTypeValue)
		_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
		_, _ = fmt.Fprintf(w.writer, "}\n")
		_, _ = fmt.Fprintf(w.writer, "\n")
		return nil
	}
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) ReadValue(stream %v) error {\n", buildTypeName, readAnyTypeValue)
	_, _ = fmt.Fprintf(w.writer, "\terr := self.%v.ReadValue(stream)\n", idlObjectValue)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\tvar n uint32\n")
	_, _ = fmt.Fprintf(w.writer, "\tn, err = %v.Read(stream)\n", UnsignedLongValue)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\tif n > 0 {\n")

	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlStruct:
		s := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), memberType)
		_, _ = fmt.Fprintf(w.writer, "\t\tself.Array = make([]*%v, n)\n", s)
	default:
		var v2 Service.GoLangTypeReference
		find, b := self.LanguageTypeService.Find(memberType.GetPrimitiveType())
		if b {
			if find.GetLexemData().GetSourceLexemData() == self.LanguageTypeService.NoLexem {
				v2 = Service.GoLangTypeReference(find.GetName())
			} else {
				v2 = self.idlToGoTranslation.AbsTypeName(
					v.GetLexemData(),
					find)
			}
		} else {
			v2 = self.idlToGoTranslation.AbsTypeName(
				v.GetLexemData(),
				memberType)
		}
		_, _ = fmt.Fprintf(w.writer, "\t\tself.Array = make([]%v, n)\n", v2)

	}
	_, _ = fmt.Fprintf(w.writer, "\t\tvar i uint32\n")
	_, _ = fmt.Fprintf(w.writer, "\t\tfor i = 0; i < n; i++ {\n")
	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlStruct:
		s := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), memberType)
		_, _ = fmt.Fprintf(w.writer, "\t\t\tself.Array[i] = &%v{}\n", s)
		_, _ = fmt.Fprintf(w.writer, "\t\t\terr = self.Array[i].ReadValue(stream)\n")
	default:
		helper, b := self.LanguageTypeService.FindHelper(memberType.GetPrimitiveType())
		if !b {
			panic(memberType.GetPrimitiveType())
		}
		s := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), helper)
		_, _ = fmt.Fprintf(w.writer, "\t\t\tself.Array[i], err = %v.Read(stream)\n", s)
	}
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSequenceDcl) WriteInsert(w *FileGenerationContext, v objects.IStructDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	writeAnyTypeValue := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), self.LanguageTypeService.WriteAnyType)
	idlObjectValue := self.LanguageTypeService.IdlObjectType.GetName()
	UnsignedLongValue := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), self.LanguageTypeService.UInt32TypeHelper)

	structMembers, err := v.GetStructMembers()
	if err != nil {
		return err
	}
	member := structMembers[0]
	memberType, err := self.DclHelpers.GetActualType(member.MemberType, knownTypes)
	if err != nil {
		return err
	}

	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlInterface:
		_, _ = fmt.Fprintf(w.writer, "func (self *%v) Write(stream %v) error {\n", buildTypeName, writeAnyTypeValue)
		_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
		_, _ = fmt.Fprintf(w.writer, "}\n")
		_, _ = fmt.Fprintf(w.writer, "\n")
		return nil
	}

	_, _ = fmt.Fprintf(w.writer, "func (self *%v) Write(stream %v) error {\n", buildTypeName, writeAnyTypeValue)
	_, _ = fmt.Fprintf(w.writer, "\terr := self.%v.Write(stream)\n", idlObjectValue)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\terr = %v.Write(stream, uint32(len(self.Array)))\n", UnsignedLongValue)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\tif len(self.Array) > 0 {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\tfor _, item := range self.Array {\n")

	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlStruct:
		_, _ = fmt.Fprintf(w.writer, "\t\t\terr = item.Write(stream)\n")
	default:
		helper, b := self.LanguageTypeService.FindHelper(memberType.GetPrimitiveType())
		if !b {
			panic(memberType.GetPrimitiveType())
		}
		s := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), helper)
		_, _ = fmt.Fprintf(w.writer, "\t\t\terr = %v.Write(stream, item)\n", s)
	}

	_, _ = fmt.Fprintf(w.writer, "\t\t\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\t\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSequenceDcl) Start(ctx context.Context) error {
	err := self.writeBaseType.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (self *WriteStructSequenceDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func (self *WriteStructSequenceDcl) WriteStruct(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
	var err error
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v struct {\n", buildTypeName)

	structBaseDcl := self.LanguageTypeService.IdlObjectType
	base := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		structBaseDcl)
	_, _ = fmt.Fprintf(w.writer, "\t%v\n", base)

	var structMembers []*objects.StructMember
	structMembers, err = v.GetStructMembers()
	if err != nil {
		return err
	}
	_member := structMembers[0]
	memberType, err := self.DclHelpers.GetActualType(_member.MemberType, types)
	if err != nil {
		return err
	}

	v1 := self.idlToGoTranslation.ExportMemberName(_member.MemberName)
	var v2 Service.GoLangTypeReference
	find, b := self.LanguageTypeService.Find(memberType.GetPrimitiveType())
	if b {
		if find.GetLexemData().GetSourceLexemData() == self.LanguageTypeService.NoLexem {
			v2 = Service.GoLangTypeReference(find.GetName())
		} else {
			v2 = self.idlToGoTranslation.AbsTypeName(
				v.GetLexemData(),
				find)
		}
	} else {
		v2 = self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			memberType)
	}

	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlStruct:
		_, _ = fmt.Fprintf(w.writer, "\t%v []*%v `json:\"%v\"`\n", v1, v2, v1)
	default:
		_, _ = fmt.Fprintf(w.writer, "\t%v []%v `json:\"%v\"`\n", v1, v2, v1)
	}

	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructSequenceDcl) WriteStringFunction(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
	buildtypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) String() string {\n", buildtypeName)
	_, _ = fmt.Fprintf(w.writer, "\tb, err := __json__.Marshal(self)\n")
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn \"\"\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn string(b)\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self *WriteStructSequenceDcl) WriteGoStringFunction(w *FileGenerationContext, v objects.IStructDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) GoString() string {\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "\treturn self.String()\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func NewWriteStructSequenceDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteStructSequenceDcl {
	return &WriteStructSequenceDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
