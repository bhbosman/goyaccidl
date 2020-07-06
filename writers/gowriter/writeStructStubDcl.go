package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
)

//noinspection GoSnakeCaseUsage
type WriteStructStubDcl struct {
	writeBaseType
	helper *WriteStructHelper
}

func (self *WriteStructStubDcl) Write(
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
	err = self.WriteStruct(w, v, knownTypes)
	if err != nil {
		return err
	}


	return nil
}



func (self *WriteStructStubDcl) Start(ctx context.Context) error {
	err := self.writeBaseType.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (self *WriteStructStubDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}


func (self *WriteStructStubDcl) WriteStruct(w *FileGenerationContext, v objects.IBaseStructDcl, knownTypes objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v struct {\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}


func NewWriteStructStubDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	helper *WriteStructHelper,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteStructStubDcl {
	return &WriteStructStubDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		helper: helper,
	}
}
