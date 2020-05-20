package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSLLoad(t *testing.T) {
	dsl := &DSL{}
	err := dsl.Load("./testdata/test_dsl.yaml")
	assert.NoError(t, err)
}
