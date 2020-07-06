package yaccSpecProcessing

import (
	"context"
	"github.com/bhbosman/gocommon/Services/implementations"
	"github.com/bhbosman/gocommon/constants"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
)

type YaccSpecToDeclForInterface struct {
	logger              *log2.SubSystemLogger
	ctx                 *ctx2.GoYaccAppCtx
	uniqueSessionNumber *implementations.UniqueSessionNumber
	processAll          IYaccSpecToDeclAll
}

func (self *YaccSpecToDeclForInterface) OnStart(ctx context.Context) error {
	b := true
	b = b && self.processAll != nil
	if !b {
		return constants.InvalidParam
	}
	return nil

}

func (self *YaccSpecToDeclForInterface) OnStop(ctx context.Context) error {
	self.processAll = nil
	return nil
}

func (self *YaccSpecToDeclForInterface) init(processAll IYaccSpecToDeclAll) error {
	b := true
	b = b && processAll != nil
	if !b {
		return constants.InvalidParam
	}

	self.processAll = processAll
	return nil
}

func NewYaccSpecToDeclForInterface(
	logger *log2.SubSystemLogger,
	ctx *ctx2.GoYaccAppCtx,
	uniqueSessionNumber *implementations.UniqueSessionNumber) *YaccSpecToDeclForInterface {
	return &YaccSpecToDeclForInterface{
		logger:              logger,
		ctx:                 ctx,
		uniqueSessionNumber: uniqueSessionNumber,
	}
}
