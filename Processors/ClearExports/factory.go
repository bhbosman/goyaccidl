package ClearExports

import (
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
)

type Factory struct {
	logFactory *log2.LogFactory
	ctx        *ctx2.GoYaccAppCtx
}

func (f Factory) Create() (process.IProcessWrapper, error) {
	var err error
	var step01 *process.ProcessorWrapper
	step01, err = process.NewProcessWrapper(f.logFactory, NewClearExport(f.ctx, f.logFactory.Create("ClearExport")))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return step01, err
}

func NewFactory(
	loggerFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx) *Factory {
	return &Factory{
		logFactory: loggerFactory,
		ctx:        ctx,
	}
}
