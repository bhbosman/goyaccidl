package objects

type IConstantValue interface {
	IDcl
	GetConstantType() IDcl
	IsConstantValue() bool
	GetValue() interface{}
}
