package processPrimitiveTypes

import (
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
)

type Factory struct {
	logFactory      *log2.LogFactory
	ctx             *ctx2.GoYaccAppCtx
	GoWriterFactory *gowriter.Factory
	defaultTypes    *Service.IdlDefaultTypes
	fileResolve     *ctx2.ResolveFileName
}

func (f Factory) Create() (process.IProcessWrapper, error) {
	var err error
	var step01, step02 *process.ProcessorWrapper
	step01, err = process.NewProcessWrapper(f.logFactory, NewProcessorStep01(f.defaultTypes))
	if err != nil {
		return nil, err
	}
	step02Internal, err := NewProcessorStep02(
		f.logFactory,
		f.ctx,
		f.GoWriterFactory,
		f.defaultTypes,
		f.fileResolve)
	if err != nil {
		return nil, err
	}
	step02, err = process.NewProcessWrapper(f.logFactory, step02Internal)
	if err != nil {
		return nil, err
	}

	step01.SetNext(step02)

	return step01, err
}

func NewFactory(
	loggerFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	GoWriterFactory *gowriter.Factory,
	defaultTypes *Service.IdlDefaultTypes,
	fileResolve *ctx2.ResolveFileName) *Factory {
	return &Factory{
		logFactory:      loggerFactory,
		ctx:             ctx,
		GoWriterFactory: GoWriterFactory,
		defaultTypes:    defaultTypes,
		fileResolve:     fileResolve,
	}
}
