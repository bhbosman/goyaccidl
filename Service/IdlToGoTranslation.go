package Service

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/objects"
	"github.com/bhbosman/yaccidl"
	"strings"
	"unicode"
)

type IdlToGoTranslation struct {
	DclHelpers          *DclHelpers
	LanguageTypeService *LanguageTypeService
	logger              *log2.SubSystemLogger
}

func (self IdlToGoTranslation) BuildTypeName(name objects.ScopeIdentifier) objects.ScopeIdentifier {
	ss := strings.Split(string(name), "::")
	for i := range ss {
		underScoreSplit := strings.Split(ss[i], "_")
		for i2 := range underScoreSplit {
			ddddd := underScoreSplit[i2]
			ddddd = self.RemoveAllCaps(ddddd)
			ddddd = self.CamelCase(ddddd)
			underScoreSplit[i2] = ddddd
		}
		ss[i] = strings.Join(underScoreSplit, "")
	}
	return objects.ScopeIdentifier(strings.Join(ss, ""))
}

type GoLangTypeReference string

func (self GoLangTypeReference) Scope() string {
	ss := strings.Split(string(self), ".")
	if len(ss) <= 1 {
		return ""
	}
	return strings.Join(ss[:len(ss)-1], ".")
}

func (self GoLangTypeReference) Name() string {
	ss := strings.Split(string(self), ".")
	if len(ss) <= 1 {
		return string(self)
	}
	return ss[len(ss)-1]
}

func (self GoLangTypeReference) NamePrePost(prefix string, postfix string) GoLangTypeReference {
	ss := strings.Split(string(self), ".")
	if len(ss) <= 1 {
		return GoLangTypeReference(fmt.Sprintf("%v%v%v", prefix, ss[0], postfix))
	}
	if len(ss) == 2 {
		return GoLangTypeReference(fmt.Sprintf("%v.%v%v%v", ss[0], prefix, ss[1], postfix))
	}
	return self
}

func (self IdlToGoTranslation) AbsTypeName(
	currentLexem objects.IDstSrcInformation,
	dcl objects.IDcl) GoLangTypeReference {
	return self.AbsTypeNameArrayName(
		currentLexem,
		dcl,
		nil,
		nil)
}

func (self IdlToGoTranslation) AbsTypeNameArrayName(
	currentLexem objects.IDstSrcInformation,
	dcl objects.IDcl,
	paramArrayValue []int64,
	typeArrayValue []int64) GoLangTypeReference {
	return self.AbsTypeNameArrayNameFull(
		currentLexem,
		dcl,
		paramArrayValue,
		typeArrayValue)
}

func (self IdlToGoTranslation) TypeNameFromBuildTypeName(
	currentLexem objects.IDstSrcInformation,
	lexemFromDcl objects.IDstSrcInformation,
	buildTypeName objects.ScopeIdentifier,
	paramArrayContext string,
	typeArrayContext string) GoLangTypeReference {
	//ns := lexemFromDcl.GetSourceFolderId()

	//b := currentLexem.GetDestinationFolderId() == lexemFromDcl.GetSourceFolderId() ||
	//	currentLexem.GetDestinationFolderId() == lexemFromDcl.GetDestinationFolderId()

	currentId, currentForlderName := currentLexem.DclResolveFolderUsage()
	lexemFromDclId, lexemFromDclForlderName := lexemFromDcl.DclResolveFolderUsage()
	if currentForlderName == lexemFromDclForlderName {

	}

	ns := ""
	if currentId != "" {
		if currentId != lexemFromDclId {
			ns = lexemFromDclId
		}
	}

	//if lexemFromDcl.GetSourceLexemData().GetCheckTarget() {
	//	var currentLexemTargetFolderName string
	//	dclLexemTargetFolderName := lexemFromDcl.GetTargetFolderName()
	//	b = b || currentLexemTargetFolderName == dclLexemTargetFolderName
	//}

	if ns != "" {
		return GoLangTypeReference(fmt.Sprintf("%v%v%v.%v", paramArrayContext, typeArrayContext, ns, buildTypeName))
	}
	return GoLangTypeReference(fmt.Sprintf("%v%v%v", paramArrayContext, typeArrayContext, buildTypeName))

}

func (self IdlToGoTranslation) AbsTypeNameArrayNameFull(
	currentLexem objects.IDstSrcInformation,
	dcl objects.IDcl,
	paramArray []int64,
	typeArray []int64) GoLangTypeReference {

	buildTypeName := self.BuildTypeName(dcl.GetName())
	paramArrayContext := self.BuildArrayContext(paramArray)
	typeArrayContext := self.BuildArrayContext(typeArray)
	lexemFromDcl := dcl.GetLexemData()
	return self.TypeNameFromBuildTypeName(
		currentLexem,
		lexemFromDcl,
		buildTypeName,
		paramArrayContext,
		typeArrayContext)
}

func (self IdlToGoTranslation) RemoveUnderScores(name string) string {
	runeArray := []rune(name)
	b := false
	bb := false
	readIndex := 0
	writeIndex := 0
	l := len(name)
	for readIndex < l {
		if runeArray[readIndex] != '_' {
			if b {
				bb = true
			} else {
				bb = false
			}
			b = false
		} else {
			b = true
			readIndex++
			continue
		}
		if bb {
			runeArray[writeIndex] = unicode.ToUpper(runeArray[readIndex])
		} else {
			runeArray[writeIndex] = runeArray[readIndex]
		}
		bb = false

		writeIndex++
		readIndex++
	}
	return string(runeArray[0:writeIndex])
}

