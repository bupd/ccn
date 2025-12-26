package cmd

import (
	"fmt"

	"github.com/bupd/ccn/internal/tmux"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach <session>",
	Short: "Switch to a tmux session",
	Long:  "Switches to the specified tmux session. If inside tmux, uses switch-client. If outside, attaches directly.",
	Args:  cobra.ExactArgs(1),
	RunE:  runAttach,
}

func init() {
	rootCmd.AddCommand(attachCmd)
}

func runAttach(cmd *cobra.Command, args []string) error {
	session := args[0]

	if err := tmux.SwitchTo(session); err != nil {
		return fmt.Errorf("failed to switch to session %s: %w", session, err)
	}

	return nil
}
