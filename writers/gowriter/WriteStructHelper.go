package gowriter

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	"strings"
)

type WriteStructHelper struct {
	writeBaseType
}

func (self *WriteStructHelper) findBase(w *FileGenerationContext, v *objects.StructDcl, knownTypes objects.KnownTypes) objects.IDcl {
	if v.StructBase != nil {
		return v.StructBase
	}
	return self.defaultBase(knownTypes)
}

func (self *WriteStructHelper) findBaseRepresentation(
	w *FileGenerationContext,
	currentLexem objects.IDstSrcInformation,
	dcl objects.IDcl,
	knownTypes objects.KnownTypes) (Service.GoLangTypeReference, bool) {
	if dcl != nil {
		return self.idlToGoTranslation.AbsTypeName(currentLexem, dcl), true
	}
	defaultBase := self.defaultBase(knownTypes)
	return self.idlToGoTranslation.AbsTypeName(currentLexem, defaultBase), false
}

func (self WriteStructHelper) Start(ctx context.Context) error {
	return self.writeBaseType.Start(ctx)
}

func (self WriteStructHelper) Stop(ctx context.Context) error {
	return self.writeBaseType.Stop(ctx)
}

func (self *WriteStructHelper) WriteStruct(w *FileGenerationContext, v objects.IBaseStructDcl, knownTypes objects.KnownTypes) error {
	var structExceptionDcl bool
	var err error
	if strucDcl, ok := v.(objects.IStructDcl); ok {
		structExceptionDcl, err = strucDcl.GetStructExceptionDcl()
		if err != nil {
			return err
		}
	}
	_, _ = fmt.Fprintf(w.writer, "// Exception Decl: %v\n", structExceptionDcl)
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "type %v struct {\n", buildTypeName)

	var structBase objects.ScopeIdentifier
	structBase, err = v.GetStructBase()
	if err != nil {
		return err
	}
	structBaseDcl, _ := knownTypes[structBase]
	base, _ := self.findBaseRepresentation(
		w,
		v.GetLexemData(),
		structBaseDcl,
		knownTypes)
	_, _ = fmt.Fprintf(w.writer, "\t%v\n", base)

	var structMembers []*objects.StructMember
	structMembers, err = v.GetStructMembers()
	if err != nil {
		return err
	}
	err = self.WriteStructMembers(w, v, structMembers, knownTypes)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	//
	return nil
}

func (self *WriteStructHelper) WriteStructMembers(
	w *FileGenerationContext,
	v objects.IBaseStructDcl,
	structMembers []*objects.StructMember,
	knownTypes objects.KnownTypes) error {

	for _, member := range structMembers {
		v1 := self.idlToGoTranslation.ExportMemberName(member.MemberName)
		memberDcl := member.MemberType
		var memberDclArray []int64
		if knownType, okKnownType := knownTypes[memberDcl.GetName()]; okKnownType {
			memberDclArray = knownType.GetArrayValue()
		}
		buildTypeName := self.idlToGoTranslation.BuildTypeName(memberDcl.GetName())
		ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(memberDcl, knownTypes)
		if err != nil {
			return err
		}
		memberDcl = ans.Dcl
		if findDcl, b := self.LanguageTypeService.Find(ans.PrimitiveType); b {
			memberDcl = findDcl
			buildTypeName = memberDcl.GetName()
		}
		paramArrayContext := self.idlToGoTranslation.BuildArrayContext(member.ArrayValue)
		typeArrayContext := self.idlToGoTranslation.BuildArrayContext(memberDclArray)
		lexemFromDcl := memberDcl.GetLexemData()
		v2 := self.idlToGoTranslation.TypeNameFromBuildTypeName(
			v.GetLexemData(),
			lexemFromDcl,
			buildTypeName,
			paramArrayContext,
			typeArrayContext)
		switch memberDcl.GetPrimitiveType() {
		case yaccidl.IdlObjectKind:
			v2 = self.idlToGoTranslation.AbsTypeNameArrayName(
				v.GetLexemData(),
				self.LanguageTypeService.CorbaIIdlObject,
				member.ArrayValue,
				memberDclArray)
		}
		_, _ = fmt.Fprintf(w.writer, "\t%v %v `json:\"%v\"`\n", v1, v2, v1)
	}
	return nil
}

