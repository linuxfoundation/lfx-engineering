# Snowflake DBT MCP

The following guide illustrates how to set up and configure the dbt Model
Context Protocol (MCP) for connecting to Snowflake.

## Claude Code Setup

```json
{
  ...
  "mcpServers": {
    "snowflake-dbt": {
      "command": "uvx",
      "args": [
        "--env-file",
        "/Users/ddeal/projects/go/src/github.com/dbt-labs/dbt-mcp/.env",
        "dbt-mcp",
      ]
    }
  }
}
```

Setup guide is available [on the dbt-mcp website](https://github.com/dbt-labs/dbt-mcp).
