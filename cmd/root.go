package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccn",
	Short: "Claude Code Notification CLI",
	Long:  "A simple CLI tool that sends desktop notifications when Claude Code needs input and navigates to the correct tmux session.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
