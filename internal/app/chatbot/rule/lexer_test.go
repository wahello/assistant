package rule

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer([]rune("subs  open abc"))
	token, err := l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, CharacterToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, CharacterToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, CharacterToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, EOFToken, token.Type)
}

func TestParseCommand(t *testing.T) {
	c, err := ParseCommand("subs")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 1)

	c, err = ParseCommand("subs list")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 2)

	c, err = ParseCommand("subs open abc")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 3)

	require.Equal(t, "subs", c[0].Value)
	require.Equal(t, "open", c[1].Value)
	require.Equal(t, "abc", c[2].Value)
}