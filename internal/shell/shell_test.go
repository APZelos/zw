package shell

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIntegration(t *testing.T) {
	tests := []struct {
		name      string
		shell     string
		wantErr   bool
		errSubstr string
		contains  []string
	}{
		{
			name:    "bash returns valid script",
			shell:   "bash",
			wantErr: false,
			contains: []string{
				"zw()",
				"command zw --query",
				"cd \"$result\"",
			},
		},
		{
			name:    "zsh returns valid script",
			shell:   "zsh",
			wantErr: false,
			contains: []string{
				"zw()",
				"command zw --query",
				"cd \"$result\"",
			},
		},
		{
			name:    "fish returns valid script",
			shell:   "fish",
			wantErr: false,
			contains: []string{
				"function zw",
				"command zw --query",
				"cd $result",
			},
		},
		{
			name:      "unsupported shell returns error",
			shell:     "powershell",
			wantErr:   true,
			errSubstr: "unsupported shell",
		},
		{
			name:      "empty shell returns error",
			shell:     "",
			wantErr:   true,
			errSubstr: "unsupported shell",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetIntegration(tt.shell)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errSubstr)
				return
			}

			require.NoError(t, err)
			for _, substr := range tt.contains {
				assert.Contains(t, result, substr, "script should contain %q", substr)
			}
		})
	}
}

func TestBashIntegration(t *testing.T) {
	script := BashIntegration()

	assert.True(t, strings.HasPrefix(script, "#"), "should start with a comment")
	assert.Contains(t, script, "zw()", "should define zw function")
	assert.Contains(t, script, "command zw", "should call command zw")
	assert.Contains(t, script, "--query", "should use --query flag")
	assert.Contains(t, script, "init", "should pass through init command")
	assert.Contains(t, script, "version", "should pass through version command")
	assert.Contains(t, script, "--help", "should pass through --help flag")
}

func TestZshIntegration(t *testing.T) {
	script := ZshIntegration()

	assert.True(t, strings.HasPrefix(script, "#"), "should start with a comment")
	assert.Contains(t, script, "zw()", "should define zw function")
	assert.Contains(t, script, "command zw", "should call command zw")
	assert.Contains(t, script, "--query", "should use --query flag")
}

func TestFishIntegration(t *testing.T) {
	script := FishIntegration()

	assert.True(t, strings.HasPrefix(script, "#"), "should start with a comment")
	assert.Contains(t, script, "function zw", "should define zw function")
	assert.Contains(t, script, "command zw", "should call command zw")
	assert.Contains(t, script, "--query", "should use --query flag")
	assert.Contains(t, script, "end", "should close function with end")
}
