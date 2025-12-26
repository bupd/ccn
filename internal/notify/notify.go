package notify

import (
	"os/exec"
	"strings"
)

type Urgency string

const (
	UrgencyLow      Urgency = "low"
	UrgencyNormal   Urgency = "normal"
	UrgencyCritical Urgency = "critical"
)

type Notification struct {
	Title   string
	Body    string
	Urgency Urgency
	Actions []Action
}

type Action struct {
	ID    string
	Label string
}

// Send sends a notification and waits for action if actions are provided
// Returns the action ID that was clicked, or empty string if dismissed
func Send(n Notification) (string, error) {
	args := []string{}

	if n.Urgency != "" {
		args = append(args, "-u", string(n.Urgency))
	}

	for _, a := range n.Actions {
		args = append(args, "-A", a.ID+"="+a.Label)
	}

	if len(n.Actions) > 0 {
		args = append(args, "--wait")
	}

	args = append(args, n.Title, n.Body)

	cmd := exec.Command("notify-send", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// UrgencyFromType maps notification type to urgency
func UrgencyFromType(notifType string) Urgency {
	switch notifType {
	case "permission_prompt":
		return UrgencyCritical
	case "idle_prompt":
		return UrgencyNormal
	default:
		return UrgencyNormal
	}
}
