if command_exists bat; then
    alias cat="bat -p --theme=ansi"
    echo "\033[32m\033[0;39m  bat activated"
else
    if command_exists batcat; then
        alias cat="batcat -p --theme=ansi"
        echo "\033[32m\033[0;39m  batcat activated"
    else
        echo "\033[31m\033[0;39m  batcat not found"
        brew install bat
    fi
fi