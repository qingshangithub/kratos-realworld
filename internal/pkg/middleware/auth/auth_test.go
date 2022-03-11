package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("secret", "2233")
	spew.Dump(tk)
}
