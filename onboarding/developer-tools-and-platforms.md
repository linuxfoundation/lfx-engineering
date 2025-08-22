# Developer Tools & Platforms

Welcome to the Linux Foundation development team! This guide provides an overview of the tools and platforms we use.

## 🚀 Getting Started

1. **Review this entire document** to understand our tech tools and platforms
2. **Request access** to all relevant platforms (see each section for account setup instructions)
3. **Access platforms** with your Linux Foundation email account, unless otherwise noted
4. **Contact your manager or IT team** if you need help with access to any platform
5. **Bookmark this page** for quick reference 😀

---

## 🔐 Security and Access Management

### 1Password

- **Purpose**: Password management and secure credential storage
- **URL**: [https://1password.com/](https://1password.com/)
- **What it does**:
  - Store and share team passwords, API keys, and sensitive information securely
  - Generate strong passwords
  - Secure notes and documents
  - Two-factor authentication backup

---

## GitHub Repositories

### Linux Foundation

- **Purpose**: Non-profit consortium dedicated to fostering the growth of Linux
- **URL**: [https://github.com/linuxfoundation](https://github.com/linuxfoundation)
- **Account Access**: You must have your personal GitHub account added to the
  Linux Foundation organization. To request access, ask your manager to submit a
  support ticket on your behalf. Access is managed through single sign-on (SSO).

### Engineering Artifact and Notes

- **Purpose**: Collection of artifacts and notes for engineers
- **URL**: [https://github.com/linuxfoundation/lfx-engineering](https://github.com/linuxfoundation/lfx-engineering)
- **Account Access**: You must have your personal GitHub account added to the
  Linux Foundation organization. To request access, ask your manager to submit a
  support ticket on your behalf. Access is managed through single sign-on (SSO).

---

## 🔧 Development Tools

### Integrated Development Environments (IDEs)

#### Cursor

- **Purpose**: AI-powered code editor with built-in AI assistance
- **URL**: [https://cursor.sh/](https://cursor.sh/)
- **What it does**:
  - AI-powered code completion and generation
  - Intelligent code refactoring and debugging
  - Natural language to code translation
  - Integration with popular extensions and themes
- **Account Access**: Ask your manager to submit a support ticket on your behalf
  or send a message in the `#lfx-ai` Slack channel.
- **Extensions**: Consider installing the Jira and GitHub Model Context Protocol
  servers.

#### Zed

- **Purpose**: High-performance, collaborative code editor
- **URL**: [https://zed.dev/](https://zed.dev/)
- **What it does**:
  - Lightning-fast performance with native speed
  - Real-time collaborative editing
  - Built-in terminal and Git integration
  - Minimalist design with powerful features
- **Account Access**: Requesting access not required. Follow the online
  documentation to download and install. Most users connect to The Linux
  Foundation GitHub account to access the CoPilot AI models.
- **Extensions**: Consider installing the Jira and GitHub Model Context Protocol
  servers.

#### Neovim

- **Purpose**: Modern, extensible terminal-based text editor
- **URL**: [https://neovim.io/](https://neovim.io/)
- **What it does**:
  - Highly customizable and extensible
  - Powerful keyboard-driven editing
  - Extensive plugin ecosystem
  - Lua scripting for configuration
- **Account Access**: Requesting access required. Download and install.  Most
  users connect to The Linux Foundation GitHub account to access the CoPilot AI
  models.

---

## 🗃️ API Documentation

### API Gateway

- **Purpose**: Central API gateway for routing and managing API requests
- **URL**: [https://api-gw.dev.platform.linuxfoundation.org/](https://api-gw.dev.platform.linuxfoundation.org/)
- **What it does**:
  - Centralized API endpoint management
  - API documentation

---

## 📋 Project Management & Collaboration

### Jira

- **Purpose**: Issue tracking, project management, and agile development workflows
- **URL**: [https://linuxfoundation.atlassian.net/jira/](https://linuxfoundation.atlassian.net/jira/)
- **What it does**:
  - Track bugs, features, and technical debt
  - Sprint planning and agile workflows
  - Project roadmap and release planning

---

## 📊 Monitoring and Analytics

### DataDog

- **Purpose**: Application performance monitoring, logging, and infrastructure metrics
- **URL**: [https://datadog.linuxfoundation.org/](https://datadog.linuxfoundation.org/)
- **What it does**:
  - Monitor application performance and uptime, including many LFX products and services
  - View logs and screen recorded sessions for debugging
  - Track infrastructure metrics and alerts
  - Review Real User Monitoring (RUM) data for user experience insights

### Snowflake

- **Purpose**: Cloud data platform for data warehousing and analytics
- **URL**: [https://www.snowflake.com/en/](https://www.snowflake.com/en/)
- **What it does**:
  - Data storage and processing
  - Business intelligence queries
  - Analytics and reporting
- **Account Access**:
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
- **URL**:
  [https://github.com/linuxfoundation/lfx-architecture/blob/main/diagrams/data-flow.md](https://github.com/linuxfoundation/lfx-architecture/blob/main/diagrams/data-flow.md)

---

## 📁 Documents

### Google Workspace

#### Google Docs

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://drive.google.com/](https://drive.google.com/)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

#### Google Spreadsheets

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/spreadsheets/](https://docs.google.com/spreadsheets/)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

#### Google Slides

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/presentation/](https://docs.google.com/presentation/)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

#### Google Videos

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/videos](https://docs.google.com/video)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

#### Google Forms

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://docs.google.com/forms](https://docs.google.com/forms)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

#### Google Drive

- **Purpose**: Linux Foundation storage of files shared with you
- **URL**: [https://drive.google.com/](https://drive.google.com/)
- **Account Access**: Google Workspace is managed through single sign-on (SSO).

---

## 🔦 Internal Tools

### Individual Dashboard

- **Purpose**: Personal profile, view your events and meetings and technical contribution and training enrollment tracking
- **URL**: [https://openprofile.dev/](https://openprofile.dev/)

### Calamari

- **Purpose**: Internal Linux Foundation tool for contractor time off tracking
- **URL**: [http://lfx.calamari.io/](http://lfx.calamari.io/)

---

## 🆘 Getting Help

If you need access to any of these tools or have questions about their usage:

1. **For tool-specific questions**: Ask in the appropriate Slack channel or reach out to team leads
2. **For urgent access issues**: Contact system administrators or use the [IT Help Center](https://jira.linuxfoundation.org/plugins/servlet/desk)

---

## 📝 Notes for New Developers

- **Bookmark this page** for quick reference
- **Ask questions** - the team is here to help you get up to speed
- **Keep credentials secure** - always use 1Password for storing sensitive information
- **Follow environment protocols** - always test in dev/staging before production

---

## 🔄 Keeping This File Updated

This document should be updated whenever:

- New tools are adopted by the team
- URLs or access procedures change
- Tools are deprecated or replaced

---

*For questions about this documentation or updates to this document, please submit a pull request
or contact the development team.*
