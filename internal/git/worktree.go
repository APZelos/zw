package git

import (
	"bufio"
	"os/exec"
	"strings"
)

// Worktree represents a git worktree.
type Worktree struct {
	Path   string
	Branch string
	Head   string
}

// ListWorktrees returns all worktrees in the current repository.
func ListWorktrees() ([]Worktree, error) {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parseWorktreeOutput(string(output))
}

func parseWorktreeOutput(output string) ([]Worktree, error) {
	var worktrees []Worktree
	var current Worktree

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if current.Path != "" {
				worktrees = append(worktrees, current)
				current = Worktree{}
			}
			continue
		}

		if strings.HasPrefix(line, "worktree ") {
			current.Path = strings.TrimPrefix(line, "worktree ")
		} else if strings.HasPrefix(line, "HEAD ") {
			current.Head = strings.TrimPrefix(line, "HEAD ")
		} else if strings.HasPrefix(line, "branch ") {
			branch := strings.TrimPrefix(line, "branch ")
			// Remove refs/heads/ prefix if present
			branch = strings.TrimPrefix(branch, "refs/heads/")
			current.Branch = branch
		}
	}

	// Don't forget the last worktree if there's no trailing newline
	if current.Path != "" {
		worktrees = append(worktrees, current)
	}

	return worktrees, scanner.Err()
}

// GetRepoRoot returns the root directory of the git repository.
func GetRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
