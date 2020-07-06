package gowriter

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
)

type WriteScopeDcl struct {
	writeBaseType
}

func (self WriteScopeDcl) Write(w *FileGenerationContext, v *objects.ScopeDcl, knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "Scope", "WriteScopeDcl", v)
	if err != nil {
		return nil
	}
	return nil
}

func (self WriteScopeDcl) FindFolderInUser(v *objects.ScopeDcl, result objects.FileUsage) {

}

func (self WriteScopeDcl) WriteGlobalVariables(w FileGenerationContext, v *objects.ScopeDcl) error {
	return nil
}

func (self WriteScopeDcl) WritePackageInit(w *FileGenerationContext, v *objects.ScopeDcl) error {
	return nil
}

func (self WriteScopeDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteScopeDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewScopeDclWriter(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteScopeDcl {
	return &WriteScopeDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
