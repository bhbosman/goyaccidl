package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type WriteConstantValue struct {
	writeBaseType
}

func (self WriteConstantValue) Write(w *FileGenerationContext, v objects.IConstantValue, knownTypes objects.KnownTypes) error {
	err := self.WriteHeader(w, "Constant", "WriteConstantValue", v)
	if err != nil {
		return nil
	}
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoUnusedConst\n")
	vv := fmt.Sprintf("%v", v.GetValue())
	if v.GetConstantType().GetPrimitiveType() == yaccidl.StringType {
		buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
		ans, err := self.idlToGoTranslation.CreateTypeReference(
			v.GetLexemData(),
			nil,
			v.GetConstantType(),
			knownTypes)
		if err != nil {
			return err
		}

		_, _ = fmt.Fprintf(w.writer, "const %v %v = \"%v\"\n",
			buildTypeName,
			ans.Reference,
			vv)

	} else {
		ans, err := self.idlToGoTranslation.CreateTypeReference(
			v.GetLexemData(),
			nil,
			v.GetConstantType(),
			knownTypes)
		if err != nil {
			return err
		}

		buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
		_, _ = fmt.Fprintf(w.writer, "const %v %v = %v\n",
			buildTypeName,
			ans.Reference,
			vv)

	}
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}

func (self WriteConstantValue) WriteGlobalVariables(w FileGenerationContext, v *objects.ConstantValue) error {
	return nil
}

func (self WriteConstantValue) WritePackageInit(w *FileGenerationContext, v *objects.ConstantValue) error {
	return nil
}

func (self WriteConstantValue) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteConstantValue) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewConstantValueWriter(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) *WriteConstantValue {
	return &WriteConstantValue{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
