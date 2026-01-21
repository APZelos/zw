package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func executeCommand(args ...string) (string, string, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(args)

	// Reset flags to default values before each test
	queryFlag = ""

	err := rootCmd.Execute()
	return stdout.String(), stderr.String(), err
}

func TestRootCmd_Help(t *testing.T) {
	stdout, _, err := executeCommand("--help")

	require.NoError(t, err)
	assert.Contains(t, stdout, "zw")
	assert.Contains(t, stdout, "git worktree")
	assert.Contains(t, stdout, "--query")
}

func TestRootCmd_QueryFlag(t *testing.T) {
	// Test that --query flag is recognized and doesn't error
	// The actual functionality is tested when implemented
	_, _, err := executeCommand("--query", "test")
	require.NoError(t, err, "--query flag should be recognized")
}

func TestRootCmd_NoArgs(t *testing.T) {
	// With no args and no flags, should show help (not error)
	stdout, _, err := executeCommand()
	require.NoError(t, err)
	assert.Contains(t, stdout, "Usage:")
}

func TestVersionCmd(t *testing.T) {
	stdout, _, err := executeCommand("version")

	require.NoError(t, err)
	assert.Contains(t, stdout, "zw")
	assert.Contains(t, stdout, "commit:")
	assert.Contains(t, stdout, "built:")
}

func TestInitCmd_Bash(t *testing.T) {
	stdout, _, err := executeCommand("init", "bash")

	require.NoError(t, err)
	assert.Contains(t, stdout, "zw()")
	assert.Contains(t, stdout, "command zw --query")
}

func TestInitCmd_Zsh(t *testing.T) {
	stdout, _, err := executeCommand("init", "zsh")

	require.NoError(t, err)
	assert.Contains(t, stdout, "zw()")
	assert.Contains(t, stdout, "command zw --query")
}

func TestInitCmd_Fish(t *testing.T) {
	stdout, _, err := executeCommand("init", "fish")

	require.NoError(t, err)
	assert.Contains(t, stdout, "function zw")
	assert.Contains(t, stdout, "command zw --query")
}

func TestInitCmd_UnsupportedShell(t *testing.T) {
	_, _, err := executeCommand("init", "powershell")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported shell")
}

func TestInitCmd_NoArgs(t *testing.T) {
	_, _, err := executeCommand("init")

	assert.Error(t, err, "should require shell argument")
}

func TestVersionVariables(t *testing.T) {
	assert.NotEmpty(t, Version)
	assert.NotEmpty(t, Commit)
	assert.NotEmpty(t, Date)
}
