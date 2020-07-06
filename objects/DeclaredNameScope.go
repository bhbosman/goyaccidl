package objects

import (
	"context"
	"fmt"
	rxgo "github.com/ReactiveX/RxGo"
	"github.com/bhbosman/gocommon/log"
	"strings"
)

type KnownTypes map[ScopeIdentifier]IDcl

func (t KnownTypes) ToItemsObs() rxgo.Observable {
	return rxgo.Defer(
		[]rxgo.Producer{
			func(ctx context.Context, next chan<- rxgo.Item) {
				for _, v := range t {
					next <- rxgo.Of(v)
				}
			},
		})
}

type KnownTypeNotFoundError struct {
	Missing ScopeIdentifier
}

func (k KnownTypeNotFoundError) Error() string {
	return fmt.Sprintf("%v not found", k.Missing)
}

func NewKnownTypeNotFoundError(missing ScopeIdentifier) *KnownTypeNotFoundError {
	return &KnownTypeNotFoundError{Missing: missing}
}

func (t KnownTypes) Find(name ScopeIdentifier) (IDcl, error) {
	if dcl, b := t[name]; b {
		return dcl, nil
	}
	return nil, NewKnownTypeNotFoundError(name)
}

type KeyMapData struct {
	IdlReference ScopeIdentifier
	LexemData    IDstSrcInformation
}

func NewKeyMapData(idlReference ScopeIdentifier, lexemData IDstSrcInformation) *KeyMapData {
	return &KeyMapData{
		IdlReference: idlReference,
		LexemData:    lexemData,
	}
}

type LocalKeyMap map[ScopeIdentifier]*KeyMapData

type DeclaredNameScope struct {
	index         int
	Scope         ScopeIdentifier
	DeclaredTypes map[ScopeIdentifier]IDcl
	TypePrefixes  map[ScopeIdentifier]string
	Errors        []string
	log           log.SubSystemLogger
	next          *DeclaredNameScope
}

func (s *DeclaredNameScope) SetError(err error) {
	s.Errors = ErrorToString(err)
}
func (s *DeclaredNameScope) PossDuplicateAdd(dcl IDcl) error {
	key := dcl.GetName()
	if _, ok := s.DeclaredTypes[key]; !ok {
		s.DeclaredTypes[key] = dcl
		return nil
	}
	return nil
}

func (s *DeclaredNameScope) Add(dcl IDcl) error {
	internalAdd := func(key ScopeIdentifier, dcl IDcl) {
		s.DeclaredTypes[key] = dcl
	}
	var errResult error = nil
	previousDecl, err := s.Find(dcl, false)
	if err != nil {
		if declaredNameScopeFindError, ok := err.(*DeclaredNameScopeFindError); ok {
			if declaredNameScopeFindError.Id != TypeNotFound {
				return err
			}
		} else {
			return err
		}
	}
	if previousDecl != nil {
		if previousDecl.GetForward() && dcl.GetForward() {
			//s.log.LogWithLevel(0, func(logger *log2.Logger) {
			//	logger.Printf("%v is a duplicate forward decl. Was already declared at %v", dcl, previousDecl.GetLexemData())
			//})
			return nil
		} else if previousDecl.GetForward() && !dcl.GetForward() {
			key := dcl.GetName()
			internalAdd(key, dcl)
			return nil
		} else if !previousDecl.GetForward() && dcl.GetForward() {
			//s.log.LogWithLevel(0, func(logger *log2.Logger) {
			//	logger.Printf("%v is a forward decl, after a declaration %v", dcl, previousDecl.GetLexemData())
			//})
			return errResult
		} else if !previousDecl.GetForward() && !dcl.GetForward() {
			return s.log.Error(fmt.Errorf("%v already added. Previous decl: %v", dcl.GetName(), previousDecl.GetLexemData()))
		}
	}
	key := dcl.GetName()
	internalAdd(key, dcl)

	return errResult
}

type DeclaredNameScopeFindErrorId int8

const (
	TypeNotFound DeclaredNameScopeFindErrorId = iota
)

type DeclaredNameScopeFindError struct {
	errorMessage string
	Id           DeclaredNameScopeFindErrorId
}

func NewDeclaredNameScopeFindError(errorMessage string, id DeclaredNameScopeFindErrorId) *DeclaredNameScopeFindError {
	return &DeclaredNameScopeFindError{errorMessage: errorMessage, Id: id}
}

func (self DeclaredNameScopeFindError) Error() string {
	return fmt.Sprintf("%v (%v)", self.errorMessage, self.Id)
}

func (s *DeclaredNameScope) Find(dcl IDcl, recursive bool) (IDcl, error) {
	returnType, ok := s.DeclaredTypes[dcl.GetName()]
	if ok {
		return returnType, nil
	}
	if recursive {
		if s.next != nil {
			return s.next.Find(dcl, true)
		}
	}

	return nil, NewDeclaredNameScopeFindError(fmt.Sprintf("type not defined: %v (%v)", dcl.GetName(), dcl.GetLexemData()), TypeNotFound)
}

func (s DeclaredNameScope) CreateDescendent(scope ScopeIdentifier) *DeclaredNameScope {
	var ss []string = nil
	if s.Scope != "" {
		ss = strings.Split(string(s.Scope), "::")
	}

	if scope != "" {
		ss = append(ss, string(scope))
	}
	scope = ScopeIdentifier(strings.Join(ss, "::"))
	result := NewDeclaredNameScope(s.log, scope)
	result.next = &s
	result.index = result.index + 1
	return result
}

func (s *DeclaredNameScope) GetAllTypes() []*KeyMapData {
	var result []*KeyMapData = nil
	if s.next != nil {
		result = s.next.GetAllTypes()
	}
	for k, v := range s.DeclaredTypes {
		result = append(result, NewKeyMapData(k, v.GetLexemData()))
	}
	return result
}

func (s *DeclaredNameScope) AddTypePrefix(name ScopeIdentifier, value string) {
	s.TypePrefixes[name] = value
}

func (s *DeclaredNameScope) FindTypePrefixes(v ScopeIdentifier) (string, bool) {
	if v, ok := s.TypePrefixes[v]; ok {
		return v, true
	}
	if s.next != nil {
		return s.next.FindTypePrefixes(v)
	}
	return "", false
}

func NewDeclaredNameScope(log log.SubSystemLogger, scope ScopeIdentifier) *DeclaredNameScope {
	return &DeclaredNameScope{
		index:         0,
		Scope:         scope,
		DeclaredTypes: make(map[ScopeIdentifier]IDcl),
		TypePrefixes:  make(map[ScopeIdentifier]string),
		Errors:        nil,
		log:           log,
		next:          nil,
	}
}
