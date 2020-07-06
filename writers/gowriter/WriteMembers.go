package gowriter

import (
	"bufio"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"os"
	"path/filepath"
	"syscall"
)

type WriteMembers struct {
	Ctx             *ctx2.GoYaccAppCtx
	Log             *log2.SubSystemLogger
	GoWriterFactory *Factory
	FileResolve     *ctx2.ResolveFileName
	IdlDefaultTypes *Service.IdlDefaultTypes
}

func NewWriteMembers(
	ctx *ctx2.GoYaccAppCtx,
	log *log2.SubSystemLogger,
	goWriterFactory *Factory,
	fileResolve *ctx2.ResolveFileName,
	idlDefaultTypes *Service.IdlDefaultTypes) (*WriteMembers, error) {
	WriteMembers := &WriteMembers{
		Ctx:             ctx,
		Log:             log,
		GoWriterFactory: goWriterFactory,
		FileResolve:     fileResolve,
		IdlDefaultTypes: idlDefaultTypes,
	}
	return WriteMembers, nil
}

func (self WriteMembers) Render(
	knownTypes objects.KnownTypes,
	fileUsage *objects.FileData) error {
	targetFile := filepath.Join(self.Ctx.OutputFolder, fmt.Sprintf("%v.go", fileUsage.TargetFileName))

	err := os.MkdirAll(filepath.Dir(targetFile), os.ModePerm)
	if err != nil {
		return err
	}

	if self.Ctx.ForceOutput {
		_ = os.Remove(targetFile)
	}
	if _, err2 := os.Stat(targetFile); err2 == nil {
		return self.Log.Error(fmt.Errorf("file (%v) already exist", targetFile))
	}
	var fileHandle *os.File = nil
	if fileHandle, err = os.Create(targetFile); err != nil {
		return self.Log.Error(err)
	}
	defer func(fileHandle *os.File) {
		_ = self.Log.Error(fileHandle.Close())
	}(fileHandle)

	buffer := bufio.NewWriter(fileHandle)
	defer func() {
		_ = self.Log.Error(buffer.Flush())
	}()

	packageName := filepath.Base(filepath.Dir(targetFile))

	writer := self.GoWriterFactory.Create(self.Ctx, knownTypes, buffer, targetFile)
	err = writer.WritePackageName(packageName)
	if err != nil {
		return err
	}
	err = writer.FoldersInUse(fileUsage, targetFile)
	if err != nil {
		return err
	}
	err = writer.WriteMembers(knownTypes, fileUsage.Members)
	if err != nil {
		return err
	}
	err = writer.WriteGlobalVariables(knownTypes, fileUsage.Members)
	if err != nil {
		return err
	}
	err = writer.WritePackageInit(knownTypes, fileUsage.Members, knownTypes)
	if err != nil {
		return err
	}

	return nil
}

func (self WriteMembers) GetCreateOutputFolder(path string) (string, error) {
	folderName := ""
	var err error = nil
	if folderName, err = self.FileResolve.Resolve(path); err != nil {
		return "", err
	}
	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(folderName); err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			if errNo, ok := pathError.Err.(syscall.Errno); ok {
				if errNo.Is(os.ErrNotExist) {
					err = os.MkdirAll(folderName, os.ModePerm)
					if err != nil {
						return "", err
					}
					return folderName, nil
				}
			}
		}
		return "", err
	}
	if !fileInfo.IsDir() {
		err = fmt.Errorf("output folder(%v) is not a folder", folderName)
		return "", self.Log.Error(err)
	}
	return folderName, nil
}
