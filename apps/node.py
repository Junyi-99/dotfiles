
if command_exists node; then
    echo "\033[32m\033[0;39m  node.js found"
else
    echo "\033[31m\033[0;39m  node.js not found"
    brew install node
fi
