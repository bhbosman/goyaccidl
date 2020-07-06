package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type WriteEnumDcl struct {
	writeBaseType
}

func (self WriteEnumDcl) Write(w *FileGenerationContext, v *objects.EnumDcl, knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "Enum", "WriteEnumDcl", v)
	if err != nil {
		return nil
	}

	err = self.WriteEnum(w, v, knownTypes)
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
	err = self.WriteUniqueId02(w, v)
	if err != nil {
		return nil
	}
	err = self.WriteExtract(w, v, knownTypes)
	if err != nil {
		return nil
	}
	err = self.WriteInsert(w, v, knownTypes)
	if err != nil {
		return nil
	}

	return nil
}

func (self WriteEnumDcl) FindFolderInUser(v *objects.EnumDcl, result objects.FileUsage) {

}

func (self WriteEnumDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.EnumDcl) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedGlobalVariable\n")
	_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	return nil
}

func (self WriteEnumDcl) WritePackageInit(w *FileGenerationContext, v *objects.EnumDcl) error {
	return nil
}

func (self WriteEnumDcl) WriteExtract(w *FileGenerationContext, v *objects.EnumDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	ReadAnyTypeName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.ReadAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Read(stream %v) (uint32, error) {\n", buildTypeName, ReadAnyTypeName)

	helper, _ := self.LanguageTypeService.FindHelper(yaccidl.Uint32Type)
	v2 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		helper)
	_, _ = fmt.Fprintf(w.writer, "\tresult, err := %v.Read(stream)\n", v2)
	_, _ = fmt.Fprintf(w.writer, "\treturn result, err\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteEnumDcl) WriteInsert(w *FileGenerationContext, v *objects.EnumDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	WriteAnyName := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		self.LanguageTypeService.WriteAnyType)
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Write(stream %v, v uint32) error {\n", buildTypeName, WriteAnyName)

	helper, _ := self.LanguageTypeService.FindHelper(yaccidl.Uint32Type)
	v2 := self.idlToGoTranslation.AbsTypeName(
		v.GetLexemData(),
		helper)

	_, _ = fmt.Fprintf(w.writer, "\treturn %v.Write(stream, v)\n", v2)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WriteEnumDcl) WriteEnum(w *FileGenerationContext, v *objects.EnumDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.Name)
	_, _ = fmt.Fprintf(w.writer, "type %v uint32\n", buildTypeName)
	_, _ = fmt.Fprint(w.writer, "\n")
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedConst\n")
	_, _ = fmt.Fprint(w.writer, "const (\n")

	vLen := 0
	for _, s := range v.Members {
		v := self.idlToGoTranslation.ExportMemberName(s)
		vLen = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(vLen, len(v))

	}

	for i, s := range v.Members {
		template := fmt.Sprintf("\t%%v%%-%dv %%v = %%v\n", vLen)
		_, _ = fmt.Fprintf(w.writer, template, buildTypeName, self.idlToGoTranslation.ExportMemberName(s), buildTypeName, i)
	}
	_, _ = fmt.Fprint(w.writer, ")\n")
	_, _ = fmt.Fprint(w.writer, "\n")
	return nil
}

func (self WriteEnumDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteEnumDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewEnumDclWriter(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService) *WriteEnumDcl {
	return &WriteEnumDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
