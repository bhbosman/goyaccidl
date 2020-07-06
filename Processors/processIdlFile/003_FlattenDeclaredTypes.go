package processIdlFile

import (
	"fmt"
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/goyaccidl/typeHelpers"

	"log"
	"reflect"
	"sort"
	"strings"
)

type FlattenDeclaredTypes struct {
	log                 log2.SubSystemLogger
	ctx                 *ctx2.GoYaccAppCtx
	logFactory          *log2.LogFactory
	importedNameScope   *objects.DeclaredNameScope
	declaredNameScope   *objects.DeclaredNameScope
	uniqueSessionNumber *implementations.UniqueSessionNumber
	StructDclHelper     *typeHelpers.StructDcl
	InterfaceDclHelper  *typeHelpers.InterfaceDcl
	TypeDefDclHelper    *typeHelpers.TypeDefDcl
	ConstantDclHelper   *typeHelpers.ConstantDcl
	NativeDclHelper     *typeHelpers.NativeDcl
	ScopeDclHelper      *typeHelpers.ScopeDcl
	PrimitiveDclHelper  *typeHelpers.PrimitiveDcl
	SequenceDcl         *typeHelpers.SequenceDcl
	UnionDcl            *typeHelpers.UnionDcl
}

func (self *FlattenDeclaredTypes) Name() string {
	return self.log.Name()
}

func (self *FlattenDeclaredTypes) Run(input interface{}) (interface{}, error) {
	if inputStreamDcl, ok := input.(*objects.InputStreamDcl); ok {
		temp := self.declaredNameScope
		defer func() {
			self.declaredNameScope = temp
		}()
		self.declaredNameScope = self.declaredNameScope.CreateDescendent("")
		for _, dcl := range inputStreamDcl.Members {
			_, _, err := self.processGeneratedObjects(dcl, true)
			if err != nil {
				return nil, self.log.Error(err)
			}
		}
		declaredMembers := self.CreateMembers(self.declaredNameScope)
		inputStreamDcl := objects.NewInputStreamDcl(
			inputStreamDcl.Name,
			self.uniqueSessionNumber.Next(),
			inputStreamDcl.GetLexemData().GetSourceLexemData(),
			declaredMembers)

		importedMembers := self.CreateMembers(self.importedNameScope)
		importedDeclarationItems := self.CreateImportedScopeItems(importedMembers)

		//scopeDeclarationItems := self.CreateDeclaredScopeItems(declaredMembers, importedDeclarationItems)
		output := objects.NewFlattenDeclaredTypesOut(
			inputStreamDcl,
			//scopeDeclarationItems,
			importedDeclarationItems)
		return output, nil
	}
	return nil, self.log.Error(fmt.Errorf("input type wrong"))
}

