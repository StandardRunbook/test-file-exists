package plugin_test

import (
	"testing"

	"github.com/StandardRunbook/test-file-exists/pkg/config"
	"github.com/StandardRunbook/test-file-exists/pkg/plugin"
	"github.com/stretchr/testify/require"
)

func TestTemplate_Run(t *testing.T) {
	t.Parallel()
	cfg := &config.FileCheckConfig{
		Name:            "test-file-exists",
		Version:         "v1.0.0",
		ExpectedOutput:  "File exists: ./run.sh",
		ScriptArguments: []string{"./run.sh"},
	}
	plugin := plugin.NewFileCheckPlugin(cfg)
	require.Equal(t, plugin.Name(), "test-file-exists")
	require.Equal(t, plugin.Version(), "v1.0.0")
	err := plugin.Run()
	require.NoError(t, err)
	require.Equal(t, plugin.ParseOutput(), "success")
}
