# Claude Code Notify (CCN)

![Claude Code Notify](claude-code-notify.png)

Desktop notifications for Claude Code CLI with tmux integration.

## For Claude Code Power Users

Running multiple Claude Code sessions in parallel? CCN notifies you when any instance needs attention.

No more checking each tmux pane. Work on something else. Get notified. Switch instantly.

## How It Works

1. You run Claude Code inside tmux
2. Claude stops and waits for input
3. You get a desktop notification
4. Click or press Enter to jump to that pane

## The Problem

You're running 3+ Claude Code instances in tmux. Each works autonomously on different tasks. When one finishes and needs input, you have no way to know without manually checking each pane.

## The Solution

CCN uses Claude Code's hook system to trigger desktop notifications. When Claude stops or sends a notification event, you see it immediately.

Click the notification to switch directly to the exact tmux session, window, and pane.

## Quick Start

```bash
git clone https://github.com/bupd/ccn.git
cp ccn/scripts/claude-notify ~/.local/bin/
chmod +x ~/.local/bin/claude-notify
```

Then add hooks to `~/.claude/settings.json`. See [Installation](installation.md).

## Requirements

- Linux (X11/Wayland)
- tmux
- dunst (notification daemon)
- Claude Code CLI
