package ctx

import (
	"bufio"
	"flag"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/process"
	"go.uber.org/fx"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type CreateProcessIdlAppCallback func(
	flags *GoYaccAppCtx,
	logger *log.Logger,
	resolveFileName *ResolveFileName) (*fx.App, process.IFactory, *log2.LogFactory)
type GoYaccAppCtx struct {
	AppFolder        string
	resolveFileName  *ResolveFileName
	OutputFolder     string
	ForceOutput      bool
	Verbose          bool
	hflag            bool
	InputFileName    string
	InputReader      *bufio.Reader
	IdlDefinitions   IdlDefinitions
	Stdin            *bufio.Reader
	Stdout           *bufio.Writer
	Stderr           *bufio.Writer
	GoSrcPath        string
	processors       map[string]CreateProcessIdlAppCallback
	DefaultProcessor string
}

func NewGoYaccAppCtx(
	defaultProcessor string,
	resolveFileName *ResolveFileName, stdin *bufio.Reader, stdout *bufio.Writer, stderr *bufio.Writer) *GoYaccAppCtx {
	return &GoYaccAppCtx{
		AppFolder:       "",
		resolveFileName: resolveFileName,
		//Oflag:            "",
		OutputFolder: "",
		ForceOutput:  false,
		Verbose:      false,
		hflag:        false,
		//Tflag:         false,
		InputFileName: "",
		InputReader:   nil,
		//OutputWriter:     nil,
		IdlDefinitions:   NewIdlDefinitions(),
		Stdin:            stdin,
		Stdout:           stdout,
		Stderr:           stderr,
		GoSrcPath:        "",
		processors:       make(map[string]CreateProcessIdlAppCallback),
		DefaultProcessor: defaultProcessor,
	}
}

func (self *GoYaccAppCtx) Run() (CreateProcessIdlAppCallback, error) {
	err := self.ReadFlags()
	if err != nil {
		return nil, err
	}
	cb, ok := self.processors[self.DefaultProcessor]
	if !ok {
		return nil, fmt.Errorf("could not find process %v\n", self.DefaultProcessor)
	}
	return cb, nil
}

func (self *GoYaccAppCtx) RegisterProcess(s string, cb CreateProcessIdlAppCallback) error {
	if _, ok := self.processors[s]; ok {
		return fmt.Errorf("already in the list")
	}
	self.processors[s] = cb
	return nil
}

func (self *GoYaccAppCtx) ReadFlags() error {
	self.GoSrcPath = filepath.Join(os.Getenv("GOPATH"), "./src")

	flag.StringVar(&self.DefaultProcessor, "processor", self.DefaultProcessor, "")
	flag.BoolVar(&self.Verbose, "v", self.Verbose, "verbose")
	flag.BoolVar(&self.hflag, "h", false, "show help and exit")
	flag.StringVar(&self.OutputFolder, "of", "", "The output folder of the parsed Idl")
	flag.BoolVar(&self.ForceOutput, "ff", false, "Force and override existing files")
	flag.Var(&self.IdlDefinitions, "idldef", "list of #definitions you want to activate, \nex: \"-idldef a,b,c\" will \n\tflag a, b and c as set and can be used in #ifdef statements")
	flag.Parse()
	self.InputFileName = flag.Arg(0)
	if self.hflag || flag.NArg() > 1 {
		flag.Usage()
		_, _ = fmt.Fprintf(self.Stderr, "\n%s [-o out_name] [other_options] [in_name]\n", os.Args[0])
		_, _ = fmt.Fprintln(self.Stderr, "  If no in_name is given then read from stdin.")
		self.Stderr.Flush()
		return fmt.Errorf("exit as help was required")
	}

	if self.OutputFolder == "" {
		self.OutputFolder = "./"
	}
	self.OutputFolder, _ = self.resolveFileName.Resolve(self.OutputFolder)
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	userHomeDir = filepath.Join(userHomeDir, ".goyaccidl")
	_, err = os.Stat(userHomeDir)
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			if pathError.Err == syscall.ENOENT {
				err = os.MkdirAll(userHomeDir, os.ModePerm)
				return err
			}
		}
	}
	self.AppFolder = userHomeDir
	return nil
}

func (self *GoYaccAppCtx) RunProcessor() error {
	_, ok := self.processors[self.DefaultProcessor]
	if !ok {
		return fmt.Errorf("could not find process %v", self.DefaultProcessor)
	}

	return nil
}

func ProvideGoYaccAppCtx(flags *GoYaccAppCtx) fx.Option {
	return fx.Provide(
		func() *GoYaccAppCtx {
			return flags
		})
}
