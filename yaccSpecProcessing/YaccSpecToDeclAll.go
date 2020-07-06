package yaccSpecProcessing

import (
	"context"
	"fmt"
	"github.com/bhbosman/gocommon/constants"
	"go.uber.org/multierr"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"encoding/json"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
)

type YaccSpecToDeclAll struct {
	logger                     *log2.SubSystemLogger
	ctx                        *ctx2.GoYaccAppCtx
	uniqueSessionNumber        *implementations.UniqueSessionNumber
	YaccSpecToDeclForInterface *YaccSpecToDeclForInterface
}

func (self *YaccSpecToDeclAll) IsIYaccSpecToDeclAll() bool {
	return true
}

func (self *YaccSpecToDeclAll) Process(node yaccidl.IYaccNode) (interface{}, error) {
	if dcl, ok := node.(yaccidl.IImportDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("ImportDcl Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processImportDcl(dcl)
	} else if dcl, ok := node.(yaccidl.ITypePrefix); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("TypePrefix Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processTypePrefix(dcl)
	} else if dcl, ok := node.(yaccidl.IUnionDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("UnionDcl Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processUnion(dcl)
	} else if dcl, ok := node.(yaccidl.IModuleDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Module Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processModule(dcl)
	} else if dcl, ok := node.(yaccidl.IInterfaceDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Interface Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processInterfaceDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IValueDefDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("ValueDef Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processValueDef(dcl)
	} else if dcl, ok := node.(yaccidl.ITypeDeclaratorDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("TypeDeclarator Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processTypeDeclarator(dcl)
	} else if dcl, ok := node.(yaccidl.IPrimitiveTypeDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("GetPrimitiveType Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processPrimitiveType(dcl)
	} else if dcl, ok := node.(yaccidl.IConstDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Const Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processConstDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IEnumDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Enum Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processEnumDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IStructDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Struct Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processStructDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IExceptionDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Exception Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processExceptionDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IAttributeDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("AttributeDcl Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processAttributeDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IOperationDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("OperationDcl Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processOperationDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IStructMemberDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("StructMember Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processStructMemberDcl(dcl)
	} else if dcl, ok := node.(yaccidl.IScopeNameDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Scope Decl found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processScopeDcl(dcl)
	} else if dcl, ok := node.(yaccidl.ISequenceTypeDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Scope SequenceType found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processSequenceTypeDcl(dcl)
	} else if dcl, ok := node.(yaccidl.INativeDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Scope SequenceType found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processNativeDcl(dcl)
	} else if dcl, ok := node.(yaccidl.INativeDcl); ok {
		self.logger.LogWithLevel(10, func(logger *log.Logger) {
			logger.Printf("Scope SequenceType found. MemberName: %v, lexemData: %v", dcl.Identifier(), dcl.LexemData())
		})
		return self.processNativeDcl(dcl)
	} else {
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(node).String(), node.Identifier(), node.LexemData())
		self.logger.LogWithLevel(0, func(logger *log.Logger) {
			logger.Print(s)
		})
		return nil, self.logger.Error(fmt.Errorf(s))
	}
}

func (self *YaccSpecToDeclAll) processModule(dcl yaccidl.IModuleDcl) (interface{}, error) {
	var node yaccidl.IYaccNode

	var members objects.IDclArray = nil

	for node = dcl.ChildDecls(); node != nil; node = node.GetNextNode() {
		member, err := self.Process(node)
		if err != nil {
			return nil, self.logger.Error(err)
		}
		switch ans := member.(type) {
		case objects.IDclArray:
			members = append(members, ans...)
		default:
			if memberDcl, ok := member.(objects.IDcl); ok {
				members = append(members, memberDcl)
			}
		}
	}
	moduleDcl := objects.NewModuleDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		members)

	return moduleDcl, nil
}

func (self *YaccSpecToDeclAll) processInterfaceDcl(dcl yaccidl.IInterfaceDcl) (*objects.InterfaceDclContainer, error) {
	var node yaccidl.IYaccNode
	var inheritedInterfaces objects.IDclArray
	for node = dcl.InterfaceHeader().Inheritance(); node != nil; node = node.GetNextNode() {
		inheritedInterfaces = append(
			inheritedInterfaces,
			objects.NewScopeDcl2(
				node,
				self.uniqueSessionNumber.Next(),
				node.LexemData()))
	}
	var (
		members    objects.IDclArray       = nil
		operations []*objects.OperationDcl = nil
		attributes []*objects.AttributeDcl = nil
	)
	for node = dcl.Body(); node != nil; node = node.GetNextNode() {
		child, err := self.Process(node)
		if err != nil {
			return nil, self.logger.Error(err)
		}
		switch ans := child.(type) {
		case objects.IDclArray:
			members = append(members, ans...)
		case *objects.OperationDcl:
			operations = append(operations, ans)
		case *objects.AttributeDcl:
			attributes = append(attributes, ans)
		case *objects.ConstantValue:
			self.logger.LogWithLevel(0, func(logger *log.Logger) {
				logger.Println("Implement constant on Interface")
			})
		default:
			if childDcl, ok := child.(objects.IDcl); ok {
				members = append(members, childDcl)
			}
		}
	}
	return objects.NewInterfaceDclContainer(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		inheritedInterfaces,
		dcl.InterfaceHeader().InterfaceKind().Local(),
		dcl.InterfaceHeader().InterfaceKind().Abstract(),
		dcl.InterfaceHeader().InterfaceKind().Custom(),
		dcl.Forward(),
		false,
		members,
		operations,
		attributes), nil
}

func (self *YaccSpecToDeclAll) processValueDef(dcl yaccidl.IValueDefDcl) (interface{}, error) {
	var node yaccidl.IYaccNode
	var inheritedInterfaces objects.IDclArray
	if dcl.ValueHeader().ValueInheritanceSpec() != nil {
		for node = dcl.ValueHeader().ValueInheritanceSpec().ValueName(); node != nil; node = node.GetNextNode() {
			inheritedInterfaces = append(
				inheritedInterfaces,
				objects.NewScopeDcl2(
					node,
					self.uniqueSessionNumber.Next(),
					node.LexemData()))
		}

		for node = dcl.ValueHeader().ValueInheritanceSpec().SupportedInterfaceName(); node != nil; node = node.GetNextNode() {
			inheritedInterfaces = append(
				inheritedInterfaces,
				objects.NewScopeDcl2(
					node,
					self.uniqueSessionNumber.Next(),
					node.LexemData()))
		}
	}

	var (
		members    objects.IDclArray       = nil
		operations []*objects.OperationDcl = nil
		attributes []*objects.AttributeDcl = nil
	)

	for node = dcl.Body(); node != nil; node = node.GetNextNode() {
		child, err := self.Process(node)
		if err != nil {
			return nil, self.logger.Error(err)
		}
		switch ans := child.(type) {
		case *objects.OperationDcl:
			operations = append(operations, ans)
		case objects.IDclArray:
			members = append(members, ans...)
		case *objects.AttributeDcl:
			attributes = append(attributes, ans)
		default:
			if childDcl, ok := child.(objects.IDcl); ok {
				members = append(members, childDcl)
			}
		}
	}

	result := objects.NewInterfaceDclContainer(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		inheritedInterfaces,
		dcl.ValueHeader().ValueKind().Local(),
		dcl.ValueHeader().ValueKind().Abstract(),
		dcl.ValueHeader().ValueKind().Custom(),
		dcl.Forward(),
		true,
		members,
		operations,
		attributes)

	return result, nil
}

func (self *YaccSpecToDeclAll) processTypeDeclarator(dcl yaccidl.ITypeDeclaratorDcl) (interface{}, error) {
	resultType, err := self.Process(dcl.TypeSpec())
	if err != nil {
		return nil, self.logger.Error(err)
	}
	if resultTypeDcl, ok := resultType.(objects.IDcl); ok {
		var result objects.IDclArray = nil
		var node yaccidl.IYaccNode
		for node = dcl.Declarator(); node != nil; node = node.GetNextNode() {
			switch v := resultType.(type) {
			case *objects.StructDcl:
				typeDecl := objects.NewStructDcl(
					objects.ScopeIdentifier(node.Identifier()),
					node.IsArray(),
					node.Array(),
					self.uniqueSessionNumber.Next(),
					node.LexemData(),
					v,
					false,
					false,
					objects.StructTypeAsDefined,
					nil,
					nil)
				result = append(result, typeDecl)

			default:
				typeDecl := objects.NewTypeDefDcl(
					objects.ScopeIdentifier(node.Identifier()),
					node.IsArray(),
					node.Array(),
					self.uniqueSessionNumber.Next(),
					node.LexemData(),
					resultTypeDcl)
				result = append(result, typeDecl)

			}
		}
		return result, nil
	}
	return nil, self.logger.Error(fmt.Errorf("TypeDeclarator return type not IDeclaredType"))
}

func (self *YaccSpecToDeclAll) processPrimitiveType(dcl yaccidl.IPrimitiveTypeDcl) (interface{}, error) {
	return objects.NewPrimitiveDcl(dcl, false), nil
}

func (self *YaccSpecToDeclAll) processConstDcl(dcl yaccidl.IConstDcl) (interface{}, error) {
	constType, err := self.Process(dcl.TypeDef())
	if err != nil {
		return nil, self.logger.Error(err)
	}

	if constTypeDcl, ok := constType.(objects.IDcl); ok {
		constantValue := objects.NewConstantValue(
			objects.ScopeIdentifier(dcl.Identifier()),
			self.uniqueSessionNumber.Next(),
			dcl.LexemData(), constTypeDcl, dcl.Value())
		return constantValue, nil
	}
	return nil, self.logger.Error(fmt.Errorf("return type not IDeclaredType"))
}

func (self *YaccSpecToDeclAll) processEnumDcl(dcl yaccidl.IEnumDcl) (*objects.EnumDcl, error) {
	var errResult error = nil
	result := objects.NewEnumDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData())
	for node := dcl.Enumerator(); node != nil; node = node.GetNextNode() {
		result.AddMember(node.Identifier())
	}
	return result, errResult
}

func (self *YaccSpecToDeclAll) processUnion(dcl yaccidl.IUnionDcl) (interface{}, error) {
	var err error
	var switchType interface{}
	var unionBodyArray []*objects.UnionBody
	var switchTypeDcl objects.IDcl
	var ok bool
	for node := dcl.UnionBody(); node != nil; node, err = node.GetNextUnionBody() {
		if err != nil {
			return nil, self.logger.Error(err)
		}
		defaultCase := false
		var caseValues []*objects.UnionCaseValue = nil
		for cl := node.CaseLabels(); cl != nil; cl, err = cl.GetNextConstValue() {
			if err != nil {
				return nil, self.logger.Error(err)
			}
			defaultCase = false
			if _, ok := cl.(yaccidl.IDefaultValue); ok {
				defaultCase = true
				caseValues = nil
				break
			}
			if constValue, ok := cl.(yaccidl.IConstValue); ok {
				switchType, err = self.Process(dcl.SwitchType())
				switchTypeDcl, ok := switchType.(objects.IDcl)
				if !ok {
					return nil, self.logger.Error(fmt.Errorf("no swith type"))
				}

				switch v := switchType.(type) {
				case *objects.ScopeDcl:
					caseValue := objects.NewUnionCaseValue(
						objects.NewScopeDcl(
							objects.ScopeIdentifier(constValue.Value().String()),
							false,
							nil,
							-1,
							v.GetLexemData().GetSourceLexemData()))
					caseValues = append(caseValues, caseValue)
				case *objects.PrimitiveDcl:
					caseValue := objects.NewUnionCaseValue(
						objects.NewConstantValue(
							objects.ScopeIdentifier(constValue.Value().String()),
							-1,
							v.GetLexemData().GetSourceLexemData(),
							switchTypeDcl,
							constValue.Value()))
					caseValues = append(caseValues, caseValue)
				default:
					return nil, fmt.Errorf("implement this now")
				}
			}
		}
		v, err := self.Process(node.ElementSpec().TypeSpec())
		if err != nil {
			return nil, self.logger.Error(err)
		}
		if resultTypeDcl, ok := v.(objects.IDcl); ok {
			unionBody := objects.NewUnionBody(
				defaultCase,
				caseValues,
				node.ElementSpec().Declarator().Identifier(),
				node.ElementSpec().Declarator().IsArray(),
				node.ElementSpec().Declarator().Array(),
				resultTypeDcl)
			unionBodyArray = append(unionBodyArray, unionBody)
		}
	}

	switchType, err = self.Process(dcl.SwitchType())
	if err != nil {
		return nil, self.logger.Error(err)
	}
	if switchTypeDcl, ok = switchType.(objects.IDcl); ok {
		return objects.NewUnionDcl(
			objects.ScopeIdentifier(dcl.Identifier()),
			dcl.IsArray(),
			dcl.Array(),
			self.uniqueSessionNumber.Next(),
			dcl.LexemData(),
			switchTypeDcl,
			unionBodyArray), nil
	}
	return nil, self.logger.Error(fmt.Errorf("no swith type"))
}

func (self *YaccSpecToDeclAll) processStructDcl(dcl yaccidl.IStructDcl) (*objects.StructDcl, error) {
	var node yaccidl.IYaccNode
	var members []*objects.StructMember = nil
	for node = dcl.GetMember(); node != nil; node = node.GetNextNode() {
		value, err := self.Process(node)
		if err != nil {
			return nil, self.logger.Error(err)
		}
		if ar, ok := value.([]*objects.StructMember); ok {
			for _, member := range ar {
				members = append(members, member)
			}
		}
	}

	var inheritedScope objects.IDcl = nil
	if dcl.InheritsFrom() != nil {
		inherited := dcl.InheritsFrom()
		inheritedScope = objects.NewScopeDcl2(
			inherited,
			self.uniqueSessionNumber.Next(),
			inherited.LexemData())
	}

	return objects.NewStructDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(), dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		inheritedScope,
		false,
		dcl.Forward(),
		objects.StructTypeAsDefined,
		members,
		nil), nil
}

func (self *YaccSpecToDeclAll) processExceptionDcl(dcl yaccidl.IExceptionDcl) (*objects.StructDcl, error) {
	var node yaccidl.IYaccNode
	var members []*objects.StructMember = nil
	for node = dcl.GetMember(); node != nil; node = node.GetNextNode() {
		value, err := self.Process(node)
		if err != nil {
			return nil, self.logger.Error(err)
		}
		if ar, ok := value.([]*objects.StructMember); ok {
			for _, member := range ar {
				members = append(members, member)
			}
		}
	}
	return objects.NewStructDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		false,
		nil,
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		nil,
		true,
		false,
		objects.StructTypeAsDefined,
		members,
		nil), nil
}

func (self *YaccSpecToDeclAll) processAttributeDcl(dcl yaccidl.IAttributeDcl) (*objects.AttributeDcl, error) {

	attrType, err := self.Process(dcl.TypeSpec())
	if err != nil {
		return nil, self.logger.Error(err)
	}
	if attrTypeDcl, ok := attrType.(objects.IDcl); ok {
		var result *objects.AttributeDcl = nil
		result = objects.NewAttribute(dcl.Identifier(), attrTypeDcl, !dcl.Readonly(), true)
		return result, nil
	}
	return nil, self.logger.Error(fmt.Errorf("attribute return type not IOperationalParameter"))
}

func (self *YaccSpecToDeclAll) processOperationDcl(dcl yaccidl.IOperationDcl) (*objects.OperationDcl, error) {
	var node yaccidl.IYaccNode
	var errResult error = nil
	var result *objects.OperationDcl = nil

	returnType, err := self.Process(dcl.ReturnType())
	errResult = multierr.Append(errResult, err)
	if returnTypeDcl, ok := returnType.(objects.IDcl); ok {
		var operationsParams []*objects.OperationsParam
		for node = dcl.ParamDcl(); node != nil; node = node.GetNextNode() {
			if member, ok := node.(yaccidl.IOperationalParameter); ok {
				paramReturnType, err := self.Process(member.ParamType())
				errResult = multierr.Append(errResult, err)
				if paramReturnTypeDcl, ok := paramReturnType.(objects.IDcl); ok {
					operationsParam := objects.NewOperationsParam(member.Identifier(), paramReturnTypeDcl, member.Direction())
					operationsParams = append(operationsParams, operationsParam)
				} else {
					errResult = multierr.Append(errResult, fmt.Errorf("parameter return type not IIDCL"))
				}
			} else {
				errResult = multierr.Append(errResult, fmt.Errorf("operation return type not IOperationalParameter"))
			}
		}
		var exceptions objects.IDclArray
		for node = dcl.ExceptionDcl(); node != nil; node = node.GetNextNode() {
			exceptions = append(
				exceptions,
				objects.NewScopeDcl2(
					node,
					self.uniqueSessionNumber.Next(),
					node.LexemData()))
		}
		result = objects.NewOperation(dcl.Identifier(), returnTypeDcl, operationsParams, exceptions)
	} else {
		errResult = multierr.Append(errResult, fmt.Errorf("return type not IDeclaredType"))
	}
	return result, errResult
}

func (self *YaccSpecToDeclAll) processScopeDcl(dcl yaccidl.IScopeNameDcl) (*objects.ScopeDcl, error) {
	return objects.NewScopeDcl2(
		dcl,
		self.uniqueSessionNumber.Next(),
		dcl.LexemData()), nil
}

func (self *YaccSpecToDeclAll) processSequenceTypeDcl(dcl yaccidl.ISequenceTypeDcl) (objects.ISequenceTypeDcl, error) {

	resultType, err := self.Process(dcl.TypeSpec())
	if err != nil {
		return nil, self.logger.Error(err)
	}
	if resultTypeDcl, ok := resultType.(objects.IDcl); ok {

		result := objects.NewSequenceTypeDcl(
			objects.ScopeIdentifier(dcl.Identifier()),
			dcl.IsArray(),
			dcl.Array(),
			self.uniqueSessionNumber.Next(),
			dcl.LexemData(),
			resultTypeDcl)
		return result, nil
	}
	return nil, self.logger.Error(fmt.Errorf("requence return type not IDeclaredType"))
}

func (self *YaccSpecToDeclAll) processStructMemberDcl(dcl yaccidl.IStructMemberDcl) ([]*objects.StructMember, error) {

	v, err := self.Process(dcl.TypeDcl())
	if err != nil {
		return nil, self.logger.Error(err)
	}

	if typeDcl, ok := v.(objects.IDcl); ok {
		var result []*objects.StructMember
		for node := dcl.Declarator(); node != nil; node = node.GetNextNode() {
			structMember := objects.NewStructMember(
				node.Identifier(),
				node.IsArray(),
				node.Array(),
				typeDcl)
			result = append(result, structMember)
		}
		return result, nil
	}
	return nil, self.logger.Error(fmt.Errorf("requence return type not objects.IDeclaredType"))
}

func (self *YaccSpecToDeclAll) processNativeDcl(dcl yaccidl.INativeDcl) (interface{}, error) {
	return objects.NewNativeDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData()), nil
}

func (self *YaccSpecToDeclAll) processTypePrefix(dcl yaccidl.ITypePrefix) (interface{}, error) {
	return objects.NewTypePrefixDcl(
		objects.ScopeIdentifier(dcl.Identifier()),
		dcl.IsArray(),
		dcl.Array(),
		self.uniqueSessionNumber.Next(),
		dcl.LexemData(),
		dcl.Value()), nil
}

func (self *YaccSpecToDeclAll) processImportDcl(dcl yaccidl.IImportDcl) (objects.IDclArray, error) {
	if scopeNameDcl, ok := dcl.ImportedScope().(yaccidl.IScopeNameDcl); ok {
		scopeId := scopeNameDcl.Identifier()
		if strings.Index(scopeId, "::") == 0 {
			// global
			importFileName := filepath.Join(self.ctx.AppFolder, fmt.Sprintf("%v_export.json", scopeId[2:]))
			fileInfo, err := os.Stat(importFileName)
			if err != nil {
				return nil, err
			}
			if fileInfo.IsDir() {
				return nil, self.logger.Error(fmt.Errorf("import trying to open folder"))
			}
			importFileNameHandle, err := os.Open(importFileName)
			if err != nil {
				return nil, self.logger.Error(err)
			}
			defer func() {
				_ = importFileNameHandle.Close()
			}()
			value := objects.NewScopeDeclaration(nil)
			decoder := json.NewDecoder(importFileNameHandle)
			err = decoder.Decode(value)
			if err != nil {
				return nil, self.logger.Error(err)
			}
			var result objects.IDclArray

			for _, v := range value.DeclaredItems {
				baseDcl := objects.NewImportedDcl(v.Dcl)
				result = append(result, baseDcl)
			}
			for _, v := range value.ImportedItems {
				baseDcl := objects.NewImportedDcl(v.Dcl)
				result = append(result, baseDcl)
			}

			return result, nil
		}
	}

	return nil, nil
}

func (self *YaccSpecToDeclAll) OnStart(ctx context.Context) error {
	b := true
	b = b && self.YaccSpecToDeclForInterface != nil
	if !b {
		return constants.InvalidParam
	}
	return nil
}

func (self *YaccSpecToDeclAll) OnStop(ctx context.Context) error {
	self.YaccSpecToDeclForInterface = nil
	return nil
}

func (self *YaccSpecToDeclAll) init(YaccSpecToDeclForInterface *YaccSpecToDeclForInterface) error {
	b := true
	b = b && YaccSpecToDeclForInterface != nil
	if !b {
		return constants.InvalidParam
	}

	self.YaccSpecToDeclForInterface = YaccSpecToDeclForInterface
	return nil
}

func NewYaccSpecToDeclAll(
	logger *log2.SubSystemLogger,
	ctx *ctx2.GoYaccAppCtx,
	uniqueSessionNumber *implementations.UniqueSessionNumber) *YaccSpecToDeclAll {
	return &YaccSpecToDeclAll{
		logger:              logger,
		ctx:                 ctx,
		uniqueSessionNumber: uniqueSessionNumber,
	}
}
