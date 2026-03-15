# Codex Installation

## Prerequisites

Install dependencies:

```bash
# Arch Linux
sudo pacman -S dunst tmux jq

# Ubuntu/Debian
sudo apt install dunst tmux jq

# Fedora
sudo dnf install dunst tmux jq
```

Make sure `dunst` is running.

## Install Scripts

```bash
git clone https://github.com/bupd/ccn.git
cd ccn

cp scripts/codex-notify ~/.local/bin/
cp scripts/codex-watch ~/.local/bin/
cp scripts/codex-with-notify ~/.local/bin/
chmod +x ~/.local/bin/codex-notify ~/.local/bin/codex-watch ~/.local/bin/codex-with-notify
```

## Recommended Usage (tmux-aware pane jump)

Run Codex with the wrapper from inside tmux:

```bash
codex-with-notify
```

Optional shell alias:

```bash
alias codexn="$HOME/.local/bin/codex-with-notify"
```

Now use `codexn` whenever you want notifications.

## Why `codex-with-notify` (and why this differs from Claude)

Claude can call notification commands directly through built-in hooks in `~/.claude/settings.json` (for example, on `Stop`/`Notification` events). That means plain `claude` can trigger notifications without a wrapper.

Codex does not currently expose the same hook flow for this use case, so plain `codex` does not automatically run a notification command when a task completes.

`codex-with-notify` solves this by starting `codex-watch` in the background while Codex runs, then cleaning it up when Codex exits. This gives you Claude-like completion notifications for Codex.

When you launch multiple Codex panes at once, the wrapper binds its watcher to the exact session file opened by that Codex process. That keeps pane-jump notifications attached to the correct tmux pane instead of racing across all active sessions.

## Optional Always-on Watcher (global alerts)

This mode notifies for Codex completions across sessions, even when you do not use the wrapper.

```bash
mkdir -p ~/.config/systemd/user
cp systemd/user/codex-watch.service ~/.config/systemd/user/
systemctl --user daemon-reload
systemctl --user enable --now codex-watch.service
```

Check status:

```bash
systemctl --user status codex-watch.service
```

Stop/disable:

```bash
systemctl --user disable --now codex-watch.service
```

## How Detection Works

`codex-watch` monitors `~/.codex/sessions/**/*.jsonl` and triggers notifications when it sees:

- `type = event_msg`
- `payload.type = task_complete`

## Verify

1. Start tmux
2. Run `codex-with-notify`
3. Let Codex finish a turn and wait for your input
4. You should see a notification
5. Click `Go to pane` to jump back

## Troubleshooting

### No notifications

```bash
pgrep dunst
```

```bash
dunstify "Test" "Codex notification test"
```

### No pane switch on click

- Use `codex-with-notify` inside tmux (recommended)
- Ensure a tmux client is attached

### Non-default sessions directory

If your Codex sessions live somewhere other than `~/.codex/sessions`, set:

```bash
export CODEX_SESSIONS_DIR=/path/to/codex/sessions
```

### Watcher debug mode

```bash
codex-watch --verbose
```
