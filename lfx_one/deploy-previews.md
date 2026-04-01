# LFX V2 Deploy Preview Architecture

## Overview

This document outlines the automated deploy preview system for the LFX
V2 UI, which enables developers to test pull request changes in isolated
Kubernetes environments before merging. The system leverages ArgoCD for
GitOps-based deployments, GitHub for source control and CI/CD, and
Kubernetes for hosting the preview environments.

**NOTE:** All deployment previews use the same LFX backend API.
Additional changes are needed in the future to support deploying
different versions of the API and its components.

## Key Components

### Repository Configurations

#### lfx-v2-argocd

**Purpose:** Manages ArgoCD configurations and ApplicationSets for
  automated PR deployments

**Key Features:**

- **ExternalSecrets Annotation:** Annotates base UI application external
    secrets with kubernetes-replicator `replicate-to` labels to
    automatically create secrets in `ui-pr-{number}` namespaces
- **ApplicationSet Controller:** Watches LFX V2 UI Pull Requests on
    GitHub and automatically creates new UI applications in dedicated
    namespaces (`ui-pr-{number}`)

#### lfx-v2-ui

**Purpose:** Main UI repository with CI/CD pipeline for deploy previews

**Key Features:**

- **CI/CD Pipeline:** Automated container builds triggered by the
  `deploy-preview` label on pull requests
- **ArgoCD Webhook Integration:** Sends notifications to ArgoCD
  ApplicationSet controller when labels are added (prevents up to
  30-minute delay from polling interval)

#### lfx-v2-opentofu

**Purpose:** Infrastructure as Code for Kubernetes operators and
supporting services

**Key Features:**

- **Kubernetes Replicator Deployment:** Helm-based deployment of the
  kubernetes-replicator operator
- **Secret Replication:** Enables automatic replication of secrets from
  the PCC namespace to dynamically created `ui-pr-*` namespaces

#### auth0_terraform

**Purpose:** Auth0 identity provider configuration for deploy previews

**Key Features:**

- **Dynamic Callback URLs:** PCC application configured with wildcard
  callback and audience URLs supporting
  `ui-pr-*.dev.v2.cluster.linuxfound.info`
- **Authentication:** Ensures authentication works seamlessly across all PR
  preview environments

## Architecture Diagrams

### ArgoCD and GitHub Integration

```mermaid
graph TB
    subgraph GitHub
        PR[Pull Request]
        Label[deploy-preview label]
        GHA[GitHub Actions]
        Container[Container Registry]
    end

    subgraph ArgoCD["ArgoCD (argocd.dev.v2.cluster.linuxfound.info)"]
        AppSet[ApplicationSet Controller]
        App[PR Application]
        Sync[Sync Engine]
    end

    subgraph Kubernetes["Kubernetes Cluster"]
        NS[ui-pr-# namespace]
        Deploy[Deployment]
        Svc[Service]
        Ing[Ingress]
    end

    PR --> Label
    Label --> GHA
    GHA --> Container
    Label -.webhook.-> AppSet

    AppSet -->|Creates| App
    App -->|Monitors| Container
    App -->|Deploys to| NS

    Sync -->|Creates| Deploy
    Sync -->|Creates| Svc
    Sync -->|Creates| Ing

    Ing -->|Routes to| Domain[ui-pr-#.dev.v2.cluster.linuxfound.info]
```

### Secrets Management Architecture

```mermaid
graph TB
    subgraph AWS
        SM[AWS Secrets Manager]
        IAM[IAM Roles]
    end

    subgraph "Kubernetes Cluster"
        subgraph "PCC Namespace"
            ESO[External Secrets Operator]
            ES[ExternalSecret Resources]
            Secret[Kubernetes Secrets]
            Annotation[replicate-to annotation]
        end

        subgraph "Kubernetes Replicator"
            Rep[Replicator Controller]
            Watch[Namespace Watcher]
        end

        subgraph "ui-pr-1 Namespace"
            Secret1[Replicated Secret]
            App1[Application]
        end

        subgraph "ui-pr-2 Namespace"
            Secret2[Replicated Secret]
            App2[Application]
        end
    end

    SM -->|Fetches| ESO
    IAM -->|Authorizes| ESO
    ESO -->|Creates| ES
    ES -->|Generates| Secret
    Secret -->|Has| Annotation

    Rep -->|Watches| Annotation
    Watch -->|Detects| NS1[ui-pr-* namespaces]

    Rep -->|Replicates to| Secret1
    Rep -->|Replicates to| Secret2

    Secret1 -->|Used by| App1
    Secret2 -->|Used by| App2
```

