# Jira MCP Configuration

The following guide illustrates how to set up and configure the Jira Model Context Protocol (MCP) for various AI tools.

## Claude Code Jira MCP Setup

To setup a Jira MCP for Claude Code (or commonly called Claude CLI), a MCP
configuration entry should be added to the Claude `mcpServers` configuration
section of the `~/.claude.json` file. Once added, Claude Code must be restarted
on the command line. The user must authorize the connection the first time the
MCP connection is used.

The MCP set up in the `~/.claude.json` file for the LF Atlassian Jira instance
would look like:

```jsonc
{
  // ...
  "mcpServers": {
    "mcp-atlassian": {
      "type": "sse",
      "url": "https://mcp.atlassian.com/v1/sse"
    }
  }
}
```

### Claude Code Atlassian Authorization

Once the configuration is complete, relaunch Claude Code by running `claude` in
a terminal window. When the MCP initializes for the first time, it will prompt
you to click on a link to authorize the configuration. This will take you to an
Atlassian configuration page. Ensure the ‘linuxfoundation’ instance is selected,
review the permissions, and approve it. Once approved, the MCP should be
accessible.

### Claude Code Validation

To validate the setup, start Claude Code and issue the `/mcp` command. This
should return details of the configured MCP servers.

Additionally, you can generate a prompt to inquire about a specific Jira ticket.
See the Validation section below for an example.

## Cursor AI Jira MCP Setup

To configure Cursor AI to use the Jira MCP, the MCP server needs to be added to
the global MCP configuration. Unlike Claude Code, Cursor AI does not need a
separate API token. The authorization workflow is handled when Cursor AI is
launched/restarted.

### Cursor Global MCP Configuration

Here is an example of the Cursor AI MCP configuration for the Jira MCP:

Edit the `~/.cursor/mcp.json` file to add the Jira MCP server.

```json
{
  "mcpServers": {
    "mcp-atlassian": {
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.atlassian.com/v1/sse"]
    }
  }
}
```

## Google Gemini Jira MCP Setup

To setup a Jira MCP for Gemini CLI, a MCP configuration entry should be added to
the Claude `mcpServers` configuration section of the `~/.gemini/settings.json`
file. Once added, Gemini CLI must be restarted on the command line. The user
must authorize the connection the first time the MCP connection is used.

The MCP set up in the `~/.gemini/settings.json` file for the LF Atlassian Jira
instance would look like:

```jsonc
{
  // ...
  "mcpServers": {
    "mcp-atlassian": {
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.atlassian.com/v1/sse"]
    }
  }
}
```

A full reference of the Gemini CLI MCP configuration can be found in the
[Gemini CLI MCP Configuration](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/configuration.md)
documentation.

### Gemini CLI Atlassian Authorization

Once the configuration is complete, relaunch Gemini CLI by running `gemini` in a
terminal window. When the MCP initializes for the first time, it will prompt you
to click on a link to authorize the configuration (if not already done). This
will take you to an Atlassian configuration page.  Ensure the ‘linuxfoundation’
instance is selected, review the permissions, and approve it.  Once approved,
the MCP should be accessible.

Optionally, you may need to provide the `cloudId` value for the MCP
configuration. If prompted, enter: 'linuxfoundation.atlassian.net' for the
value.

### Gemini CLI Validation

To validate the setup, start Gemini CLI and issue the `/mcp` command. This
should return details of the configured MCP servers.

Additionally, you can generate a prompt to inquire about a specific Jira ticket.
See the Validation section below for an example.

## Zed

Zed supports agentic editing including tools. Like Cursor, it does not support SSE MCP servers natively, and needs a `stdio->sse` proxy. Use `Cmd-,` to bring up the configuration, and add the following:

```jsonc
{
  // ...
  "context_servers": {
    // ...
    "mcp-atlassian": {
      "source": "custom",
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.atlassian.com/v1/sse"]
    }
  }
}
```
### Zed validation

The agent panel settings show the status of each enabled MCP server (green or red). If the Jira MCP times out, this has been connected to an issue where global `npx` invocations of `mcp-remote` fail on missing dependencies (e.g. `iconv-lite`) both when run via the CLI or in Zed. For a workaround, install mcp-remote globally:

```bash
npm i -g mcp-remote
which mcp-remote
```

If `which` does not return the location of mcp-remote, ensure that the global Node `bin` folder is in your `PATH`.

Then, modify the Zed MCP command and args to use the globally-installed `mcp-remote`:

```jsonc
    "mcp-atlassian": {
      "enabled": true,
      "source": "custom",
      "command": "mcp-remote",
      "args": ["https://mcp.atlassian.com/v1/sse"]
    }
```

## Validation

To validate the setup, start Claude Code or Gemini CLI and issue the `/mcp`
command. This should return details of the configured MCP servers.

Additionally, you can generate a prompt to inquire about a specific Jira ticket.
The AI should query Jira and return a valid response. An example prompt might
be:

```text
What is the current status of the Jira ticket number: ME-36?
```

Claude Code or Gemini CLI may ask for permission to make the MCP Jira request.
You can approve it the first time or choose to approve and not ask again,
allowing the tool to make the request.

Here's an example output from Claude Code:

```code
What is the current status of the Jira ticket: ME-36?

⏺ I'll help you check the status of Jira ticket ME-36. First, let me get your Atlassian
   cloud information and then look up the ticket.
  ⎿  [
       {
         "id": "c497ae62-1596-4157-81fb-96205956683e",
     … +26 lines (ctrl+r to expand)

⏺ Now let me get the details of ticket ME-36:

⏺ mcp-atlassian:getJiraIssue (MCP)(cloudId: "c497ae62-1596-4157-81fb-96205956683e",
                                  issueIdOrKey: "ME-36")
  ⎿ {
      "expand": "renderedFields,names,schema,operations,editmeta,changelog,versionedRep
    resentations,customfield_10010.requestTypePractice",
    … +1498 lines (ctrl+r to expand)


⏺ Status: Released

  Jira ticket ME-36 is currently in Released status and has been resolved. The ticket
  was about adding a link for next steps in RISC-V Individual Membership to ensure
  members complete Schedule A forms. The work was completed on June 18, 2025, with the
  solution deployed to production.
```
