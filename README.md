# Dotfiles

> [!TIP]
> There isn’t a one-size-fits-all configuration file. You should explore dotfiles shared by others and continuously build your own personalized dotfile.

<img width="1005" alt="image" src="https://github.com/user-attachments/assets/36147a29-6c22-40e1-85b4-875cfa773b63">

<img width="1098" alt="image" src="https://github.com/user-attachments/assets/30b98735-9bf8-475d-acff-6f7716a3c16f">

## Features
- **out-of-the-box**
- **minimal side effects**
  - most operations happen in `~/.oh-my-dotfiles`
  - tmux config is symlinked to `~/.config/tmux/`
  - third-party binaries will be downloaded into `~/.oh-my-dotfiles/bin` or `~/.oh-my-dotfiles/deps` (See [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/index.html))
- neovim
  - AI assistant ([copilot](https://github.com/github/copilot.vim))
  - lsp + lsp-manager ([mason](https://github.com/williamboman/mason.nvim))
  - plugin-manager ([lazy-nvim](https://github.com/folke/lazy.nvim))
  - notification-manager ([nvim-notify](https://github.com/rcarriga/nvim-notify))
  - code highlighting ([nvim-treesitter](https://github.com/nvim-treesitter/nvim-treesitter))
  - shortcut key hint ([which-key](https://github.com/folke/which-key.nvim))
- [batcat](https://github.com/sharkdp/bat)
- [tmux](https://github.com/tmux/tmux)
  - True color (256color + RGB)
  - <kbd>CTRL + A</kbd> Prefix
  - Mouse enabled
  - Vim-style pane navigation (<kbd>h</kbd><kbd>j</kbd><kbd>k</kbd><kbd>l</kbd>) and resizing (<kbd>H</kbd><kbd>J</kbd><kbd>K</kbd><kbd>L</kbd>)
  - Intuitive split keys: <kbd>|</kbd> horizontal, <kbd>-</kbd> vertical
  - Pane sync toggle (<kbd>e</kbd>)
  - Vi copy mode with system clipboard integration
  - Auto dark/light theme following macOS appearance
- [fzf](https://github.com/junegunn/fzf)
- [zoxide](https://github.com/ajeetdsouza/zoxide)
- [zsh](https://www.zsh.org/) + [oh-my-zsh](https://github.com/ohmyzsh/ohmyzsh) + [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions) + [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)

## Get Started

1. Delete your `.zshrc` (Optional)
```sh
rm -f ~/.zshrc
```

2. Clone the latest repository
```sh
rm -rf ~/.oh-my-dotfiles
git clone https://github.com/Junyi-99/dotfiles.git ~/.oh-my-dotfiles
echo "source ~/.oh-my-dotfiles/etc/zsh/zshenv" >> ~/.zshenv
echo "source ~/.oh-my-dotfiles/etc/zsh/zshrc" >> ~/.zshrc
```

3. Set up tmux config
```sh
mkdir -p ~/.config/tmux
ln -sf ~/.oh-my-dotfiles/etc/tmux/tmux.conf ~/.config/tmux/tmux.conf
ln -sf ~/.oh-my-dotfiles/etc/tmux/apply-theme.sh ~/.config/tmux/apply-theme.sh
```

4. Setting up SSH configs
```sh
echo "Include ~/.oh-my-dotfiles/etc/ssh/*.local" >> ~/.ssh/config
```

5. Activate!
```sh
source ~/.zshrc
```
