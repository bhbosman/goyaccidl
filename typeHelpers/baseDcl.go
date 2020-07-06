package typeHelpers

import (
	"context"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	yaccToken "github.com/bhbosman/yaccidl"
	"path/filepath"
	"strings"

	//"github.com/bhbosman/yaccidl"
	"log"
)

type BaseDcl struct {
	appCtx              *ctx2.GoYaccAppCtx
	logger              *log2.SubSystemLogger
	IdlDefaultTypes     *Service.IdlDefaultTypes
	LanguageTypeService *Service.LanguageTypeService
	DclHelpers          *Service.DclHelpers
	fileInformation     *yaccToken.FileInformation
}

func (self *BaseDcl) FindType(
	currentScope CurrentScope,
	keyMap objects.LocalKeyMap,
	key objects.ScopeIdentifier,
	lexemData objects.IDstSrcInformation) (*objects.KeyMapData, bool, error) {

	v, ok := keyMap[key]
	if ok {
		return v, ok, nil
	}

	scope := key.Scope()
	if _, ok := currentScope[scope]; ok {
		return nil, false, nil
	}

	return nil, false, self.logger.Error(fmt.Errorf("can not find type %v(%v)", key, lexemData))
}

func (self *BaseDcl) Stop(context.Context) error {
	self.appCtx = nil
	self.logger = nil
	self.IdlDefaultTypes = nil
	return nil
}

func (self *BaseDcl) Start(context.Context) error {
	b := true
	b = b && self.appCtx != nil
	b = b && self.logger != nil
	b = b && self.IdlDefaultTypes != nil
	if !b {
		return fmt.Errorf("BaseDcl not initialized")
	}
	return nil
}

func (self *BaseDcl) ResolveType(currentScope CurrentScope, keyMap objects.LocalKeyMap, dcl objects.IDcl) error {
	if dcl == nil {
		return nil
	}
	name := dcl.GetName()
	v, ok, err := self.FindType(currentScope, keyMap, name, dcl.GetLexemData())

	if err != nil {
		self.logger.LogWithLevel(0, func(logger *log.Logger) {
			logger.Printf("Ignore error as I need to fix issues where declares are used and not yet declared")
		})
	}
	if ok {
		err := dcl.UpdateIdlReference(v.IdlReference)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (self *BaseDcl) InternalResolveFolderUsage(
	usage objects.FileUsage,
	folderId, folderName string,
	what string) error {
	if folderId != "" {
		d, ok := usage[folderId]
		if !ok {
			if filepath.IsAbs(folderName) {
				if strings.Index(folderName, self.appCtx.GoSrcPath) == 0 {
					targetFileArray := strings.Split(folderName, string(filepath.Separator))
					pathFromGoSrc := filepath.Dir(filepath.Join(targetFileArray[len(strings.Split(self.appCtx.GoSrcPath, string(filepath.Separator))):]...))
					inUse := objects.NewFindFoldersInUse(folderId, pathFromGoSrc, what)
					usage[inUse.FolderId] = inUse
				}
			}
			inUse := objects.NewFindFoldersInUse(folderId, folderName, what)
			usage[inUse.FolderId] = inUse
		} else {
			d.Add(fmt.Sprintf("%v", what))
		}
	}
	return nil
}
func (self *BaseDcl) DclResolveFolderUsage(dcl objects.IDcl) (id string, folderName string) {
	return dcl.DclResolveFolderUsage()
}

func (self *BaseDcl) ResolveFolderUsage(
	usage objects.FileUsage,
	fromWhere string,
	dcl objects.IDcl,
	knownTypes objects.KnownTypes) error {
	folderId, folderName := self.DclResolveFolderUsage(dcl)
	return self.InternalResolveFolderUsage(
		usage,
		folderId,
		folderName,
		fmt.Sprintf("%v(%v)", fromWhere, dcl.GetName()))
}

func NewBaseDcl(
	appCtx *ctx2.GoYaccAppCtx,
	logger *log2.SubSystemLogger,
	IdlDefaultTypes *Service.IdlDefaultTypes,
	LanguageTypeService *Service.LanguageTypeService,
	DclHelpers *Service.DclHelpers,
	fileInformation *yaccToken.FileInformation) BaseDcl {
	return BaseDcl{
		appCtx:              appCtx,
		logger:              logger,
		IdlDefaultTypes:     IdlDefaultTypes,
		LanguageTypeService: LanguageTypeService,
		DclHelpers:          DclHelpers,
		fileInformation:     fileInformation,
	}
}
