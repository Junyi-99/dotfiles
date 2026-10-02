package main

import "strings"

// Package describes selection, detection, installation, and shell integration.
// Commands use OMD_HOME; generated shell snippets use USER_CONFIG_HOME.
type Package struct {
	ID          string
	Name        string
	Description string
	Tab         string
	Category    string
	Presets     []string
	Check       string
	Env         string
	Installers  map[string]Installer
}

// Installer pairs shell commands with the prerequisites they need.
type Installer struct {
	Commands []string
	Requires []string
}

func brewInstaller(formula string) Installer {
	return Installer{
		Commands: []string{"brew install " + formula},
		Requires: []string{"brew"},
	}
}

func (p Package) installerFor(platform string) Installer {
	if installer, ok := p.Installers[platform]; ok {
		return installer
	}
	return p.Installers["all"]
}

func registry() []Package {
	allPresets := []string{"minimal", "agent", "full"}
	fullPreset := []string{"full"}
	agentPresets := []string{"agent", "full"}

	return []Package{
		{
			ID:          "zsh-autosuggestions",
			Name:        "zsh-autosuggestions",
			Description: "Suggestions from your shell history",
			Tab:         "Basics",
			Category:    "Shell",
			Presets:     allPresets,
			Check:       `test -f "$OMD_HOME/dep/zsh-autosuggestions/zsh-autosuggestions.zsh"`,
			Installers: map[string]Installer{
				"all": {
					Commands: []string{`git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions "$OMD_HOME/dep/zsh-autosuggestions"`},
					Requires: []string{"git"},
				},
			},
			Env: `source "$USER_CONFIG_HOME/dep/zsh-autosuggestions/zsh-autosuggestions.zsh"
ZSH_AUTOSUGGEST_STRATEGY=(match_prev_cmd completion)
ZSH_AUTOSUGGEST_CLEAR_WIDGETS+=(bracketed-paste)
bindkey '^B' autosuggest-toggle
`,
		},
		{
			ID:          "zsh-syntax-highlighting",
			Name:        "zsh-syntax-highlighting",
			Description: "Highlight commands as you type",
			Tab:         "Basics",
			Category:    "Shell",
			Presets:     allPresets,
			Check:       `test -f "$OMD_HOME/dep/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh"`,
			Installers: map[string]Installer{
				"all": {
					Commands: []string{`git clone --depth=1 https://github.com/zsh-users/zsh-syntax-highlighting "$OMD_HOME/dep/zsh-syntax-highlighting"`},
					Requires: []string{"git"},
				},
			},
			Env: `source "$USER_CONFIG_HOME/dep/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh"
`,
		},
		{
			ID:          "starship",
			Name:        "Starship",
			Description: "A fast, context-aware prompt",
			Tab:         "Basics",
			Category:    "Shell",
			Presets:     fullPreset,
			Check:       "command -v starship",
			Installers: map[string]Installer{
				"darwin": brewInstaller("starship"),
				"linux": {
					Commands: []string{`curl -fsSL https://starship.rs/install.sh | sh -s -- -y -b "$OMD_HOME/bin"`},
					Requires: []string{"curl"},
				},
			},
			Env: `command -v starship >/dev/null && eval "$(starship init zsh)"
`,
		},
		{
			ID:          "homebrew",
			Name:        "Homebrew",
			Description: "Package manager for macOS and Linux",
			Tab:         "Basics",
			Category:    "Package managers",
			Presets:     fullPreset,
			Check:       "command -v brew",
			Installers: map[string]Installer{
				"darwin": {
					Commands: []string{`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`},
					Requires: []string{"curl", "git"},
				},
				"linux": {
					Commands: []string{`mkdir -p "$OMD_HOME/dep/homebrew"
curl -fL https://github.com/Homebrew/brew/tarball/master | tar xz --strip-components=1 -C "$OMD_HOME/dep/homebrew"`},
					Requires: []string{"curl", "git", "tar"},
				},
			},
			Env: `for _omd_brew in /opt/homebrew/bin/brew /usr/local/bin/brew "$USER_CONFIG_HOME/dep/homebrew/bin/brew" /home/linuxbrew/.linuxbrew/bin/brew; do
    if [[ -x "$_omd_brew" ]]; then
        eval "$("$_omd_brew" shellenv)"
        break
    fi
done
unset _omd_brew
`,
		},
		{
			ID:          "pyenv",
			Name:        "pyenv",
			Description: "Manage Python versions",
			Tab:         "Basics",
			Category:    "Package managers",
			Presets:     fullPreset,
			Check:       "command -v pyenv",
			Installers: map[string]Installer{
				"darwin": brewInstaller("pyenv"),
				"linux": {
					Commands: []string{"curl -fsSL https://pyenv.run | bash"},
					Requires: []string{"curl", "git"},
				},
			},
			Env: `export PYENV_ROOT="$HOME/.pyenv"
[[ -d "$PYENV_ROOT/bin" ]] && path=("$PYENV_ROOT/bin" $path)
command -v pyenv >/dev/null && eval "$(pyenv init - zsh)"
`,
		},
		{
			ID:          "neovim",
			Name:        "Neovim",
			Description: "Terminal editor; selection also aliases vim to nvim",
			Tab:         "Basics",
			Category:    "Tools",
			Presets:     fullPreset,
			Check:       "command -v nvim",
			Installers: map[string]Installer{
				"darwin": brewInstaller("neovim"),
				"linux": {
					Commands: []string{`case "$(uname -m)" in
    x86_64) arch=x86_64 ;;
    aarch64|arm64) arch=arm64 ;;
    *) echo 'Unsupported Neovim architecture' >&2; exit 1 ;;
esac
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "https://github.com/neovim/neovim/releases/latest/download/nvim-linux-$arch.tar.gz" -o "$tmp/nvim.tar.gz"
tar -xf "$tmp/nvim.tar.gz" -C "$tmp"
cp -R "$tmp/nvim-linux-$arch/". "$OMD_HOME/"`},
					Requires: []string{"curl", "tar"},
				},
			},
			Env: `alias vim="nvim"
`,
		},
		{
			ID:          "fzf",
			Name:        "fzf",
			Description: "Interactive fuzzy finder",
			Tab:         "Basics",
			Category:    "Tools",
			Presets:     fullPreset,
			Check:       "command -v fzf",
			Installers:  map[string]Installer{"all": brewInstaller("fzf")},
		},
		{
			ID:          "zoxide",
			Name:        "zoxide",
			Description: "Jump to frequently used directories",
			Tab:         "Basics",
			Category:    "Tools",
			Presets:     fullPreset,
			Check:       "command -v zoxide",
			Installers:  map[string]Installer{"all": brewInstaller("zoxide")},
			Env: `command -v zoxide >/dev/null && eval "$(zoxide init zsh)"
`,
		},
		{
			ID:          "batcat",
			Name:        "bat",
			Description: "Read files with syntax highlighting",
			Tab:         "Basics",
			Category:    "Tools",
			Presets:     fullPreset,
			Check:       "command -v bat || command -v batcat",
			Installers:  map[string]Installer{"all": brewInstaller("bat")},
			// Debian and Ubuntu ship the binary as batcat.
			Env: `if command -v bat >/dev/null; then
    alias cat="bat -p"
elif command -v batcat >/dev/null; then
    alias cat="batcat -p"
fi
`,
		},
		{
			ID:          "duf",
			Name:        "duf",
			Description: "Disk usage at a glance",
			Tab:         "Basics",
			Category:    "Tools",
			Presets:     fullPreset,
			Check:       "command -v duf",
			Installers:  map[string]Installer{"all": brewInstaller("duf")},
		},
		{
			ID:          "kitty-config-link",
			Name:        "Kitty config symlink",
			Description: "Link etc/kitty to ~/.config/kitty; backs up an existing config first",
			Tab:         "Basics",
			Category:    "Configuration",
			Check:       `test -L "$HOME/.config/kitty" && test "$(readlink "$HOME/.config/kitty")" = "$OMD_HOME/etc/kitty"`,
		},
		{
			ID:          "claude-code",
			Name:        "Claude Code",
			Description: "Anthropic's command-line assistant",
			Tab:         "CLI Agents",
			Category:    "Agents",
			Presets:     agentPresets,
			Check:       "command -v claude",
			Installers: map[string]Installer{
				"all": {
					Commands: []string{"npm install -g @anthropic-ai/claude-code"},
					Requires: []string{"npm"},
				},
			},
		},
		{
			ID:          "codex-cli",
			Name:        "Codex CLI",
			Description: "OpenAI's command-line assistant",
			Tab:         "CLI Agents",
			Category:    "Agents",
			Presets:     agentPresets,
			Check:       "command -v codex",
			Installers: map[string]Installer{
				"all": {
					Commands: []string{"npm install -g @openai/codex"},
					Requires: []string{"npm"},
				},
			},
		},
		{
			ID:          "apptainer",
			Name:        "Apptainer",
			Description: "Linux container runtime; Linux only",
			Tab:         "Basics",
			Category:    "Tools",
			Check:       `test "$(uname -s)" = Linux && command -v apptainer`,
			Installers: map[string]Installer{
				"darwin": {
					Commands: []string{`echo 'Apptainer can only be installed on Linux.' >&2; exit 1`},
				},
				"linux": brewInstaller("apptainer"),
			},
		},
	}
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
}
