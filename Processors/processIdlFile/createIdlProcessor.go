package processIdlFile

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
	"github.com/bhbosman/goyaccidl/typeHelpers"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
	"github.com/bhbosman/goyaccidl/yaccSpecProcessing"
	"github.com/bhbosman/yaccidl"
	"go.uber.org/fx"
	"log"
	"os"
)

func CreateIdlProcessor(
	flags *ctx.GoYaccAppCtx,
	logger *log.Logger,
	resolveFileName *ctx.ResolveFileName) (*fx.App, process.IFactory, *log2.LogFactory) {
	var factory process.IFactory
	var logFactory *log2.LogFactory
	app := fx.New(
		fx.Logger(logger),
		log2.ProvideLogFactory(logger, func(logFactory *log2.LogFactory) {
			logFactory.SetLogLevel("YaccSpecToDeclTypes", 1)
			logFactory.SetLogLevel("YaccUserData", 0)
		}),

		gowriter.ProvideAllWriters(),

		gowriter.ProvideGoWriterFactory(),

		log2.ProvideLogger(logger),
		typeHelpers.ProvideTypeHelpers(),
		implementations.ProvideUniqueSessionNumber(),
		ProvideIdlProcess(),
		Service.Provide(),
		yaccidl.ProvideFileInformation(),
		ProvideCreateKnownTypeService(),
		yaccSpecProcessing.Provide(),
		ProvideProcessFactory(),
		ctx.ProvideGoYaccAppCtx(flags),
		ctx.ProvideResolveFileName(resolveFileName),
		fx.Invoke(
			func(lc fx.Lifecycle, item *ctx.ResolveFileName, flags *ctx.GoYaccAppCtx) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						var err error
						flags.InputFileName, err = item.Resolve(flags.InputFileName)
						return err
					},
					OnStop: nil,
				})
			}),
		fx.Invoke(
			func(lc fx.Lifecycle, flags *ctx.GoYaccAppCtx) {
				var l *os.File
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						var err error
						if flags.InputFileName != "" {
							stat, err := os.Stat(flags.InputFileName)
							if err != nil {
								return err
							}
							if stat.IsDir() {
								return fmt.Errorf("can not open folder")
							}
						}
						if flags.InputFileName == "" {
							flags.InputReader = flags.Stdin
							flags.InputFileName = "(stdin)"
						} else {
							l, err = os.Open(flags.InputFileName)
							if err != nil {
								return err
							}
							flags.InputReader = bufio.NewReader(l)
						}
						return err
					},
					OnStop: func(ctx context.Context) error {
						if l != nil {
							return l.Close()
						}
						return nil
					},
				})
			}),
		fx.Invoke(
			func(lc fx.Lifecycle, flags *ctx.GoYaccAppCtx, log *log.Logger) {
				//var g *os.File
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				})
			}),
		fx.Populate(&resolveFileName),
		fx.Populate(&factory),
		fx.Populate(&logFactory))
	return app, factory, logFactory

}
