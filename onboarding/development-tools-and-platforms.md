# Developer Tools & Platforms

Welcome to the Linux Foundation development team! This guide provides an overview of the tools and platforms we use.

## 🚀 Getting Started

1. **Review this entire document** to understand our tech tools and platforms
2. **Access platforms** with your Linux Foundation email account, unless otherwise noted
3. **Contact your manager or tech lead** if you need help with access to any platform
4. **Bookmark this page** for quick reference 😀

---

## 🔐 Security and Access Management

### 1Password

- **Purpose**: Password management and secure credential storage
- **URL**: [https://1password.com/](https://1password.com/)

**What you'll use 1Password for:**

  - Store and share team passwords, API keys, and sensitive information securely
  - Generate strong passwords for all development accounts
  - Access shared team vaults for common credentials
  - Store secure notes and documents related to projects
  - Backup codes for two-factor authentication across all platforms

**How to get access:**

  1. Access is typically granted
  2. If you cannot access, contact your team lead


**Setup Steps:**

  1. Install 1Password browser extension for seamless login experience
  2. Accept team invitation via email and create your master password
  3. Configure 1Password app on mobile device for 2FA backup codes

---

## GitHub Repositories

### Linux Foundation

- **Purpose**: Non-profit consortium dedicated to fostering the growth of Linux
- **URL**: [https://github.com/linuxfoundation](https://github.com/linuxfoundation)

**What you'll use GitHub for:**

  - Clone and contribute to LFX platform repositories
  - Create feature branches
  - Submit pull requests
  - Participate in code reviews and provide feedback to team members
  - Track issues and collaborate on technical discussions
  - Access documentation, coding standards, and architecture decisions

**How to get access:**

  1. You must have your personal GitHub account added to the Linux Foundation organization. 
  2. To request access, ask your manager to submit a support ticket on your behalf. 
  3. Access is managed through single sign-on (SSO).

**Setup Steps:**

  1. Enable two-factor authentication (required for all organization members)
  2. Configure SSH keys following our security guidelines
  3. Authorize the organization under SSH and GPG keys -> Configure SSO
  4. Set up Git configuration with your personal GitHub email: `git config --global user.email "your.email@gmail.com"`
  5. Clone your project's repository and follow it's `Getting Started` guide

### Engineering Artifact and Notes

- **Purpose**: Collection of artifacts and notes for engineers
- **URL**: [https://github.com/linuxfoundation/lfx-engineering](https://github.com/linuxfoundation/lfx-engineering)

**What you'll use lfx-engineering for:**

  - Access onboarding materials and training resources
  - Review architecture documentation and technical specifications
  - Learn AI integrations and mcp clients
  - Contribute to internal documentation and knowledge sharing

**How to get access:**

  1. You must have your personal GitHub account added to the Linux Foundation organization. 
  2. To request access, ask your manager to submit a support ticket on your behalf. 
  3. Access is managed through single sign-on (SSO).

---

## 🔧 Development Tools

### Integrated Development Environments (IDEs)

#### Cursor

- **Purpose**: AI-powered code editor with built-in AI assistance
- **URL**: [https://cursor.sh/](https://cursor.sh/)

**What you'll use Cursor for:**

  - AI-powered code completion and generation for faster development
  - Intelligent code refactoring and debugging assistance
  - Natural language to code translation for complex logic
  - Integration with our team's extensions and coding standards
  - Collaborative debugging sessions with AI assistance

**How to get access:**

  1. Ask your manager to submit a support ticket on your behalf or send a message
  in the `#lfx-ai` Slack channel.

**Setup Steps:**

  1. Download and install Cursor
  2. Sign in with your Linux Foundation email
  3. Install the [recommended extensions](https://github.com/linuxfoundation/lfx-engineering/tree/main/onboarding/code-editor-extensions-and-plugins/cursor.md)
  4. Consider installing the Jira and GitHub Model Context Protocol servers.

#### Zed

- **Purpose**: High-performance, collaborative code editor
- **URL**: [https://zed.dev/](https://zed.dev/)

**What you'll use Zed for:**

  - Lightning-fast performance for large codebases
  - Real-time collaborative editing during pair programming sessions
  - Built-in terminal integration for development workflows
  - Git integration for seamless version control
  - Minimalist interface that reduces cognitive load during coding

**How to get access:**

  1. Requesting access not required. Follow the online documentation to download and install.
  2. Most users connect to The Linux Foundation GitHub account to access the CoPilot AI models.

**Setup Steps:**

  1. Download and install Zed
  2. Sign in with your choice of email
  3. Install the [recommended extensions](https://github.com/linuxfoundation/lfx-engineering/tree/main/onboarding/code-editor-extensions-and-plugins/zed.md)

#### Neovim

- **Purpose**: Modern, extensible terminal-based text editor
- **URL**: [https://neovim.io/](https://neovim.io/)

**What you'll use Neovim for:**

  - Highly efficient keyboard-driven editing for experienced developers
  - Customizable development environment tailored to your workflow
  - Powerful plugin ecosystem for language-specific development
  - Terminal-based editing for remote development and server work
  - Lua scripting for advanced configuration and automation

**How to get access:**

  1. Requesting access not required. Follow the online documentation to download and install.
  2. Most users connect to The Linux Foundation GitHub account to access the CoPilot AI models.

**Setup Steps:**

  1. Install Neovim using your system's package manager
  2. Install the [recommended extensions](https://github.com/linuxfoundation/lfx-engineering/tree/main/onboarding/code-editor-extensions-and-plugins/neovim.md)

---

## 🗃️ API Documentation

### API Gateway

- **Purpose**: Central API gateway for routing and managing API requests
- **URL**: [https://api-gw.dev.platform.linuxfoundation.org/](https://api-gw.dev.platform.linuxfoundation.org/)

**What you'll use API Gateway for:**

  - Access comprehensive API documentation for all LFX services
  - Test API endpoints during development and debugging
  - Understand request/response formats and authentication requirements

**How to get access:**

  1. Access is typically granted
  2. If you cannot access, contact your team lead

---

## 📋 Project Management & Collaboration

### Jira

- **Purpose**: Issue tracking, project management, and agile development workflows
- **URL**: [https://linuxfoundation.atlassian.net/jira/](https://linuxfoundation.atlassian.net/jira/)

**What you'll use Jira for:**

  - Track and update progress on assigned bugs, features, and technical debt
  - Create detailed bug reports with reproduction steps and severity levels
  - Manage feature requests through our workflow: Backlog → In Progress → Code Review → Testing → Done
  - Estimate story points and track velocity metrics
  - Generate reports for stakeholders and project planning

**How to get access:**

  1. You should receive an invitation to Atlassian via email
  2. Accept invitation and complete account setup

**Setup Steps:**

  1. Sign in to Jira using your Linux Foundation credentials
  2. Set up email notifications to your preference 
  3. View your team's project boards and understand the workflow stages
  4. Install Jira mobile app for updates on the go

---

## 📊 Monitoring and Analytics

### DataDog

- **Purpose**: Application performance monitoring, logging, and infrastructure metrics
- **URL**: [https://datadog.linuxfoundation.org/](https://datadog.linuxfoundation.org/)

**What you'll use DataDog for:**

  - Monitor application performance and uptime, including many LFX products and services
  - View logs and screen recorded sessions for debugging
  - Review Real User Monitoring (RUM) data for user experience insights
  - View infrastructure metrics to understand system resource usage

**How to get access:**

  1. You should receive an invitation to Datadog via email
  2. Accept invitation and head to Datadog to sign in

### Snowflake

- **Purpose**: Cloud data platform for data warehousing and analytics
- **URL**: [https://app.snowflake.com/](https://app.snowflake.com/)

**What you'll use Snowflake for:**

  - Query large datasets for business intelligence and reporting needs
  - Analyze user behavior and application metrics for product decisions
  - Extract data for machine learning models and analytics projects
  - Generate custom reports for stakeholders and project planning
  - Collaborate with data team on complex analytical queries

**How to get access:**

  - Step 1: Datalake Team: IT/Ops Team Account Setup
    1. Determine which roles should be added to the new user(s).
    2. Create a pull request to add the users and roles to the GitHub Snowflake
       Terraform repository. Here’s a link [to the GitHub
       repository](https://github.com/linuxfoundation/lfx-snowflake-terraform) and
       the user configuration file. Anyone can create a pull request.
    3. Request an IT/Ops team member review the pull request in the #lfx-devops
       Slack channel. They will review and merge the pull request changes wich
       will trigger a deployment to the Snowflake production environment.
  - Step 2: Developer: User Account Setup
    1. Once Step 1 above is complete, users can now log into Snowflake.
    2. [Direct Link to The Linux
       Foundation](https://app.snowflake.com/jnmhvwd/xpb85243/) Snowflake account.
    3. Users should log in using their LF email (e.g.,
       [your_lf_email@linuxfoundation.org](your_lf_email@linuxfoundation.org)) via
       the Google SSO option.
    4. Once logged in, the user should be redirected back to the Snowflake
       landing page.
    5. Users should then notify the Datalake Team. The Datalake Team will share
       any relevant dashboards or interactive views/applications
       (sharing is done after the user has logged in the first time).
    6. Once the Datalake Team has granted access from the Snowflake console,
       navigate to: Menu -> Projects -> Dashboards -> “Shared with me"

---

## Organization Diagram

### LFX-Datalake

- **Purpose**: Architecture diagram showing LFX services, data flows, and system dependencies
- **URL**: [https://whimsical.com/lfx-datalake-Qgy1wT6KC4RtrhCbfVQ37X](https://whimsical.com/lfx-datalake-Qgy1wT6KC4RtrhCbfVQ37X)

**What you'll use LFX-Datalake diagram for:**

  - Understand overall organization architecture and platform relationships

**How to get access:**

  1. Request access to the specific diagram from your director or team lead

### Linux Product & Engineering organization

- **Purpose**: Organizational chart with team structure, roles and reporting relationships
- **URL**: [https://lucid.app/lucidchart/a39693e8-9f93-4f6e-9f31-77cb28de4f81/edit?page=0_0#](https://lucid.app/lucidchart/a39693e8-9f93-4f6e-9f31-77cb28de4f81/edit?page=0_0#)

**What you'll use the organization chart for:**

  - Understand team structure, roles, and reporting relationships

**How to get access:**

  1. Request access to the specific diagram from your director or team lead


TODO: Leaving for David Deal to update

### LFX-Datalake

- **Purpose**: Architecture diagram showing LFX services, data flows, and system dependencies
- **URL**: [https://github.com/linuxfoundation/lfx-architecture/blob/main/diagrams/data-flow.md](https://github.com/linuxfoundation/lfx-architecture/blob/main/diagrams/data-flow.md)

---

## 📁 Documents

### Google Workspace

#### Google Docs

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://drive.google.com/](https://drive.google.com/)

**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)
  - If you cannot access specific documents, request sharing permissions from document owner

#### Google Spreadsheets

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/spreadsheets/](https://docs.google.com/spreadsheets/)


**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)
  - If you cannot access specific documents, request sharing permissions from spreadsheet owner

#### Google Slides

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/presentation/](https://docs.google.com/presentation/)

**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)
  - If you cannot access specific documents, request sharing permissions from presentation owner

#### Google Vids

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/videos](https://docs.google.com/videos)

**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)
  - If you cannot access specific documents, request sharing permissions from video owner

#### Google Forms

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/forms](https://docs.google.com/forms)

**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)
  - If you cannot access specific documents, request sharing permissions from form owner

#### Google Drive

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://drive.google.com/](https://drive.google.com/)

**How to get access:**

  - Google Workspace is managed through single sign-on (SSO)

---

## 🔦 Internal Tools

### Individual Dashboard

- **Purpose**: Personal profile, view your events and meetings and technical contribution and training enrollment tracking
- **URL**: [https://openprofile.dev/](https://openprofile.dev/)

**What you'll use Individual Dashboard for:**

  - Track your technical contributions across projects and repositories
  - Monitor your training progress and certifications 
  - View upcoming meetings and events relevant to your role
  - Access performance metrics and goal tracking information

**Setup Steps:**

  1. Log in using your Linux Foundation credentials

### Calamari

- **Purpose**: Internal Linux Foundation tool for time off tracking
- **URL**: [http://lfx.calamari.io/](http://lfx.calamari.io/)

**What you'll use Calamari for:**

  - Request vacation time, sick leave, and other time off
  - View your time off balance and accrual rates
  - Submit time off requests for manager approval
  - View team calendar to coordinate time off with colleagues
  - Track public holidays and company-wide time off policies

**How to get access:**

  1. You should receive an invitation to Calamari via email
  2. Accept invitation and complete account setup

---

## 🆘 Getting Help

If you need access to any of these tools or have questions about their usage:

1. **For tool-specific questions**: Ask in the appropriate Slack channel or reach out to a fellow team member
2. **For access issues**: Contact your manager or tech lead for assistance

---

## 📝 Notes for New Developers

- **Bookmark this page** for quick reference during your first few weeks
- **Ask questions** - the team is here to help you get up to speed quickly
- **Keep credentials secure** - always use 1Password for storing sensitive information
- **Follow environment protocols** - always test in dev/staging before production
- **Complete setup checklists** - use the setup steps for each tool to ensure proper configuration
- **Join relevant Slack channels** - stay connected with your team for real-time support and updates

---

## 🔄 Keeping This File Updated

This document should be updated whenever:

- New tools are adopted by the team
- URLs or access procedures change
- Tools are deprecated or replaced
- Setup procedures are modified or improved
- New onboarding feedback suggests improvements

---

*For questions about this documentation or updates to this document, please submit a pull request
or contact the development team.*
