
if command_exists duf; then
    alias df="duf"
    echo "\033[32m\033[0;39m  duf activated"
else
    echo "\033[31m\033[0;39m  duf not found"
    brew install duf
fi
