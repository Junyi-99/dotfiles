#!/bin/bash
# display-popup 浮层里输入窗口名并新建窗口。$1 = 起始路径
# Enter 建窗口 · Esc 取消 · Backspace 删字符
start="$1"
name=""

# 主题色（由 apply-theme.sh 写入的 tmux 用户选项，取不到则用回退值）
opt() { tmux show -gv "$1" 2>/dev/null; }
ACCENT=$(opt @popup_accent); ACCENT=${ACCENT:-#bb9af7}
DIM=$(opt @popup_dim);       DIM=${DIM:-#565f89}
FG=$(opt @popup_fg);         FG=${FG:-#a9b1d6}

# #RRGGBB -> 24bit 前景色转义
fg() { local h=${1#\#}; printf '\033[38;2;%d;%d;%dm' "$((16#${h:0:2}))" "$((16#${h:2:2}))" "$((16#${h:4:2}))"; }
R='\033[0m'                       # reset
A=$(fg "$ACCENT"); D=$(fg "$DIM"); T=$(fg "$FG")

W=32                              # 输入框内宽
bar() { printf '─%.0s' $(seq 1 "$W"); }

render() {
  local shown="$name" avail=$((W-4)) fill
  (( ${#shown} > avail )) && shown="…${name: -$((avail-1))}"   # 太长显示尾部
  fill=$((W - 4 - ${#shown}))
  clear
  printf '\n'
  printf "  ${A}New window${R}\n\n"
  printf "  ${D}┌$(bar)┐${R}\n"
  printf "  ${D}│${R} ${A}❯${R} ${T}%s${R}\033[7m \033[0m%*s${D}│${R}\n" "$shown" "$fill" ""
  printf "  ${D}└$(bar)┘${R}\n\n"
  printf "  ${A}⏎${R} ${D}confirm${R}   ${A}^C${R} ${D}clear${R}   ${A}esc${R} ${D}cancel${R}"
}

if [ -t 0 ]; then
  _stty=$(stty -g)                       # 保存终端设置
  stty -isig                             # 让 Ctrl-C 作为字节 ^C 读入，而非发信号
  trap 'stty "$_stty"; printf "\033[?25h"' EXIT   # 退出时恢复终端
fi
printf '\033[?25l'                       # 隐藏真实光标

render
while IFS= read -rsn1 key; do
  case "$key" in
    $'\x1b')                       # ESC —— 裸 Esc 取消，方向键等转义序列忽略
      if ! read -rsn1 -t 0.001 _; then exit 0; fi
      read -rsn3 -t 0.001 _ ;;
    ''|$'\n'|$'\r') break ;;       # Enter → 建窗口
    $'\x03') name=""; render ;;    # Ctrl-C → 清空输入
    $'\x7f'|$'\x08') name="${name%?}"; render ;;   # Backspace
    *) name="$name$key"; render ;;
  esac
done

[ -n "$DRYRUN" ] && exit 0
if [ -n "$name" ]; then
  exec tmux new-window -c "$start" -n "$name"
else
  exec tmux new-window -c "$start"   # 留空用默认名
fi
