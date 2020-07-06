package gowriter

import (
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"io"
)

type Factory struct {
	logFactory          *log2.LogFactory
	WriteStructDcl      *WriteStructDcl
	UnionWriter         *WriteUnionDcl
	InterfaceWriter     *WriteInterface
	TypeDefDclWriter    *WriteTypeDefDcl
	ConstantValueWriter *WriteConstantValue
	EnumDclWriter       *WriteEnumDcl
	ScopeDclWriter      *WriteScopeDcl
	NativeDclWriter     *WriteNativeDcl
	WritePrimitiveDcl   *WritePrimitiveDcl
	fileResolve         *ctx2.ResolveFileName
	WriteSeqStructDcl   *WriteStructSequenceDcl
	StructSkeletonWriter *WriteStructSkeletonDcl
	StructStubWriter     *WriteStructStubDcl
}

func NewGoWriterFactory(
	logFactory *log2.LogFactory,
	UnionWriter *WriteUnionDcl,
	WriteStructDcl *WriteStructDcl,
	InterfaceWriter *WriteInterface,
	TypeDefDclWriter *WriteTypeDefDcl,
	ConstantValueWriter *WriteConstantValue,
	EnumDclWriter *WriteEnumDcl,
	ScopeDclWriter *WriteScopeDcl,
	NativeDclWriter *WriteNativeDcl,
	WritePrimitiveDcl *WritePrimitiveDcl,
	WriteSeqStructDcl *WriteStructSequenceDcl,
	fileResolve *ctx2.ResolveFileName,
	StructSkeletonWriter *WriteStructSkeletonDcl,
	StructStubWriter     *WriteStructStubDcl) *Factory {
	return &Factory{
		logFactory:           logFactory,
		WriteStructDcl:       WriteStructDcl,
		UnionWriter:          UnionWriter,
		InterfaceWriter:      InterfaceWriter,
		TypeDefDclWriter:     TypeDefDclWriter,
		ConstantValueWriter:  ConstantValueWriter,
		EnumDclWriter:        EnumDclWriter,
		ScopeDclWriter:       ScopeDclWriter,
		NativeDclWriter:      NativeDclWriter,
		WritePrimitiveDcl:    WritePrimitiveDcl,
		fileResolve:          fileResolve,
		WriteSeqStructDcl:    WriteSeqStructDcl,
		StructSkeletonWriter: StructSkeletonWriter,
		StructStubWriter:     StructStubWriter,
	}
}

func (self Factory) Create(
	appCtx *ctx2.GoYaccAppCtx,
	knownTypes map[objects.ScopeIdentifier]objects.IDcl,
	writer io.Writer,
	targetFile string) FileGenerationContext {
	return NewFileGenerationContext(
		targetFile,
		appCtx,
		knownTypes,
		self.logFactory,
		writer,
		self.WriteStructDcl,
		self.UnionWriter,
		self.InterfaceWriter,
		self.TypeDefDclWriter,
		self.ConstantValueWriter,
		self.EnumDclWriter,
		self.ScopeDclWriter,
		self.NativeDclWriter,
		self.WritePrimitiveDcl,
		self.WriteSeqStructDcl,
		self.fileResolve,
		self.StructSkeletonWriter,
		self.StructStubWriter)
}
