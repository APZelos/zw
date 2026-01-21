package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseWorktreeOutput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Worktree
		wantErr bool
	}{
		{
			name: "single worktree",
			input: `worktree /home/user/project
HEAD abc123def456
branch refs/heads/main

`,
			want: []Worktree{
				{
					Path:   "/home/user/project",
					Head:   "abc123def456",
					Branch: "main",
				},
			},
		},
		{
			name: "multiple worktrees",
			input: `worktree /home/user/project
HEAD abc123def456
branch refs/heads/main

worktree /home/user/project-feature
HEAD def456abc123
branch refs/heads/feature/auth

`,
			want: []Worktree{
				{
					Path:   "/home/user/project",
					Head:   "abc123def456",
					Branch: "main",
				},
				{
					Path:   "/home/user/project-feature",
					Head:   "def456abc123",
					Branch: "feature/auth",
				},
			},
		},
		{
			name: "worktree with detached HEAD",
			input: `worktree /home/user/project
HEAD abc123def456
detached

`,
			want: []Worktree{
				{
					Path:   "/home/user/project",
					Head:   "abc123def456",
					Branch: "",
				},
			},
		},
		{
			name: "worktree without trailing newline",
			input: `worktree /home/user/project
HEAD abc123def456
branch refs/heads/main`,
			want: []Worktree{
				{
					Path:   "/home/user/project",
					Head:   "abc123def456",
					Branch: "main",
				},
			},
		},
		{
			name:    "empty output",
			input:   "",
			want:    nil,
			wantErr: false,
		},
		{
			name: "bare repository worktree",
			input: `worktree /home/user/project.git
bare

`,
			want: []Worktree{
				{
					Path:   "/home/user/project.git",
					Head:   "",
					Branch: "",
				},
			},
		},
		{
			name: "branch without refs/heads prefix",
			input: `worktree /home/user/project
HEAD abc123
branch main

`,
			want: []Worktree{
				{
					Path:   "/home/user/project",
					Head:   "abc123",
					Branch: "main",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseWorktreeOutput(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWorktree_Fields(t *testing.T) {
	wt := Worktree{
		Path:   "/home/user/project",
		Branch: "feature/test",
		Head:   "abc123",
	}

	assert.Equal(t, "/home/user/project", wt.Path)
	assert.Equal(t, "feature/test", wt.Branch)
	assert.Equal(t, "abc123", wt.Head)
}
