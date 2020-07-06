package objects

type FlattenDeclaredTypesOut struct {
	InputStreamDcl           *InputStreamDcl `json:"input_stream_dcl"`
	ImportedDeclarationItems []*ScopeItem    `json:"imported_declaration_items"`
}

func NewFlattenDeclaredTypesOut(
	inputStreamDcl *InputStreamDcl,
	importedDeclarationItems []*ScopeItem) *FlattenDeclaredTypesOut {
	return &FlattenDeclaredTypesOut{
		InputStreamDcl:           inputStreamDcl,
		ImportedDeclarationItems: importedDeclarationItems,
	}
}
