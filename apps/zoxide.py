
# Zoxide must be placed at the end of the zshrc
if command_exists zoxide; then
    eval "$(zoxide init zsh)"
    alias cd="z"
    echo "\033[32m\033[0;39m  zoxide activated"
else
    echo "\033[31m\033[0;39m  zoxide not found"
    brew install zoxide
fi