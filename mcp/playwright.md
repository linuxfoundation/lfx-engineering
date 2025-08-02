# Playwright MCP

The Playwright Model Context Protocol (MCP) server provides browser automation
capabilities using Playwright. The MCP enables LLMs to interact with web pages
through structured accessibility snapshots, bypassing the need for screenshots
or visually-tuned models.

## Features

- Accessible snapshots of web pages
- Browser automation
- Structured data output
- No need for screenshots or visually-tuned models

## Requirements

- Node.js 18 or newer
- VS Code, Cursor, Windsurf, Claude Desktop, Goose or any other MCP client

## Setup

Standard config works in most of the tools:

```jsonc
{
  // Other configurations...

  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

The MCP server [GitHub repository](https://github.com/microsoft/playwright-mcp)
includes additional instructions for adding the MCP to various code editors and
AI tools.

## Configuration

Configuration for the Playwright MCP service can be found on [the GitHub
repository README
documentation](https://github.com/microsoft/playwright-mcp?tab=readme-ov-file#configuration).

## References

1. [Playwright MCP](https://github.com/microsoft/playwright-mcp)