## Workflow

1. **Developer creates PR:** A pull request is opened in the lfx-v2-ui repository
2. **Label Application:** Developer or reviewer adds the `deploy-preview` label
3. **Container Build:** GitHub Actions builds and pushes a container image tagged with the PR number
4. **ArgoCD Notification:** Webhook notifies ArgoCD ApplicationSet controller (avoids polling delay)
5. **Application Creation:** ApplicationSet creates a new ArgoCD application for the PR
6. **Namespace Setup:** Kubernetes namespace `ui-pr-{number}` is created
7. **Secret Replication:** Kubernetes-replicator copies necessary secrets from PCC namespace
8. **Deployment:** ArgoCD deploys the application with the PR-specific container image
9. **Access:** Preview available at `ui-pr-{number}.dev.v2.cluster.linuxfound.info`

## Values Override Mechanism

Deploy previews use ArgoCD's multi-source feature to layer Helm values
from multiple sources. The ApplicationSet
(`apps/dev/lfx-v2-ui-branch.yaml`) defines two sources that work
together to configure each `ui-pr-{number}` deployment.

### Values Hierarchy (lowest to highest priority)

1. **Chart defaults** (`charts/lfx-v2-ui/values.yaml` in the PR branch)
   — The Helm chart's built-in defaults from the `lfx-v2-ui` repository.
   Since the chart source is `pull/{number}/head`, developers can modify
   chart defaults in their PR branch.

2. **Branch deployment values**
   (`values/dev/lfx-v2-ui-branch.yaml` in `lfx-v2-argocd`) — Shared
   configuration for all deploy previews including resource limits,
   environment variables, and ingress settings. This file is referenced
   via the `$values` multi-source reference.

3. **ApplicationSet parameters** (inline in
   `apps/dev/lfx-v2-ui-branch.yaml`) — PR-specific overrides injected
   by the ApplicationSet template using `{{.number}}` templating:
   - `ingress.hosts[0].host` and `ingress.tls[0].hosts[0]` — sets the
     preview domain to `ui-pr-{number}.dev.v2.cluster.linuxfound.info`
   - `environment.PCC_BASE_URL.value` — sets the base URL for
     authentication callbacks
   - `image.tag` — pins the container image to `ui-pr-{number}`

### Branch Deployment Values

The `values/dev/lfx-v2-ui-branch.yaml` file in `lfx-v2-argocd`
configures deploy previews with reduced resources compared to the main
development environment:

- **Resources:** 100m/128Mi requests, 500m/512Mi limits
- **Scaling:** Single replica, autoscaling disabled
- **Environment:** Shared backend API endpoints, Auth0 dev tenant,
  Snowflake analytics, LaunchDarkly feature flags, and Datadog RUM
- **Secrets:** Referenced from `pcc-secrets` (replicated from PCC
  namespace via kubernetes-replicator)
- **Ingress:** Traefik with TLS via cert-manager and external-dns

### Modifying Deploy Preview Configuration

To change the shared configuration for **all** deploy previews, update
`values/dev/lfx-v2-ui-branch.yaml` in the `lfx-v2-argocd` repository.
Common modifications include adding new environment variables, adjusting
resource limits, or updating service endpoints.

To change configuration for a **single** deploy preview, modify the
chart's `values.yaml` in the PR branch of `lfx-v2-ui`. Note that values
from `lfx-v2-ui-branch.yaml` and ApplicationSet parameters take
precedence over chart defaults.

## Security Considerations

- All deploy previews use the same Auth0 tenant with wildcard configurations
- Secrets are replicated only to authorized `ui-pr-*` namespaces
- Each PR environment is isolated in its own Kubernetes namespace
- AWS IAM roles restrict secret access to authorized services only
- Deploy preview environments are automatically cleaned up when PRs are closed
