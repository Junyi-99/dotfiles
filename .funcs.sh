#!/bin/bash

function log_message() {
    local message="$1"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $message"
}

# 解压文件的函数
extract() {
    local file="$1"
    if [[ -f $file ]]; then
        case $file in
        *.tar.bz2) tar xjvf "$file" ;;
        *.tar.gz) tar xzvf "$file" ;;
        *.bz2) bunzip2 "$file" ;;
        *.rar) rar x "$file" ;;
        *.gz) gunzip "$file" ;;
        *.tar) tar xvf "$file" ;;
        *.tbz2) tar xjvf "$file" ;;
        *.tgz) tar xzvf "$file" ;;
        *.zip) unzip "$file" ;;
        *.Z) uncompress "$file" ;;
        *) log_message "'$file' cannot be extracted via extract()" ;;
        esac
    else
        log_message "'$file' is not a valid file"
    fi
}

function print_greetings() {
    export FORTUNE_FILE="${XDG_CONFIG_HOME}/fortunes/ubuntu-server-tips"
    if command -v fortune &>/dev/null; then
        if command -v cowsay &>/dev/null; then
            if command -v lolcat &>/dev/null; then
                # fortune | cowsay | lolcat
                cowsay -t "$(fortune)" | lolcat
            else
                cowsay -t "$(fortune)"
            fi
        else
            : # pass
        fi
    else
        : # fortune not installed, skip greeting
    fi
}
