package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bhbosman/gocommon"
	"github.com/bhbosman/goyaccidl/Processors/ClearExports"
	"github.com/bhbosman/goyaccidl/Processors/processIdlFile"
	"github.com/bhbosman/goyaccidl/Processors/processPrimitiveTypes"
	"github.com/bhbosman/goyaccidl/ctx"
	process2 "github.com/bhbosman/goyaccidl/process"
	"io"
	"log"
	"os"
)

func main() {
	resolveFileName := &ctx.ResolveFileName{}
	defaultProcessor := "IDLProcessor"
	flags := ctx.NewGoYaccAppCtx(defaultProcessor,
		resolveFileName,
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
		bufio.NewWriter(os.Stderr))
	_ = flags.RegisterProcess(defaultProcessor, processIdlFile.CreateIdlProcessor)
	_ = flags.RegisterProcess("PrimitiveProcessor", processPrimitiveTypes.CreatePrimitiveProcessor)
	_ = flags.RegisterProcess("ClearExports", ClearExports.CreatePrimitiveProcessor)

	cb, err := flags.Run()
	if err != nil {
		s := fmt.Sprintf("error: %v", err.Error())
		os.Stderr.Write([]byte(s))
		os.Exit(1)
	}
	getLogger := func(verbose bool) io.Writer {
		if verbose {
			return os.Stdout
		}
		return &gocommon.NullWriter{}
	}
	logger := log.New(getLogger(flags.Verbose), "goyaccidl: ", log.Ldate|log.Ltime|log.Lmicroseconds)

	app, factory, logFactory := cb(flags, logger, resolveFileName)
	subSystemLogger := logFactory.Create("main")

	exitCode := 0
	defer func() {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	subSystemLogger.LogWithLevel(0, func(logger *log.Logger) { logger.Println("App start...") })
	startError := app.Start(context.TODO())
	if startError != nil {
		exitCode = 1
		return
	}
	defer func() {
		subSystemLogger.LogWithLevel(0, func(logger *log.Logger) { logger.Println("App stop...") })
		_ = app.Stop(context.TODO())
	}()

	subSystemLogger.LogWithLevel(0, func(logger *log.Logger) { logger.Println("Create process") })
	process, err := factory.Create()
	if err != nil {
		subSystemLogger.Error(fmt.Errorf("error occured. Error: %v", err.Error()))
		exitCode = 4
		return
	}

	inputData := process2.NewProcessWrapperInput(nil, nil)
	subSystemLogger.LogWithLevel(0, func(logger *log.Logger) { logger.Println("Resolve process") })
	_, err = process.Run(inputData)
	if err != nil {
		exitCode = 10
		return
	}
}
