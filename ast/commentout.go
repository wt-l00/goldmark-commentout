package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

// Commentout は//text//を <!-- text -->に変換する
type Commentout struct {
	gast.BaseInline
}

// Dump helper
func (n *Commentout) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindCommentout = gast.NewNodeKind("Commentout")

// Kind implements Node.Kind.
func (n *Commentout) Kind() gast.NodeKind {
	return KindCommentout
}

func NewCommentout() *Commentout {
	return &Commentout{}
}
