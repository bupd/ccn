# Claude Code Notify (CCN)

![Claude Code Notify](docs/claude-code-notify.png)

Desktop notifications for Claude Code with tmux integration.

[View Documentation](https://cc-notify.bupd.xyz)

## The Problem

You're running multiple Claude Code sessions in tmux. Each instance works autonomously on different tasks. But when Claude finishes and needs your input, you have no way to know which session needs attention without manually checking each one.

## The Solution

This script sends desktop notifications whenever Claude Code:
- Stops and waits for input
- Sends a notification event

Click the notification to jump directly to the tmux pane that needs you.

## User Story

> As a developer running multiple Claude Code instances in tmux,
> I want to receive desktop notifications when any instance needs my attention,
> so that I can context-switch efficiently without polling each session.

## Quick Start

```bash
# 1. Copy the notify script
cp scripts/claude-notify ~/.local/bin/
chmod +x ~/.local/bin/claude-notify

# 2. Add hooks to your Claude Code settings (~/.claude/settings.json)
# See https://cc-notify.bupd.xyz/#/installation for details
```

## Requirements

- Linux with X11/Wayland
- tmux
- dunst (notification daemon)
- Claude Code CLI

## Documentation

- [Installation Guide](https://cc-notify.bupd.xyz/#/installation)
- [Configuration](https://cc-notify.bupd.xyz/#/configuration)
- [Keyboard Navigation](https://cc-notify.bupd.xyz/#/keyboard-navigation)

## Why Not a Full Orchestration Tool?

There are great multi-agent orchestration tools out there:

- [claude-squad](https://github.com/smtg-ai/claude-squad) - manage multiple AI terminal agents
- [claude-colony](https://github.com/MakingJamie/claude-colony) - tmux-based multi-agent orchestration
- [claude-code-agent-farm](https://github.com/Dicklesworthstone/claude_code_agent_farm) - run 20+ agents in parallel
- [ccmanager](https://github.com/ryanhoangt/ccmanager) - Claude Code session manager

CCN is intentionally minimal. No meta layer. No orchestration. Just notifications.

I work alongside Claude across multiple features running in parallel. I prefer manual approval over yolo mode. When context gets full, Claude starts yak shaving. These notifications help me monitor all my instances without the overhead of orchestrators. Simple and minimal.

## How It Works

```
Claude Code stops/notifies
        |
        v
Hook triggers claude-notify
        |
        v
Captures tmux session/window/pane info
        |
        v
Sends dunstify notification with action
        |
        v
User clicks "Go to window"
        |
        v
Switches to exact tmux pane
```

## License

Apache-2.0
