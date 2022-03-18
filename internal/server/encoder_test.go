package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"kratos-realworld/internal/errors"
	"testing"
)

func TestHttpError(t *testing.T) {
	e := &errors.HTTPError{
		Errors: make(map[string][]string),
	}
	e.Errors["body"] = []string{"this is error body"}
	b, err := json.Marshal(e)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
}
