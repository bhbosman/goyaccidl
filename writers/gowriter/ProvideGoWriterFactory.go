package gowriter

import (
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/ctx"
	"go.uber.org/fx"
)

type FactoryParams struct {
	fx.In
	LogFactory           *log2.LogFactory
	WriteStructDcl       *WriteStructDcl
	UnionWriter          *WriteUnionDcl
	WriteInterface       *WriteInterface
	WriteTypeDefDcl      *WriteTypeDefDcl
	ConstantValueWriter  *WriteConstantValue
	WriteEnumDcl         *WriteEnumDcl
	ScopeDclWriter       *WriteScopeDcl
	NativeDclWriter      *WriteNativeDcl
	WritePrimitiveDcl    *WritePrimitiveDcl
	FileResolve          *ctx.ResolveFileName
	WriteSeqStructDcl    *WriteStructSequenceDcl
	StructSkeletonWriter *WriteStructSkeletonDcl
	StructStubWriter     *WriteStructStubDcl
}

func ProvideGoWriterFactory() fx.Option {
	return fx.Provide(
		func(params FactoryParams) *Factory {
			return NewGoWriterFactory(
				params.LogFactory,
				params.UnionWriter,
				params.WriteStructDcl,
				params.WriteInterface,
				params.WriteTypeDefDcl,
				params.ConstantValueWriter,
				params.WriteEnumDcl,
				params.ScopeDclWriter,
				params.NativeDclWriter,
				params.WritePrimitiveDcl,
				params.WriteSeqStructDcl,
				params.FileResolve,
				params.StructSkeletonWriter,
				params.StructStubWriter)
		})
}
