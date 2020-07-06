package gowriter

import (
	"context"
	"fmt"
	"github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"strings"
)

type writeBaseType struct {
	idlToGoTranslation  *Service.IdlToGoTranslation
	logger              *log.SubSystemLogger
	LanguageTypeService *Service.LanguageTypeService
	DclHelpers          *Service.DclHelpers
}

func (self *writeBaseType) IterativeArrayForLoop(
	w *FileGenerationContext,
	tabs int,
	index, n int,
	array []int64,
	cb func(n int, s string) error) error {
	if n > 0 {
		_, _ = fmt.Fprintf(w.writer, "%vfor i%v := 0; i%v < %v; i%v++ {\n",
			strings.Repeat("\t", tabs+1),
			index,
			index,
			array[index],
			index)
		err := self.IterativeArrayForLoop(w, tabs+1, index+1, n-1, array, cb)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", tabs+1))
	} else {
		sb := strings.Builder{}
		for i, _ := range array {
			sb.WriteString(fmt.Sprintf("[i%v]", i))
		}
		return cb(index, sb.String())
	}
	return nil
}

func (self writeBaseType) WriteHeader(w *FileGenerationContext, typeName, generatedBy string, v objects.IDcl) error {
	_, _ = fmt.Fprintf(w.writer, "// %v declaration: \"%v\", generatedBy by: \"%v\"\n", typeName, v.GetName(), generatedBy)
	//
	return nil
}

func (self writeBaseType) WriteHelperStruct(w *FileGenerationContext, v objects.IDcl, knownTypes objects.KnownTypes) error {
	_, _ = fmt.Fprintf(w.writer, "//noinspection GoSnakeCaseUsage\n")
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v_Helper struct {\n", buildTypeName)

	if structDcl, ok := v.(objects.IStructDcl); ok {
		if structDcl.IsIStructDcl() {
			structBaseScopeIdentifier, err := structDcl.GetStructBase()
			if err != nil {
				return err
			}
			if structBaseDcl, ok := knownTypes[structBaseScopeIdentifier]; ok {
				typeName := self.idlToGoTranslation.AbsTypeName(v.GetLexemData(), structBaseDcl)
				_ = fmt.Sprintf("\t%v\n", typeName)

			}
		}
	}
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}

func (self writeBaseType) WriteUniqueId01(w *FileGenerationContext, v objects.IDcl) error {
	constName := self.BuildConstName(
		w,
		func() objects.ScopeIdentifier {
			return v.GetName()
		})
	constValue, err := v.BuildIdValue()
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(w.writer, "//noinspection GoSnakeCaseUsage\n")
	_, _ = fmt.Fprintf(w.writer, "const %v = \"%v\"\n", constName, constValue)
	_, _ = fmt.Fprintf(w.writer, "\n")

	return nil
}

func (self writeBaseType) WriteUniqueId02(w *FileGenerationContext, v objects.IDcl) error {
	constName := self.BuildConstName(
		w,
		func() objects.ScopeIdentifier {
			return v.GetName()
		})
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self %v_Helper) Id() string {\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "\treturn %v\n", constName)
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self *writeBaseType) defaultBase(knownTypes objects.KnownTypes) objects.IDcl {
	return self.LanguageTypeService.IdlObjectType
}

func (self writeBaseType) Start(ctx context.Context) error {
	if self.logger == nil {
		return fmt.Errorf("logger not initialized")
	}
	if self.idlToGoTranslation == nil {
		return self.logger.Error(fmt.Errorf("idlToGoTranslation not initialized"))
	}
	//if self.IdlDefaultTypes == nil {
	//	return self.logger.Error(fmt.Errorf("IdlDefaultTypes not initialized"))
	//}

	return nil
}

func (self writeBaseType) Stop(ctx context.Context) error {
	self.idlToGoTranslation = nil
	self.logger = nil
	//self.IdlDefaultTypes = nil
	return nil
}

func (self writeBaseType) BuildConstName(w *FileGenerationContext, cb func() objects.ScopeIdentifier) string {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(cb())
	return fmt.Sprintf("%vId_Const", buildTypeName)
}

func NewBaseTypeWriter(
	idlToGoTranslation *Service.IdlToGoTranslation,
	SubSystemLogger *log.SubSystemLogger,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers) writeBaseType {

	return writeBaseType{
		idlToGoTranslation:  idlToGoTranslation,
		logger:              SubSystemLogger,
		LanguageTypeService: LanguageTypeService,
		DclHelpers:          DclHelpers,
	}
}
