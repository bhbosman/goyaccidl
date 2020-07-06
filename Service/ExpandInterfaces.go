package Service

import (
	"fmt"

	"github.com/bhbosman/goyaccidl/objects"
)

type InterfaceDclMap map[objects.ScopeIdentifier]*objects.InterfaceDcl

type ExpandInterfaces struct {
	LanguageTypeService *LanguageTypeService
	DclHelpers          *DclHelpers
	InterfaceHelper     *InterfaceHelper
}

func (self *ExpandInterfaces) Expand(dcls objects.IDclArray, types objects.KnownTypes) (objects.IDclArray, error) {
	result := make(objects.IDclArray, 0, len(dcls))

	for _, dcl := range dcls {
		switch v := dcl.(type) {
		case *objects.InterfaceDcl:
			ans, err := self.expandInterface(v, types)
			if err != nil {
				return nil, err
			}
			result = append(result, ans...)
		default:
			result = append(result, v)
		}
	}
	return result, nil
}

func (self *ExpandInterfaces) expandInterface(v *objects.InterfaceDcl, types objects.KnownTypes) (objects.IDclArray, error) {
	obs := self.InterfaceHelper.GetStructToRepresentOperation(v, types)

	name := v.GetName().Append("Stub")
	stub := objects.NewStructDcl(
		name,
		false,
		nil, -1,
		v.DstSrcInformation.SourceLexemData,
		nil, false,
		false,
		objects.StructTypeAsIdlStub,
		nil,
		objects.IDclArray{v})

	name = v.GetName().Append("Skeleton")
	//skeleton := objects.NewStructDcl(
	//	name,
	//	false,
	//	nil, -1,
	//	v.DstSrcInformation.SourceLexemData,
	//	nil, false,
	//	false,
	//	objects.StructTypeAsIdlSkeleton,
	//	nil,
	//	objects.IDclArray{v})

	result := objects.IDclArray{
		v,
		//skeleton,
		stub}
	m := make(map[objects.ScopeIdentifier]bool)
	for i := range obs.Observe() {
		s, ok := i.V.(*objects.StructDcl)
		if ok {
			if _, ok := m[s.GetName()]; !ok {
				m[s.GetName()] = true
				result = append(result, s)
				continue
			}
		}
		return nil, fmt.Errorf("duplicate operation %v on interface %v",
			s.GetName(),
			v.GetName())
	}
	return result, nil
}

func NewExpandInterfaces(
	languageTypeService *LanguageTypeService,
	DclHelpers *DclHelpers) *ExpandInterfaces {
	return &ExpandInterfaces{
		LanguageTypeService: languageTypeService,
		DclHelpers:          DclHelpers,
		InterfaceHelper: &InterfaceHelper{
			DclHelpers: DclHelpers,
		},
	}
}
