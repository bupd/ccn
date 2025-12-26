package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bupd/ccn/internal/notify"
	"github.com/bupd/ccn/internal/tmux"
	"github.com/spf13/cobra"
)

type hookInput struct {
	SessionID        string `json:"session_id"`
	Cwd              string `json:"cwd"`
	HookEventName    string `json:"hook_event_name"`
	Message          string `json:"message"`
	NotificationType string `json:"notification_type"`
}

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Read Claude hook JSON from stdin and send notification",
	Long:  "Reads JSON from stdin (sent by Claude Code hook), finds the tmux session, sends a notification, and attaches on click.",
	RunE:  runNotify,
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}

func runNotify(cmd *cobra.Command, args []string) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	var input hookInput
	if err := json.Unmarshal(data, &input); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if input.Cwd == "" {
		return fmt.Errorf("cwd not provided in hook input")
	}

	panes, err := tmux.FindByPath(input.Cwd)
	if err != nil {
		return fmt.Errorf("failed to find tmux pane: %w", err)
	}

	if len(panes) == 0 {
		return fmt.Errorf("no tmux pane found for path: %s", input.Cwd)
	}

	target := panes[0].Target
	session := tmux.SessionFromTarget(target)

	title := session
	if len(panes) > 1 {
		title = fmt.Sprintf("%s (+%d more)", session, len(panes)-1)
	}

	body := input.Message
	if body == "" {
		body = "Claude needs input"
	}

	n := notify.Notification{
		Title:   title,
		Body:    body,
		Urgency: notify.UrgencyFromType(input.NotificationType),
		Actions: []notify.Action{
			{ID: "open", Label: "Open"},
		},
	}

	action, err := notify.Send(n)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	if action == "open" {
		return tmux.SwitchTo(target)
	}

	return nil
}
