package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	spew.Dump(s)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$veDRvkG.p4Jpx934w0p0bOJD0ScvpFPug2JQlWHy4tLKSeqju6.y.", "abc32"))
	//spew.Dump(bl)
}
