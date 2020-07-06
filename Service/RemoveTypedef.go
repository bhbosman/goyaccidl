package Service

import (
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type RemoveTypedef struct {
	LanguageTypeService *LanguageTypeService
}

func NewRemoveTypedef(languageTypeService *LanguageTypeService) *RemoveTypedef {
	return &RemoveTypedef{LanguageTypeService: languageTypeService}
}

func (self *RemoveTypedef) Remove(members objects.IDclArray, knownTypes objects.KnownTypes) (objects.IDclArray, error) {
	var result objects.IDclArray
	for _, dcl := range members {
		switch v := dcl.(type) {
		case *objects.TypeDefDcl:
			typeDcl, err := self.RemoveTypeDcl(v, knownTypes)
			if err != nil {
				return nil, err
			}
			if typeDcl != nil {
				result = append(result, typeDcl)
			}
		default:
			result = append(result, v)
			continue
		}
	}
	return result, nil
}

func (self *RemoveTypedef) RemoveTypeDcl(incomingType *objects.TypeDefDcl, knownTypes objects.KnownTypes) (objects.IDcl, error) {
	var last objects.IDcl
	forOk := true
	for typeDcl := incomingType; forOk && typeDcl.IsTypeDefDcl(); typeDcl, forOk = typeDcl.TypeDecl.(*objects.TypeDefDcl) {
		last = typeDcl.TypeDecl
		if last.GetPrimitiveType() == yaccidl.IdlScope {
			var findDcl objects.IDcl
			var err error
			findDcl, err = knownTypes.Find(last.GetName())
			if err != nil {
				return nil, err
			}
			if td, ok := findDcl.(*objects.TypeDefDcl); ok {
				typeDcl = td
				last = td.TypeDecl
			}
		}
	}
	switch v2 := last.(type) {
	case *objects.ScopeDcl:
		var findDcl objects.IDcl
		var err error
		findDcl, err = knownTypes.Find(v2.GetName())
		if err != nil {
			return nil, err
		}
		switch findDcl.(type) {
		case *objects.EnumDcl:
			return incomingType, nil
		case *objects.StructDcl:
			return objects.NewStructDcl(
				incomingType.GetName(),
				incomingType.GetIsArray(),
				incomingType.GetArrayValue(),
				incomingType.GetOrderId(),
				incomingType.GetLexemData().GetSourceLexemData(),
				findDcl,
				false,
				false,
				objects.StructTypeAsDefined,
				nil,
				nil), nil
		case *objects.SeqStructDcl:
			return objects.NewStructDcl(
				incomingType.GetName(),
				incomingType.GetIsArray(),
				incomingType.GetArrayValue(),
				incomingType.GetOrderId(),
				incomingType.GetLexemData().GetSourceLexemData(),
				findDcl,
				false,
				false,
				objects.StructTypeAsDefined,
				nil,
				nil), nil
		default:
			return incomingType, nil
		}
	case *objects.SequenceTypeDcl:
		var findDcl objects.IDcl
		var err error
		findDcl, err = knownTypes.Find(v2.SequenceType.GetName())
		if err != nil {
			return nil, err
		}
		switch SeqType := findDcl.(type) {
		case *objects.TypeDefDcl:
			dcl, err := self.RemoveTypeDcl(SeqType, knownTypes)
			if err != nil {
				return nil, err
			}
			return objects.NewSeqStructDcl(
				incomingType.GetName(),
				incomingType.GetIsArray(),
				incomingType.GetArrayValue(),
				incomingType.GetOrderId(),
				incomingType.GetLexemData().GetSourceLexemData(),
				nil,
				false,
				false,
				[]*objects.StructMember{
					objects.NewStructMember(
						"Array",
						false,
						nil,
						dcl),
				}), nil
		default:
			return objects.NewSeqStructDcl(
				incomingType.GetName(),
				incomingType.GetIsArray(),
				incomingType.GetArrayValue(),
				incomingType.GetOrderId(),
				incomingType.GetLexemData().GetSourceLexemData(),
				nil,
				false,
				false,
				[]*objects.StructMember{
					objects.NewStructMember(
						"Array",
						false,
						nil,
						findDcl),
				}), nil
		}
	case *objects.PrimitiveDcl:
		switch v2.GetPrimitiveType() {
		case yaccidl.IdlObjectKind:
			return objects.NewStructDcl(
				incomingType.GetName(),
				incomingType.GetIsArray(),
				incomingType.GetArrayValue(),
				incomingType.GetOrderId(),
				incomingType.GetLexemData().GetSourceLexemData(),
				nil,
				false,
				false,
				objects.StructTypeAsDefined,
				nil,
				nil), nil
		default:
			return incomingType, nil
		}

	default:
		panic(v2)
		return incomingType, nil
	}
}
