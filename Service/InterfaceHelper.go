package Service

import (
	"context"
	rxgo "github.com/ReactiveX/RxGo"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	"sort"
)

type InterfaceHelper struct {
	DclHelpers         *DclHelpers
	IdlToGoTranslation *IdlToGoTranslation
}

func NewInterfaceHelper(dclHelpers *DclHelpers, IdlToGoTranslation *IdlToGoTranslation) *InterfaceHelper {
	return &InterfaceHelper{
		DclHelpers:         dclHelpers,
		IdlToGoTranslation: IdlToGoTranslation,
	}

}

func (self *InterfaceHelper) GetStructToRepresentOperation(v *objects.InterfaceDcl, types objects.KnownTypes) rxgo.Observable {
	return rxgo.Defer(
		[]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
			for i := range self.GetOperationsSorted(v, types).Observe() {
				if operation, ok := i.V.(*objects.OperationDcl); ok {
					var InMembers []*objects.StructMember
					var OutMembers []*objects.StructMember
					for _, param := range operation.OperationsParams{
						member := objects.NewStructMember(param.ParameterName, false, nil, param.ParameterDcl)
						switch param.Direction {
						case yaccidl.ParamDirectionIn:
							InMembers = append(InMembers, member)
						case yaccidl.ParamDirectionOut:
							OutMembers = append(OutMembers, member)
						case yaccidl.ParamDirectionIn|yaccidl.ParamDirectionOut:
							InMembers = append(InMembers, member)
							OutMembers = append(OutMembers, member)
						}
					}
					name := v.GetName().Append(operation.OperationName).Append("In")
					strucDcl := objects.NewStructDcl(
						name,
						false,
						nil, -1,
						v.DstSrcInformation.SourceLexemData,
						nil, false,
						false,
						objects.StructTypeAsMethodParamIn,
						InMembers,
						nil)
					next <- rxgo.Of(strucDcl)
					name = v.GetName().Append(operation.OperationName).Append("Out")
					strucDcl = objects.NewStructDcl(
						name,
						false,
						nil, -1,
						v.DstSrcInformation.SourceLexemData,
						nil, false,
						false,
						objects.StructTypeAsMethodParamOut,
						OutMembers,
						nil)
					next <- rxgo.Of(strucDcl)
				}
			}
		}})
}

func (self *InterfaceHelper) FindAllInterfacesUsed(InterfaceDclMap InterfaceDclMap, v *objects.InterfaceDcl, types objects.KnownTypes) error {
	if _, ok := InterfaceDclMap[v.GetName()]; !ok {
		InterfaceDclMap[v.GetName()] = v
		for _, baseInterface := range v.BaseInterfaces {
			actualType, err := self.DclHelpers.GetActualType(baseInterface, types)
			if err != nil {
				return err
			}
			switch ans := actualType.(type) {
			case *objects.InterfaceDcl:
				err := self.FindAllInterfacesUsed(InterfaceDclMap, ans, types)
				if err != nil {
					return err
				}
			default:
				continue
			}
		}
	}
	return nil
}

func (self *InterfaceHelper) GetOperations(v *objects.InterfaceDcl, types objects.KnownTypes) rxgo.Observable {
	return self.GetAllInterfaces(v, types).FlatMap(
		func(item rxgo.Item) rxgo.Observable {
			return rxgo.Defer([]rxgo.Producer{
				func(ctx context.Context, next chan<- rxgo.Item) {
					if dcl, ok := item.V.(*objects.InterfaceDcl); ok {
						for _, operation := range dcl.Operations {
							next <- rxgo.Of(operation)
						}
					}
				},
			})
		})
}

func (self *InterfaceHelper) GetOperationsSorted(v *objects.InterfaceDcl, types objects.KnownTypes) rxgo.Observable {
	return rxgo.Defer([]rxgo.Producer{
		func(ctx context.Context, next chan<- rxgo.Item) {
			m := make(map[string]*objects.OperationDcl)
			for i := range self.GetOperations(v, types).Observe() {
				if operation, ok := i.V.(*objects.OperationDcl); ok {
					m[operation.OperationName] = operation
				}
			}
			var ss []string
			for k, _ := range m {
				ss = append(ss, k)
			}
			sort.Strings(ss)
			for _, s := range ss {
				next <- rxgo.Of(m[s])
			}
		},
	})
}

func (self *InterfaceHelper) GetAttributes(v *objects.InterfaceDcl, types objects.KnownTypes) rxgo.Observable {
	return self.GetAllInterfaces(v, types).FlatMap(
		func(item rxgo.Item) rxgo.Observable {
			return rxgo.Defer([]rxgo.Producer{
				func(ctx context.Context, next chan<- rxgo.Item) {
					if dcl, ok := item.V.(*objects.InterfaceDcl); ok {
						for _, operation := range dcl.Attributes {
							next <- rxgo.Of(operation)
						}
					}
				},
			})
		})
}

func (self *InterfaceHelper) GetAllInterfaces(v *objects.InterfaceDcl, types objects.KnownTypes) rxgo.Observable {
	return rxgo.Defer(
		[]rxgo.Producer{
			func(ctx context.Context, next chan<- rxgo.Item) {
				if v.Forward {
					return
				}
				InterfaceDclMap := make(InterfaceDclMap)
				err := self.FindAllInterfacesUsed(InterfaceDclMap, v, types)
				if err != nil {
					return
				}
				var ss []string
				for s, _ := range InterfaceDclMap {
					ss = append(ss, string(s))
				}
				sort.Strings(ss)
				for _, s := range ss {
					next <- rxgo.Of(InterfaceDclMap[objects.ScopeIdentifier(s)])
				}
			},
		})
}
