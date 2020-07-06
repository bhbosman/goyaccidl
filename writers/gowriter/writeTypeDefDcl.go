package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"

	"log"
	"reflect"
)

type WriteTypeDefDcl struct {
	writeBaseType
	WriteTypeDefOfPrimitiveDcl *WriteTypeDefOfPrimitiveDcl
	WriteTypeDefOfNativeDcl    *WriteTypeDefOfNativeDcl
}

func (self *WriteTypeDefDcl) Write(w *FileGenerationContext, v *objects.TypeDefDcl, knownTypes objects.KnownTypes) error {

	usedInSwitch := v.TypeDecl
	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(usedInSwitch, knownTypes)
	if err != nil {
		return err
	}
	usedInSwitch = ans.Dcl
	switch dcl := usedInSwitch.(type) {
	case *objects.EnumDcl:
	case *objects.PrimitiveDcl:
		switch dcl.PrimitiveType {
		default:
			err := self.WriteTypeDefOfPrimitiveDcl.Write(w, v, dcl, knownTypes)
			if err != nil {
				return err
			}
			break
		}
	case *objects.NativeDcl:
		err := self.WriteTypeDefOfNativeDcl.Write(w, v, dcl, knownTypes)
		if err != nil {
			return err
		}
		break
	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(dcl).String(), dcl.GetName(), dcl.GetLexemData())
		self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
	}
	return nil
}

func (self WriteTypeDefDcl) WriteGlobalVariables(
	w FileGenerationContext,
	v *objects.TypeDefDcl,
	knownTypes objects.KnownTypes) error {

	usedInSwitch := v.TypeDecl
	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(usedInSwitch, knownTypes)
	if err != nil {
		return err
	}
	usedInSwitch = ans.Dcl
	switch dcl := usedInSwitch.(type) {

	case *objects.EnumDcl:
		break
	case *objects.PrimitiveDcl:
		switch dcl.PrimitiveType {
		default:
			err := self.WriteTypeDefOfPrimitiveDcl.WriteGlobalVariables(w, v)
			if err != nil {
				return nil
			}
			break
		}
	case *objects.NativeDcl:
		err := self.WriteTypeDefOfNativeDcl.WriteGlobalVariables(w, v)
		if err != nil {
			return nil
		}
		break
	case *objects.ScopeDcl:
		underlyingStructDcl := false
		if scopeDclRef, ok := knownTypes[dcl.GetName()]; ok {
			if structDclInterface, ok := scopeDclRef.(objects.IStructDcl); ok {
				underlyingStructDcl = structDclInterface.IsIStructDcl()
			}
		}
		switch underlyingStructDcl {
		case false:
			//err := self.WriteTypeDefOfScopeDcl.WriteGlobalVariables(w, v)
			//if err != nil {
			//	return nil
			//}
		}
		break

	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(dcl).String(), dcl.GetName(), dcl.GetLexemData())
		self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
	}
	return nil
}

func (self WriteTypeDefDcl) WritePackageInit(w *FileGenerationContext, v *objects.TypeDefDcl, knownTypes objects.KnownTypes) error {
	usedInSwitch := v.TypeDecl
	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(usedInSwitch, knownTypes)
	if err != nil {
		return err
	}
	usedInSwitch = ans.Dcl
	switch dcl := usedInSwitch.(type) {
	case *objects.EnumDcl:
		return nil
	//case *objects.SequenceTypeDcl:
	//	return self.WriteTypeDefOfSequenceDcl.WritePackageInit(w, v)
	case *objects.PrimitiveDcl:
		switch dcl.PrimitiveType {
		default:
			return self.WriteTypeDefOfPrimitiveDcl.WritePackageInit(w, v)
		}
	case *objects.NativeDcl:
		return self.WriteTypeDefOfNativeDcl.WritePackageInit(w, v)
	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(dcl).String(), dcl.GetName(), dcl.GetLexemData())
		self.logger.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
		return nil
	}
}

func (self WriteTypeDefDcl) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteTypeDefDcl) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func NewWriteTypeDefDcl(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	WriteTypeDefOfPrimitiveDcl *WriteTypeDefOfPrimitiveDcl,
	WriteTypeDefOfNativeDcl *WriteTypeDefOfNativeDcl,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService) *WriteTypeDefDcl {
	return &WriteTypeDefDcl{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
		WriteTypeDefOfPrimitiveDcl: WriteTypeDefOfPrimitiveDcl,
		WriteTypeDefOfNativeDcl:    WriteTypeDefOfNativeDcl,
	}
}
