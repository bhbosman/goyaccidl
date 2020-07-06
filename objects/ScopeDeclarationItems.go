package objects

type ScopeDeclarationItems map[string]*ScopeDeclaration

func NewScopeDeclarationItems() ScopeDeclarationItems {
	return make(ScopeDeclarationItems)
}
