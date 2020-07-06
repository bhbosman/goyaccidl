package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
)

type WritePrimitiveDcl struct {
	writeBaseType
}

func (self WritePrimitiveDcl) Write(w *FileGenerationContext, v objects.IDcl, types objects.KnownTypes) error {
	if v.GetName() != "IdlVoidType" {
		err := self.WriteHelperStruct(w, v, types)
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
		err = self.WriteExtract(w, v, types)
		if err != nil {
			return nil
		}
		err = self.WriteInsert(w, v, types)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (self WritePrimitiveDcl) WriteExtract(w *FileGenerationContext, v objects.IDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) ReadByte(stream IReadAny) (%v, error) {\n", buildTypeName, buildTypeName)
	buildTypeName = self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "\treturn stream.ReadByte%v()\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WritePrimitiveDcl) WriteInsert(w *FileGenerationContext, v objects.IDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Write(stream IWriteAny, v %v) error {\n", buildTypeName, buildTypeName)
	buildTypeName = self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "\treturn stream.Write%v(v)\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self WritePrimitiveDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.PrimitiveDcl) error {
	if v.GetName() != "IdlVoidType" {
		buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
		_, _ = fmt.Fprintf(w.writer, "var %vHelper = %v_Helper{}\n", buildTypeName, buildTypeName)
	}
	return nil
}

func (self *WritePrimitiveDcl) WritePackageInit(w *FileGenerationContext, v *objects.PrimitiveDcl) error {
	return nil
}

func (self WritePrimitiveDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WritePrimitiveDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewWritePrimitiveDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WritePrimitiveDcl {
	return &WritePrimitiveDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
