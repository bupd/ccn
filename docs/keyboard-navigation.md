# Keyboard Navigation

Interact with Claude Code notifications entirely from the keyboard using dunst's built-in shortcuts.

## Default Dunst Keybindings

These work when a notification is visible:

| Key | Action |
|-----|--------|
| `Space` or `Enter` | Invoke default action (switches to tmux pane) |
| `Escape` | Dismiss notification |
| `Shift+Space` | Dismiss all notifications |

## Configure Dunst Keybindings

Add to `~/.config/dunst/dunstrc`:

```ini
[global]
    # Close notification
    close = ctrl+space

    # Close all notifications
    close_all = ctrl+shift+space

    # Show history
    history = ctrl+grave

    # Context menu (shows action buttons)
    context = ctrl+shift+period
```

## Recommended Setup

For keyboard-driven workflow:

```ini
[global]
    # Mouse clicks
    mouse_left_click = do_action, close_current
    mouse_middle_click = close_all
    mouse_right_click = close_current

    # Keyboard shortcuts
    close = ctrl+space
    close_all = ctrl+shift+space
    history = ctrl+grave
    context = ctrl+period
```

## Workflow

1. Notification appears: "Claude Code - session:window"
2. Press `Enter` or `Space` to switch to that tmux pane
3. Or press `Escape` to dismiss

## Using with Window Manager Keybindings

### i3/Sway

```
# Focus notification and trigger action
bindsym $mod+n exec --no-startup-id dunstctl action 0
bindsym $mod+Shift+n exec --no-startup-id dunstctl close-all
```

### Hyprland

```
bind = $mod, N, exec, dunstctl action 0
bind = $mod SHIFT, N, exec, dunstctl close-all
```

## dunstctl Commands

Control notifications programmatically:

```bash
# Trigger action on most recent notification
dunstctl action 0

# Close current notification
dunstctl close

# Close all notifications
dunstctl close-all

# Show notification history
dunstctl history-pop

# Toggle do-not-disturb
dunstctl set-paused toggle
```

## Tips

- Set `timeout = 0` in the script so notifications persist until you act
- Use `dunstctl count` in your status bar to show pending notification count
- Bind `dunstctl action 0` to a convenient key for quick switching
