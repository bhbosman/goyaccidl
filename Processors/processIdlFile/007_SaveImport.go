package processIdlFile

import (
	"encoding/json"
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/process"

	"os"
	"path/filepath"
)

type SaveImport struct {
	logger *log2.SubSystemLogger
	appCtx *ctx2.GoYaccAppCtx
}

func (self SaveImport) Name() string {
	return self.logger.Name()
}

func (self *SaveImport) CreateScopeItem(member objects.IDcl) *objects.ScopeItem {
	name := member.GetName()

	return objects.NewScopeItem(
		member,
		member.GetPrimitiveType(),
		member.GetOrderId(),
		name,
		member.GetLexemData().GetSourceLexemData())
}

func (self SaveImport) Run(incoming interface{}) (interface{}, error) {

	if FileAllocationToGolangOut, ok := incoming.(*FileAllocationToGolangOut); ok {
		m := make(map[string]*objects.ScopeDeclaration)

		for _, dcl := range FileAllocationToGolangOut.Files.ToDclArray() {
			f := dcl.GetName().First()
			if v, b := m[f]; b {
				v.DeclaredItems[dcl.GetName()] = self.CreateScopeItem(dcl)
			} else {
				v := objects.NewScopeDeclaration(nil)
				v.DeclaredItems[dcl.GetName()] = self.CreateScopeItem(dcl)
				m[f] = v
			}
		}
		for k, v := range m {
			self.SaveData(k, v)
		}
	}

	return incoming, nil
}

func (self SaveImport) SaveData(key string, value *objects.ScopeDeclaration) error {
	keyFileName := filepath.Join(self.appCtx.AppFolder, fmt.Sprintf("%v_export.json", key))
	keyFileHandle, err := os.Create(keyFileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = keyFileHandle.Close()
	}()

	encoder := json.NewEncoder(keyFileHandle)
	encoder.SetIndent("", "\t")
	return encoder.Encode(value)
}

func NewSaveImport(logFactory *log2.LogFactory, appCtx *ctx2.GoYaccAppCtx) process.IProcess {
	return &SaveImport{
		logger: logFactory.Create("SaveImport"),
		appCtx: appCtx,
	}
}
