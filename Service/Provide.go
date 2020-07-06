package Service

import (
	log2 "github.com/bhbosman/gocommon/log"
	"go.uber.org/fx"
)

type ProvideInParams struct {
	fx.In
	LogFactory *log2.LogFactory
}

type ProvideOutParams struct {
	fx.Out
	RemoveTypedef       *RemoveTypedef
	DclHelpers          *DclHelpers
	IdlDefaultTypes     *IdlDefaultTypes
	LanguageTypeService *LanguageTypeService
	IdlToGoTranslation  *IdlToGoTranslation
	ExpandInterfaces    *ExpandInterfaces
	InterfaceHelper     *InterfaceHelper
}

func Provide() fx.Option {
	return fx.Provide(func(params ProvideInParams) ProvideOutParams {
		LanguageTypeService := newLanguageTypeService()
		RemoveTypedef := NewRemoveTypedef(LanguageTypeService)
		DclHelpers := NewDclHelpers(LanguageTypeService)
		IdlDefaultTypes := NewIdlDefaultTypes(LanguageTypeService)
		ExpandInterfaces := NewExpandInterfaces(LanguageTypeService, DclHelpers)
		IdlToGoTranslation := NewIdlToGoTranslation(
			params.LogFactory.Create("IdlToGoTranslation"),
			DclHelpers,
			LanguageTypeService)
		InterfaceHelper := NewInterfaceHelper(DclHelpers, IdlToGoTranslation)
		return ProvideOutParams{
			Out:                 fx.Out{},
			RemoveTypedef:       RemoveTypedef,
			DclHelpers:          DclHelpers,
			IdlDefaultTypes:     IdlDefaultTypes,
			LanguageTypeService: LanguageTypeService,
			IdlToGoTranslation:  IdlToGoTranslation,
			ExpandInterfaces:    ExpandInterfaces,
			InterfaceHelper:     InterfaceHelper,
		}
	})
}
