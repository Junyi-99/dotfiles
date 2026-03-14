#!/bin/bash
# 检测 macOS 系统主题并应用对应 tmux 配色

if defaults read -g AppleInterfaceStyle &>/dev/null; then
    MODE="dark"
else
    MODE="light"
fi

if [ "$MODE" = "dark" ]; then
    # ── 深色主题 ──
    BG="#1a1b26"           # 深蓝黑背景
    FG="#a9b1d6"           # 柔和前景
    ACCENT="#7aa2f7"       # 蓝色强调
    ACCENT_FG="#1a1b26"    # 强调色上的文字
    DIM="#3b4261"          # 暗淡色
    BORDER="#3b4261"
    ACTIVE_BORDER="#7aa2f7"
    MSG_BG="#1a1b26"
    MSG_FG="#7aa2f7"
else
    # ── 浅色主题 ──
    BG="#e1e2e7"           # 浅灰背景
    FG="#3760bf"           # 深蓝前景
    ACCENT="#2e7de9"       # 蓝色强调
    ACCENT_FG="#e1e2e7"    # 强调色上的文字
    DIM="#8990b3"          # 暗淡色
    BORDER="#c4c8da"
    ACTIVE_BORDER="#2e7de9"
    MSG_BG="#e1e2e7"
    MSG_FG="#2e7de9"
fi

# 应用配色
tmux set -g status-style "bg=$BG,fg=$FG"
tmux set -g status-left-style "bg=$BG,fg=$ACCENT,bold"
tmux set -g status-right-style "bg=$BG,fg=$DIM"
tmux setw -g window-status-style "bg=$BG,fg=$DIM"
tmux setw -g window-status-current-style "bg=$ACCENT,fg=$ACCENT_FG,bold"
tmux set -g pane-border-style "fg=$BORDER"
tmux set -g pane-active-border-style "fg=$ACTIVE_BORDER"
tmux set -g message-style "bg=$MSG_BG,fg=$MSG_FG"
tmux set -g message-command-style "bg=$MSG_BG,fg=$MSG_FG"
tmux setw -g mode-style "bg=$ACCENT,fg=$ACCENT_FG"
