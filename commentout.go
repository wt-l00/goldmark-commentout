package commentout

import (
	"github.com/wt-l00/goldmark-commentout/ast"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"

	gast "github.com/yuin/goldmark/ast"
)

type commentoutDelimiterProcessor struct {
}

func (p *commentoutDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '/'
}

func (p *commentoutDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *commentoutDelimiterProcessor) OnMatch(consumes int) gast.Node {
	return ast.NewCommentout()
}

var defaultCommentoutDelimiterProcessor = &commentoutDelimiterProcessor{}

type commentoutParser struct {
}

var defaultCommentoutParser = &commentoutParser{}

func NewCommentoutParser() parser.InlineParser {
	return defaultCommentoutParser
}

func (s *commentoutParser) Trigger() []byte {
	return []byte{'/'}
}

func (s *commentoutParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultCommentoutDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *commentoutParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

type CommentoutHTMLRenderer struct {
	html.Config
}

func NewCommentoutHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &CommentoutHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *CommentoutHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindCommentout, r.renderCommentout)
}

var CommentoutAttributeFilter = html.GlobalAttributeFilter

func (r *CommentoutHTMLRenderer) renderCommentout(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<!-- ")
	} else {
		_, _ = w.WriteString(" -->")
	}
	return gast.WalkContinue, nil
}

type commentoutASTTransformer struct {
}

var defaultCommentoutASTTransformer = &commentoutASTTransformer{}

func NewCommentoutASTTransformer() parser.ASTTransformer {
	return defaultCommentoutASTTransformer
}

func (a *commentoutASTTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
	gast.Walk(node, func(node gast.Node, entering bool) (gast.WalkStatus, error) {
		if commentoutNode, ok := node.(*ast.Commentout); ok && entering && gast.IsParagraph(node.Parent()) {
			paragrahphNode := commentoutNode.Parent()

			paragrahphNode.Parent().AppendChild(paragrahphNode.Parent(), commentoutNode)

			if paragrahphNode.ChildCount() == 0 {
				paragrahphNode.Parent().RemoveChild(paragrahphNode.Parent(), paragrahphNode)
			}
		}
		return gast.WalkContinue, nil
	})
}

type commentout struct {
}

var Commentout = &commentout{}

func (e *commentout) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewCommentoutParser(), 500),
		),
		parser.WithASTTransformers(
			util.Prioritized(NewCommentoutASTTransformer(), 500),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewCommentoutHTMLRenderer(), 500),
	))
}
