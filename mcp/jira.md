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

```json
{
  ...
  "mcpServers": {
    "mcp-atlassian": {
      "type": "sse",
      "url": "https://mcp.atlassian.com/v1/sse"
    }
  }
}
```

### Atlassian Authorization

Once the configuration is complete, relaunch Claude Code by running `claude` in
a terminal window. When the MCP initializes for the first time, it will prompt
you to click on a link to authorize the configuration. This will take you to an
Atlassian configuration page. Ensure the ‘linuxfoundation’ instance is selected,
review the permissions, and approve it. Once approved, the MCP should be
accessible.

### Validation

To validate the setup, start Claude Code and issue the `/mcp` command. This
should return details of the configured MCP servers.

Additionally, you can generate a prompt to inquire about a specific Jira ticket.
The AI should query Jira and return a valid response. An example prompt might
be:

```text
What is the current status of the Jira ticket number: ME-36?
```

Claude Code may ask for permission to make the MCP Jira request. You can approve
it the first time or choose to approve and not ask again, allowing Claude Code
to make the request.

Example output:

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
