package processIdlFile

import (
	"context"
	"errors"
	"fmt"
	rxgo "github.com/ReactiveX/RxGo"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/objects"
	"go.uber.org/fx"
)

type CreateKnownTypeService struct {
	logger          *log2.SubSystemLogger
	DclHelpers      *Service.DclHelpers
	IdlDefaultTypes *Service.IdlDefaultTypes
}

var CreateKnownTypesError = errors.New("error creating CreateKnownTypes")

func (self *CreateKnownTypeService) CreateKnownTypes(
	dclArray objects.IDclArray,
	ImportedDeclarationItems []*objects.ScopeItem) (objects.KnownTypes, error) {

	result := rxgo.Defer(
		[]rxgo.Producer{
			func(ctx context.Context, next chan<- rxgo.Item) {
				knownTypes001 := make(objects.KnownTypes)
				for _, member := range dclArray {
					knownTypes001[member.GetName()] = member

				}
				for _, v := range ImportedDeclarationItems {
					baseDcl := objects.NewImportedDcl(v.Dcl)
					knownTypes001[v.Dcl.GetName()] = baseDcl
				}
				for _, item := range self.IdlDefaultTypes.IdlPrimitivesInScopedFormat() {
					knownTypes001[item.GetName()] = item
				}
				next <- rxgo.Of(knownTypes001)
			},
		}).
		Map(func(ctx context.Context, i interface{}) (interface{}, error) {
			if knownTypes, ok := i.(objects.KnownTypes); ok {
				result := make(objects.KnownTypes)
				for _, dcl := range knownTypes {
					ans, err := self.DclHelpers.RemoveImportsScope(dcl, knownTypes)
					if err != nil {
						return nil, err
					}
					name := ans.GetName()
					result[name] = ans
				}
				return result, nil
			}
			return nil, self.logger.Error(CreateKnownTypesError)
		}).
		Map(func(ctx context.Context, i interface{}) (interface{}, error) {
			if knownTypes, ok := i.(objects.KnownTypes); ok {
				result := make(objects.KnownTypes)
				for _, dcl := range knownTypes {
					ans, err := self.DclHelpers.RemoveScope(dcl, knownTypes)
					if err != nil {
						return nil, err
					}
					name := ans.GetName()
					result[name] = ans
				}
				return result, nil
			}
			return nil, self.logger.Error(CreateKnownTypesError)
		}).FirstOrDefault(nil)

	v, err := result.Get()
	if err != nil {
		return nil, err
	}
	if v.V != nil {
		if knownTypes, ok := v.V.(objects.KnownTypes); ok {
			firstScopedOrImportedItem := knownTypes.ToItemsObs().Filter(func(i interface{}) bool {
				switch i.(type) {
				case *objects.ImportedDcl:
					return true
				case *objects.ScopeDcl:
					return true
				}
				return false
			}).FirstOrDefault(nil)
			item, err := firstScopedOrImportedItem.Get()
			if err != nil {
				return nil, err
			}
			if item.V != nil {
				return nil, fmt.Errorf("not all Scope items and imported items are removed")
			}

			return knownTypes, nil
		}
	}
	return nil, self.logger.Error(CreateKnownTypesError)
}

func NewCreateKnownTypeService(
	logger *log2.SubSystemLogger,
	DclHelpers *Service.DclHelpers,
	IdlDefaultTypes *Service.IdlDefaultTypes) *CreateKnownTypeService {
	return &CreateKnownTypeService{
		logger:          logger,
		DclHelpers:      DclHelpers,
		IdlDefaultTypes: IdlDefaultTypes,
	}
}

func ProvideCreateKnownTypeService() fx.Option {
	return fx.Provide(
		func(factory *log2.LogFactory,
			DclHelpers *Service.DclHelpers,
			IdlDefaultTypes *Service.IdlDefaultTypes) *CreateKnownTypeService {
			return NewCreateKnownTypeService(
				factory.Create("CreateKnownTypeService"),
				DclHelpers,
				IdlDefaultTypes)
		})
}
