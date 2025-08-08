# Claude Context MCP

## Overview

The Claude Context Model Context Protocol (MCP) server provides semantic code search
capabilities for AI development tools. It enables searching code by meaning rather than
exact text matches, using vector embeddings to find conceptually related code patterns,
implementations, and logic.

## Features

- Semantic code search using vector embeddings
- Support for local (Ollama) or cloud (OpenAI) embedding models
- Integration with Milvus vector database (local or Zilliz Cloud)
- Intelligent code chunking and indexing
- Cross-file conceptual pattern matching

## Requirements

- Node.js 18 or newer
- Claude Code, Cursor, or any other MCP client
- For local setup:
  - Docker for running Milvus
  - Ollama for local embeddings
- For cloud setup:
  - Zilliz Cloud account (optional, $100 free credits)
  - OpenAI API key (optional)

## Local Prerequisites Setup

### Install Milvus Vector Database

Milvus is an open source vector database (LF AI & Data Foundation graduated project) required for storing code embeddings.

Follow the [official Milvus standalone Docker instructions](https://milvus.io/docs/install_standalone-docker.md).

**Note**: Local Milvus can be used with either Ollama local embeddings (free) or
OpenAI cloud embeddings (paid). Choose based on your needs for cost vs. quality.

### Optional: Attu Web UI for Milvus

[Attu](https://github.com/zilliztech/attu) provides a web interface to visualize your
indexed code embeddings and collections in Milvus:

```bash
# Run Attu alongside local Milvus
docker run -p 8000:3000 \
  -e MILVUS_URL=host.docker.internal:19530 \
  zilliz/attu:latest
```

Access Attu at `http://localhost:8000` to:

- View indexed collections and their embeddings
- Monitor indexing progress
- Debug search queries
- Manage collections and data

### Install Ollama for Local Embeddings

Ollama enables free, local embedding generation without API costs.

```bash
# Install Ollama from https://ollama.ai

# Pull the recommended embedding model
ollama pull nomic-embed-text:v1.5

# Verify installation
ollama list
```

Expected output:

```text
NAME                    ID              SIZE      MODIFIED
nomic-embed-text:v1.5   0a109f422b47    274 MB    2 hours ago
```

## Claude Code Setup

### Project-Level Configuration

Create `.mcp.json` in your project root:

```jsonc
{
  "mcpServers": {
    "claude-context": {
      "type": "stdio",
      "command": "npx",
      "args": ["@zilliz/claude-context-mcp@latest"],
      "env": {}
    }
  }
}
```

Create `.envrc` in your project root for local setup:

```bash
#!/bin/bash

# Use Ollama for embeddings (free, local)
export EMBEDDING_PROVIDER=Ollama
export EMBEDDING_MODEL=nomic-embed-text:v1.5

# Point to local Milvus
export MILVUS_ADDRESS=http://localhost:19530

# Optional: Configure batch size for smaller local model context sizes
export EMBEDDING_BATCH_SIZE=10
```

If using direnv:

```bash
direnv allow
```

### Global Configuration

Alternatively, add to `~/.claude.json`:

```jsonc
{
  // ...
  "mcpServers": {
    "claude-context": {
      "command": "npx",
      "args": ["-y", "@zilliz/claude-context-mcp@latest"],
      "env": {
        "EMBEDDING_PROVIDER": "Ollama",
        "EMBEDDING_MODEL": "nomic-embed-text:v1.5",
        "MILVUS_ADDRESS": "http://localhost:19530"
      }
    }
  }
}
```

### Validation

After restarting your MCP client

1. Verify MCP connection:

   ```text
   /mcp
   ```

   Expected output:

   ```text
   > /mcp
       (no content)

    Manage MCP servers

     1. claude-context   connected · Enter to view details

    MCP Config locations (by scope):
     • User config (available in all your projects):
       • /Users/jld/.claude.json
     • Project config (shared via .mcp.json):
       • /Users/jld/Repositories/LF/linuxfoundation/lfx-engineering/.mcp.json
     • Local config (private to you in this project):
       • /Users/jld/.claude.json [project: /Users/jld/Repositories/LF/linuxfoundation/lfx-engineering]

    For help configuring MCP servers, see: https://docs.anthropic.com/en/docs/claude-code/mcp

      Esc to exit
   ```

## Cloud Configuration Options

### Using Zilliz Cloud

Instead of running Milvus locally, use the managed Zilliz Cloud service:

1. Sign up at [Zilliz Cloud](https://zilliz.com/cloud) ($100 free credits)
2. Create a cluster and obtain credentials
3. Update environment configuration:

```bash
# Replace local Milvus address with Zilliz Cloud
export MILVUS_TOKEN=your-api-key
```

### Using OpenAI Embeddings

For potentially higher quality embeddings (with costs), you can use OpenAI with either local Milvus or Zilliz Cloud:

```bash
# Use OpenAI with local Milvus
export EMBEDDING_PROVIDER=OpenAI
export OPENAI_API_KEY=your-api-key
export EMBEDDING_MODEL=text-embedding-3-small  # or text-embedding-3-large
export MILVUS_ADDRESS=http://localhost:19530  # Local Milvus

# Or use OpenAI with Zilliz Cloud
export EMBEDDING_PROVIDER=OpenAI
export OPENAI_API_KEY=your-api-key
export EMBEDDING_MODEL=text-embedding-3-small
export MILVUS_TOKEN=your-zilliz-api-key  # Cloud Milvus
```

## Configuration

### Embedding Model Dimensions

**Critical**: Claude Context expects specific embedding dimensions. Mismatched
dimensions cause silent failures, this is applicable if you venture to use different
embedding models with something like LM studio and OpenAI Like API. You'll need to
ensure the dimensions match the hardcoded dimensions in the codebase.

## Usage Examples

### Finding MCP Server Configurations

1. Index a codebase:

   ```text
   Index your $CWD

   Expected output:

   ```text
    I'll index the current working directory to enable semantic search capabilities.

    claude-context - index_codebase (MCP)(path: "/Users/jld/Repositories/LF/linuxfoundation/lfx-engineering", splitter: "ast")


    Tool use

      claude-context - index_codebase(path:
        "/Users/jld/Repositories/LF/linuxfoundation/lfx-engineering", splitter: "ast") (MCP)

      Index a codebase directory to enable semantic search using a configurable code splitter.

       **IMPORTANT**:
      - You MUST provide an absolute path to the target codebase.

       **Usage Guidance**:
      - This tool is typically used when search fails due to an unindexed codebase.
      - If indexing is attempted on an already indexed path, and a conflict is detected,
        you MUST prompt the user to confirm whether to proceed with a force index
        (i.e., re-indexing and overwriting the previous index).


    Do you want to proceed?
     1. Yes
      2. Yes, and don't ask again for claude-context - index_codebase commands in
         /Users/jld/Repositories/LF/linuxfoundation/lfx-engineering
      3. No, and tell Claude what to do differently (esc)


   ```

2. Search the codebase:

    ```text
    what MCP servers I can use with claude code
    ```

    Expected output shows documented MCP servers:

    ```text
    Available MCP Servers

    1. Jira (Atlassian) - For Jira ticket management
      - SSE endpoint: https://mcp.atlassian.com/v1/sse
      - No API token needed (uses Atlassian authorization)
    2. GitHub - For repository, issue, and PR operations
      - Package: @modelcontextprotocol/server-github
      - Requires GitHub Personal Access Token
    3. Playwright - For browser automation
      - Package: @playwright/mcp
      - No authentication required
    4. Snowflake DBT - For dbt/Snowflake data operations
      - Package: dbt-mcp
      - Requires environment configuration
    ```

## References

1. [Claude Context MCP GitHub](https://github.com/zilliztech/claude-context)
2. [Milvus Documentation](https://milvus.io/docs) - LF AI & Data Foundation project
3. [Zilliz Cloud](https://zilliz.com/cloud) - Managed Milvus service
4. [Ollama](https://ollama.ai) - Local LLM and embedding models
5. [MCP Documentation](https://modelcontextprotocol.io/docs)
