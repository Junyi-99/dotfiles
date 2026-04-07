#!/bin/bash
# Auto-detect macOS light/dark mode and apply tmux theme

mode=$(defaults read -g AppleInterfaceStyle 2>/dev/null || echo "Light")

if [ "$mode" = "Dark" ]; then
    # --- Tokyo Night Dark ---
    bg="#1a1b26"
    fg="#a9b1d6"
    accent="#7aa2f7"
    secondary="#bb9af7"
    dim="#565f89"
    border="#3b4261"
    alert="#f7768e"
else
    # --- Tokyo Night Light ---
    bg="#d5d6db"
    fg="#343b58"
    accent="#34548a"
    secondary="#5a4a78"
    dim="#9699a3"
    border="#b4b5b9"
    alert="#c0392b"
fi

tmux set -g status-style "bg=$bg,fg=$fg"
tmux set -g status-left "#[fg=$bg,bg=$accent,bold]  #S #[fg=$accent,bg=$bg] "
tmux set -g status-right "#{?pane_synchronized,#[bg=$alert]#[fg=#ffffff]#[bold] SYNC #[default],}#{?window_zoomed_flag,#[bg=#2ecc71]#[fg=#ffffff]#[bold] ZOOM #[default],}#[fg=$accent,bg=$bg]#[fg=$bg,bg=$accent,bold] #h "
tmux setw -g window-status-format "#{?window_bell_flag,#[fg=#ffffff bg=$alert bold] 🔔 #I:#W #[default bg=$bg],#[fg=$dim] #I:#W }"
tmux setw -g window-status-current-format "#[fg=$bg,bg=$secondary,bold] #I:#W #[fg=$secondary,bg=$bg]"
tmux set -g pane-border-style "fg=$border"
tmux set -g pane-active-border-style "fg=$accent"
tmux set -g message-style "fg=$accent,bg=$bg,bold"
tmux set -g message-command-style "fg=$secondary,bg=$bg"
tmux setw -g mode-style "fg=$bg,bg=$secondary,bold"
tmux set -g menu-style "default"
tmux set -g menu-selected-style "bg=$accent,fg=$bg,bold"
tmux set -g menu-border-style "fg=$dim"
tmux set -g menu-border-lines rounded
