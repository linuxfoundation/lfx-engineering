# LFX Engineering Agent Guidelines

This markdown file defines guidelines and standards for both human contributors
and AI-powered agents interacting with the opencode command line tool. Agents
should use this file to ensure all documentation, code style, repository
structure, git conventions, and security practices are followed when creating,
modifying, or reviewing content in this repository.

## Build/Lint/Test Commands

- Lint markdown: `npx markdownlint-cli2 "**/*.md"` or
  use `act -W .github/workflows/markdown-lint.yml` (requires container runtime)
- Validate documentation: Run markdown lint locally before committing
- CI checks: All markdown changes will be validated in CI via GitHub Actions

## Code Style Guidelines

### Documentation Standards

- Use ATX style headers (# H1, ## H2)
- Indent lists with 2 spaces
- Keep line length under 120 characters
- Headers with same content allowed in different nesting
- HTML tags are permitted in markdown
- Documents don't require a top-level header

### Repository Structure

- All documentation files must use Markdown format
- Software components use MIT license
- Documentation uses Creative Commons Attribution 4.0 license
- Follow existing directory structure and naming patterns

### Git Conventions

- Use `--signoff` for DCO and `-S` for signed commits
- Include Jira ticket ID in branch names: `feature/JIRA-123-description`
- Include Jira ticket in commit messages when applicable
- Properly link GitHub PRs to Jira tickets

### Security Practices

- Store API tokens securely (use password managers)
- Follow proper authorization workflows for tool integrations
