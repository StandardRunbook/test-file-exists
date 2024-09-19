package plugin

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"

	pluginInterface "github.com/StandardRunbook/plugin-interface"
	"github.com/StandardRunbook/plugin-template-go/app/config"
)

//go:embed run.sh
var runScript []byte

// Template is a placeholder - please change to be unique to your plugin name
type Template struct {
	name           string
	version        string
	arguments      []string
	output         string
	expectedOutput string
}

func (t *Template) Name() string {
	return t.name
}

func (t *Template) Version() string {
	return t.version
}

func (t *Template) Run() error {
	// Step 1: Create a temporary file
	tmpFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Ensure the file is removed after execution

	// Step 2: Write the embedded script to the temporary file
	_, err = tmpFile.Write(runScript)
	if err != nil {
		return fmt.Errorf("failed to write '%s' script to temporary file: %w", t.Name(), err)
	}

	// Step 3: Close the file to flush writes and prepare it for execution
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	// Step 4: Set the appropriate permissions to make the script executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		return fmt.Errorf("failed to set executable permissions on file: %w", err)
	}

	// Step 5: Execute the script
	cmd := exec.Command(tmpFile.Name(), t.arguments...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing script: %w", err)
	}
	t.output = string(output)
	return nil
}

func (t *Template) ParseOutput() string {
	if strings.Contains(t.output, t.expectedOutput) {
		return "success"
	}
	return "failure"
}

func NewPluginTemplate(cfg *config.Config) pluginInterface.IPlugin {
	return &Template{
		name:           cfg.Name,
		version:        cfg.Version,
		arguments:      cfg.ScriptArguments,
		expectedOutput: cfg.ExpectedOutput,
	}
}
