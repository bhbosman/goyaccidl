package objects

type IContainer interface {
	IGetName
	GetList() IDclArray
	Clear()
}
