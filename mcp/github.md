# GitHub MCP Server

## Overview

The GitHub MCP Server is a server that provides a MCP interface for the GitHub API. It is a simple server that can be
used to interact with the GitHub API. Users can use this server to get information about their GitHub repositories,
issues, pull requests, and more.

## GitHub MCP Personal Access Token Setup

To use the GitHub MCP, you need to create a personal access token. GitHub users can create a personal access token
[by navigating to the GitHub settings page](https://github.com/settings/tokens).

Once on the Personal Access Tokens page, click on 'Generate new token'.

As the first step, provide a descriptive name for your token, such as 'GitHub MCP Token'.

Then, select the required scopes for the token:

- For full access to private repositories, select repo.
- For access to public repositories only, select public_repo.
- Scroll down and click on Generate token.

Copy the generated token and save it securely. This will be used in the various tooling configurations.

As a last step, users should select the 'Configure SSO' option adjacent to the newly generated token.
From the drop-down, select the 'linuxfoundation' to Authorize the token for use under the Linux
Foundation GitHub organization.

## Claude Code GitHub MCP Setup

> Note: Before proceeding, ensure that the Personal Access Token has been
> configured for use under the Linux Foundation GitHub organization. See above.

To setup the GitHub MCP for Claude Code (or commonly called Claude CLI), a MCP
configuration entry should be added to the Claude `mcpServers` configuration
section of the `~/.claude.json` file. Once added, Claude Code must be restarted
on the command line. The user must authorize the connection the first time the
MCP connection is used.

The MCP set up in the `~/.claude.json` file for the LF Atlassian Jira instance
would look like:

```jsonc
{
  //...
  "mcpServers": {
    "github": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-github"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "YOUR_PERSONAL_ACCESS_TOKEN"
      }
    }
  }
}
```

Once this is added, users can use the GitHub MCP by restarting Claude Code.

### Claude Code Validation

To validate the setup, start Claude Code and issue the `/mcp` command. This
should return details of the configured MCP servers. The GitHub MCP should
show as being connected.

Additionally, you can generate a prompt to inquire about a specific GitHub
pull request, issue, or repository. See [validation section below](#validation)for an
example.

## Cursor AI GitHub MCP Setup

To configure Cursor AI to use the GitHub MCP, the MCP server needs to be added to
the global MCP configuration.

### Cursor Global GitHub MCP Configuration

> Note: Before proceeding, ensure that the Personal Access Token has been
> configured for use under the Linux Foundation GitHub organization. See above.

Here is an example of the Cursor AI MCP configuration for the GitHub MCP:

Edit the `~/.cursor/mcp.json` file to add the GitHub MCP server.

```jsonc
{
  // ...
  "mcpServers": {
    "github": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-github"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "YOUR_PERSONAL_ACCESS_TOKEN"
      }
    }
  }
}
```

### Cursor MCP Setup Validation

Once this is added, users can use the GitHub MCP by viewing the MCP server configuration under
the Cursor AI MCP settings page (View: Open MCP Settings). Under MCP Tools section, the GitHub MCP
should be listed and enabled. There should be a green dot next to the GitHub MCP entry.

Additionally, you can generate a prompt to inquire about a specific GitHub
pull request, issue, or repository. See the Validation section below for an
example. Hit '<CMD> + i' to open the AI assistant within Cursor and enter a prompt to query GitHub
(see [Validation section below](#validation)).

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

## References

1. [How to set up Github MCP server for use with Claude Desktop on Windows and
   Mac](https://allthings.how/how-to-set-up-github-mcp-server-for-use-with-claude-desktop-on-windows-and-mac/)
