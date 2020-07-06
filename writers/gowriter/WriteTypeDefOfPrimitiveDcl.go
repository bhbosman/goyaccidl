package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"strings"
)

type WriteTypeDefOfPrimitiveDcl struct {
	writeBaseType
}

func (self WriteTypeDefOfPrimitiveDcl) Write(
	w *FileGenerationContext,
	v *objects.TypeDefDcl,
	dcl *objects.PrimitiveDcl,
	knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "TypeDef", "WriteTypeDefOfPrimitiveDcl", v)
	if err != nil {
		return nil
	}
	typeDecl := v.TypeDecl

	buildTypeName01 := self.idlToGoTranslation.BuildTypeName(v.GetName())
	buildTypeName02 := self.idlToGoTranslation.BuildTypeName(typeDecl.GetName())

	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(typeDecl, knownTypes)
	if err != nil {
		return err
	}

	if find, b := self.LanguageTypeService.Find(ans.PrimitiveType); b {
		buildTypeName02 = find.GetName()
		typeDecl = find
	}

	arrayContext := self.idlToGoTranslation.BuildArrayContext(v.GetArrayValue())

	v2 := self.idlToGoTranslation.TypeNameFromBuildTypeName(
		v.GetLexemData(),
		typeDecl.GetLexemData(),
		buildTypeName02,
		"",
		arrayContext)
	_, _ = fmt.Fprintf(w.writer, "//Typedef Primitive declaration: \"%v\" from: \"%v\"\n", v.GetName(), v2)
	_, _ = fmt.Fprintf(w.writer, "type %v %v\n", buildTypeName01, v2)
	_, _ = fmt.Fprintf(w.writer, "\n")
	err = self.WriteHelperStruct(w, v, knownTypes)
	if err != nil {
		return err
	}
	err = self.WriteUniqueId01(w, v)
	if err != nil {
		return err
	}
	err = self.WriteUniqueId02(w, v)
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
	return nil
}

func (self *WriteTypeDefOfPrimitiveDcl) WriteExtract(
	w *FileGenerationContext,
	v *objects.TypeDefDcl,
	types objects.KnownTypes) error {

	ans, err := self.idlToGoTranslation.CreateTypeReference(
		v.GetLexemData(),
		nil,
		v,
		types)
	if err != nil {
		return err
	}

	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())

	helper, b := self.LanguageTypeService.FindHelper(ans.PrimitiveType)
	if !b {
		return self.logger.Error(fmt.Errorf("could not find %v's helper", ans.PrimitiveType))
	}

	v2 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		helper)
	readAnyType := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Read(stream %v) (%v, error) {\n",
		buildTypeName,
		readAnyType, ans.Reference)
	switch v.GetIsArray() {
	case true:
		_, _ = fmt.Fprintf(w.writer, "\tvar data %v\n", buildTypeName)
		_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")
		err := self.IterativeArrayForLoop(
			w,
			0,
			0,
			len(v.GetArrayValue()),
			v.GetArrayValue(),
			func(n int, index string) error {
				_, _ = fmt.Fprintf(w.writer, "%vdata%v, err = %v.Read(stream)\n",
					strings.Repeat("\t", len(v.GetArrayValue())+1),
					index,
					v2)
				_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", len(v.GetArrayValue())+1))
				_, _ = fmt.Fprintf(w.writer, "%vreturn data, err\n", strings.Repeat("\t", len(v.GetArrayValue())+2))
				_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", len(v.GetArrayValue())+1))
				return nil
			})
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "\treturn data, nil\n")
		break
	case false:
		_, _ = fmt.Fprintf(w.writer, "\tresult, err := %v.Read(stream)\n", v2)
		if ans.MapToLanguageType {
			_, _ = fmt.Fprintf(w.writer, "\treturn result, err\n")
		} else {
			_, _ = fmt.Fprintf(w.writer, "\treturn %v(result), err\n", buildTypeName)
		}
		break
	}

	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteTypeDefOfPrimitiveDcl) WriteInsert(
	w *FileGenerationContext,
	v *objects.TypeDefDcl,
	types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())

	ans, err := self.idlToGoTranslation.CreateTypeReference(
		v.GetLexemData(),
		nil,
		v,
		types)
	if err != nil {
		return err
	}
	helper, b := self.LanguageTypeService.FindHelper(ans.PrimitiveType)
	if !b {
		return self.logger.Error(fmt.Errorf("could not find %v's helper", ans.PrimitiveType))
	}

	v2 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		helper)

	writeAnyType := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.WriteAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Write(stream %v, v %v) error {\n",
		buildTypeName,
		writeAnyType,
		ans.Reference)

	switch v.GetIsArray() {
	case true:
		_, _ = fmt.Fprintf(w.writer, "\tvar data %v\n", buildTypeName)
		_, _ = fmt.Fprintf(w.writer, "\tvar err error\n")
		err := self.IterativeArrayForLoop(
			w,
			0,
			0,
			len(v.GetArrayValue()),
			v.GetArrayValue(),
			func(n int, index string) error {
				_, _ = fmt.Fprintf(w.writer, "%verr = %v.Write(stream, data%v)\n",
					strings.Repeat("\t", len(v.GetArrayValue())+1),
					v2,
					index)
				_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", len(v.GetArrayValue())+1))
				_, _ = fmt.Fprintf(w.writer, "%vreturn err\n", strings.Repeat("\t", len(v.GetArrayValue())+2))
				_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", len(v.GetArrayValue())+1))
				return nil
			})
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "\treturn nil\n")
		break
	case false:
		_, _ = fmt.Fprintf(w.writer, "\treturn %v.Write(stream, v)\n", v2)
		//if ans.MapToLanguageType {
		//	_, _ = fmt.Fprintf(w.writer, "\treturn %v.Write(stream, v)\n", v2)
		//} else {
		//	_, _ = fmt.Fprintf(w.writer, "\treturn %v.Write(stream, %v(v))\n", v2, v2)
		//}
		break
	}

	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteTypeDefOfPrimitiveDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.TypeDefDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())

	_, _ = fmt.Fprintf(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self WriteTypeDefOfPrimitiveDcl) WritePackageInit(w *FileGenerationContext, v *objects.TypeDefDcl) error {
	return nil
}

func (self WriteTypeDefOfPrimitiveDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteTypeDefOfPrimitiveDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}
func NewWriteTypeDefOfPrimitiveDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService) *WriteTypeDefOfPrimitiveDcl {
	return &WriteTypeDefOfPrimitiveDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
