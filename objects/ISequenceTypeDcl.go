package objects

type ISequenceTypeDcl interface {
	IDcl
	IsISequenceTypeDcl() bool
	GetSequenceType() IDcl
	SetSequenceType(v IDcl) error
}
