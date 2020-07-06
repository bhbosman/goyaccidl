package ClearExports

import (
	"fmt"
	"github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"io"
	log2 "log"
	"os"
	"path/filepath"
	"regexp"
)

type ClearExport struct {
	ctx    *ctx2.GoYaccAppCtx
	logger *log.SubSystemLogger
}

type ClearExportResult struct {
}

func NewClearExport(ctx *ctx2.GoYaccAppCtx, logger *log.SubSystemLogger) *ClearExport {
	return &ClearExport{
		ctx:    ctx,
		logger: logger,
	}
}

func (self ClearExport) Name() string {
	return "get primitive types"
}

func (self ClearExport) Run(_ interface{}) (interface{}, error) {
	stat, err := os.Stat(self.ctx.AppFolder)
	if err != nil {
		return nil, self.logger.Error(err)
	}
	if !stat.IsDir() {
		return nil, self.logger.Error(fmt.Errorf("%v is not a folder must be a folder", self.ctx.AppFolder))
	}

	dirRead, err := os.Open(self.ctx.AppFolder)
	if err != nil {
		return nil, self.logger.Error(err)
	}

	defer func(file io.Closer) {
		_ = self.logger.Error(file.Close())
	}(dirRead)

	dirFiles, err := dirRead.Readdir(0)
	if err != nil {
		return nil, self.logger.Error(err)
	}

	compile, err := regexp.Compile("([:a-zA-Z0-9])+(_export.json)$")
	if err != nil {
		return nil, self.logger.Error(err)
	}

	for index := range dirFiles {
		fileHere := dirFiles[index]
		name := fileHere.Name()
		n := compile.FindStringIndex(name)
		if n == nil {
			continue
		}
		if n[0] == 0 && n[1] == len(fileHere.Name()) {
			fullPath := filepath.Join(self.ctx.AppFolder, fileHere.Name())
			err = os.Remove(fullPath)
			if err != nil {
				return nil, self.logger.Error(err)
			}
			self.logger.LogWithLevel(0, func(logger *log2.Logger) {
				logger.Printf("Remove file: %v", fullPath)
			})
		}
	}
	return &ClearExportResult{}, nil
}
