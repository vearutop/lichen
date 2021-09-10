package module_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vearutop/lichen/internal/model"
	"github.com/vearutop/lichen/internal/module"
)

func TestModuleFetchNoModules(test *testing.T) {
	modules, err := module.Fetch(context.Background(), []model.ModuleReference{})

	assert.NoError(test, err)
	assert.Empty(test, modules)
}
