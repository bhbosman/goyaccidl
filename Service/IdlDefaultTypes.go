package Service

import (
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type IdlDefaultTypes struct {
	LanguageTypeService  *LanguageTypeService
	shortType            yaccidl.IPrimitiveTypeDcl
	int16Type            yaccidl.IPrimitiveTypeDcl
	int32Type            yaccidl.IPrimitiveTypeDcl
	int64Type            yaccidl.IPrimitiveTypeDcl
	unsignedShortType    yaccidl.IPrimitiveTypeDcl
	uint16Type           yaccidl.IPrimitiveTypeDcl
	uint32Type           yaccidl.IPrimitiveTypeDcl
	uint64Type           yaccidl.IPrimitiveTypeDcl
	longType             yaccidl.IPrimitiveTypeDcl
	longLongType         yaccidl.IPrimitiveTypeDcl
	unsignedLongType     yaccidl.IPrimitiveTypeDcl
	unsignedLongLongType yaccidl.IPrimitiveTypeDcl
	charType             yaccidl.IPrimitiveTypeDcl
	wideCharType         yaccidl.IPrimitiveTypeDcl
	stringType           yaccidl.IPrimitiveTypeDcl
	wideStringType       yaccidl.IPrimitiveTypeDcl
	booleanType          yaccidl.IPrimitiveTypeDcl
	floatType            yaccidl.IPrimitiveTypeDcl
	doubleType           yaccidl.IPrimitiveTypeDcl
	longDoubleType       yaccidl.IPrimitiveTypeDcl
	voidType             yaccidl.IPrimitiveTypeDcl
	int8Type             yaccidl.IPrimitiveTypeDcl
	uInt8Type            yaccidl.IPrimitiveTypeDcl
	anyType              yaccidl.IPrimitiveTypeDcl
	objectType           yaccidl.IPrimitiveTypeDcl
	valueBaseType        yaccidl.IPrimitiveTypeDcl
	octetType            yaccidl.IPrimitiveTypeDcl
	idlPrimitives        objects.IDclArray
}

func NewIdlDefaultTypes(LanguageTypeService *LanguageTypeService) *IdlDefaultTypes {
	NoLexem := yaccidl.NewLexemDataForPrimitives(
		"",
		"",
		"",
		0,
		0)

	shortType := yaccidl.NewPrimitiveNode("IdlShort", yaccidl.ShortType, NoLexem)
	int16Type := yaccidl.NewPrimitiveNode("IdlInt16", yaccidl.Int16Type, NoLexem)
	int32Type := yaccidl.NewPrimitiveNode("IdlInt32", yaccidl.Int32Type, NoLexem)
	int64Type := yaccidl.NewPrimitiveNode("IdlInt64", yaccidl.Int64Type, NoLexem)
	unsignedShortType := yaccidl.NewPrimitiveNode("IdlUnsignedShort", yaccidl.UnsignedShortType, NoLexem)
	uint16Type := yaccidl.NewPrimitiveNode("IdlUInt16", yaccidl.Uint16Type, NoLexem)
	uint32Type := yaccidl.NewPrimitiveNode("IdlUInt32", yaccidl.Uint32Type, NoLexem)
	uint64Type := yaccidl.NewPrimitiveNode("IdlUint64", yaccidl.Uint64Type, NoLexem)
	longType := yaccidl.NewPrimitiveNode("IdlLong", yaccidl.LongType, NoLexem)
	longLongType := yaccidl.NewPrimitiveNode("IdlLongLong", yaccidl.LongLongType, NoLexem)
	unsignedLongType := yaccidl.NewPrimitiveNode("IdlUnsignedLong", yaccidl.UnsignedLongType, NoLexem)
	unsignedLongLongType := yaccidl.NewPrimitiveNode("IdlUnsignedLongLong", yaccidl.UnsignedLongLongType, NoLexem)
	charType := yaccidl.NewPrimitiveNode("IdlChar", yaccidl.CharType, NoLexem)
	wideCharType := yaccidl.NewPrimitiveNode("IdlWideChar", yaccidl.WideCharType, NoLexem)
	stringType := yaccidl.NewPrimitiveNode("IdlString", yaccidl.StringType, NoLexem)
	wideStringType := yaccidl.NewPrimitiveNode("IdlWideString", yaccidl.WideStringType, NoLexem)
	booleanType := yaccidl.NewPrimitiveNode("IdlBoolean", yaccidl.BooleanType, NoLexem)
	floatType := yaccidl.NewPrimitiveNode("IdlFloatType", yaccidl.FloatType, NoLexem)
	doubleType := yaccidl.NewPrimitiveNode("IdlDoubleType", yaccidl.DoubleType, NoLexem)
	longDoubleType := yaccidl.NewPrimitiveNode("IdlLongDoubleType", yaccidl.LongDoubleType, NoLexem)
	voidType := yaccidl.NewPrimitiveNode("IdlVoidType", yaccidl.VoidType, NoLexem)
	int8Type := yaccidl.NewPrimitiveNode("IdlInt8", yaccidl.Int8Type, NoLexem)
	uInt8Type := yaccidl.NewPrimitiveNode("IdlUInt8", yaccidl.UInt8Type, NoLexem)
	anyType := yaccidl.NewPrimitiveNode("IdlAny", yaccidl.AnyType, NoLexem)
	objectType := yaccidl.NewPrimitiveNode("IdlObject", yaccidl.IdlObjectKind, LanguageTypeService.IdlObjectType.GetLexemData().GetSourceLexemData())
	valueBaseType := yaccidl.NewPrimitiveNode("IdlValueBase", yaccidl.IdlValueBaseKind, NoLexem)
	octetType := yaccidl.NewPrimitiveNode("IdlOctet", yaccidl.IdlOctetKind, NoLexem)
	shortTypeScoped := objects.NewPrimitiveDcl(shortType, false)
	int16TypeScoped := objects.NewPrimitiveDcl(int16Type, false)
	int32TypeScoped := objects.NewPrimitiveDcl(int32Type, false)
	int64TypeScoped := objects.NewPrimitiveDcl(int64Type, false)
	unsignedShortTypeScoped := objects.NewPrimitiveDcl(unsignedShortType, false)
	uint16TypeScoped := objects.NewPrimitiveDcl(uint16Type, false)
	uint32TypeScoped := objects.NewPrimitiveDcl(uint32Type, false)
	uint64TypeScoped := objects.NewPrimitiveDcl(uint64Type, false)
	longTypeScoped := objects.NewPrimitiveDcl(longType, false)
	longLongTypeScoped := objects.NewPrimitiveDcl(longLongType, false)
	unsignedLongTypeScoped := objects.NewPrimitiveDcl(unsignedLongType, false)
	unsignedLongLongTypeScoped := objects.NewPrimitiveDcl(unsignedLongLongType, false)
	charTypeScoped := objects.NewPrimitiveDcl(charType, false)
	wideCharTypeScoped := objects.NewPrimitiveDcl(wideCharType, false)
	stringTypeScoped := objects.NewPrimitiveDcl(stringType, false)
	wideStringTypeScoped := objects.NewPrimitiveDcl(wideStringType, false)
	booleanTypeScoped := objects.NewPrimitiveDcl(booleanType, false)
	floatTypeScoped := objects.NewPrimitiveDcl(floatType, false)
	doubleTypeScoped := objects.NewPrimitiveDcl(doubleType, false)
	longDoubleTypeScoped := objects.NewPrimitiveDcl(longDoubleType, false)
	voidTypeScoped := objects.NewPrimitiveDcl(voidType, false)
	int8TypeScoped := objects.NewPrimitiveDcl(int8Type, false)
	uInt8TypeScoped := objects.NewPrimitiveDcl(uInt8Type, false)
	anyTypeScoped := objects.NewPrimitiveDcl(anyType, false)
	valueBaseTypeScoped := objects.NewPrimitiveDcl(valueBaseType, false)
	octetTypeScoped := objects.NewPrimitiveDcl(octetType, false)
	objectTypeScoped := objects.NewPrimitiveDcl(objectType, true)
	array := objects.IDclArray{
		shortTypeScoped,
		int16TypeScoped,
		int32TypeScoped,
		int64TypeScoped,
		unsignedShortTypeScoped,
		uint16TypeScoped,
		uint32TypeScoped,
		uint64TypeScoped,
		longTypeScoped,
		longLongTypeScoped,
		unsignedLongTypeScoped,
		unsignedLongLongTypeScoped,
		charTypeScoped,
		wideCharTypeScoped,
		stringTypeScoped,
		wideStringTypeScoped,
		booleanTypeScoped,
		floatTypeScoped,
		doubleTypeScoped,
		longDoubleTypeScoped,
		voidTypeScoped,
		int8TypeScoped,
		uInt8TypeScoped,
		anyTypeScoped,
		objectTypeScoped,
		valueBaseTypeScoped,
		octetTypeScoped,
	}

	return &IdlDefaultTypes{
		LanguageTypeService:  LanguageTypeService,
		shortType:            shortType,
		int16Type:            int16Type,
		int32Type:            int32Type,
		int64Type:            int64Type,
		unsignedShortType:    unsignedShortType,
		uint16Type:           uint16Type,
		uint32Type:           uint32Type,
		uint64Type:           uint64Type,
		longType:             longType,
		longLongType:         longLongType,
		unsignedLongType:     unsignedLongType,
		unsignedLongLongType: unsignedLongLongType,
		charType:             charType,
		wideCharType:         wideCharType,
		stringType:           stringType,
		wideStringType:       wideStringType,
		booleanType:          booleanType,
		floatType:            floatType,
		doubleType:           doubleType,
		longDoubleType:       longDoubleType,
		voidType:             voidType,
		int8Type:             int8Type,
		uInt8Type:            uInt8Type,
		anyType:              anyType,
		objectType:           objectType,
		valueBaseType:        valueBaseType,
		octetType:            octetType,
		idlPrimitives:        array,
	}
}

func (self *IdlDefaultTypes) ShortType() yaccidl.IYaccNode {
	return self.shortType
}

func (self *IdlDefaultTypes) Int16Type() yaccidl.IYaccNode {
	return self.int16Type
}

func (self *IdlDefaultTypes) Int32Type() yaccidl.IYaccNode {
	return self.int32Type
}

func (self *IdlDefaultTypes) Int64Type() yaccidl.IYaccNode {
	return self.int64Type
}

func (self *IdlDefaultTypes) UnsignedShortType() yaccidl.IYaccNode {
	return self.unsignedShortType
}

func (self *IdlDefaultTypes) Uint16Type() yaccidl.IYaccNode {
	return self.uint16Type
}

func (self *IdlDefaultTypes) Uint32Type() yaccidl.IYaccNode {
	return self.uint32Type
}

func (self *IdlDefaultTypes) Uint64Type() yaccidl.IYaccNode {
	return self.uint64Type
}

func (self *IdlDefaultTypes) LongType() yaccidl.IYaccNode {
	return self.longType
}

func (self *IdlDefaultTypes) LongLongType() yaccidl.IYaccNode {
	return self.longLongType
}

func (self *IdlDefaultTypes) UnsignedLongType() yaccidl.IYaccNode {
	return self.unsignedLongType
}

func (self *IdlDefaultTypes) UnsignedLongLongType() yaccidl.IYaccNode {
	return self.unsignedLongLongType
}

func (self *IdlDefaultTypes) CharType() yaccidl.IYaccNode {
	return self.charType
}

func (self *IdlDefaultTypes) WideCharType() yaccidl.IYaccNode {
	return self.wideCharType
}

func (self *IdlDefaultTypes) StringType() yaccidl.IYaccNode {
	return self.stringType
}

func (self *IdlDefaultTypes) WideStringType() yaccidl.IYaccNode {
	return self.wideStringType
}

func (self *IdlDefaultTypes) BooleanType() yaccidl.IYaccNode {
	return self.booleanType
}

func (self *IdlDefaultTypes) FloatType() yaccidl.IYaccNode {
	return self.floatType
}

func (self *IdlDefaultTypes) DoubleType() yaccidl.IYaccNode {
	return self.doubleType
}

func (self *IdlDefaultTypes) LongDoubleType() yaccidl.IYaccNode {
	return self.longDoubleType
}

func (self *IdlDefaultTypes) VoidType() yaccidl.IYaccNode {
	return self.voidType
}

func (self *IdlDefaultTypes) Int8Type() yaccidl.IYaccNode {
	return self.int8Type
}

func (self *IdlDefaultTypes) UInt8Type() yaccidl.IYaccNode {
	return self.uInt8Type
}

func (self *IdlDefaultTypes) AnyType() yaccidl.IYaccNode {
	return self.anyType
}

func (self *IdlDefaultTypes) ObjectType() yaccidl.IYaccNode {
	return self.objectType
}

func (self *IdlDefaultTypes) ValueBaseType() yaccidl.IYaccNode {
	return self.valueBaseType
}

func (self *IdlDefaultTypes) OctetType() yaccidl.IYaccNode {
	return self.octetType
}

func (self *IdlDefaultTypes) IdlPrimitivesInScopedFormat() objects.IDclArray {
	return self.idlPrimitives
}
