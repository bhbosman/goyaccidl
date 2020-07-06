package process

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"log"
	"reflect"
)

type ProcessorWrapper struct {
	logger        log2.SubSystemLogger
	processorName string
	processor     IProcess
	nextProcessor IProcessWrapper
}

func (pw *ProcessorWrapper) SetNext(next IProcessWrapper) {
	pw.nextProcessor = next
}

func (pw ProcessorWrapper) Name() string {
	return "(ProcessWrapper)"
}

func NewProcessWrapper(logger *log2.LogFactory, processor IProcess) (*ProcessorWrapper, error) {
	if processor == nil {
		return nil, fmt.Errorf("processor can not be nil")
	}
	return &ProcessorWrapper{
		logger:        *logger.Create(fmt.Sprintf("PW-%v", processor.Name())),
		processorName: processor.Name(),
		processor:     processor,
	}, nil
}

func (pw ProcessorWrapper) Run(input IProcessWrapperInput) (interface{}, error) {
	if pw.processor != nil {
		pw.logger.LogWithLevel(0, func(logger *log.Logger) {
			logger.Printf("Before run %v", pw.processorName)
		})

		result, err := pw.processor.Run(input.Input())
		if err != nil {
			return nil, pw.logger.Error(err)
		}
		if result == nil {
			return nil, pw.logger.Error(fmt.Errorf("data is nil"))
		}

		var resultType string
		resultType = reflect.TypeOf(result).String()

		pw.logger.LogWithLevel(0, func(logger *log.Logger) {
			logger.Printf("After run %v, result type: %v", pw.processorName, resultType)
		})

		input.SetResult(pw.processor.Name(), result)
		if pw.nextProcessor != nil {
			var nextInput IProcessWrapperInput = NewProcessWrapperInput(result, input)

			return pw.nextProcessor.Run(nextInput)
		}
		return input, nil
	}
	pw.logger.LogWithLevel(0, func(logger *log.Logger) {
		logger.Printf("No IProcess defined for %v", pw.processorName)
	})

	return nil, pw.logger.Error(fmt.Errorf("IProcess not assigned %v", pw.processorName))
}
