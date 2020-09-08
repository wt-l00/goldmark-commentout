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
	var buf bytes.Buffer

	source := []byte("//TODO: abc//")
	err := md.Convert(source, &buf)
	assert.NoError(t, err)
	assert.Equal(t, "<p><!-- TODO: abc --></p>\n", buf.String())
}