func (self *FlattenDeclaredTypes) processGeneratedObjects(dcl objects.IDcl, add bool) (bool, objects.IDcl, error) {
	switch dclType := dcl.(type) {
	case *objects.ModuleDcl:
		temp := self.declaredNameScope
		defer func() {
			self.declaredNameScope = temp
		}()
		self.declaredNameScope = self.declaredNameScope.CreateDescendent(dclType.Name)
		err := self.processContainer(dclType)
		if err != nil {
			return false, nil, err
		}
		err = self.Flatten(dclType, self.declaredNameScope, temp, true)
		if err != nil {
			return false, nil, err
		}
		return false, nil, nil

	case *objects.InterfaceDclContainer:
		temp := self.declaredNameScope
		interfaceDcl := objects.NewInterfaceDcl(
			dclType.GetName(),
			dclType.GetIsArray(),
			dclType.GetArrayValue(),
			dclType.GetOrderId(),
			dclType.GetLexemData().GetSourceLexemData(),
			dclType.BaseInterfaces,
			dclType.Local,
			dclType.Abstract,
			dclType.Custom,
			dclType.Forward,
			dclType.ValueType,
			dclType.Operations,
			dclType.Attributes)
		defer func() {
			self.declaredNameScope = temp
			_ = self.declaredNameScope.Add(interfaceDcl)
		}()
		name := interfaceDcl.GetName()
		fmt.Sprintln(name)

		self.declaredNameScope = self.declaredNameScope.CreateDescendent(dclType.Name)
		interfaceDcl.SetTypePrefix(self.FindTypePrefix())
		err := self.processContainer(dclType)
		if err != nil {
			return false, nil, err
		}
		err = self.Flatten(interfaceDcl, self.declaredNameScope, temp, true)
		if err != nil {
			return false, nil, err
		}
		return false, nil, nil

	case *objects.StructDcl:
		temp := self.declaredNameScope
		defer func() {
			self.declaredNameScope = temp
			_ = self.declaredNameScope.Add(dclType)
		}()
		self.declaredNameScope = self.declaredNameScope.CreateDescendent(dclType.GetName())
		dclType.SetTypePrefix(self.FindTypePrefix())
		err := self.processContainer(dclType)
		if err != nil {
			return false, nil, err
		}
		err = self.Flatten(dclType, self.declaredNameScope, temp, false)
		if err != nil {
			return false, nil, err
		}
		return false, nil, nil

	case *objects.UnionDcl:
		temp := self.declaredNameScope
		defer func() {
			self.declaredNameScope = temp
			_ = self.declaredNameScope.Add(dclType)
		}()
		self.declaredNameScope = self.declaredNameScope.CreateDescendent(dclType.Name)
		dclType.SetTypePrefix(self.FindTypePrefix())
		err := self.processContainer(dclType)
		if err != nil {
			return false, nil, err
		}
		err = self.Flatten(dclType, self.declaredNameScope, temp, false)
		if err != nil {
			return false, nil, err
		}
		return false, nil, nil

	case *objects.TypePrefixDcl:
		self.declaredNameScope.AddTypePrefix(dclType.Name, dclType.Value)
		return false, nil, nil

	case *objects.StructMember:
		b, newDclType, err := self.processGeneratedObjects(dclType.MemberType, false)
		if b {
			dclType.MemberType = newDclType
		}
		return false, nil, err

	case *objects.TypeDefDcl:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil

	case *objects.SequenceTypeDcl:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		} else {
			// unnamed object
			//number := self.uniqueSessionNumber.Next()
			//name := objects.ScopeIdentifier(fmt.Sprintf("UnnamedSequenceOf%v%d", dclType.SequenceType().GetName(), number))
			//unnamed := objects.NewTypeDefDcl(
			//	name,
			//	self.uniqueSessionNumber.Next(),
			//	dclType.LexemData,
			//	dclType)
			//seq := objects.NewSequenceTypeDcl(
			//	name, self.uniqueSessionNumber.Next(),dclType.LexemData,dclType)
			//if seq != nil {
			//
			//}
			//return true, unnamed, self.declaredNameScope.Add(unnamed)
			return false, nil, fmt.Errorf("no unnamed sequence types. deprecated. %v", dclType.GetLexemData())
		}

	case *objects.ConstantValue:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil

	case *objects.PrimitiveDcl:
		if add {
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil

	case *objects.EnumDcl:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil

	case *objects.ScopeDcl:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil

	case *objects.NativeDcl:
		if add {
			dclType.SetTypePrefix(self.FindTypePrefix())
			return false, nil, self.declaredNameScope.Add(dclType)
		}
		return false, nil, nil
	case *objects.ImportedDcl:
		return false, nil, self.importedNameScope.PossDuplicateAdd(dclType)

	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(dcl).String(), dcl.GetName(), dcl.GetLexemData())
		self.log.LogWithLevel(0, func(logger *log.Logger) {
			logger.Print(s)
		})
		return false, nil, fmt.Errorf(s)
	}
}

