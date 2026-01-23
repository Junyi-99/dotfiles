if command_exists fzf; then
    source <(fzf --zsh)
    echo "\033[32m\033[0;39m  fzf was found and activated"
else
    echo "\033[31m\033[0;39m  fzf not found"
    brew install fzf
fi