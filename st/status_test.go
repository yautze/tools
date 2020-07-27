package st

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_NewErr(t *testing.T) {
	err := NewError(0, "success", OK)
	spew.Dump(err)
	assert.NoError(t, err)

	err = NewError(1000, "NotFound", NotFound)
	spew.Dump(err)
	assert.Error(t, err, "NotFound")
}
