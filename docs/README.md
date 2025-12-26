# Claude Code Notify

Desktop notifications for Claude Code with tmux integration.

## The Problem

Running multiple Claude Code sessions in tmux. Each instance works autonomously. When Claude stops and needs input, you have no way to know which session needs attention.

## The Solution

Get desktop notifications when any Claude Code instance needs you. Click to jump directly to that tmux pane.

## Quick Start

```bash
cp scripts/claude-notify ~/.local/bin/
chmod +x ~/.local/bin/claude-notify
```

Then add hooks to `~/.claude/settings.json`. See [Installation](installation.md).
