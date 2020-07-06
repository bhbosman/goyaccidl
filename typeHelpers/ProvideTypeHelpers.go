package typeHelpers

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/ctx"
	yaccToken "github.com/bhbosman/yaccidl"
	"go.uber.org/fx"
)

type TypeHelpersOut2 struct {
	fx.Out
	ScopeDcl          *ScopeDcl
	PrimitiveDcl      *PrimitiveDcl
	SequenceDcl       *SequenceDcl
	TypeDefDcl        *TypeDefDcl
	Attributes        *Attributes
	BaseInterfaces    *InheritedInterfaces
	Operations        *Operations
	InterfaceDcl      *InterfaceDcl
	NativeDcl         *NativeDcl
	ConstantDclHelper *ConstantDcl
	StructDcl         *StructDcl
	StructStubDcl     *StructStubDcl
	StructSequenceDcl *StructSequenceDcl
	UnionDcl          *UnionDcl
	EnumDcl           *EnumDcl
}

type ScopeDclParams struct {
	fx.In
	LoggerFactory       *log2.LogFactory
	IdlDefaultTypes     *Service.IdlDefaultTypes
	AppCtx              *ctx.GoYaccAppCtx
	DclHelpers          *Service.DclHelpers
	LanguageTypeService *Service.LanguageTypeService
	FileInformation     *yaccToken.FileInformation
}

func ProvideTypeHelpers() fx.Option {
	return fx.Provide(
		func(lc fx.Lifecycle, params ScopeDclParams) (TypeHelpersOut2, error) {
			operationsDcl := NewOperationsDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			scopeDcl := NewScopeDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			primitiveDcl := NewPrimitiveDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			sequenceDcl := NewSequenceDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			attributesDcl := NewAttributesDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			inheritedInterfacesDcl := NewInheritedInterfacesDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			StructSequenceDcl := NewSeqStructDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			structDcl := NewStructDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)

			structStubDcl := NewStructStubDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)

			enumDcl := NewEnumDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			typeDefDcl := NewTypeDefDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				sequenceDcl,
				primitiveDcl,
				scopeDcl,
				params.DclHelpers,
				structDcl,
				params.LanguageTypeService,
				enumDcl,
				params.FileInformation)
			interfaceDcl := NewInterfaceDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				operationsDcl,
				attributesDcl,
				inheritedInterfacesDcl,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation,
				structDcl)
			nativeDcl := NewNative(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			constantDcl := NewConstantDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)
			unionDcl := NewUnionDcl(
				params.AppCtx,
				params.LoggerFactory,
				params.IdlDefaultTypes,
				params.LanguageTypeService,
				params.DclHelpers,
				params.FileInformation)

			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					// possible circular reference
					return unionDcl.Init(interfaceDcl, typeDefDcl, constantDcl, nativeDcl, scopeDcl, primitiveDcl, sequenceDcl, structDcl)
				},
				OnStop: func(ctx2 context.Context) error {
					return nil
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					// possible circular reference
					return structDcl.Init(
						params.FileInformation)
				},
				OnStop: func(ctx2 context.Context) error {
					return nil
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					// possible circular reference
					return structStubDcl.Init(
						params.FileInformation)
				},
				OnStop: func(ctx2 context.Context) error {
					return nil
				},
			})

			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					// possible circular reference
					return StructSequenceDcl.Init(
						//interfaceDcl,
						//typeDefDcl,
						//constantDcl,
						//nativeDcl,
						//scopeDcl,
						//primitiveDcl,
						//sequenceDcl,
						//unionDcl,
						//importedDcl,
						//enumDcl,
						params.FileInformation)
				},
				OnStop: func(ctx2 context.Context) error {
					return nil
				},
			})

			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return scopeDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return scopeDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return primitiveDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return primitiveDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return sequenceDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return sequenceDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return typeDefDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return typeDefDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return typeDefDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return typeDefDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return inheritedInterfacesDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return inheritedInterfacesDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return operationsDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return operationsDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return interfaceDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return interfaceDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return nativeDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return nativeDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return constantDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return constantDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return structDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return structDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return unionDcl.Start(ctx2)

				},
				OnStop: func(ctx2 context.Context) error {
					return unionDcl.Stop(ctx2)
				},
			})

			return TypeHelpersOut2{
				ScopeDcl:          scopeDcl,
				PrimitiveDcl:      primitiveDcl,
				SequenceDcl:       sequenceDcl,
				TypeDefDcl:        typeDefDcl,
				Attributes:        attributesDcl,
				BaseInterfaces:    inheritedInterfacesDcl,
				Operations:        operationsDcl,
				InterfaceDcl:      interfaceDcl,
				NativeDcl:         nativeDcl,
				ConstantDclHelper: constantDcl,
				StructDcl:         structDcl,
				StructStubDcl:     structStubDcl,
				StructSequenceDcl: StructSequenceDcl,
				UnionDcl:          unionDcl,
				EnumDcl:           enumDcl,
			}, nil
		})
}
