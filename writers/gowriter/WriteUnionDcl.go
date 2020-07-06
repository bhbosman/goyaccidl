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

type WriteUnionDcl struct {
	writeBaseType
	helper *WriteStructHelper
}

func (self *WriteUnionDcl) Write(w *FileGenerationContext, v *objects.UnionDcl, knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "Union", "WriteUnionDcl", v)
	if err != nil {
		return nil
	}
	err = self.helper.WriteStruct(w, v, knownTypes)
	if err != nil {
		return nil
	}
	err = self.helper.WriteStringFunction(w, v, knownTypes)
	if err != nil {
		return nil
	}
	err = self.helper.WriteGoStringFunction(w, v, knownTypes)
	if err != nil {
		return nil
	}

	err = self.WriteHelperStruct(w, v, knownTypes)
	if err != nil {
		return nil
	}
	err = self.WriteUniqueId01(w, v)
	if err != nil {
		return nil
	}
	//err = self.writeStructDefaultConstructor.WriteConstructor(
	//	w,
	//	v,
	//	knownTypes,
	//	"Default",
	//	self.StructDclConstructorParts)
	//if err != nil {
	//	return err
	//}
	err = self.WriteExtractValue(w, v, knownTypes)
	if err != nil {
		return err
	}
	err = self.WriteInsert(w, v, knownTypes)
	if err != nil {
		return nil
	}
	//
	return nil

}

func (self WriteUnionDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.UnionDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self WriteUnionDcl) WritePackageInit(w *FileGenerationContext, v *objects.UnionDcl) error {
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
		self.LanguageTypeService.TypeDefType)

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

func (self WriteUnionDcl) WriteExtractValue(w *FileGenerationContext, v *objects.UnionDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	val201 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "// WriteUnionDcl::WriteExtractValue\n")
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) Read(stream %v) error {\n", buildTypeName, val201)
	_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")
	v2 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		v.SwitchTypeDcl)
	_, _ = fmt.Fprintf(w.writer, "\tself.Switch, err = %vHelper.Read(stream)\n", v2)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\tswitch self.Switch {\n")
	defaultDefined := false
	for _, unionBody := range v.UnionBodies {
		if unionBody.DefaultCase {
			_, _ = fmt.Fprintf(w.writer, "\tdefault")
			defaultDefined = true
		} else {
			err := self.WriteCaseStatement(w, v, unionBody, knownTypes)
			if err != nil {
				return err
			}
		}
		_, _ = fmt.Fprintf(w.writer, ":\n")

		member := objects.NewStructMember(
			unionBody.Identifier,
			unionBody.Array,
			unionBody.Int64s,
			unionBody.TypeSpec)
		err := self.helper.WriteStructMemberExtractValue(w, v, member, 1, knownTypes)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "\t\tbreak\n")
	}
	if !defaultDefined {
		_, _ = fmt.Fprintf(w.writer, "\tdefault:\n")
		invalidSwitchName := self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			self.LanguageTypeService.InvalidSwitch)
		_, _ = fmt.Fprintf(w.writer, "\t\treturn %v\n", invalidSwitchName)
	}
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self WriteUnionDcl) WriteInsert(
	w *FileGenerationContext,
	v *objects.UnionDcl,
	types objects.KnownTypes) error {

	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	val201 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.WriteAnyType)
	_, _ = fmt.Fprintf(w.writer, "// WriteUnionDcl::WriteInsert\n")
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) Write(stream %v) error {\n", buildTypeName, val201)
	_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")

	valSwith := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		v.SwitchTypeDcl)
	_, _ = fmt.Fprintf(w.writer, "\terr = %vHelper.Write(stream, self.Switch)\n", valSwith)
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn err\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")

	_, _ = fmt.Fprintf(w.writer, "\tswitch self.Switch {\n")
	defaultDefined := false
	for _, unionBody := range v.UnionBodies {
		if unionBody.DefaultCase {
			_, _ = fmt.Fprintf(w.writer, "\tdefault ")
			defaultDefined = true
		} else {
			err := self.WriteCaseStatement(w, v, unionBody, types)
			if err != nil {
				return err
			}
		}
		_, _ = fmt.Fprintf(w.writer, ":\n")
		member := objects.NewStructMember(
			unionBody.Identifier,
			unionBody.Array,
			unionBody.Int64s,
			unionBody.TypeSpec)
		err := self.helper.WriteStructMemberInsert(w, v, member, 1, types)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "\t\tbreak\n")
	}
	if defaultDefined {

	}
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self WriteUnionDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteUnionDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func (self *WriteUnionDcl) WriteCaseStatement(
	w *FileGenerationContext,
	v *objects.UnionDcl,
	unionBody *objects.UnionBody,
	types objects.KnownTypes) error {

	_, _ = fmt.Fprintf(w.writer, "\tcase ")
	for i, caseValue := range unionBody.CaseValues {
		if i > 0 {
			_, _ = fmt.Fprintf(w.writer, ", ")
		}
		switch vv := caseValue.Value.(type) {
		case *objects.ConstantValue:
			if vv.ConstantValue != nil {
				if int64Value, ok := vv.ConstantValue.(*yaccidl.Int64Value); ok {
					ans, err := self.idlToGoTranslation.CreateTypeReference(
						v.GetLexemData(),
						nil,
						v.SwitchTypeDcl,
						types)
					if err != nil {
						return err
					}

					if ans.MapToLanguageType {
						_, _ = fmt.Fprintf(w.writer, "%v", int64Value.Int64Value())
					} else {
						val202 := self.idlToGoTranslation.AbsTypeName(
							v.GetLexemData(),
							v.SwitchTypeDcl)
						_, _ = fmt.Fprintf(w.writer, "%v(%v)", val202, int64Value.Int64Value())
					}
					continue
				}
			}
		}

		label := self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			caseValue.Value)
		ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(v.SwitchTypeDcl, types)
		if err != nil {
			return err
		}
		switch ans.PrimitiveType {
		case yaccidl.IdlEnum:
			val202 := self.idlToGoTranslation.AbsTypeName(
				v.GetLexemData(),
				v.SwitchTypeDcl)
			_, _ = fmt.Fprintf(w.writer, "uint32(%v%v)", val202, label)
		default:
			ans, err := self.idlToGoTranslation.CreateTypeReference(
				v.GetLexemData(),
				nil,
				v.SwitchTypeDcl,
				types)
			if err != nil {
				return err
			}

			if ans.MapToLanguageType {
				_, _ = fmt.Fprintf(w.writer, "%v", label)
			} else {
				val202 := self.idlToGoTranslation.AbsTypeName(
					v.GetLexemData(),
					v.SwitchTypeDcl)
				_, _ = fmt.Fprintf(w.writer, "%v(%v)", val202, label)
			}
		}
	}
	return nil
}

func NewWriteUnionDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	helper *WriteStructHelper,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService) *WriteUnionDcl {
	return &WriteUnionDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		helper: helper,
	}
}
