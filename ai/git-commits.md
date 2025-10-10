# AI Attribution Guide for Git Commits

## Overview

This guide establishes standards for attributing AI assistance in git commits
when AI tools generate or help generate code, documentation, or other project
artifacts. Proper attribution ensures transparency in our development process
and helps track the usage and effectiveness of AI tools across our codebase.

## Attribution Standards

### When to Include AI Attribution

Include AI attribution in your commit message when:

- AI tools generate complete files or substantial portions of code
- AI tools assist with significant refactoring or code transformation
- AI tools help generate documentation, configuration files, mock data, or scripts
- AI tools provide substantial assistance in debugging or problem-solving that results in code changes

### Attribution Format

Add AI attribution lines **above** the standard (required) DCO `Signed-off-by`
line in your commit message. Use the following format:

```text
<commit title>

<commit description>

<[Resolves | Jira Ticket]: <jira_ticket_link> >

Generated with [AI Tool Name](URL)

Signed-off-by: Your Name <your.email@company.com>
```

### Supported AI Tools

Use these exact attribution lines for common AI tools:

#### Development Assistants

```text
Generated with [Claude Code](https://claude.ai/code)
Generated with [OpenCode](https://opencode.ai)
Generated with [GitHub Copilot](https://github.com/features/copilot)
Generated with [Cursor](https://cursor.com)
Generated with [Gemini](https://gemini.google.com)
```

The goal of this attribution is to enable a high-level understanding of which AI
tools are being used throughout the organization, rather than tracking
individual model usage. This information helps the leadership team gather
overall metrics on AI tool adoption and make informed decisions about where to
allocate future resources.

#### Additional AI Services

As new AI tools are released, the same convention should be followed. Here are some additional examples:

```text
Generated with [ChatGPT](https://chat.openai.com)
Generated with [Amazon CodeWhisperer](https://aws.amazon.com/codewhisperer)
Generated with [Tabnine](https://www.tabnine.com)
Generated with [Grok](https://grok.com)
Generated with [Kiro](https://kiro.dev)
Generated with [Figma](https://figma.com)
```

### Multiple AI Tools

When multiple AI tools contribute to a single commit, list each tool on a separate line:

```text
Generated with [Claude Code](https://claude.ai/code)
Generated with [GitHub Copilot](https://github.com/features/copilot)

Signed-off-by: Your Name <your.email@company.com>
```

### Partial Attribution

For commits where AI provides partial assistance, consider using more specific language:

```text
Assisted by [Claude Code](https://claude.ai/code)
Assisted by [Cursor](https://cursor.com)
Reviewed with [GitHub Copilot](https://github.com/features/copilot)
```

## Example Commit Messages

### Complete AI Generation

```text
feat: add user authentication middleware

Implements JWT-based authentication middleware with proper error handling
and token validation. Includes unit tests and documentation.

Generated with [Claude Code](https://claude.ai/code)

Signed-off-by: Your Name <your.email@company.com>
```

### AI-Assisted Development

```text
refactor: optimize database query performance

Refactored user lookup queries to use indexed fields and reduced
N+1 query problems in user profile endpoints.

Assisted by [GitHub Copilot](https://github.com/features/copilot)

Signed-off-by: John Developer <john@company.com>
```

### Multiple Tools

```text
docs: update API documentation

Added comprehensive API documentation with examples and error codes.
Updated OpenAPI specification to match current implementation.

Generated with [Claude Code](https://claude.ai/code)
Assisted by [Cursor](https://cursor.com)

Signed-off-by: Alex Developer <alex@company.com>
```

## Automation with Git Templates and Hooks

### Option 1: Git Commit Message Template

Create a commit message template to remind developers to include AI attribution:

First, create a template file `.gitmessage` in your repository root:

```bash
# .gitmessage
# <type>: <description>
#
# Explain what and why (not how)
#
# If AI tools were used, add attribution above the Signed-off-by line:
# Generated with [AI Tool Name](URL)
# 
# Resolves: https://linuxfoundation.atlassian.net/browse/SWS-111
#
# Signed-off-by: Your Name <your.email@company.com>
```

Then, configure git to use the template:

```bash
git config commit.template .gitmessage
```

For team-wide adoption, add to your repository setup scripts:

```bash
#!/bin/bash
# setup-repo.sh
git config commit.template .gitmessage
echo "Git commit template configured"
```

### Option 2: Prepare-Commit-Msg Hook

Create a git hook that automatically adds AI attribution prompts:

First, create `.git/hooks/prepare-commit-msg`:

