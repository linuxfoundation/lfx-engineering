# GitHub MCP Server

## Overview

The GitHub MCP Server provides an MCP interface for the GitHub API. Users can
use this server to interact with GitHub repositories, issues, pull requests, and
more through AI-powered tools like Claude Code and Cursor AI.

> **Note:** The npm package `@modelcontextprotocol/server-github` is
> **deprecated as of April 2025**. Follow the updated instructions below to use
> the official GitHub MCP server (`ghcr.io/github/github-mcp-server` via Docker
> or the remote hosted endpoint).

## GitHub MCP Personal Access Token Setup

To use the GitHub MCP, you need to create a personal access token. GitHub users
can create a personal access token [by navigating to the GitHub settings
page](https://github.com/settings/tokens).

Once on the Personal Access Tokens page, click on 'Generate new token'.

As the first step, provide a descriptive name for your token, such as
'GitHub MCP Token'.

Then, select the required scopes for the token:

- For full access to private repositories, select `repo`.
- For access to public repositories only, select `public_repo`.
- Scroll down and click on Generate token.

Copy the generated token and save it securely. At TLF, we recommend storing this
value into 1Password in your personal vault. This will be used in the various
tooling configurations.

As a last step, users should select the 'Configure SSO' option adjacent to the
newly generated token. From the drop-down, select 'linuxfoundation' to authorize
the token for use under the Linux Foundation GitHub organization.

## Claude Code GitHub MCP Setup

> Note: Before proceeding, ensure that the Personal Access Token has been
> configured for use under the Linux Foundation GitHub organization. See above.

Claude Code 2.1.1+ supports connecting directly to the hosted GitHub MCP remote
server over HTTP. This is the recommended approach as it requires no local
Docker installation.

### Option 1: Remote Server (Recommended, Claude Code 2.1.1+)

> **Prerequisite:** The remote endpoint (`https://api.githubcopilot.com/mcp`)
> requires an active **GitHub Copilot** subscription (Individual, Business, or
> Enterprise). If your account does not have Copilot access, authentication will
> fail with a 401 or 403 error. Use Option 2 or Option 3 instead.

Run the following command, replacing `YOUR_GITHUB_PAT` with your personal
access token:

```bash
claude mcp add-json github '{"type":"http","url":"https://api.githubcopilot.com/mcp","headers":{"Authorization":"Bearer YOUR_GITHUB_PAT"}}'
```

Use the `--scope` flag to control where the configuration is stored:

- `--scope local` — current project only (default)
- `--scope project` — shared with the team via `.mcp.json`
- `--scope user` — all projects on this machine

> **Security warning:** Do **not** store personal access tokens or secrets in
> `--scope project` configs (`.mcp.json`). Project-scoped files are typically
> committed to version control and will expose your token to everyone with
> repository access. Prefer `--scope user` or `--scope local` for tokens.
> Use environment variables or a secret manager rather than embedding
> `Authorization` headers with real tokens in JSON. If a project-scoped config
> is unavoidable, add `.mcp.json` to `.gitignore` and use a placeholder (e.g.,
> `YOUR_GITHUB_PAT`) in any committed examples.

Example with user scope:

```bash
claude mcp add-json --scope user github '{"type":"http","url":"https://api.githubcopilot.com/mcp","headers":{"Authorization":"Bearer YOUR_GITHUB_PAT"}}'
```

> Note: if you have an existing GitHub MCP installed and configured, you will need
> to disable and remove the old MCP before configuring the new MCP.

### Option 2: Local Docker Deployment

If you prefer to run the server locally, ensure Docker is installed and running.
There are two ways to supply your token:

**Option A — inline value** (replace `YOUR_GITHUB_PAT` with your actual token):

```bash
claude mcp add github -e GITHUB_PERSONAL_ACCESS_TOKEN=YOUR_GITHUB_PAT -- \
  docker run -i --rm -e GITHUB_PERSONAL_ACCESS_TOKEN ghcr.io/github/github-mcp-server
```

**Option B — inherit from shell** (export the variable first, then omit the
value so Docker inherits it from the parent environment):

```bash
export GITHUB_PERSONAL_ACCESS_TOKEN=YOUR_GITHUB_PAT
claude mcp add github -e GITHUB_PERSONAL_ACCESS_TOKEN -- \
  docker run -i --rm -e GITHUB_PERSONAL_ACCESS_TOKEN ghcr.io/github/github-mcp-server
```

In Option A, `-e GITHUB_PERSONAL_ACCESS_TOKEN=VALUE` passes the token directly
in the `claude mcp` command. In Option B, `-e GITHUB_PERSONAL_ACCESS_TOKEN`
(without `=VALUE`) tells Docker to inherit the variable from your shell
environment, keeping the token out of your shell history and command line.

> **Note:** For Option B, `GITHUB_PERSONAL_ACCESS_TOKEN` must be exported (not
> just set) in your shell so Docker can inherit it as an environment variable.
> A variable set with `VAR=value` (without `export`) is only available to the
> current shell and will not be passed to child processes like Docker.

### Option 3: Binary (No Docker)

Download a release binary from the
[github-mcp-server releases page](https://github.com/github/github-mcp-server/releases),
add it to your `PATH`, then run:

```bash
claude mcp add-json github '{"command":"github-mcp-server","args":["stdio"],"env":{"GITHUB_PERSONAL_ACCESS_TOKEN":"YOUR_GITHUB_PAT"}}'
```

### Claude Code Validation

To validate the setup, run:

```bash
claude mcp list
claude mcp get github
```

You can also start Claude Code and issue the `/mcp` command. The GitHub MCP
should show as connected.

Additionally, you can generate a prompt to inquire about a specific GitHub
pull request, issue, or repository. See the [Validation section below](#validation)
for an example.

## Cursor AI GitHub MCP Setup

To configure Cursor AI to use the GitHub MCP, add the MCP server to the global
MCP configuration.

### Cursor Global GitHub MCP Configuration

> Note: Before proceeding, ensure that the Personal Access Token has been
> configured for use under the Linux Foundation GitHub organization. See above.

Edit the `~/.cursor/mcp.json` file to add the GitHub MCP server using the
official Docker image:

```jsonc
{
  // ...
  "mcpServers": {
    "github": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "GITHUB_PERSONAL_ACCESS_TOKEN",
        "ghcr.io/github/github-mcp-server",
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "YOUR_GITHUB_PAT",
      },
    },
  },
}
```

### Cursor MCP Setup Validation

Once this is added, users can use the GitHub MCP by viewing the MCP server
configuration under the Cursor AI MCP settings page (View: Open MCP Settings).
Under the MCP Tools section, the GitHub MCP should be listed and enabled with a
green dot next to the entry.

Additionally, you can generate a prompt to inquire about a specific GitHub
pull request, issue, or repository. See the Validation section below for an
example. Hit `CMD+I` to open the AI assistant within Cursor and enter a prompt
to query GitHub (see [Validation section below](#validation)).

## Validation

To validate the MCP configuration, ask the AI tool to perform a simple GitHub
operation, such as listing repositories or creating a new issue. If everything
is set up correctly, the AI tool should be able to return the proper results.

Example prompts:

```text
How many open pull requests are on the lfx-engineering repository?
```

This should return the number of open pull requests on the lfx-engineering
repository.

## Troubleshooting

**Authentication Issues:**

- Verify the token includes the `repo` scope.
- Confirm the token has not expired.
- Ensure the token is authorized for the `linuxfoundation` GitHub organization
  via SSO (Configure SSO → Authorize).

**Remote Server Problems:**

- Validate the endpoint URL: `https://api.githubcopilot.com/mcp`

**Docker Issues:**

- Confirm Docker Desktop is running.
- Pull the image manually: `docker pull ghcr.io/github/github-mcp-server`
- Log out from GitHub Container Registry (clears stored credentials): `docker logout ghcr.io`

**Server Startup Failures:**

- Run `claude mcp list` to review your configuration.
- Check JSON syntax validity.
- Remove and reconfigure: `claude mcp remove github`, then re-add.
- Run the `/mcp` command inside Claude Code to review server status and logs.

## References

1. [GitHub MCP Server — Official Installation Guide for
   Claude](https://github.com/github/github-mcp-server/blob/main/docs/installation-guides/install-claude.md)
2. [GitHub MCP Server Repository](https://github.com/github/github-mcp-server)
