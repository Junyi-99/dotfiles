


if command_exists nvim; then
    alias vim="nvim"
    echo "\033[32m\033[0;39m  neovim activated"
else
    echo "\033[31m\033[0;39m  neovim not found"
fi
