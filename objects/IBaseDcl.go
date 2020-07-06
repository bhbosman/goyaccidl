package objects

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/yaccidl"
	"reflect"
	"strings"
)

type IGetName interface {
	GetName() ScopeIdentifier
	GetIsArray() bool
	GetArrayValue() []int64
}

type IDcl interface {
	IGetName
	GetPrimitiveType() yaccidl.PrimitiveType
	GetLexemData() IDstSrcInformation
	GetOrderId() int
	GetForward() bool
	GetImported() bool
	UpdateIdlReference(s ScopeIdentifier) error
	BuildIdValue() (string, error)
	SetTypePrefix(s string)
	GetTypePrefix() string
	SetDestination(folderId, folder string) error
	DclResolveFolderUsage() (id string, folderName string)
}

type IDclStream struct {
	Assigned bool            `json:"assigned"`
	Name     string          `json:"name"`
	Data     json.RawMessage `json:"data"`
}

func NewIDclStream(dcl IDcl) (IDclStream, error) {
	if dcl == nil {
		return IDclStream{
			Assigned: false,
			Name:     "",
			Data:     nil,
		}, nil
	}
	memberType := reflect.TypeOf(dcl)
	s := memberType.String()
	key := strings.TrimFunc(
		s,
		func(r rune) bool {
			switch r {
			case '&', '*':
				return true
			default:
				return false
			}
		})

	codec, ok := IDclCodecMap[key]
	if !ok {
		return IDclStream{}, fmt.Errorf("no IDcl codec found for %v", key)
	}

	if codec.JsonEncoder == nil {
		return IDclStream{}, fmt.Errorf("no %v JsonEncoder for IDeclaredType is nil", key)
	}

	rawData, err := codec.JsonEncoder(dcl)
	if err != nil {
		return IDclStream{}, err
	}

	return IDclStream{
		Assigned: true,
		Name:     key,
		Data:     rawData,
	}, nil
}

func (self *IDclStream) GetDcl() (IDcl, error) {
	if !self.Assigned {
		return nil, nil
	}
	codec, ok := IDclCodecMap[self.Name]
	if !ok {
		return nil, fmt.Errorf("no IDeclaredType codec found for %v", self.Name)
	}
	if codec.JsonDecoder == nil {
		return nil, fmt.Errorf("no %v JsonDecoder for IDeclaredType is nil", self.Name)
	}
	result, err := codec.JsonDecoder(self.Data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//type IBaseDclArray IDclArray

type DclJsonCodec struct {
	JsonEncoder func(data IDcl) ([]byte, error)
	JsonDecoder func(b []byte) (IDcl, error)
}

var IDclCodecMap = make(map[string]DclJsonCodec)

func RegisterIDclCodec(objType reflect.Type, newCb func() IDcl) {
	key := strings.TrimFunc(
		objType.String(),
		func(r rune) bool {
			switch r {
			case '&', '*':
				return true
			default:
				return false
			}
		})
	IDclCodecMap[key] = DclJsonCodec{
		JsonEncoder: func(data IDcl) ([]byte, error) {
			return json.MarshalIndent(data, "", "\t")
		},
		JsonDecoder: func(b []byte) (IDcl, error) {
			v := newCb()
			err := json.Unmarshal(b, v)
			return v, err
		},
	}
}