func (self *WriteStructHelper) WriteStringFunction(w *FileGenerationContext, v objects.IBaseStructDcl, types objects.KnownTypes) error {
	buildtypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) String() string {\n", buildtypeName)
	_, _ = fmt.Fprintf(w.writer, "\tb, err := __json__.Marshal(self)\n")
	_, _ = fmt.Fprintf(w.writer, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w.writer, "\t\treturn \"\"\n")
	_, _ = fmt.Fprintf(w.writer, "\t}\n")
	_, _ = fmt.Fprintf(w.writer, "\treturn string(b)\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self *WriteStructHelper) WriteGoStringFunction(w *FileGenerationContext, v objects.IBaseStructDcl, types objects.KnownTypes) error {
	buildTypeName := self.idlToGoTranslation.BuildTypeName(v.GetName())
	_, _ = fmt.Fprintf(w.writer, "func (self *%v) GoString() string {\n", buildTypeName)
	_, _ = fmt.Fprintf(w.writer, "\treturn self.String()\n")
	_, _ = fmt.Fprintf(w.writer, "}\n")
	_, _ = fmt.Fprintf(w.writer, "\n")
	return nil
}

func (self *WriteStructHelper) WriteStructMemberInsert(
	w *FileGenerationContext,
	v objects.IBaseStructDcl,
	member *objects.StructMember,
	tabs int,
	knownTypes objects.KnownTypes) error {

	if iDcl, ok01 := knownTypes[member.MemberType.GetName()]; ok01 {
		v1 := self.idlToGoTranslation.ExportMemberName(member.MemberName)
		ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(iDcl, knownTypes)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w.writer, "\t// WriteStructHelper::WriteStructMemberInsert(%v)\n", ans.PrimitiveType)
		switch ans.PrimitiveType {
		case yaccidl.IdlNative, yaccidl.IdlValueBaseKind:

		case yaccidl.IdlStruct, yaccidl.IdlUnion:
			propertyWriter := func(n int, s string) error {
				_, _ = fmt.Fprintf(w.writer, "%verr = self.%v%v.Write(stream)\n", strings.Repeat("\t", tabs+n+1), v1, s)
				_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", tabs+n+1))
				_, _ = fmt.Fprintf(w.writer, "%vreturn err\n", strings.Repeat("\t", tabs+n+2))
				_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", tabs+n+1))
				return nil
			}
			if member.IsArray {
				err := self.IterativeArrayForLoop(
					w,
					tabs,
					0,
					len(member.ArrayValue),
					member.ArrayValue,
					propertyWriter)
				if err != nil {
					return err
				}
			} else {
				err := propertyWriter(0, "")
				if err != nil {
					return err
				}
			}
		default:
			dcl, _ := knownTypes[member.MemberType.GetName()]
			v2 := fmt.Sprintf("%vHelper", self.idlToGoTranslation.AbsTypeName(
				v.GetLexemData(),
				dcl))

			ans, err = self.DclHelpers.FindPrimitiveTypeForWriters(member.MemberType, knownTypes)
			if err != nil {
				return err
			}
			if ans.Array == nil {
				helper, b := self.LanguageTypeService.FindHelper(ans.PrimitiveType)
				if b {
					v2 = fmt.Sprintf("%v", self.idlToGoTranslation.AbsTypeName(
						v.GetLexemData(),
						helper))
				}
			}

			propertyWriter := func(n int, s string) error {
				_, _ = fmt.Fprintf(w.writer, "%verr = %v.Write(stream, self.%v%v)\n", strings.Repeat("\t", tabs+n+1), v2, v1, s)
				_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", tabs+n+1))
				_, _ = fmt.Fprintf(w.writer, "%vreturn err\n", strings.Repeat("\t", tabs+n+2))
				_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", tabs+n+1))
				return nil
			}
			if member.IsArray {
				err := self.IterativeArrayForLoop(
					w,
					tabs,
					0,
					len(member.ArrayValue),
					member.ArrayValue,
					propertyWriter)
				if err != nil {
					return err
				}
			} else {
				err := propertyWriter(0, "")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (self *WriteStructHelper) WriteStructMemberExtractValue(
	w *FileGenerationContext,
	v objects.IBaseStructDcl,
	member *objects.StructMember,
	tabs int,
	knownTypes objects.KnownTypes) error {
	v1 := self.idlToGoTranslation.ExportMemberName(member.MemberName)

	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(member.MemberType, knownTypes)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(w.writer, "\t// WriteStructHelper::WriteStructMemberExtractValue(%v)\n", ans.PrimitiveType)
	switch ans.PrimitiveType {
	case yaccidl.IdlNative, yaccidl.IdlValueBaseKind:

	case yaccidl.IdlStruct, yaccidl.IdlUnion:
		propertyReader := func(n int, s string) error {
			_, _ = fmt.Fprintf(w.writer, "%verr = self.%v%v.Read(stream)\n", strings.Repeat("\t", tabs+n+1), v1, s)
			_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", tabs+n+1))
			_, _ = fmt.Fprintf(w.writer, "%vreturn err\n", strings.Repeat("\t", tabs+n+2))
			_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", tabs+n+1))
			return nil
		}
		if member.IsArray {
			err := self.IterativeArrayForLoop(
				w,
				tabs,
				0,
				len(member.ArrayValue),
				member.ArrayValue,
				propertyReader)
			if err != nil {
				return err
			}
		} else {
			err = propertyReader(0, "")
			if err != nil {
				return err
			}
		}
		break
	default:
		dcl, _ := knownTypes[member.MemberType.GetName()]
		v2 := fmt.Sprintf("%vHelper", self.idlToGoTranslation.AbsTypeName(
			v.GetLexemData(),
			dcl))

		ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(member.MemberType, knownTypes)
		if err != nil {
			return err
		}
		if ans.Array == nil {
			helper, b := self.LanguageTypeService.FindHelper(ans.PrimitiveType)
			if b {
				v2 = fmt.Sprintf("%v", self.idlToGoTranslation.AbsTypeName(
					v.GetLexemData(),
					helper))
			}
		}
		//switch ans.PrimitiveType {
		//case yaccidl.IdlEnum:
		//	v2 = self.idlToGoTranslation.AbsTypeName(
		//		v.GetLexemData(),
		//		ans.Dcl)
		//}

		propertyReader := func(n int, s string) error {
			_, _ = fmt.Fprintf(w.writer, "%vself.%v%v, err = %v.Read(stream)\n",
				strings.Repeat("\t", tabs+n+1), v1, s, v2)
			_, _ = fmt.Fprintf(w.writer, "%vif err != nil {\n", strings.Repeat("\t", tabs+n+1))
			_, _ = fmt.Fprintf(w.writer, "%vreturn err\n", strings.Repeat("\t", tabs+n+2))
			_, _ = fmt.Fprintf(w.writer, "%v}\n", strings.Repeat("\t", tabs+n+1))
			return nil
		}
		if member.IsArray {
			err := self.IterativeArrayForLoop(
				w,
				tabs,
				0,
				len(member.ArrayValue),
				member.ArrayValue, propertyReader)
			if err != nil {
				return err
			}
		} else {
			err = propertyReader(0, "")
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func NewWriteStructHelper(
	translation *Service.IdlToGoTranslation,
	subSystemLogger *log2.SubSystemLogger,
	DclHelpers *Service.DclHelpers,
	LanguageTypeService *Service.LanguageTypeService) *WriteStructHelper {
	return &WriteStructHelper{
		writeBaseType: NewBaseTypeWriter(
			translation,
			subSystemLogger,
			LanguageTypeService,
			DclHelpers),
	}
}
