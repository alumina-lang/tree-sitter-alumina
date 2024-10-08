package tree_sitter_alumina_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_alumina "github.com/alumina-lang/tree-sitter-alumina/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_alumina.Language())
	if language == nil {
		t.Errorf("Error loading Alumina grammar")
	}
}
