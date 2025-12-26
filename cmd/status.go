package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/bupd/ccn/internal/tmux"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Claude Code instances in tmux",
	Long:  "Scans tmux panes for running Claude Code processes and shows their status.",
	RunE:  runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatus(cmd *cobra.Command, args []string) error {
	panes, err := tmux.ListPanes()
	if err != nil {
		return fmt.Errorf("failed to list panes: %w", err)
	}

	fmt.Printf("%-15s %-50s %s\n", "SESSION", "PATH", "CLAUDE")
	fmt.Printf("%-15s %-50s %s\n", "-------", "----", "------")

	for _, p := range panes {
		hasClaude := checkClaudeInPane(p.Target)
		status := "-"
		if hasClaude {
			status = "running"
		}

		path := p.Path
		if len(path) > 50 {
			path = "..." + path[len(path)-47:]
		}

		session := tmux.SessionFromTarget(p.Target)
		fmt.Printf("%-15s %-50s %s\n", session, path, status)
	}

	return nil
}

func checkClaudeInPane(target string) bool {
	cmd := exec.Command("tmux", "list-panes", "-t", target, "-F", "#{pane_current_command}")
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	command := strings.TrimSpace(string(out))
	return strings.Contains(strings.ToLower(command), "claude")
}
