package Service

import (
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type LanguageTypeService struct {
	//mappedToLanguage               map[yaccidl.PrimitiveType]IDcl
	data                           map[yaccidl.PrimitiveType]objects.IDcl
	helperData                     map[yaccidl.PrimitiveType]objects.IDcl
	ReadAnyType                    objects.IDcl
	WriteAnyType                   objects.IDcl
	CorbaIIdlObject                objects.IDcl
	OctetType                      objects.IDcl
	CharType                       objects.IDcl
	ShortType                      objects.IDcl
	Int16Type                      objects.IDcl
	LongType                       objects.IDcl
	Int32Type                      objects.IDcl
	UnsignedLongType               objects.IDcl
	UInt32Type                     objects.IDcl
	Int64Type                      objects.IDcl
	UInt64Type                     objects.IDcl
	LongLongType                   objects.IDcl
	UnsignedShortType              objects.IDcl
	UInt16Type                     objects.IDcl
	BooleanType                    objects.IDcl
	AddRegistrationFunc            objects.IDcl
	NewRegistrationInformationFunc objects.IDcl
	RandomDataGenerator            objects.IDcl
	InvalidSwitch                  objects.IDcl
	LexemDataForPrimitives         *yaccidl.LexemValue
	UnsignedLongLongType           objects.IDcl
	WideCharType                   objects.IDcl
	FloatType                      objects.IDcl
	DoubleType                     objects.IDcl
	LongDoubleType                 objects.IDcl
	Int8Type                       objects.IDcl
	UInt8Type                      objects.IDcl
	StringType                     objects.IDcl
	WideStringType                 objects.IDcl
	IdlObjectType                  objects.IDcl
	IdlAnyType                     objects.IDcl
	ByteTypeHelper                 objects.IDcl
	CharTypeHelper                 objects.IDcl
	ShortTypeHelper                objects.IDcl
	Int16TypeHelper                objects.IDcl
	Int8TypeHelper                 objects.IDcl
	LongTypeHelper                 objects.IDcl
	Int32TypeHelper                objects.IDcl
	UInt8TypeHelper                objects.IDcl
	DoubleTypeHelper               objects.IDcl
	LongDoubleTypeHelper           objects.IDcl
	UnsignedLongTypeHelper         objects.IDcl
	StringTypeHelper               objects.IDcl
	WideStringTypeHelper           objects.IDcl
	UInt32TypeHelper               objects.IDcl
	EnumTypeHelper                 objects.IDcl
	Int64TypeHelper                objects.IDcl
	FloatTypeHelper                objects.IDcl
	UInt64TypeHelper               objects.IDcl
	UnsignedLongLongHelper         objects.IDcl
	LongLongTypeHelper             objects.IDcl
	WideCharTypeHelper             objects.IDcl
	UInt16TypeHelper               objects.IDcl
	BooleanTypeHelper              objects.IDcl
	UnsignedShortTypeHelper        objects.IDcl
	IdlAnyHelperTypeHelper         objects.IDcl
	VoidType                       objects.IDcl
	ValueBaseType                  objects.IDcl
	NoLexem                        *yaccidl.LexemValue
	StructConst                    objects.IDcl
	SequenceConst                  objects.IDcl
	TypeDefType                    objects.IDcl
}

//func (s *LanguageTypeService) MappedToLanguage(primitiveType yaccidl.PrimitiveType) (IDcl, bool) {
//	dcl, b := s.mappedToLanguage[primitiveType]
//	return dcl, b
//}

func (s *LanguageTypeService) Find(primitiveType yaccidl.PrimitiveType) (objects.IDcl, bool) {
	dcl, b := s.data[primitiveType]
	return dcl, b
}

func (s *LanguageTypeService) FindHelper(primitiveType yaccidl.PrimitiveType) (objects.IDcl, bool) {
	dcl, b := s.helperData[primitiveType]
	return dcl, b
}

func (self *LanguageTypeService) IdlPrimitivesInScopedFormat() objects.IDclArray {
	return objects.IDclArray{
		self.ShortType,
		self.Int16Type,
		self.Int32Type,
		self.Int64Type,
		self.UnsignedShortType,
		self.UInt16Type,
		self.UInt32Type,
		self.UInt64Type,
		self.LongType,
		self.LongLongType,
		self.UnsignedLongType,
		self.UnsignedLongLongType,
		self.CharType,
		self.WideCharType,
		self.StringType,
		self.WideStringType,
		self.BooleanType,
		self.FloatType,
		self.DoubleType,
		self.LongDoubleType,
		self.VoidType,
		self.Int8Type,
		self.UInt8Type,
		self.IdlAnyType,
		self.IdlObjectType,
		self.ValueBaseType,
		self.OctetType,
	}
}

func newLanguageTypeService() *LanguageTypeService {

	NewBaseDcl := func(
		primitiveType yaccidl.PrimitiveType,
		name objects.ScopeIdentifier,
		lexemData *yaccidl.LexemValue) *objects.BaseDcl {
		return objects.NewBaseDcl(
			primitiveType,
			name,
			false,
			nil,
			-1,
			lexemData,
			false,
			true)
	}

	NoLexem := yaccidl.NewLexemDataForPrimitives(
		"",
		"",
		"",
		0,
		0)

	OctetType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlOctetKind, "byte", NoLexem))
	CharType := objects.NewImportedDcl(NewBaseDcl(yaccidl.CharType, "byte", NoLexem))
	ShortType := objects.NewImportedDcl(NewBaseDcl(yaccidl.ShortType, "int16", NoLexem))
	Int16Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int16Type, "int16", NoLexem))
	Int8Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int8Type, "int8", NoLexem))
	LongType := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongType, "int32", NoLexem))
	Int32Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int32Type, "int32", NoLexem))
	UInt8Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.UInt8Type, "uint8", NoLexem))
	DoubleType := objects.NewImportedDcl(NewBaseDcl(yaccidl.DoubleType, "float64", NoLexem))
	LongDoubleType := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongDoubleType, "float64", NoLexem))
	UnsignedLongType := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedLongType, "uint32", NoLexem))
	StringType := objects.NewImportedDcl(NewBaseDcl(yaccidl.StringType, "string", NoLexem))
	WideStringType := objects.NewImportedDcl(NewBaseDcl(yaccidl.WideStringType, "string", NoLexem))
	UInt32Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint32Type, "uint32", NoLexem))
	EnumType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlEnum, "uint32", NoLexem))
	Int64Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int64Type, "int64", NoLexem))
	FloatType := objects.NewImportedDcl(NewBaseDcl(yaccidl.FloatType, "float32", NoLexem))
	UInt64Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint64Type, "uint64", NoLexem))
	UnsignedLongLongType := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedLongLongType, "uint64", NoLexem))
	LongLongType := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongLongType, "int64", NoLexem))
	WideCharType := objects.NewImportedDcl(NewBaseDcl(yaccidl.WideCharType, "rune", NoLexem))
	UInt16Type := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint16Type, "uint16", NoLexem))
	BooleanType := objects.NewImportedDcl(NewBaseDcl(yaccidl.BooleanType, "bool", NoLexem))
	UnsignedShortType := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedShortType, "uint16", NoLexem))
	LexemDataForPrimitives := yaccidl.NewLexemDataForPrimitives("", "github.com/bhbosman/orb/src/omg", "__omg__", 0, 0)

	LexemDataForGoIdl := yaccidl.NewLexemDataForPrimitives("", "github.com/bhbosman/goidl", "__goidl__", 0, 0)
	InvalidSwitch := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "InvalidSwitch", LexemDataForGoIdl))
	ReadAnyType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "IReadAny", LexemDataForGoIdl))
	WriteAnyType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "IWriteAny", LexemDataForGoIdl))
	IdlAnyType := objects.NewImportedDcl(NewBaseDcl(yaccidl.AnyType, "IdlAny", LexemDataForGoIdl))
	IdlObjectType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlStruct, "IdlObject", LexemDataForGoIdl))
	CorbaIIdlObject := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "IIdlObject", LexemDataForGoIdl))
	ByteTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlOctetKind, "IdlOctetHelper", LexemDataForGoIdl))
	CharTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.CharType, "IdlOctetHelper", LexemDataForGoIdl))
	ShortTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.ShortType, "IdlInt16Helper", LexemDataForGoIdl))
	Int16TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int16Type, "IdlInt16Helper", LexemDataForGoIdl))
	Int8TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int8Type, "IdlInt8Helper", LexemDataForGoIdl))
	LongTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongType, "IdlInt32Helper", LexemDataForGoIdl))
	Int32TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int32Type, "IdlInt32Helper", LexemDataForGoIdl))
	UInt8TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.UInt8Type, "IdlUInt8Helper", LexemDataForGoIdl))
	DoubleTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.DoubleType, "IdlFloat64Helper", LexemDataForGoIdl))
	LongDoubleTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongDoubleType, "IdlFloat64Helper", LexemDataForGoIdl))
	UnsignedLongTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedLongType, "IdlUInt32Helper", LexemDataForGoIdl))
	StringTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.StringType, "IdlStringHelper", LexemDataForGoIdl))
	WideStringTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.WideStringType, "IdlWideStringHelper", LexemDataForGoIdl))
	UInt32TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint32Type, "IdlUInt32Helper", LexemDataForGoIdl))
	EnumTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlEnum, "IdlUInt32Helper", LexemDataForGoIdl))
	Int64TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Int64Type, "IdlInt64Helper", LexemDataForGoIdl))
	FloatTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.FloatType, "IdlFloat32Helper", LexemDataForGoIdl))
	UInt64TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint64Type, "IdlUInt64Helper", LexemDataForGoIdl))
	UnsignedLongLongHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedLongLongType, "IdlUInt64Helper", LexemDataForGoIdl))
	LongLongTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.LongLongType, "IdlInt64Helper", LexemDataForGoIdl))
	WideCharTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.WideCharType, "IdlRuneHelper", LexemDataForGoIdl))
	UInt16TypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.Uint16Type, "IdlUInt16Helper", LexemDataForGoIdl))
	BooleanTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.BooleanType, "IdlBooleanHelper", LexemDataForGoIdl))
	UnsignedShortTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.UnsignedShortType, "IdlUInt16Helper", LexemDataForGoIdl))
	IdlAnyHelperTypeHelper := objects.NewImportedDcl(NewBaseDcl(yaccidl.AnyType, "IdlAnyHelper", LexemDataForGoIdl))
	RandomDataGenerator := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "IRandomDataGenerator", LexemDataForGoIdl))
	AddRegistrationFunc := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "AddRegistration", LexemDataForGoIdl))
	NewRegistrationInformationFunc := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlInterface, "NewRegistrationInformation", LexemDataForGoIdl))
	VoidType := objects.NewImportedDcl(NewBaseDcl(yaccidl.VoidType, "InvalidSwitch", LexemDataForGoIdl))
	ValueBaseType := objects.NewImportedDcl(NewBaseDcl(yaccidl.IdlValueBaseKind, "IdlValueBase", LexemDataForGoIdl))

	StructConst := objects.NewImportedDcl(NewBaseDcl(yaccidl.VoidType, "StructType", LexemDataForGoIdl))
	SequenceConst := objects.NewImportedDcl(NewBaseDcl(yaccidl.VoidType, "SequenceType", LexemDataForGoIdl))
	TypeDefType := objects.NewImportedDcl(NewBaseDcl(yaccidl.VoidType, "TypeDefType", LexemDataForGoIdl))

	data := make(map[yaccidl.PrimitiveType]objects.IDcl)
	data[OctetType.GetPrimitiveType()] = OctetType
	data[CharType.GetPrimitiveType()] = CharType
	data[ShortType.GetPrimitiveType()] = ShortType
	data[Int16Type.GetPrimitiveType()] = Int16Type
	data[Int32Type.GetPrimitiveType()] = Int32Type
	data[LongType.GetPrimitiveType()] = LongType
	data[UnsignedLongType.GetPrimitiveType()] = UnsignedLongType
	data[UInt32Type.GetPrimitiveType()] = UInt32Type
	data[EnumType.GetPrimitiveType()] = EnumType
	data[Int64Type.GetPrimitiveType()] = Int64Type
	data[LongLongType.GetPrimitiveType()] = LongLongType
	data[UnsignedShortType.GetPrimitiveType()] = UnsignedShortType
	data[UInt16Type.GetPrimitiveType()] = UInt16Type
	data[BooleanType.GetPrimitiveType()] = BooleanType
	data[UInt64Type.GetPrimitiveType()] = UInt64Type
	data[UnsignedLongLongType.GetPrimitiveType()] = UnsignedLongLongType
	data[WideCharType.GetPrimitiveType()] = WideCharType
	data[FloatType.GetPrimitiveType()] = FloatType
	data[DoubleType.GetPrimitiveType()] = DoubleType
	data[Int8Type.GetPrimitiveType()] = Int8Type
	data[UInt8Type.GetPrimitiveType()] = UInt8Type
	data[StringType.GetPrimitiveType()] = StringType
	data[WideStringType.GetPrimitiveType()] = WideStringType
	data[IdlAnyType.GetPrimitiveType()] = IdlAnyType
	data[LongDoubleType.GetPrimitiveType()] = LongDoubleType
	data[ValueBaseType.GetPrimitiveType()] = ValueBaseType

	helperData := make(map[yaccidl.PrimitiveType]objects.IDcl)
	helperData[ByteTypeHelper.GetPrimitiveType()] = ByteTypeHelper
	helperData[CharTypeHelper.GetPrimitiveType()] = CharTypeHelper
	helperData[ShortTypeHelper.GetPrimitiveType()] = ShortTypeHelper
	helperData[Int16TypeHelper.GetPrimitiveType()] = Int16TypeHelper
	helperData[Int32TypeHelper.GetPrimitiveType()] = Int32TypeHelper
	helperData[LongTypeHelper.GetPrimitiveType()] = LongTypeHelper
	helperData[UnsignedLongTypeHelper.GetPrimitiveType()] = UnsignedLongTypeHelper
	helperData[UInt32TypeHelper.GetPrimitiveType()] = UInt32TypeHelper
	helperData[EnumTypeHelper.GetPrimitiveType()] = EnumTypeHelper
	helperData[Int64TypeHelper.GetPrimitiveType()] = Int64TypeHelper
	helperData[LongLongTypeHelper.GetPrimitiveType()] = LongLongTypeHelper
	helperData[UnsignedShortTypeHelper.GetPrimitiveType()] = UnsignedShortTypeHelper
	helperData[UInt16TypeHelper.GetPrimitiveType()] = UInt16TypeHelper
	helperData[BooleanTypeHelper.GetPrimitiveType()] = BooleanTypeHelper
	helperData[UInt64TypeHelper.GetPrimitiveType()] = UInt64TypeHelper
	helperData[UnsignedLongLongHelper.GetPrimitiveType()] = UnsignedLongLongHelper
	helperData[WideCharTypeHelper.GetPrimitiveType()] = WideCharTypeHelper
	helperData[FloatTypeHelper.GetPrimitiveType()] = FloatTypeHelper
	helperData[DoubleTypeHelper.GetPrimitiveType()] = DoubleTypeHelper
	helperData[Int8TypeHelper.GetPrimitiveType()] = Int8TypeHelper
	helperData[UInt8TypeHelper.GetPrimitiveType()] = UInt8TypeHelper
	helperData[StringTypeHelper.GetPrimitiveType()] = StringTypeHelper
	helperData[WideStringTypeHelper.GetPrimitiveType()] = WideStringTypeHelper
	helperData[LongDoubleTypeHelper.GetPrimitiveType()] = LongDoubleTypeHelper
	helperData[IdlAnyHelperTypeHelper.GetPrimitiveType()] = IdlAnyHelperTypeHelper

	return &LanguageTypeService{
		//mappedToLanguage: mappedToLanguage,
		data:                           data,
		helperData:                     helperData,
		ReadAnyType:                    ReadAnyType,
		WriteAnyType:                   WriteAnyType,
		CorbaIIdlObject:                CorbaIIdlObject,
		OctetType:                      OctetType,
		CharType:                       CharType,
		ShortType:                      ShortType,
		Int16Type:                      Int16Type,
		LongType:                       LongType,
		Int32Type:                      Int32Type,
		UnsignedLongType:               UnsignedLongType,
		UInt32Type:                     UInt32Type,
		Int64Type:                      Int64Type,
		UInt64Type:                     UInt64Type,
		LongLongType:                   LongLongType,
		UnsignedShortType:              UnsignedShortType,
		UInt16Type:                     UInt16Type,
		BooleanType:                    BooleanType,
		AddRegistrationFunc:            AddRegistrationFunc,
		NewRegistrationInformationFunc: NewRegistrationInformationFunc,
		RandomDataGenerator:            RandomDataGenerator,
		InvalidSwitch:                  InvalidSwitch,
		LexemDataForPrimitives:         LexemDataForPrimitives,
		UnsignedLongLongType:           UnsignedLongLongType,
		WideCharType:                   WideCharType,
		FloatType:                      FloatType,
		DoubleType:                     DoubleType,
		LongDoubleType:                 LongDoubleType,
		Int8Type:                       Int8Type,
		UInt8Type:                      UInt8Type,
		StringType:                     StringType,
		WideStringType:                 WideStringType,
		IdlObjectType:                  IdlObjectType,
		IdlAnyType:                     IdlAnyType,
		ByteTypeHelper:                 ByteTypeHelper,
		CharTypeHelper:                 CharTypeHelper,
		ShortTypeHelper:                ShortTypeHelper,
		Int16TypeHelper:                Int16TypeHelper,
		Int8TypeHelper:                 Int8TypeHelper,
		LongTypeHelper:                 LongTypeHelper,
		Int32TypeHelper:                Int32TypeHelper,
		UInt8TypeHelper:                UInt8TypeHelper,
		DoubleTypeHelper:               DoubleTypeHelper,
		LongDoubleTypeHelper:           LongDoubleTypeHelper,
		UnsignedLongTypeHelper:         UnsignedLongTypeHelper,
		StringTypeHelper:               StringTypeHelper,
		WideStringTypeHelper:           WideStringTypeHelper,
		UInt32TypeHelper:               UInt32TypeHelper,
		EnumTypeHelper:                 EnumTypeHelper,
		Int64TypeHelper:                Int64TypeHelper,
		FloatTypeHelper:                FloatTypeHelper,
		UInt64TypeHelper:               UInt64TypeHelper,
		UnsignedLongLongHelper:         UnsignedLongLongHelper,
		LongLongTypeHelper:             LongLongTypeHelper,
		WideCharTypeHelper:             WideCharTypeHelper,
		UInt16TypeHelper:               UInt16TypeHelper,
		BooleanTypeHelper:              BooleanTypeHelper,
		UnsignedShortTypeHelper:        UnsignedShortTypeHelper,
		IdlAnyHelperTypeHelper:         IdlAnyHelperTypeHelper,
		VoidType:                       VoidType,
		ValueBaseType:                  ValueBaseType,
		NoLexem:                        NoLexem,
		StructConst:                    StructConst,
		SequenceConst:                  SequenceConst,
		TypeDefType:                    TypeDefType,
	}
}
