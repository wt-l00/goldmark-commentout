package commentout

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
)

func TestCommentout(t *testing.T) {
	var md = goldmark.New(
		goldmark.WithExtensions(
			Commentout,
		),
	)

	commentoutTests := []struct {
		actual   string
		expected string
	}{
		{"//TODO: something//", "<!-- TODO: something -->"},
		{"123//TODO: something//", "<p>123</p>\n<!-- TODO: something -->"},
		{"//TODO: something//456", "<!-- TODO: something --><p>456</p>\n"},
		{"123//TODO: something//456", "<p>123456</p>\n<!-- TODO: something -->"},
	}

	for _, commentout := range commentoutTests {
		var buf bytes.Buffer
		err := md.Convert([]byte(commentout.actual), &buf)
		assert.NoError(t, err)
		assert.Equal(t, commentout.expected, buf.String())
	}
}
