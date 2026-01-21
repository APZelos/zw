package shell

import "fmt"

// GetIntegration returns the shell integration script for the specified shell.
func GetIntegration(shellName string) (string, error) {
	switch shellName {
	case "bash":
		return BashIntegration(), nil
	case "zsh":
		return ZshIntegration(), nil
	case "fish":
		return FishIntegration(), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s (supported: bash, zsh, fish)", shellName)
	}
}
