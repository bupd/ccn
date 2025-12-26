package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Pane struct {
	Session string
	Window  int
	Pane    int
	Path    string
	Target  string // session:window.pane format
}

// ListPanes returns all tmux panes with their paths
func ListPanes() ([]Pane, error) {
	cmd := exec.Command("tmux", "list-panes", "-a", "-F", "#{session_name}:#{window_index}.#{pane_index} #{pane_current_path}")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list panes: %w", err)
	}

	var panes []Pane
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}
		panes = append(panes, Pane{
			Target: parts[0],
			Path:   parts[1],
		})
	}
	return panes, nil
}

// FindByPath finds panes matching the given path
func FindByPath(path string) ([]Pane, error) {
	panes, err := ListPanes()
	if err != nil {
		return nil, err
	}

	var matches []Pane
	for _, p := range panes {
		if p.Path == path {
			matches = append(matches, p)
		}
	}
	return matches, nil
}

// IsInsideTmux checks if we're running inside tmux
func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

// SwitchTo switches to the given session/target
// When called from background (no TTY), switches the most recent client
func SwitchTo(target string) error {
	// Always use switch-client - it works even from background
	// -t specifies the target session, the current/most-recent client will switch
	cmd := exec.Command("tmux", "switch-client", "-t", target)
	if err := cmd.Run(); err != nil {
		// Fallback: try selecting the pane directly
		cmd = exec.Command("tmux", "select-window", "-t", target)
		return cmd.Run()
	}
	return nil
}

// SessionFromTarget extracts session name from target (session:window.pane)
func SessionFromTarget(target string) string {
	if idx := strings.Index(target, ":"); idx != -1 {
		return target[:idx]
	}
	return target
}
