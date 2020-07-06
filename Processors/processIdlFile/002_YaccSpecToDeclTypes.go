package processIdlFile

import (
	"fmt"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/yaccSpecProcessing"
	"github.com/bhbosman/yaccidl"
	"sort"
)

type YaccSpecToDeclTypes struct {
	logger              *log2.SubSystemLogger
	ctx                 *ctx2.GoYaccAppCtx
	uniqueSessionNumber *implementations.UniqueSessionNumber
	YaccSpecToDeclAll   yaccSpecProcessing.IYaccSpecToDeclAll
}

func (self YaccSpecToDeclTypes) Name() string {
	return self.logger.Name()
}

func (self YaccSpecToDeclTypes) Run(input interface{}) (interface{}, error) {
	var members objects.IDclArray = nil
	yaccNode, ok := input.(yaccidl.IYaccNode)
	if !ok {
		return nil, self.logger.Error(fmt.Errorf("input param is not of type IYaccnode"))
	} else {
		var node yaccidl.IYaccNode
		for node = yaccNode; node != nil; node = node.GetNextNode() {
			v, err := self.YaccSpecToDeclAll.Process(node)
			if err != nil {
				return nil, self.logger.Error(err)
			}
			switch ans := v.(type) {
			case objects.IDclArray:
				members = append(members, ans...)
			default:
				if dcl, ok := v.(objects.IDcl); ok {
					members = append(members, dcl)
				} else {
					return nil, self.logger.Error(fmt.Errorf("type not recognized"))
				}
			}
		}
	}
	sort.Sort(&objects.IDclArraySorter{Data: members})
	result := objects.NewInputStreamDcl(
		"",
		self.uniqueSessionNumber.Next(),
		nil, members)

	return result, nil
}

func NewYaccSpecToDeclTypes(
	log *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	uniqueSessionNumber *implementations.UniqueSessionNumber,
	YaccSpecToDeclAll yaccSpecProcessing.IYaccSpecToDeclAll) *YaccSpecToDeclTypes {
	return &YaccSpecToDeclTypes{
		logger:              log.Create("YaccSpecToDeclTypes"),
		ctx:                 ctx,
		uniqueSessionNumber: uniqueSessionNumber,
		YaccSpecToDeclAll:   YaccSpecToDeclAll,
	}
}
