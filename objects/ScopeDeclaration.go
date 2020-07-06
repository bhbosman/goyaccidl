package objects

type ScopeDeclaration struct {
	DeclaredItems map[ScopeIdentifier]*ScopeItem `json:"declared_items"`
	ImportedItems []*ScopeItem                   `json:"imported_items"`
}

func NewScopeDeclaration(items []*ScopeItem) *ScopeDeclaration {
	return &ScopeDeclaration{
		DeclaredItems: make(map[ScopeIdentifier]*ScopeItem),
		ImportedItems: items,
	}
}