```bash
#!/bin/bash
# .git/hooks/prepare-commit-msg

# Check if this is not an amend, merge, or squash
if [ -z "$2" ]; then
    # Add AI attribution template if not already present
    if ! grep -q "Generated with\|Assisted by" "$1"; then
        sed -i.bak '/^# Please enter the commit message/i\
# If AI tools were used, add attribution above Signed-off-by:\
# Generated with [AI Tool Name](URL)\
#\
' "$1"
    fi
fi
```

Next, don't forget to make it executable:

```bash
chmod +x .git/hooks/prepare-commit-msg
```

### Option 3: Commit-Msg Validation Hook

Create a hook that validates AI attribution format:

First, create `.git/hooks/commit-msg` file:

```bash
#!/bin/bash
# .git/hooks/commit-msg

commit_msg_file="$1"
commit_msg=$(cat "$commit_msg_file")

# Check for AI attribution format
if echo "$commit_msg" | grep -qE '^(Generated with|Assisted by) '; then
    # Validate format: should have proper markdown links
    if ! echo "$commit_msg" | grep -qE '^(Generated with|Assisted by) \[.+\]\(.+\)$'; then
        echo "Error: AI attribution must use format:"
        echo "Generated with [Tool Name](URL)"
        echo "or"
        echo "Assisted by [Tool Name](URL)"
        exit 1
    fi

    # Require DCO Signed-off-by when attribution is present
    if ! echo "$commit_msg" | grep -qE '^Signed-off-by:'; then
        echo "Error: Missing DCO Signed-off-by line."
        exit 1
    fi

    # Ensure attribution appears above Signed-off-by
    attrib_line=$(awk '/^(Generated with|Assisted by) /{print NR; exit}' "$commit_msg_file")
    sob_line=$(awk '/^Signed-off-by:/{print NR; exit}' "$commit_msg_file")
    if [ -n "$attrib_line" ] && [ -n "$sob_line" ] && [ "$attrib_line" -gt "$sob_line" ]; then
        echo "Error: AI attribution must be placed ABOVE the Signed-off-by line"
        exit 1
    fi
fi

exit 0
```

Then, make the script executable:

```bash
chmod +x .git/hooks/commit-msg
```

### Team Distribution of Hooks

Since git hooks are not tracked in the repository, distribute them to the team:

1. Create a `hooks/` directory in your repository
2. Store hook scripts there
3. Add a setup script to copy hooks:

```bash
#!/bin/bash
# scripts/setup-hooks.sh

# Copy hooks to .git/hooks
cp hooks/* .git/hooks/
chmod +x .git/hooks/*

echo "Git hooks installed successfully"
echo "AI attribution will now be validated in commits"
```

## Best Practices

### Developer Guidelines

- Be honest about AI assistance - both full generation and partial help
- Include attribution even for small AI-generated snippets if they're substantial
- When in doubt, error on the side of inclusion rather than omission
- Review AI-generated code carefully before committing
- Update attribution if you significantly modify AI-generated code

### Code Review Considerations

- Reviewers should pay extra attention to commits with AI attribution
- Ensure AI-generated code follows team coding standards
- Verify that tests are included for AI-generated functionality
- Check that AI-generated documentation is accurate and complete

### Repository Management

- Periodically review commits to track AI tool usage patterns
- Use git log filters to analyze AI attribution:

```bash
# Find all commits with AI attribution
git log --grep="Generated with\|Assisted by"

# Count commits by AI tool
git log --grep="Claude Code" --oneline | wc -l
```

## Compliance and Legal Considerations

- Ensure all AI-generated code complies with the Linux Foundation licensing requirements
- Be aware of potential intellectual property implications
- Maintain records of AI tool usage for audit purposes
- Review AI tool terms of service regarding code ownership and attribution

## Tools for Analysis

### Extract AI Attribution Statistics

```bash
#!/bin/bash
# scripts/ai-stats.sh

echo "AI Attribution Statistics"
echo "========================"
echo
echo "Total commits with AI attribution:"
git log --grep="Generated with\|Assisted by" --oneline | wc -l
echo
echo "By tool:"
git log --grep="Claude Code" --oneline | wc -l | xargs echo "Claude Code:"
git log --grep="GitHub Copilot" --oneline | wc -l | xargs echo "GitHub Copilot:"
git log --grep="Cursor" --oneline | wc -l | xargs echo "Cursor:"
git log --grep="Gemini" --oneline | wc -l | xargs echo "Gemini:"
```

This guide ensures consistent, transparent attribution of AI assistance in our
development process while providing the tools and automation necessary for
team-wide adoption.
