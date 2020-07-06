package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
)

type WriteNativeDcl struct {
	writeBaseType
}

func (self WriteNativeDcl) Write(w *FileGenerationContext, v *objects.NativeDcl, knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "Native", "WriteNativeDcl", v)
	if err != nil {
		return nil
	}
	buildTypeName01, buildTypeName02 :=
		self.idlToGoTranslation.BuildTypeName(v.Name),
		self.idlToGoTranslation.BuildTypeName(v.Name)
	_, _ = fmt.Fprintf(w.writer, "type %v Native%v\n",
		buildTypeName01, buildTypeName02)
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}

func (self WriteNativeDcl) FindFolderInUser(v *objects.NativeDcl, result objects.FileUsage) {

}

func (self WriteNativeDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.NativeDcl) error {
	return nil
}

func (self WriteNativeDcl) WritePackageInit(w *FileGenerationContext, v *objects.NativeDcl) error {
	return nil
}

func (self WriteNativeDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteNativeDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewNativeDclWriter(
	translation *Service.IdlToGoTranslation,
	logger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteNativeDcl {
	return &WriteNativeDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			logger,
			LanguageTypeService,
			DclHelpers),
	}
}
