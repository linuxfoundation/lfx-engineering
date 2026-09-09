# CodeRabbit CLI Setup

## Overview

[CodeRabbit CLI](https://www.coderabbit.ai/cli) is a command-line tool
that provides AI-powered code reviews locally before pushing changes.
It analyzes your code changes against the base branch and provides
detailed feedback, including potential bugs, security issues, and
improvement suggestions. Running reviews locally helps catch issues
early, reducing review cycles and improving code quality.

## LFX Organization Note

CodeRabbit is already enabled on some (not all) LFX GitHub Enterprise
repositories by the CloudOps team. Team members should reach out to the
LFX platform engineering team to enable CodeRabbit on their repository
pull request reviews.

- Contact via the **#lf-ai** Slack channel or tag a platform team member in the #lfx-devops Slack channel

Administrators can access the CodeRabbit dashboard at
[app.coderabbit.ai](https://app.coderabbit.ai/settings/repositories)
using the GitHub login option.

## Installation

### Homebrew (macOS only)

> Requires macOS 10.15 (Catalina) or later.

```bash
brew install --cask coderabbit
```

For more details, see the
[Homebrew formula page](https://formulae.brew.sh/cask/coderabbit).

### Install Script (macOS and Linux)

> This is the cross-platform option for macOS and Linux systems.

```bash
curl -fsSL https://cli.coderabbit.ai/install.sh | sh
```

After installation, confirm that `cr` is available in your `PATH`
or reload your shell. The full command `coderabbit` also works,
but `cr` is the shorthand used throughout this guide:

```bash
source ~/.zshrc
```

## Authentication

Run the following command to authenticate:

```bash
cr auth login
```

Follow the browser redirect to sign in. An access token will be
generated and shown on the screen after authentication. Copy the
access token from the web page and paste it into the terminal window
CLI when prompted.

## Basic Usage

The CLI command is `cr` (short alias for `coderabbit`). Run it from
within a Git repository to review your changes against the base branch.

- `cr` — Interactive mode with full browsable findings interface
- `cr --plain` — Plain text output with detailed fix suggestions
- `cr --prompt-only` — Minimal output optimized for AI agents
- `cr --base develop` — Specify a base branch other than main

## IDE Integration: VSCode

Install the **CodeRabbit** extension from the
[VS Code marketplace](https://docs.coderabbit.ai/ide/vscode-install).
Search for "CodeRabbit" in Extensions. The extension reviews every
commit in the IDE and provides one-click fixes.

## IDE Integration: Cursor

The same CodeRabbit extension works in
[Cursor](https://docs.coderabbit.ai/cli/cursor-integration)
(a VS Code fork). Search for "CodeRabbit" in the Cursor extensions
marketplace to install it.

## Using with Anthropic CoWork

[Anthropic CoWork](https://docs.coderabbit.ai/cli/claude-code-integration)
is Anthropic's Claude-based knowledge worker tool. Use the CodeRabbit
CLI within CoWork by running `cr --prompt-only` and having the agent
evaluate and fix findings.

**Example workflow:**

1. Ask the agent to run the review:
   "Run cr --prompt-only in the background, let it take as long as
   it needs, and check on it periodically."
2. Have the agent evaluate the results:
   "Evaluate the fixes and considerations. Fix major issues only."
3. Run a follow-up review to confirm:
   "Run cr one more time to make sure we addressed all critical
   issues."

## Using with Claude Code

Claude Code has native plugin support for CodeRabbit. Use the
`/coderabbit:review` command directly within Claude Code to trigger
a review. For more details, see the
[Claude Code integration guide](https://docs.coderabbit.ai/cli/claude-code-integration).

## Multi-Repo Analysis

CodeRabbit can detect breaking changes, API mismatches, and dependency
issues that cross repository boundaries during code review. This is
useful when your codebase is split across multiple repositories that
share contracts or dependencies.

### Use Cases

- **API contracts** — When a backend API changes, frontend or mobile
  repositories may need coordinated updates.
- **Microservices** — A change to a service's REST API may break
  consumers in other repositories.
- **Shared libraries** — Modifications to a shared utility or type
  definition can have ripple effects across multiple repositories.
- **Database schemas** — Schema changes can affect all services that
  query the same data model.

### Configuration

Add a `linked_repositories` section under `knowledge_base` in your
`.coderabbit.yaml` file:

```yaml
# yaml-language-server: $schema=https://coderabbit.ai/integrations/schema.v2.json
knowledge_base:
  linked_repositories:
    - repository: "myorg/backend-api"
      instructions: "Contains REST API endpoints and database models"
```

You can also configure linked repositories through the CodeRabbit web
interface at
[app.coderabbit.ai](https://app.coderabbit.ai/settings/repositories)
under **Knowledge Base** > **Linked Repositories**.

> **Note:** Each repository configuration currently supports one
> linked repository. The CodeRabbit GitHub App must be installed on
> all linked repositories. Contact the platform engineering team via
> **#lf-ai** Slack if you need help setting this up.

For full details, see the
[Multi-Repo Analysis documentation](https://docs.coderabbit.ai/knowledge-base/multi-repo-analysis).

## Uninstall

### Homebrew

```bash
brew remove coderabbit
```

### Manual Removal (Install Script)

Remove both the `cr` shorthand and the `coderabbit` binary if they
exist:

```bash
rm -f "$(command -v cr)" "$(command -v coderabbit)"
```

## Resources

- [CodeRabbit CLI Documentation](https://docs.coderabbit.ai/cli/)
- [CodeRabbit Website](https://www.coderabbit.ai/cli)
- [CodeRabbit Discord Community](https://discord.gg/coderabbit)
- [AI Reviews (LFX)](ai-reviews.md) — Overview of AI review tools at LFX
