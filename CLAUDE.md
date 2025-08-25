# CLAUDE.md - LFX Engineering Repository Guide

## Project Overview

This repository serves as a centralized collection of engineering artifacts,
documentation, and configuration guides for The Linux Foundation's engineering
team. It focuses on AI tool integration, team onboarding, and best practices for
development workflows.

## Repository Structure

```code
lfx-engineering/
├── README.md                 # Basic repository description
├── MAINTAINERS.md           # Team maintainers and contact information
├── LICENSE                  # MIT License for software components
├── LICENSE-docs            # Creative Commons Attribution 4.0 for documentation
├── .gitignore              # Excludes .claude/ directory
├── mcp/                    # Model Context Protocol configurations
│   └── jira.md            # Jira MCP setup guide for AI tools
├── onboarding/             # Team onboarding resources
│   └── cursor.md          # Cursor AI setup and workspace invitation
├── expectations/           # (Currently empty) Expected to contain team expectations
├── prompts/               # (Currently empty) Expected to contain AI prompts
└── training/              # (Currently empty) Expected to contain training materials
```

## Technology Stack

- **Repository Type**: Documentation and configuration repository
- **Primary Focus**: AI tooling integration and team resources
- **Licensing**: Dual licensing approach
  - MIT License for software components
  - Creative Commons Attribution 4.0 for documentation
- **Version Control**: Git with main branch as primary
- **Documentation Standards**: All markdown files must pass markdownlint validation

## Key Components

### 1. AI Tool Integration

- **Jira MCP Configuration**: Comprehensive setup guide for integrating Jira with AI tools
  - Claude Code (CLI) configuration with API tokens
  - Cursor AI integration with MCP servers
  - Atlassian authorization workflows
  - Validation and testing procedures

### 2. Team Onboarding

- **Cursor AI Setup**: Instructions for team members to join the LF AI workspace
- **Tool Access**: Process for requesting invites through #lf-ai Slack channel

### 3. Team Management

- **Maintainers**: Four primary maintainers from The Linux Foundation (see [MAINTAINERS.md](MAINTAINERS.md))
  - David Deal (dealako)
  - Jordan Evans (jordane)
  - Luis Guerra (luismoriguerra)
  - Michal Lehotsky (mlehotskylf)

## Development Workflow

### File Management

- Documentation files use Markdown format
- Configuration examples provided in JSON format
- Clear separation between software (MIT) and documentation (CC BY 4.0) licenses

### AI Tool Configuration

- MCP (Model Context Protocol) servers for external service integration
- Support for multiple AI platforms (Claude Code, Cursor AI)
- Standardized configuration patterns across tools

## Key Features

### MCP Jira Integration

- Complete MCP setup for Jira API access
- Support for both individual and team configurations
- Validation procedures with example queries
- Security considerations with API token management

### Team Collaboration

- Centralized workspace management for Cursor AI
- Slack-based invitation process (#lf-ai channel)
- Standardized onboarding procedures

## Usage Patterns

### For New Team Members

1. Review onboarding documentation in `/onboarding/`
2. Request Cursor AI workspace access via #lf-ai Slack channel
3. Follow MCP setup guides for tool integration
4. Configure personal API tokens as needed

### For AI Tool Setup

1. Reference `/mcp/jira.md` for Jira integration
2. Follow platform-specific configuration (Claude Code vs Cursor AI)
3. Complete authorization workflows
4. Validate setup with provided test procedures

### For Documentation Updates

1. Maintain dual licensing approach
2. Use Markdown for all documentation
3. Follow existing structure and formatting patterns
4. Consider both software and documentation licensing requirements
5. Ensure all markdown passes markdownlint validation
6. Run `markdownlint` to verify compliance before committing changes

## Security Considerations

- API tokens should be stored securely (password managers recommended)
- Many MCP configurations require proper authorization workflows
- Git ignores `.claude/` directory for local configurations, command approvals, etc.

## Future Development

Based on the directory structure, planned expansions likely include:

- **Expectations**: Team standards and development expectations
- **Prompts**: Standardized AI prompts for common tasks
- **Training**: Educational materials and best practices

## Best Practices

1. **Documentation**: All setup procedures include validation steps
2. **Security**: Clear guidance on secure token management
3. **Team Coordination**: Centralized communication through dedicated Slack channels
4. **Tool Integration**: Standardized MCP patterns across different AI platforms
5. **Version Control**: Proper gitignore patterns for sensitive local configurations
6. **Markdown Quality**: All markdown documentation must pass markdownlint validation

## Contact and Support

For questions or access requests:

- Join the #lf-ai Slack channel
- Contact repository maintainers listed in MAINTAINERS.md
- Follow established onboarding procedures

This repository demonstrates a mature approach to AI tool integration within an
enterprise engineering environment, with emphasis on security, standardization,
and team collaboration.

## Jira Integration

When a Jira ticket is mentioned, ask the user if we should update the
corresponding Jira ticket to indicate it is ‘In Progress’. Ask the user if a
comment should be added to the Jira ticket after successfully completing a task
or activity. If the user agrees, then build a progress report and submit this
comment to the Jira ticket.

When a Jira ticket is mentioned, review the Jira ticket description. If the Jira
ticket description does not include a ‘Acceptance Criteria’ section, then prompt
the user to add one. Each ticket should have some completion criteria.

When a Jira ticket is mentioned, review the Jira ticket description. If the Jira
ticket description does not include a ‘Reviewers/Stakeholders’ section, then
prompt the user to add one. Each ticket should know who can validate and confirm
that the feature was implemented or the bug fix was resolved.

When a pull request is created and a Jira ticket is associated with the work,
confirm and add the GitHub pull request to the Jira as a comment.

## Git Integration

For all git commits, use the `--signoff` option for DCO and the `-S` for signed
commits. If a Jira ticket is mentioned in the prompt, ensure the Jira ticket
identifier and link are included in the commit message.

Review the git branch name. If a Jira ticket number is provided and is not
included in the branch name, prompt and suggest to the user a better branch name
that includes the Jira ticket identifier in the name. A bad branch name example
is ‘feature/update-cookie-policy’. A better branch name example would be
‘feature/LH-2-update-cookie-policy’, where the format is
‘feature/{JIRA_TICKET_IDENTIFIER}-{feature-short-name}’ or
‘bug/{JIRA_TICKET_IDENTIFIER}-{bug-short-name}.
