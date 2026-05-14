package tree_sitter_lang1_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-lang1"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_lang1.Language())
	if language == nil {
		t.Errorf("Error loading Lang1 grammar")
	}
}
