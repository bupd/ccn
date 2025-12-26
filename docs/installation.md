# Installation

## Prerequisites

Install the required dependencies:

```bash
# Arch Linux
sudo pacman -S dunst tmux

# Ubuntu/Debian
sudo apt install dunst tmux

# Fedora
sudo dnf install dunst tmux
```

Make sure dunst is running:

```bash
dunst &
```

Or add it to your window manager's autostart.

## Install the Script

```bash
# Clone the repository
git clone https://github.com/bupd/ccn.git
cd ccn

# Copy the script to your local bin
cp scripts/claude-notify ~/.local/bin/
chmod +x ~/.local/bin/claude-notify
```

## Configure Claude Code Hooks

Add the following to your `~/.claude/settings.json`:

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
    ],
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

Replace `YOUR_USERNAME` with your actual username.

## Verify Installation

1. Start a Claude Code session inside tmux
2. Run a command that makes Claude stop (e.g., ask it to do something)
3. You should see a desktop notification

## Troubleshooting

### No notification appears

Check if dunst is running:
```bash
pgrep dunst
```

Test dunstify directly:
```bash
dunstify "Test" "This is a test notification"
```

### Action button doesn't switch panes

Make sure you're running Claude Code inside tmux. The script needs tmux context to capture session info.
