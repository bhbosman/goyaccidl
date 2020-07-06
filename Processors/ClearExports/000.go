package ClearExports

import (
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
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
					params.Ctx)
			}),
		fx.Populate(&factory),
		fx.Populate(&logFactory))
	return app, factory, logFactory
}

type ProcessFactoryImplParams struct {
	fx.In
	LoggerFactory *log2.LogFactory
	Ctx           *ctx.GoYaccAppCtx
}