func (self IdlToGoTranslation) CamelCase(name string) string {
	if len(name) == 0 {
		return name
	}
	b := []byte(name)
	b[0] = byte(unicode.ToUpper(rune(name[0])))
	return string(b)
}

func (self IdlToGoTranslation) UpFirstLetter(name string) string {
	b := []byte(name)
	b[0] = byte(unicode.ToUpper(rune(name[0])))
	return string(b)
}

func (self IdlToGoTranslation) ExportMethodName(name string) string {
	return self.CamelCase(name)
}

func (self IdlToGoTranslation) ExportTypeName(name string) string {
	b := []byte(name)
	b[0] = byte(unicode.ToUpper(rune(name[0])))
	return string(b)
}

func (self IdlToGoTranslation) ExportMemberName(name string) string {
	ss := strings.Split(name, "_")
	for i := range ss {
		if ss[i] == "" {
			continue
		}
		ss[i] = self.RemoveAllCaps(ss[i])
		if ss[i] == "" {
			continue
		}
		ss[i] = self.UpFirstLetter(ss[i])
	}
	return strings.Join(ss, "")
}

func (self IdlToGoTranslation) ExportParamName(name string) interface{} {
	ss := strings.Split(name, "_")
	for i := range ss {
		if ss[i] == "" {
			continue
		}
		ss[i] = self.RemoveAllCaps(ss[i])
		if ss[i] == "" {
			continue
		}
		if i != 0 {
			ss[i] = self.UpFirstLetter(ss[i])
		}

	}
	return strings.Join(ss, "")
}

func (self IdlToGoTranslation) RemoveAllCaps(name string) string {
	name = self.RemoveUnderScores(name)
	rr := []rune(name)
	b := true
	for _, r := range rr {
		b = b && unicode.IsUpper(r)
		if !b {
			return name
		}
	}

	name2 := strings.ToLower(name)
	name2 = self.CamelCase(name2)
	return name2
}

func (self IdlToGoTranslation) BuildArrayContext(value []int64) string {
	arrayContext := ""
	if value != nil {
		sb := strings.Builder{}
		for _, v := range value {
			sb.WriteString(fmt.Sprintf("[%v]", v))
		}
		arrayContext = sb.String()
	}
	return arrayContext
}

type CreateTypeReferenceAnswer struct {
	Reference         GoLangTypeReference
	PrimitiveType     yaccidl.PrimitiveType
	MapToLanguageType bool
	//knownType         objects.IDcl
	LanguageType    objects.IDcl
	TypeArrayValues []int64
}

func NewCreateTypeReferenceAnswer(
	reference GoLangTypeReference,
	primitiveType yaccidl.PrimitiveType,
	mapToLanguageType bool,
	//knownType objects.IDcl,
	languageType objects.IDcl,
	typeArrayValues []int64) *CreateTypeReferenceAnswer {
	return &CreateTypeReferenceAnswer{
		Reference:         reference,
		PrimitiveType:     primitiveType,
		MapToLanguageType: mapToLanguageType,
		//knownType:         knownType,
		LanguageType:    languageType,
		TypeArrayValues: typeArrayValues,
	}
}
func (self IdlToGoTranslation) CreateTypeReference(
	AcurrentLexem objects.IDstSrcInformation,
	AparamArrayValues []int64,
	AincomingDcl objects.IDcl,
	AknownTypes objects.KnownTypes) (*CreateTypeReferenceAnswer, error) {

	ans, err := self.DclHelpers.FindPrimitiveTypeForWriters(AincomingDcl, AknownTypes)
	if err != nil {
		return nil, err
	}
	typeArrayValues := AincomingDcl.GetArrayValue()
	refType := AincomingDcl
	knownType, ok := AknownTypes[AincomingDcl.GetName()]
	if ok {
		typeArrayValues = knownType.GetArrayValue()
	} else {
		return nil, self.logger.Error(fmt.Errorf("could not find type %v", AincomingDcl.GetName()))
	}
	buildTypeName := self.BuildTypeName(refType.GetName())
	mapToLanguageType := false
	languagetype, b := self.LanguageTypeService.Find(ans.PrimitiveType)
	if b {
		refType = languagetype
		buildTypeName = languagetype.GetName()
		knownType = languagetype
		mapToLanguageType = true
	}

	paramArrayContext := self.BuildArrayContext(AparamArrayValues)
	typeArrayContext := self.BuildArrayContext(typeArrayValues)

	return NewCreateTypeReferenceAnswer(
		self.TypeNameFromBuildTypeName(
			AcurrentLexem,
			knownType.GetLexemData(),
			buildTypeName,
			paramArrayContext,
			typeArrayContext),
		ans.PrimitiveType,
		mapToLanguageType,
		languagetype,
		typeArrayValues), nil
}

func NewIdlToGoTranslation(
	//appCtx *ctx.GoYaccAppCtx,
	logger *log2.SubSystemLogger,
	DclHelpers *DclHelpers,
	LanguageTypeService *LanguageTypeService) *IdlToGoTranslation {
	return &IdlToGoTranslation{
		//appCtx:              appCtx,
		DclHelpers:          DclHelpers,
		LanguageTypeService: LanguageTypeService,
		logger:              logger,
	}
}
