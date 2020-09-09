package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

// Commentout は//text//を <!-- text -->に変換する
type Commentout struct {
	gast.BaseInline
}

// Dump implements Node.Dump.
func (n *Commentout) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindCommentout is a NodeKind of the Commentout node.
var KindCommentout = gast.NewNodeKind("Commentout")

// Kind implements Node.Kind.
func (n *Commentout) Kind() gast.NodeKind {
	return KindCommentout
}

// NewCommentout returns a new Commentout node
func NewCommentout() *Commentout {
	return &Commentout{}
}
