# Operations Handbook

## Approval Workflows

### Items Requiring Platform Engineering Review and Approval

The following items require formal review and approval from Platform Engineering before implementation in production environments:

- CI/CD pipeline configurations and modifications
- Production environment configuration changes
- Deployment configurations including scaling policies and load balancing rules
- Infrastructure-as-Code (IaC) templates and modifications
- Security group rules and network configurations
- Access management and permissions changes

### Items Requiring Architectural Team Review

The following items are presented at the weekly 1-hour architectural review call for team alignment and approval:

- Significant system architecture changes or new architectural patterns
- New technology stack additions or framework changes
- API design patterns and cross-service integration approaches
- Data model changes affecting multiple services or products
- Security architecture decisions and implementations
- Performance and scalability architectural considerations

### Standard Approval Process

1. **Submission**: Engineers submit change requests via designated channels (pull requests, ticketing system, or infrastructure change request forms)
2. **Review**: Platform Engineering reviews for compliance with security policies, operational standards, and best practices
3. **Feedback**: Platform Engineering provides feedback within agreed-upon SLA (typically 1-2 business days for standard changes)
4. **Approval**: Platform Engineering approves or requests modifications
5. **Implementation**: Engineers implement approved changes
6. **Verification**: Platform Engineering verifies implementation meets approved specifications

### Expedited Approval Process

For urgent production issues or security vulnerabilities:

- Engineers request expedited review via designated escalation channels
- Platform Engineering provides accelerated review (target: within 4 hours during business hours)
- Post-implementation review conducted within 24 hours for compliance verification

### Disagreement Resolution

If Engineers and Platform Engineering disagree on approval decisions:

1. Document technical rationale from both perspectives
2. Schedule joint review meeting within 24 hours
3. If unresolved, escalate to engineering management for final decision
4. Document decision and rationale for future reference

## Ticketing System Procedures

### Ticket Creation Requirements

All work items must be tracked through Jira tickets to ensure proper coordination, visibility, and accountability across teams. Tickets should be created in the appropriate product Jira board for:

- **Feature Development**: New functionality or capabilities
- **Enhancements**: Improvements to existing features
- **Bug Fixes**: Defect resolution and corrections
- **CI/CD Changes**: Pipeline modifications, build process updates, or testing framework changes
- **Deployment Updates**: Configuration changes, scaling adjustments, or deployment strategy modifications

### Platform Engineering Support Requests

When Engineers require Platform Engineering team support for infrastructure, deployment, or operational needs:

1. **Create Ticket in Product Board**: Create the support request ticket within the same product Jira board (not a separate Platform Engineering board)

2. **Assign to Platform Engineering Team Member**: Assign the ticket to one of the designated Platform Engineering team members:
   - Alan Sherman
   - Antonia Gaete
   - Robert Detjens

3. **Slack Communication**: After creating and assigning the ticket, post a message in the `#lfx-devops` Slack channel that includes:
   - A link to the Jira ticket
   - An `@mention` tag referencing the assigned team member
   - Brief context about the request if not immediately clear from the ticket title

**Example Slack Message**:
```
@alan I've created ticket LFXV2-302 for the Nats: Automated Backup Solution 
implementation: https://linuxfoundation.atlassian.net/browse/LFXV2-302
```

### Benefits of This Workflow

- **Centralized Tracking**: All product-related work remains in the product board
- **Clear Ownership**: Explicit assignment ensures accountability
- **Transparent Communication**: Slack notification provides immediate visibility
- **Audit Trail**: Complete history of requests and resolutions in one location

## Cross-Role Collaboration Guide

### Product Planning Phase
- Product Owners define requirements and success criteria
- Engineers provide technical feasibility and effort estimates
- Platform Engineering advise on infrastructure constraints and deployment considerations

### Development Phase
- Product Owners provide clarification and acceptance testing
- Engineers develop features, build deployment configurations, and implement deployments subject to Platform Engineering approval
- Platform Engineering review and approve environment and deployment configurations

### Release Phase
- Product Owners coordinate go-to-market activities and communicate with stakeholders
- Engineers execute deployments and troubleshoot deployment issues
- Platform Engineering approve deployments and monitor infrastructure health, security, and costs

### Post-Release Phase
- Product Owners gather user feedback and measure success metrics
- Engineers address bugs, resolve security issues, and optimize performance
- Platform Engineering monitor infrastructure cost and manage operational concerns

### Ongoing Architectural Governance
- Engineers present architectural changes and proposals at weekly 1-hour architectural review call
- Architectural team (cross-product representation) provides feedback, guidance, and approval
- Ensures architectural consistency and knowledge sharing across the engineering organization
- Decisions and recommendations documented for reference by all teams

## Decision Authority by Domain

- **Product Domain**: Product Owners have final authority on feature requirements and business priorities
- **Technical Domain**: Engineers have authority on technical implementation, architecture choices (with architectural team approval), and deployment configuration design
- **Operational Domain**: Platform Engineering has authority on infrastructure provisioning, deployment approvals, and operational standards/compliance
- **Cross-functional Conflicts**: Escalate to management when decisions impact multiple domains