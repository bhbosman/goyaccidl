package process

//type IDealWithResult interface {
//	Deal(data interface{}, err error)
//}

type IBaseProcess interface {
	Name() string
}
type IProcess interface {
	IBaseProcess
	Run(interface{}) (interface{}, error)
}

type FlattenData struct {
	Data interface{}
	Name string
}

type IProcessWrapperInput interface {
	Input() interface{}
	Result() interface{}
	SetResult(name string, result interface{})
	FLatten() []*FlattenData
}

type ProcessWrapperInput struct {
	input  interface{}
	name   string
	result interface{}
	prev   IProcessWrapperInput
}

func (self *ProcessWrapperInput) FLatten() []*FlattenData {
	var result []*FlattenData = nil
	if self.prev != nil {
		result = self.prev.FLatten()
	}
	result = append(result, &FlattenData{
		Data: self.result,
		Name: self.name,
	})
	return result

}

func (self *ProcessWrapperInput) SetResult(name string, result interface{}) {
	self.name = name
	self.result = result
}

func NewProcessWrapperInput(input interface{}, prev IProcessWrapperInput) *ProcessWrapperInput {
	return &ProcessWrapperInput{
		input: input,
		prev:  prev,
	}
}

func (self ProcessWrapperInput) Input() interface{} {
	return self.input
}

func (self ProcessWrapperInput) Result() interface{} {
	return self.result
}

type IProcessWrapper interface {
	IBaseProcess
	SetNext(process IProcessWrapper)
	Run(result IProcessWrapperInput) (interface{}, error)
}

type ISetError interface {
	SetError(err error)
}

type IFactory interface {
	Create() (IProcessWrapper, error)
}
