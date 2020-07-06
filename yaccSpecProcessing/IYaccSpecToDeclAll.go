package yaccSpecProcessing

import "github.com/bhbosman/yaccidl"

type IYaccSpecToDeclAll interface {
	Process(node yaccidl.IYaccNode) (interface{}, error)
	IsIYaccSpecToDeclAll() bool
}
