package Service

import (
	"errors"
	"fmt"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type DclHelpers struct {
	LanguageTypeService *LanguageTypeService
}

func NewDclHelpers(LanguageTypeService *LanguageTypeService) *DclHelpers {
	return &DclHelpers{
		LanguageTypeService: LanguageTypeService,
	}
}

type FindPrimitiveTypeAnswer struct {
	PrimitiveType yaccidl.PrimitiveType
	Dcl           objects.IDcl
}

func NewFindPrimitiveTypeAnswer(PrimitiveType yaccidl.PrimitiveType, Dcl objects.IDcl) *FindPrimitiveTypeAnswer {
	return &FindPrimitiveTypeAnswer{
		PrimitiveType: PrimitiveType,
		Dcl:           Dcl,
	}
}

func (self DclHelpers) FindPrimitiveTypeForTypeReference(
	dcl objects.IDcl,
	knownTypes objects.KnownTypes) (*FindPrimitiveTypeAnswer, error) {

	v := dcl.GetPrimitiveType()
	switch v {
	case yaccidl.IdlTypedef:
		if _, okIncomingTypeDcl := dcl.(objects.ITypeDefDcl); okIncomingTypeDcl {
			for typeDcl, forOk := dcl.(objects.ITypeDefDcl); forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = dcl.(objects.ITypeDefDcl) {
				primitiveType, err2 := typeDcl.TypeDeclPrimitiveType()
				if err2 != nil {
					return nil, err2
				}
				switch primitiveType {
				case yaccidl.IdlSequence:
					if obj, typeOk := typeDcl.(*objects.TypeDefDcl); typeOk {
						return NewFindPrimitiveTypeAnswer(
							yaccidl.IdlSequence,
							obj.TypeDecl), nil
					}
					name, err := typeDcl.TypeDeclName()
					if err != nil {
						return nil, err
					}
					if sequenceOf, foundOk := knownTypes[name]; foundOk {
						return NewFindPrimitiveTypeAnswer(yaccidl.IdlSequence, sequenceOf), nil
					}
				}

				var foundOk bool
				name, err := typeDcl.TypeDeclName()
				if err != nil {
					return nil, err
				}
				if dcl, foundOk = knownTypes[name]; foundOk {
					v = dcl.GetPrimitiveType()
					if dcl.GetPrimitiveType() == yaccidl.IdlTypedef {
						continue
					}
				}
				break
			}
			return NewFindPrimitiveTypeAnswer(v, dcl), nil
		}
		return NewFindPrimitiveTypeAnswer(v, dcl), nil
	case yaccidl.IdlScope:
		if iDcl, ok := knownTypes[dcl.GetName()]; ok {
			return self.FindPrimitiveTypeForTypeReference(iDcl, knownTypes)
		}
		return nil, fmt.Errorf("could not find %v", dcl.GetName())
	default:
		if ImportedDcl, ok := dcl.(*objects.ImportedDcl); ok {
			return self.FindPrimitiveTypeForTypeReference(ImportedDcl.BaseDcl, knownTypes)
		}
		return NewFindPrimitiveTypeAnswer(v, dcl), nil
	}
}

func (self DclHelpers) FindPrimitiveTypeForUpdateIdlFolderUsage(
	dcl objects.IDcl,
	knownTypes objects.KnownTypes) (*FindPrimitiveTypeAnswer, error) {

	v := dcl.GetPrimitiveType()
	switch v {
	case yaccidl.IdlTypedef:
		if _, okIncomingTypeDcl := dcl.(objects.ITypeDefDcl); okIncomingTypeDcl {
			for typeDcl, forOk := dcl.(objects.ITypeDefDcl); forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = dcl.(objects.ITypeDefDcl) {
				var foundOk bool
				name, err := typeDcl.TypeDeclName()
				if err != nil {
					return nil, err
				}
				if dcl, foundOk = knownTypes[name]; foundOk {
					v = dcl.GetPrimitiveType()
					if dcl.GetPrimitiveType() == yaccidl.IdlTypedef {
						continue
					}
				}
				break
			}
			return NewFindPrimitiveTypeAnswer(v, dcl), nil
		}
		return NewFindPrimitiveTypeAnswer(v, dcl), nil
	case yaccidl.IdlScope:
		if iDcl, ok := knownTypes[dcl.GetName()]; ok {
			return self.FindPrimitiveTypeForUpdateIdlFolderUsage(iDcl, knownTypes)
		}
		return nil, fmt.Errorf("could not find %v", dcl.GetName())
	default:
		if ImportedDcl, ok := dcl.(*objects.ImportedDcl); ok {
			return self.FindPrimitiveTypeForUpdateIdlFolderUsage(ImportedDcl.BaseDcl, knownTypes)
		}
		return NewFindPrimitiveTypeAnswer(v, dcl), nil
	}
}

type FindPrimitiveTypeForWritersAnswer struct {
	PrimitiveType yaccidl.PrimitiveType
	Dcl           objects.IDcl
	Array         []int64
}

func NewFindPrimitiveTypeForWritersAnswer(
	primitiveType yaccidl.PrimitiveType,
	dcl objects.IDcl,
	newArray []int64) *FindPrimitiveTypeForWritersAnswer {
	a := dcl.GetArrayValue()
	b := newArray
	var c []int64
	if a == nil && b == nil {
		c = nil
	}
	if a != nil && b != nil {
		cc := append(newArray, a...)
		c = cc
	} else if a != nil {
		c = a
	} else {
		c = b
	}

	return &FindPrimitiveTypeForWritersAnswer{
		PrimitiveType: primitiveType,
		Dcl:           dcl,
		Array:         c,
	}
}

func (self DclHelpers) FindPrimitiveTypeForWriters(
	dcl objects.IDcl,
	knownTypes objects.KnownTypes) (*FindPrimitiveTypeForWritersAnswer, error) {

	incomdingDcl := dcl

	v := dcl.GetPrimitiveType()
	switch v {
	case yaccidl.IdlTypedef:
		if incomingTypeDcl, okIncomingTypeDcl := dcl.(objects.ITypeDefDcl); okIncomingTypeDcl {
			for typeDcl, forOk := dcl.(objects.ITypeDefDcl); forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = dcl.(objects.ITypeDefDcl) {

				var foundOk bool
				name, err := typeDcl.TypeDeclName()
				if err != nil {
					return nil, err
				}
				if dcl, foundOk = knownTypes[name]; foundOk {
					v = dcl.GetPrimitiveType()
					if dcl.GetPrimitiveType() == yaccidl.IdlTypedef {
						continue
					}
				}
				break
			}
			return NewFindPrimitiveTypeForWritersAnswer(v, dcl,
				incomingTypeDcl.GetArrayValue()), nil
		}
		return NewFindPrimitiveTypeForWritersAnswer(v, dcl,
			incomdingDcl.GetArrayValue()), nil
	case yaccidl.IdlScope:
		if iDcl, ok := knownTypes[dcl.GetName()]; ok {
			recursiveCall, err := self.FindPrimitiveTypeForWriters(iDcl, knownTypes)
			if err != nil {
				return nil, err
			}
			return NewFindPrimitiveTypeForWritersAnswer(
				recursiveCall.PrimitiveType,
				recursiveCall.Dcl,
				iDcl.GetArrayValue()), nil
		}
		return nil, fmt.Errorf("could not find %v", dcl.GetName())
	default:
		if ImportedDcl, ok := dcl.(*objects.ImportedDcl); ok {
			recursiveCall, err := self.FindPrimitiveTypeForWriters(ImportedDcl.BaseDcl, knownTypes)
			if err != nil {
				return nil, err
			}
			return NewFindPrimitiveTypeForWritersAnswer(
				recursiveCall.PrimitiveType,
				recursiveCall.Dcl,
				incomdingDcl.GetArrayValue()), nil
		}
		return NewFindPrimitiveTypeForWritersAnswer(v, dcl,
			incomdingDcl.GetArrayValue()), nil
	}
}

var DclHelpersRemoveScopeError = errors.New("remove scope failure as there is still a scope")

func (self DclHelpers) RemoveScope(incomingDcl objects.IDcl, types objects.KnownTypes) (objects.IDcl, error) {
	switch v := incomingDcl.(type) {
	case *objects.ScopeDcl:
		find, err := types.Find(v.GetName())
		if err != nil {
			return nil, err
		}
		return self.RemoveScope(find, types)
	default:
		if incomingDcl.GetPrimitiveType() == yaccidl.IdlScope {
			return nil, DclHelpersRemoveScopeError
		}
		return incomingDcl, nil
	}
}

func (self DclHelpers) RemoveImportsScope(incomingDcl objects.IDcl, types objects.KnownTypes) (objects.IDcl, error) {
	switch v := incomingDcl.(type) {
	case *objects.ImportedDcl:
		last := v
		for next, ok := v.BaseDcl.(*objects.ImportedDcl); ok; next, ok = next.BaseDcl.(*objects.ImportedDcl) {
			last = next
		}
		return self.RemoveImportsScope(last.BaseDcl, types)
	default:
		return incomingDcl, nil
	}

}

type RemoveTypeDefError struct {
}

func NewRemoveTypeDefError() *RemoveTypeDefError {
	return &RemoveTypeDefError{}
}

func (r RemoveTypeDefError) Error() string {
	panic("implement me")
}

func (self DclHelpers) RemoveTypeDef(incomingDcl objects.IDcl, types001 objects.KnownTypes) (objects.IDcl, error) {
	switch incomingDcl.(type) {
	case *objects.TypeDefDcl:
		if incomingTypeDcl, okIncomingTypeDcl := incomingDcl.(objects.ITypeDefDcl); okIncomingTypeDcl {
			dcl := incomingDcl
			for typeDcl, forOk := dcl.(objects.ITypeDefDcl); forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = dcl.(objects.ITypeDefDcl) {
				name, err := typeDcl.TypeDeclName()
				if err != nil {
					return nil, err
				}
				find, err := types001.Find(name)
				if err != nil {
					return nil, err
				}
				switch v := find.(type) {
				case *objects.TypeDefDcl:
					dcl = v
					continue
				case *objects.StructDcl:
					return incomingTypeDcl, nil
				case *objects.EnumDcl:
					return incomingTypeDcl, nil
				case *objects.InterfaceDcl:
					return incomingTypeDcl, nil
				case *objects.PrimitiveDcl:
					return incomingTypeDcl, nil
				case *objects.SequenceTypeDcl:
					return incomingTypeDcl, nil
				default:
					return nil, NewRemoveTypeDefError()
				}
			}
			return nil, NewRemoveTypeDefError()
		}
		return nil, NewRemoveTypeDefError()
	default:
		return incomingDcl, nil
	}
}

func (self *DclHelpers) GetActualType(dcl objects.IDcl, types objects.KnownTypes) (objects.IDcl, error) {
	var memberType = dcl
	switch memberType.GetPrimitiveType() {
	case yaccidl.IdlScope:

		find, err := types.Find(memberType.GetName())
		if err != nil {
			return nil, err
		}
		return self.GetActualType(find, types)
	case yaccidl.IdlTypedef:
		if incomingType, ok := memberType.(*objects.TypeDefDcl); ok {
			forOk := true
			last := incomingType.TypeDecl
			for typeDcl := incomingType; forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = typeDcl.TypeDecl.(*objects.TypeDefDcl) {
				last = typeDcl.TypeDecl
				if last.GetPrimitiveType() == yaccidl.IdlScope {
					var findDcl objects.IDcl
					var err error
					findDcl, err = types.Find(last.GetName())
					if err != nil {
						return nil, err
					}
					if td, ok := findDcl.(*objects.TypeDefDcl); ok {
						typeDcl = td
						last = td.TypeDecl
					}
				}
			}
			memberType = last
		}
	}
	return memberType, nil
}
