package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
)

type WriteTypeDefOfNativeDcl struct {
	writeBaseType
}

func (self WriteTypeDefOfNativeDcl) Write(w *FileGenerationContext, v *objects.TypeDefDcl, dcl *objects.NativeDcl, types objects.KnownTypes) error {
	err := self.WriteHeader(w, "TypeDef", "WriteTypeDefOfNativeDcl", v)
	if err != nil {
		return nil
	}
	_, _ = fmt.Fprintf(w.writer, "// Typedef Native declaration: %v\n", v.GetName())
	buildTypeName01, buildTypeName02 :=
		self.idlToGoTranslation.BuildTypeName(v.GetName()),
		self.idlToGoTranslation.BuildTypeName(v.TypeDecl.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v %vRepresentative \n",
		buildTypeName01, buildTypeName02)
	return nil
}

func (self WriteTypeDefOfNativeDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.TypeDefDcl) error {
	return nil
}

func (self WriteTypeDefOfNativeDcl) WritePackageInit(w *FileGenerationContext, v *objects.TypeDefDcl) error {
	return nil
}

func (self WriteTypeDefOfNativeDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteTypeDefOfNativeDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewWriteTypeDefOfNativeDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteTypeDefOfNativeDcl {
	return &WriteTypeDefOfNativeDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
