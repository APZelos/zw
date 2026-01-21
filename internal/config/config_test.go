package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.NotNil(t, cfg)
	assert.Equal(t, "main", cfg.DefaultBranch)
}

func TestLoad_NoConfigFile(t *testing.T) {
	// Use a temp directory as HOME
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	cfg, err := Load()

	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "main", cfg.DefaultBranch, "should return default config when file doesn't exist")
}

func TestLoad_ValidConfigFile(t *testing.T) {
	// Use a temp directory as HOME
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create config directory and file
	configDir := filepath.Join(tmpDir, ".config", "zw")
	require.NoError(t, os.MkdirAll(configDir, 0755))

	configContent := `default_branch = "develop"
`
	configPath := filepath.Join(configDir, "config.toml")
	require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644))

	cfg, err := Load()

	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "develop", cfg.DefaultBranch)
}

func TestLoad_InvalidTOML(t *testing.T) {
	// Use a temp directory as HOME
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create config directory and file with invalid TOML
	configDir := filepath.Join(tmpDir, ".config", "zw")
	require.NoError(t, os.MkdirAll(configDir, 0755))

	configContent := `this is not valid toml {{{`
	configPath := filepath.Join(configDir, "config.toml")
	require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644))

	_, err := Load()

	assert.Error(t, err, "should return error for invalid TOML")
}

func TestLoad_EmptyConfigFile(t *testing.T) {
	// Use a temp directory as HOME
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create config directory and empty file
	configDir := filepath.Join(tmpDir, ".config", "zw")
	require.NoError(t, os.MkdirAll(configDir, 0755))

	configPath := filepath.Join(configDir, "config.toml")
	require.NoError(t, os.WriteFile(configPath, []byte(""), 0644))

	cfg, err := Load()

	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "main", cfg.DefaultBranch, "should use defaults for empty config")
}

func TestConfig_PartialOverride(t *testing.T) {
	// Use a temp directory as HOME
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create config with only some fields
	configDir := filepath.Join(tmpDir, ".config", "zw")
	require.NoError(t, os.MkdirAll(configDir, 0755))

	configContent := `# Only override default_branch
default_branch = "master"
`
	configPath := filepath.Join(configDir, "config.toml")
	require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644))

	cfg, err := Load()

	require.NoError(t, err)
	assert.Equal(t, "master", cfg.DefaultBranch)
}