func (self FlattenDeclaredTypes) processContainer(dcl objects.IDcl) error {
	if container, ok := dcl.(objects.IContainer); ok {
		defer func() {
			container.Clear()
		}()
		members := container.GetList()
		for _, member := range members {
			_, _, err := self.processGeneratedObjects(member, true)
			if err != nil {
				return self.log.Error(err)
			}
		}
	}
	return nil
}

func (self *FlattenDeclaredTypes) Flatten(incomingDcl objects.IDcl, current, prev *objects.DeclaredNameScope, scoped bool) error {
	localKeyMap := make(objects.LocalKeyMap)
	dclArray := objects.IDclArray{incomingDcl}

	for k, v := range current.DeclaredTypes {
		var s objects.ScopeIdentifier
		if scoped {
			s = objects.ScopeIdentifier(fmt.Sprintf("%v::%v", incomingDcl.GetName(), k))
		} else {
			s = objects.ScopeIdentifier(fmt.Sprintf("%v", k))
		}

		localKeyMap[k] = objects.NewKeyMapData(s, v.GetLexemData())
		err := v.UpdateIdlReference(s)
		if err != nil {
			return err
		}
		_ = prev.Add(v)
		dclArray = append(dclArray, v)
	}
	allPrevTypes := prev.GetAllTypes()
	for _, s := range allPrevTypes {
		localKeyMap[s.IdlReference] = s
	}

	scopes := make(typeHelpers.CurrentScope)
	scopeInPieces := strings.Split(string(current.Scope), "::")
	for i := 1; i <= len(scopeInPieces); i++ {
		scopeInPiece := objects.ScopeIdentifier(strings.Join(scopeInPieces[0:i], "::"))
		scopes[scopeInPiece] = scopeInPiece
	}

	for _, dcl := range dclArray {
		err := self.flattenGeneratedObjects(scopes, localKeyMap, dcl)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *FlattenDeclaredTypes) flattenGeneratedObjects(currentScope typeHelpers.CurrentScope, keyMap objects.LocalKeyMap, dcl objects.IDcl) error {
	switch dclType := dcl.(type) {
	case *objects.InterfaceDcl:
		return self.InterfaceDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.TypeDefDcl:
		return self.TypeDefDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.UnionDcl:
		return self.UnionDcl.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.StructDcl:
		return self.StructDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.ConstantValue:
		return self.ConstantDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.NativeDcl:
		return self.NativeDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.EnumDcl:
		return nil
	case *objects.ModuleDcl:
		if len(dclType.Members) > 0 {
			return self.log.Error(fmt.Errorf("error: At this point module members must have been flattened"))
		}
		return nil
	case *objects.SequenceTypeDcl:
		return self.SequenceDcl.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.ScopeDcl:
		return self.ScopeDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	case *objects.PrimitiveDcl:
		return self.PrimitiveDclHelper.UpdateIdlReference(currentScope, keyMap, dclType)
	default:
		s := fmt.Sprintf("No handler found for %v(%v). LexemData: %v", reflect.TypeOf(dcl).String(), dcl.GetName(), dcl.GetLexemData())
		self.log.LogWithLevel(0, func(logger *log.Logger) { logger.Print(s) })
		return fmt.Errorf(s)
	}
}

func (self *FlattenDeclaredTypes) FindTypePrefix() string {
	if self.declaredNameScope.Scope != "" {
		ss := strings.Split(string(self.declaredNameScope.Scope), "::")
		var v string
		var ok bool

		for i := len(ss) - 1; i >= 0; i-- {
			v, ok = self.declaredNameScope.FindTypePrefixes(objects.ScopeIdentifier(ss[i]))
			if ok {
				return v
			}
		}

		//ss = strings.Split(string(self.declaredNameScope.Scope), "::")
		//if len(ss) >= 2 {
		//	ss = ss[:len(ss)-1]
		//	s := strings.Join(ss, "/")
		//	return s
		//}

	}
	return ""
}

func (self *FlattenDeclaredTypes) CreateScopeItem(member objects.IDcl) *objects.ScopeItem {
	name := member.GetName()

	return objects.NewScopeItem(
		member,
		member.GetPrimitiveType(),
		member.GetOrderId(),
		name,
		member.GetLexemData().GetSourceLexemData())
}

//func (self *FlattenDeclaredTypes) CreateDeclaredScopeItems(members objects.IBaseDclArray, items []*objects.ScopeItem) objects.ScopeDeclarationItems {
//	result := objects.NewScopeDeclarationItems()
//	for _, member := range members {
//		var scopeDeclaration *objects.ScopeDeclaration = nil
//		var ok bool
//		name := member.GetName()
//		firstPart := name.Scope().First()
//		if firstPart == "" {
//			continue
//		}
//		if scopeDeclaration, ok = result[firstPart]; !ok {
//			scopeDeclaration = objects.NewScopeDeclaration(items)
//			result[firstPart] = scopeDeclaration
//		}
//		scopeDeclaration.DeclaredItems[name] = self.CreateScopeItem(member)
//	}
//	return result
//}

func (self *FlattenDeclaredTypes) CreateMembers(scope *objects.DeclaredNameScope) objects.IDclArray {
	var declaredMembers objects.IDclArray
	for _, v := range scope.DeclaredTypes {
		declaredMembers = append(declaredMembers, v)
	}
	sorter := &objects.IDclArraySorter{Data: declaredMembers}
	sort.Sort(sorter)
	return declaredMembers
}

func (self *FlattenDeclaredTypes) CreateImportedScopeItems(members objects.IDclArray) []*objects.ScopeItem {
	var result []*objects.ScopeItem
	for _, member := range members {
		name := member.GetName()
		firstPart := name.Scope().First()
		if firstPart == "" {
			continue
		}
		result = append(result, self.CreateScopeItem(member))
	}
	return result
}

func NewFlattenDeclaredTypes(
	logFactory *log2.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	uniqueSessionNumber *implementations.UniqueSessionNumber,
	flattenStructDcl *typeHelpers.StructDcl,
	flattenInterfaceDcl *typeHelpers.InterfaceDcl,
	flattenTypeDefDcl *typeHelpers.TypeDefDcl,
	flattenConstantDcl *typeHelpers.ConstantDcl,
	flattenNativeDcl *typeHelpers.NativeDcl,
	ScopeDclHelper *typeHelpers.ScopeDcl,
	PrimitiveDclHelper *typeHelpers.PrimitiveDcl,
	SequenceDcl *typeHelpers.SequenceDcl,
	UnionDcl *typeHelpers.UnionDcl,
	idlDefaultTypes *Service.IdlDefaultTypes) *FlattenDeclaredTypes {
	logger := logFactory.Create("FlattenDeclaredTypes")
	importedNameScope := objects.NewDeclaredNameScope(logger.Sub("DeclaredNameScope"), "")
	declaredNameScope := importedNameScope.CreateDescendent("")

	for _, primitive := range idlDefaultTypes.IdlPrimitivesInScopedFormat() {
		_ = declaredNameScope.Add(primitive)
	}

	result := &FlattenDeclaredTypes{
		log:                 *logger,
		ctx:                 ctx,
		logFactory:          logFactory,
		importedNameScope:   importedNameScope,
		declaredNameScope:   declaredNameScope,
		uniqueSessionNumber: uniqueSessionNumber,
		StructDclHelper:     flattenStructDcl,
		InterfaceDclHelper:  flattenInterfaceDcl,
		TypeDefDclHelper:    flattenTypeDefDcl,
		ConstantDclHelper:   flattenConstantDcl,
		NativeDclHelper:     flattenNativeDcl,
		ScopeDclHelper:      ScopeDclHelper,
		PrimitiveDclHelper:  PrimitiveDclHelper,
		SequenceDcl:         SequenceDcl,
		UnionDcl:            UnionDcl,
	}
	return result
}
