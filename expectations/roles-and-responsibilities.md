# Roles and Responsibilities Framework

## Responsibility Matrix

| Application | Product Owner | Application Engineer | Platform Engineer | Jira Project |
|---------------|---------------|----------------------|-------------------|--------------|
| LFX One | | Jordan Evans | Alan Sherman | [LFXV2](https://linuxfoundation.atlassian.net/jira/software/c/projects/LFXV2/boards/479) |
| Insights | | Joana | Antonia | [IN](https://linuxfoundation.atlassian.net/jira/software/c/projects/IN/boards/546) |
| CDP | | Joana | Alan Sherman | [CDP](https://linuxfoundation.atlassian.net/jira/software/c/projects/CM/boards/545) |
| EasyCLA | | ? | Robert Detjens | |
| Individual Dashboard | | Luis | Robert Detjens | |
| Auth0 | | Andres | Robert | |
| Mentorship | | Michal | Robert | [MENV2](https://linuxfoundation.atlassian.net/jira/software/c/projects/MENV2/boards/4) |
| Crowdfunding | | ? | Antonia | |
| LFX Security | | ? | Antonia | |
| ITX | | Trevor Bramwell | Alan Sherman | |
| Project Control Center | | Jordan Evans | Alan Sherman | [PCC](https://linuxfoundation.atlassian.net/jira/software/c/projects/PCC/boards/50) |

## Product Owners

### Core Responsibilities
- Gather, analyze, and prioritize business requirements from stakeholders
- Define user experience flows, wireframes, and interaction patterns
- Establish measurable KPIs, acceptance criteria, and business objectives
- Maintain product specifications, user stories, and feature documentation
- Validate product functionality against success criteria
- Define product vision and roadmap
- Manage stakeholder relationships and communications
- Conduct user research and analyze feedback

### Boundaries
- Cannot mandate specific technical architectures, frameworks, or database choices
- Cannot directly assign tasks to engineers or manage development timelines
- Cannot modify infrastructure, CI/CD pipelines, or deployment configurations
- Cannot bypass engineering review processes for technical decisions

## Engineers/Developers

### Core Responsibilities
- Design, develop, and implement end-to-end software solutions (UX, backend, data)
- Create and execute comprehensive testing strategies (unit, integration, functional)
- Build and maintain CI/CD pipelines and deployment configurations
- Implement security best practices and address vulnerabilities
- Design technical architecture following architectural team guidelines
- Provide release documentation and technical specifications
- Conduct code reviews and maintain code quality standards
- Define environment configuration requirements and deployment specifications
- Present significant architectural changes at weekly architectural review for approval

### Boundaries
- Cannot define business requirements, feature priorities, or product strategy
- Cannot provision cloud infrastructure or manage cloud resources without Platform Engineering
- Cannot deploy infrastructure changes to production without Platform Engineering review and approval
- Cannot directly handle customer support or negotiate business requirements

## Platform Engineering

### Core Responsibilities
- Review and approve all CI/CD pipelines, deployment configurations, and infrastructure changes
- Provision and manage cloud resources, servers, and networking components
- Implement and maintain monitoring, logging, and alerting systems
- Respond to incidents, coordinate resolution, and conduct post-mortems
- Track and report service health metrics and SLA compliance
- Implement infrastructure security controls and compliance frameworks
- Manage capacity planning, backup systems, and disaster recovery
- Optimize cloud costs and resource allocation

### Boundaries
- Cannot make product feature decisions or define business requirements
- Cannot write application code, business logic, or develop product features
- Cannot define user experience, UI workflows, or product functionality
- Cannot override architectural decisions made by engineering teams (within approved guidelines)

## Approval Requirements

### Platform Engineering Approval Required
CI/CD pipelines, production environment changes, deployment configurations, scaling policies,
security group rules, and access management changes require Platform Engineering review and
approval before production implementation.
### Architectural Team Review Required
Significant architecture changes, new technology additions, API design patterns,
cross-service integrations, and security architecture decisions must be presented at the
weekly architectural review call for team alignment and approval.
See the Operations Handbook for detailed approval workflows and procedures.
