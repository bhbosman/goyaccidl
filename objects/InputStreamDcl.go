package objects

import "github.com/bhbosman/yaccidl"

type InputStreamDcl struct {
	Dcl
	Members IDclArray `json:"members"`
	Errors  []string  `json:"errors"`
}

func (d *InputStreamDcl) SetError(err error) {
	d.Errors = ErrorToString(err)
}

func NewInputStreamDcl(
	name ScopeIdentifier,
	orderId int,
	lexemData *yaccidl.LexemValue,
	Members IDclArray) *InputStreamDcl {
	return &InputStreamDcl{
		//todo: remove this dcl registration as it is not needed
		Dcl: *NewDcl(
			name,
			false,
			nil,
			orderId,
			lexemData,
			false,
			yaccidl.IdlInvalid),
		Members: Members,
	}
}
