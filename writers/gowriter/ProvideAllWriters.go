package gowriter

import (
	"context"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/ctx"
	"go.uber.org/fx"
)

func ProvideAllWriters() fx.Option {
	return fx.Provide(
		func(lc fx.Lifecycle, params ProvideAllWritersIn) ProvideAllWritersOut {
			WriteTypeDefOfPrimitiveDcl := NewWriteTypeDefOfPrimitiveDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteTypeDefOfPrimitiveDcl"),
				params.DclHelpers,
				params.LanguageTypeService)
			WriteTypeDefOfNativeDcl := NewWriteTypeDefOfNativeDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteTypeDefOfNativeDcl"),
				params.LanguageTypeService,
				params.DclHelpers)

			WriteStructHelper := NewWriteStructHelper(
				params.IdlToGoTranslation,
				params.LogFactory.Create("StructWriter"),
				params.DclHelpers,
				params.LanguageTypeService)

			WriteStructDcl := NewWriteStructDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteStructDcl"),
				WriteStructHelper,
				params.LanguageTypeService,
				params.DclHelpers)



			WriteStructStubDcl := NewWriteStructStubDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteStructDcl"),
				WriteStructHelper,
				params.LanguageTypeService,
				params.DclHelpers)
			WriteStructSkeletonDcl := NewWriteStructSkeletonDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteStructDcl"),
				WriteStructHelper,
				params.LanguageTypeService,
				params.DclHelpers)










			WriteStructSequenceDcl := NewWriteStructSequenceDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteStructSequenceDcl"),
				params.LanguageTypeService,
				params.DclHelpers)

			WriteUnionDcl := NewWriteUnionDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteUnionDcl"),
				WriteStructHelper,
				params.DclHelpers,
				params.LanguageTypeService)

			WriteNativeDcl := NewNativeDclWriter(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteNativeDcl"),
				params.LanguageTypeService,
				params.DclHelpers)

			WriteInterface := NewWriteInterface(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteInterface"),
				params.DclHelpers,
				params.LanguageTypeService,
				params.InterfaceHelper)
			WriteTypeDefDcl := NewWriteTypeDefDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteTypeDefDcl"),
				WriteTypeDefOfPrimitiveDcl,
				WriteTypeDefOfNativeDcl,
				params.DclHelpers,
				params.LanguageTypeService)
			WriteConstantValue := NewConstantValueWriter(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteConstantValue"),
				params.LanguageTypeService,
				params.DclHelpers)
			WriteEnumDcl := NewEnumDclWriter(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteEnumDcl"),
				params.DclHelpers,
				params.LanguageTypeService)
			WriteScopeDcl := NewScopeDclWriter(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WriteScopeDcl"),
				params.LanguageTypeService,
				params.DclHelpers)
			WritePrimitiveDcl := NewWritePrimitiveDcl(
				params.IdlToGoTranslation,
				params.LogFactory.Create("WritePrimitiveDcl"),
				params.LanguageTypeService,
				params.DclHelpers)

			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WritePrimitiveDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WritePrimitiveDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteNativeDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteNativeDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteUnionDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteUnionDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteStructDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteStructDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteInterface.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteInterface.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteTypeDefDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteTypeDefDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteConstantValue.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteConstantValue.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteEnumDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteEnumDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteScopeDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteScopeDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteTypeDefOfPrimitiveDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteTypeDefOfPrimitiveDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteTypeDefOfNativeDcl.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteTypeDefOfNativeDcl.Stop(ctx2)
				},
			})
			lc.Append(fx.Hook{
				OnStart: func(ctx2 context.Context) error {
					return WriteStructHelper.Start(ctx2)
				},
				OnStop: func(ctx2 context.Context) error {
					return WriteStructHelper.Stop(ctx2)
				},
			})

			return ProvideAllWritersOut{
				Out:                        fx.Out{},
				WritePrimitiveDcl:          WritePrimitiveDcl,
				WriteNativeDcl:             WriteNativeDcl,
				WriteUnionDcl:              WriteUnionDcl,
				WriteStructDcl:             WriteStructDcl,
				WriteStructStubDcl:         WriteStructStubDcl,
				WriteStructSkeletonDcl:     WriteStructSkeletonDcl,
				WriteInterface:             WriteInterface,
				WriteTypeDefDcl:            WriteTypeDefDcl,
				WriteConstantValue:         WriteConstantValue,
				WriteEnumDcl:               WriteEnumDcl,
				WriteScopeDcl:              WriteScopeDcl,
				WriteTypeDefOfPrimitiveDcl: WriteTypeDefOfPrimitiveDcl,
				WriteTypeDefOfNativeDcl:    WriteTypeDefOfNativeDcl,
				WriteStructSequenceDcl:     WriteStructSequenceDcl,
			}
		})
}

type ProvideAllWritersIn struct {
	fx.In
	AppCtx              *ctx.GoYaccAppCtx
	LogFactory          *log2.LogFactory
	DclHelpers          *Service.DclHelpers
	LanguageTypeService *Service.LanguageTypeService
	IdlToGoTranslation  *Service.IdlToGoTranslation
	InterfaceHelper     *Service.InterfaceHelper
}

type ProvideAllWritersOut struct {
	fx.Out
	WritePrimitiveDcl          *WritePrimitiveDcl
	WriteNativeDcl             *WriteNativeDcl
	WriteUnionDcl              *WriteUnionDcl
	WriteStructDcl             *WriteStructDcl
	WriteStructStubDcl         *WriteStructStubDcl
	WriteStructSkeletonDcl     *WriteStructSkeletonDcl
	WriteInterface             *WriteInterface
	WriteTypeDefDcl            *WriteTypeDefDcl
	WriteConstantValue         *WriteConstantValue
	WriteEnumDcl               *WriteEnumDcl
	WriteScopeDcl              *WriteScopeDcl
	WriteTypeDefOfPrimitiveDcl *WriteTypeDefOfPrimitiveDcl
	WriteTypeDefOfNativeDcl    *WriteTypeDefOfNativeDcl
	WriteStructSequenceDcl     *WriteStructSequenceDcl
}
