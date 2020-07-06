package objects

type IUnionDcl interface {
	IBaseStructDcl
	IsIUnionDcl() bool
}
