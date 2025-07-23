# Adding Copilot support in NeoVim

Go to [Neovim GitHub releases](https://github.com/neovim/neovim/releases), find the latest version number,
and replace `<latest_version>` in the command below:

```bash
# Replace <latest_version> with the latest version number from the NeoVim GitHub releases page.
wget https://github.com/neovim/neovim/releases/download/<latest_version>/nvim-linux-x86_64.appimage && \
    chmod +x nvim-linux-x86_64.appimage && \
    sudo cp -iv ./nvim-linux-x86_64.appimage /usr/bin/nvim
```

Confirm that you can run nvim: `nvim`.

Install and configure `git` (I'm assuming that this is already done).

Install GitHub CLI tools (I'm providing example installation just for reference):

```bash
type -p curl >/dev/null || sudo apt install curl -y
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
sudo apt update
sudo apt install gh
gh --version
gh auth login
```

You need to authenticate via browser for Copilot plugin usage and if you're inside an SSH session,
then you need to do it on another computer with browser support.

Then copy `~/.config/gh` and `~/.config/github-copilot` directories from that machine to remote one via `sftp`.

Remember to set correct permissions: `chown -R username ~username/.config`.

Install Copilot extension inside `gh` command line tool: `gh extension install github/gh-copilot`.

Verify that Copilot is installed and working:

```bash
gh extensions list && \
    gh copilot explain 'git rebase -i HEAD~3' && \
    gh copilot suggest "convert all .mov files to .mp4 using ffmpeg"
```

Install lazy plugins support in neovim:

```bash
git clone https://github.com/folke/lazy.nvim.git ~/.local/share/nvim/lazy/lazy.nvim
```

Install Copilot plugin in neovim:

```bash
mkdir -p ~/.config/nvim/lua/plugins && \
    vim ~/.config/nvim/init.lua
```

Put the following content in `~/.config/nvim/init.lua`:

```lua
-- to disable mouse
-- vim.o.mouse = ""
-- To import your favorite vim settings
-- vim.cmd('source ~/.vimrc')
vim.opt.rtp:prepend("~/.local/share/nvim/lazy/lazy.nvim")
require("lazy").setup("plugins")
```

Put the following content in `~/.config/nvim/lua/plugins/copilot.lua`:

```lua
return {
  "github/copilot.vim",
  lazy = false,
  config = function()
    vim.g.copilot_filetypes = {
      ["*"] = true,
    }
    vim.g.copilot_no_tab_map = true
    vim.api.nvim_set_keymap("i", "<C-J>", 'copilot#Accept("<CR>")', { silent = true, expr = true })
    vim.api.nvim_set_keymap("i", "<C-H>", 'copilot#Dismiss()',      { silent = true, expr = true })
    vim.api.nvim_set_keymap("i", "<C-L>", 'copilot#Next()', { silent = true, expr = true })
    vim.api.nvim_set_keymap("i", "<C-K>", 'copilot#Previous()', { silent = true, expr = true })
    vim.api.nvim_set_keymap("i", "<C-Space>", 'copilot#Refresh()',  { silent = true, expr = true })
  end,
}
```

Then start neovim: `nvim` - you may see message that Node is too old, update as needed.

Run `:Copilot status`. It can say `Error: You are not signed into GitHub`, then do `gh auth login && gh auth status`.

Then it should return `Ready`.

Now test that Copilot works in `nvim`: `nvim test.py`.

Start writing some function, for example start typing: `def add(a, b):` and it will suggest something like:

```python
    """Returns the sum of a and b."""
    return a + b
```

In grayed/shadowed text - you can accept via: CTRL+J, reject by CTRL+H, prev/next suggestions via CTRL+L/K,
refresh copilot via CTRL+space, see: `~/.config/nvim/lua/plugins/copilot.lua` for bindings.

More info can be found in the [official documentation](https://github.com/github/copilot.vim)
and via `:Copilot help`/`:help copilot` inside `nvim`.
