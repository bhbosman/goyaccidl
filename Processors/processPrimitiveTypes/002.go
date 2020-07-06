package processPrimitiveTypes

import (
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
	///**/"go.uber.org/multierr"
)

type ProcessorStep02 struct {
	WriteMembers *gowriter.WriteMembers
}

func NewProcessorStep02(
	logFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	GoWriterFactory *gowriter.Factory,
	defaultTypes *Service.IdlDefaultTypes,
	fileResolve *ctx2.ResolveFileName) (*ProcessorStep02, error) {

	writeMembers, err := gowriter.NewWriteMembers(
		ctx,
		logFactory.Create("RenderToGo"),
		GoWriterFactory,
		fileResolve,
		defaultTypes)
	if err != nil {
		return nil, err
	}
	return &ProcessorStep02{
		WriteMembers: writeMembers,
	}, nil
}

func (self ProcessorStep02) Name() string {
	return "get primitive types"
}

func (self ProcessorStep02) Run(incoming interface{}) (interface{}, error) {
	ddd, ok := incoming.(objects.IDclArray)
	if ok && ddd != nil {
		knownTypes := make(objects.KnownTypes)
		for _, scopeItem := range self.WriteMembers.IdlDefaultTypes.IdlPrimitivesInScopedFormat() {
			knownTypes[scopeItem.GetName()] = scopeItem

		}
		data := objects.NewFileData("omg", "omg", "defaultTypes")
		for _, dclNode := range self.WriteMembers.IdlDefaultTypes.IdlPrimitivesInScopedFormat() {
			data.Add(dclNode)
		}
		self.WriteMembers.Render(knownTypes, data)
	}
	return self.WriteMembers.IdlDefaultTypes.IdlPrimitivesInScopedFormat(), nil
}
