package processPrimitiveTypes

import (
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
	"go.uber.org/fx"
	"log"
)

func CreatePrimitiveProcessor(
	flags *ctx.GoYaccAppCtx,
	logger *log.Logger,
	resolveFileName *ctx.ResolveFileName) (*fx.App, process.IFactory, *log2.LogFactory) {
	var factory process.IFactory
	var logFactory *log2.LogFactory
	app := fx.New(
		fx.Logger(logger),
		log2.ProvideLogFactory(logger, nil),
		gowriter.ProvideAllWriters(),

		Service.Provide(),

		gowriter.ProvideGoWriterFactory(),
		fx.Provide(
			func() *ctx.GoYaccAppCtx {
				return flags
			}),
		fx.Provide(
			func(flags *ctx.GoYaccAppCtx) *ctx.ResolveFileName {
				return resolveFileName
			}),
		fx.Provide(
			func(params ProcessFactoryImplParams) process.IFactory {
				return NewFactory(
					params.LoggerFactory,
					params.Ctx,
					params.GoWriterFactory,
					params.DefaultTypes,
					params.FileResolve)
			}),
		fx.Populate(&factory),
		fx.Populate(&logFactory))
	return app, factory, logFactory
}

type ProcessFactoryImplParams struct {
	fx.In
	LoggerFactory   *log2.LogFactory
	Ctx             *ctx.GoYaccAppCtx
	GoWriterFactory *gowriter.Factory
	DefaultTypes    *Service.IdlDefaultTypes
	FileResolve     *ctx.ResolveFileName
}
