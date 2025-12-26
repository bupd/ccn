# Configuration

## Claude Code Hooks

The script uses Claude Code's hook system. Two hook types are relevant:

### Stop Hook

Triggers when Claude Code stops execution and waits for user input.

```json
{
  "hooks": {
    "Stop": [
      {
        "matcher": "",
        "hooks": [
          { "type": "command", "command": "/home/YOUR_USERNAME/.local/bin/claude-notify &" }
        ]
      }
    ]
  }
}
```

### Notification Hook

Triggers when Claude Code sends a notification event.

```json
{
  "hooks": {
    "Notification": [
      {
        "matcher": "",
        "hooks": [
          { "type": "command", "command": "/home/YOUR_USERNAME/.local/bin/claude-notify &" }
        ]
      }
    ]
  }
}
```

## Dunst Configuration

Customize notification appearance in `~/.config/dunst/dunstrc`:

```ini
[claude-code]
    appname = "dunstify"
    summary = "Claude Code*"
    urgency = normal
    background = "#1a1a2e"
    foreground = "#eaeaea"
    frame_color = "#7c3aed"
    timeout = 0
```

### Notification Timeout

The script uses `-t 0` for persistent notifications. To change this, edit the script:

```bash
# Change -t 0 to desired milliseconds
dunstify -A "switch,Go to window" \
    "Claude Code - $TMUX_WINDOW_NAME" \
    "Waiting for your input" \
    -t 10000  # 10 seconds
```

## Script Customization

### Custom Notification Message

Edit the dunstify call in the script:

```bash
ACTION=$(dunstify -A "switch,Go to window" \
    "Claude Code - $TMUX_WINDOW_NAME" \
    "Your custom message here" \
    -t 0)
```

### Add Sound

Add a sound notification:

```bash
# Add before dunstify
paplay /usr/share/sounds/freedesktop/stereo/complete.oga &
```

### Add Icon

```bash
ACTION=$(dunstify -A "switch,Go to window" \
    -i "terminal" \
    "Claude Code - $TMUX_WINDOW_NAME" \
    "Waiting for your input" \
    -t 0)
```
