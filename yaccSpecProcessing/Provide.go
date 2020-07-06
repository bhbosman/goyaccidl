package yaccSpecProcessing

import (
	"context"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"go.uber.org/fx"
)

type IncomingParams = struct {
	fx.In
	LogFactory          *log2.LogFactory
	Ctx                 *ctx2.GoYaccAppCtx
	UniqueSessionNumber *implementations.UniqueSessionNumber
}

type OutgoingParams = struct {
	fx.Out
	YaccSpecToDeclAll IYaccSpecToDeclAll
}

func Provide() fx.Option {
	return fx.Provide(
		func(lifeCycle fx.Lifecycle, params IncomingParams) (OutgoingParams, error) {
			YaccSpecToDeclForInterface := NewYaccSpecToDeclForInterface(
				params.LogFactory.Create("YaccSpecToDeclForInterface"),
				params.Ctx,
				params.UniqueSessionNumber)

			processAll := NewYaccSpecToDeclAll(
				params.LogFactory.Create("YaccSpecToDeclAll"),
				params.Ctx,
				params.UniqueSessionNumber)

			err := YaccSpecToDeclForInterface.init(processAll)
			if err != nil {
				return OutgoingParams{}, err
			}

			err = processAll.init(YaccSpecToDeclForInterface)
			if err != nil {
				return OutgoingParams{}, err
			}

			lifeCycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return YaccSpecToDeclForInterface.OnStart(ctx)

					},
					OnStop: func(ctx context.Context) error {
						return YaccSpecToDeclForInterface.OnStop(ctx)
					},
				})
			lifeCycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return processAll.OnStart(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return processAll.OnStop(ctx)
					},
				})

			return OutgoingParams{
				Out:               fx.Out{},
				YaccSpecToDeclAll: processAll,
			}, nil
		})
}
