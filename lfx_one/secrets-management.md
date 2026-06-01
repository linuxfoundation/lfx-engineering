# LFX V2 Secrets Management

This document provides comprehensive guidance for managing secrets in the LFX V2 platform using the centralized
[lfx-secrets-management](https://github.com/linuxfoundation/lfx-secrets-management) repository.

## Table of Contents

- [Overview](#overview)
  - [Integration with Service Accounts](#integration-with-service-accounts)
  - [Secrets Flow Architecture](#secrets-flow-architecture)
- [Prerequisites](#prerequisites)
- [Configuration Requirements](#configuration-requirements)
  - [Tags](#tags)
  - [Service Tag Integration](#service-tag-integration)
  - [Vault Requirements](#vault-requirements)
  - [AWS Destination Accounts](#aws-destination-accounts)
  - [Region](#region)
  - [Path Format](#path-format)
- [Step-by-Step Process](#step-by-step-process)
  - [1. Add Secret to 1Password](#1-add-secret-to-1password)
  - [2. Create YAML Configuration](#2-create-yaml-configuration)
  - [3. Submit Pull Request](#3-submit-pull-request)
  - [4. Deploy Secrets](#4-deploy-secrets)
  - [5. Using the Secret in Kubernetes](#5-using-the-secret-in-kubernetes)
  - [6. Register the IAM Service Account Role](#6-register-the-iam-service-account-role)
  - [7. Add the ServiceAccount Template](#7-add-the-serviceaccount-template)
  - [8. Add the SecretStore Template](#8-add-the-secretstore-template)
  - [9. Add the ExternalSecret Template](#9-add-the-externalsecret-template)
  - [10. Add the values.yaml Defaults](#10-add-the-valuesyaml-defaults)
  - [11. Configure Per-Environment Values](#11-configure-per-environment-values)
  - [12. Configure Local Development](#12-configure-local-development)
- [Configuration Example](#configuration-example)
  - [Configuration Breakdown](#configuration-breakdown)
- [Deployment Methods](#deployment-methods)
  - [Using GitHub Actions](#using-github-actions)
- [Validation and Testing](#validation-and-testing)
  - [Verify Deployment](#verify-deployment)
  - [Test Service Access](#test-service-access)
  - [Audit Secrets](#audit-secrets)
- [Best Practices](#best-practices)
  - [Security Considerations](#security-considerations)
  - [Naming Conventions](#naming-conventions)
  - [Environment Management](#environment-management)
  - [Change Management](#change-management)
- [Troubleshooting](#troubleshooting)
  - [Common Issues](#common-issues)
  - [Getting Help](#getting-help)

## Overview

LFX V2 uses an automated secrets management system that synchronizes secrets between 1Password vaults and
AWS Secrets Manager across different environments. This approach ensures:

- **Centralized Management**: All secrets are managed through a single repository
- **Environment Isolation**: Separate vaults and AWS accounts for dev, staging, and production
- **Automated Deployment**: Secrets are deployed using GitHub Actions or local tooling
- **Security**: Secrets never leave secure storage systems during transit

### Integration with Service Accounts

LFX V2 services automatically consume secrets through a tag-based system managed by the platform's
infrastructure. Each service has a dedicated IAM role and Kubernetes service account that can only access
secrets tagged with their specific service identifier. For detailed information about how services consume
these secrets, see the [Service Account Management documentation](https://github.com/linuxfoundation/lfx-v2-opentofu/blob/main/docs/service-accounts.md).

**Key Integration Points:**

- **Tag-Based Access**: Services only access secrets tagged with their service name
- **External Secrets Operator**: Automatically discovers and merges secrets into Kubernetes
- **IRSA Authentication**: Secure role assumption without storing credentials
- **10-minute Refresh**: Secrets are automatically refreshed every 10 minutes

### Secrets Flow Architecture

The following diagram illustrates how secrets flow from 1Password through to application deployment:

```mermaid
graph TB
    subgraph "1Password Vaults"
        OP1["LFX V2 - Development<br/>📦 Secret Item"]
        OP2["LFX V2 - Staging<br/>📦 Secret Item"]
        OP3["LFX V2 - Production<br/>📦 Secret Item"]
    end

    subgraph "Secrets Management Repository"
        YAML["YAML Configuration<br/>🔧 secretsmanagement/secrets/cloud.yml"]
        GA["GitHub Actions<br/>🚀 Deploy Workflow"]
    end

    subgraph "AWS Secrets Manager"
        ASM1["us-west-2<br/>lfx-development<br/>🔐 /cloud/litellm/lfx-v2<br/>🏷️ service: pcc"]
        ASM2["us-west-2<br/>lfx-staging<br/>🔐 /cloud/litellm/lfx-v2<br/>🏷️ service: pcc"]
        ASM3["us-west-2<br/>lfx-production<br/>🔐 /cloud/litellm/lfx-v2<br/>🏷️ service: pcc"]
    end

    subgraph "Development Cluster"
        subgraph "Dev Service Account & IRSA"
            SA1["pcc-sa<br/>🔑 Service Account<br/>📋 Role: k8s-secret-access-pcc"]
        end

        subgraph "Dev External Secrets"
            SS1["pcc-secret-store<br/>🏪 SecretStore<br/>🔗 Uses IRSA"]
            ES1["pcc-secrets<br/>📊 ExternalSecret<br/>🏷️ Filters by service: pcc<br/>⏱️ Refreshes every 10m"]
        end

        subgraph "Dev Kubernetes Secrets"
            KS1["pcc-secrets<br/>🔒 K8s Secret<br/>📝 Merged secret values"]
        end

        subgraph "Dev Application"
            DEP1["pcc-deployment<br/>🚀 Deployment<br/>📌 serviceAccountName: pcc-sa<br/>🔗 secretRef: pcc-secrets"]
            POD1["Pod<br/>🏃 Running Container<br/>🌍 Environment Variables<br/>LITELLM_API_KEY=***"]
        end
    end

    subgraph "Staging Cluster"
        subgraph "Staging Service Account & IRSA"
            SA2["pcc-sa<br/>🔑 Service Account<br/>📋 Role: k8s-secret-access-pcc"]
        end

        subgraph "Staging External Secrets"
            SS2["pcc-secret-store<br/>🏪 SecretStore<br/>🔗 Uses IRSA"]
            ES2["pcc-secrets<br/>📊 ExternalSecret<br/>🏷️ Filters by service: pcc<br/>⏱️ Refreshes every 10m"]
        end

        subgraph "Staging Kubernetes Secrets"
            KS2["pcc-secrets<br/>🔒 K8s Secret<br/>📝 Merged secret values"]
        end

        subgraph "Staging Application"
            DEP2["pcc-deployment<br/>🚀 Deployment<br/>📌 serviceAccountName: pcc-sa<br/>🔗 secretRef: pcc-secrets"]
            POD2["Pod<br/>🏃 Running Container<br/>🌍 Environment Variables<br/>LITELLM_API_KEY=***"]
        end
    end

    subgraph "Production Cluster"
        subgraph "Prod Service Account & IRSA"
            SA3["pcc-sa<br/>🔑 Service Account<br/>📋 Role: k8s-secret-access-pcc"]
        end

        subgraph "Prod External Secrets"
            SS3["pcc-secret-store<br/>🏪 SecretStore<br/>🔗 Uses IRSA"]
            ES3["pcc-secrets<br/>📊 ExternalSecret<br/>🏷️ Filters by service: pcc<br/>⏱️ Refreshes every 10m"]
        end

        subgraph "Prod Kubernetes Secrets"
            KS3["pcc-secrets<br/>🔒 K8s Secret<br/>📝 Merged secret values"]
        end

        subgraph "Prod Application"
            DEP3["pcc-deployment<br/>🚀 Deployment<br/>📌 serviceAccountName: pcc-sa<br/>🔗 secretRef: pcc-secrets"]
            POD3["Pod<br/>🏃 Running Container<br/>🌍 Environment Variables<br/>LITELLM_API_KEY=***"]
        end
    end

    %% Flow connections
    OP1 --> YAML
    OP2 --> YAML
    OP3 --> YAML

    YAML --> GA
    GA --> ASM1
    GA --> ASM2
    GA --> ASM3

    %% Development cluster flows
    ASM1 --> SS1
    SA1 --> SS1
    SS1 --> ES1
    ES1 --> KS1
    KS1 --> DEP1
    SA1 --> DEP1
    DEP1 --> POD1

    %% Staging cluster flows
    ASM2 --> SS2
    SA2 --> SS2
    SS2 --> ES2
    ES2 --> KS2
    KS2 --> DEP2
    SA2 --> DEP2
    DEP2 --> POD2

    %% Production cluster flows
    ASM3 --> SS3
    SA3 --> SS3
    SS3 --> ES3
    ES3 --> KS3
    KS3 --> DEP3
    SA3 --> DEP3
    DEP3 --> POD3

    %% Styling
    classDef onepassword fill:#0f62fe,stroke:#0f62fe,stroke-width:2px,color:#fff
    classDef secrets fill:#ff6b35,stroke:#ff6b35,stroke-width:2px,color:#fff
    classDef aws fill:#ff9500,stroke:#ff9500,stroke-width:2px,color:#fff
    classDef k8s fill:#326ce5,stroke:#326ce5,stroke-width:2px,color:#fff
    classDef app fill:#2d5016,stroke:#2d5016,stroke-width:2px,color:#fff

    class OP1,OP2,OP3 onepassword
    class YAML,GA secrets
    class ASM1,ASM2,ASM3 aws
    class SA1,SS1,ES1,KS1,SA2,SS2,ES2,KS2,SA3,SS3,ES3,KS3 k8s
    class DEP1,POD1,DEP2,POD2,DEP3,POD3 app
```

**Flow Explanation:**

1. **Source**: Secrets are stored in environment-specific 1Password vaults
2. **Configuration**: YAML files define the secret mapping and deployment rules
3. **Deployment**: GitHub Actions deploy secrets to AWS Secrets Manager with appropriate tags
4. **Discovery**: External Secrets Operator uses IRSA to authenticate and discover tagged secrets
5. **Synchronization**: Secrets are merged into Kubernetes secrets and refreshed every 10 minutes
6. **Consumption**: Applications reference the service account and secret to access environment variables

## Prerequisites

Before managing secrets for LFX V2, ensure you have:

- Access to the appropriate 1Password vaults (see [Vault Requirements](#vault-requirements))
- AWS access to the target environments
- Contributor access to the [lfx-secrets-management](https://github.com/linuxfoundation/lfx-secrets-management) repository
- Understanding of the service architecture and which backend services will consume the secrets

## Configuration Requirements

### Tags

Every secret configuration **must** include these tags:

- `lfx_v2` - Identifies the secret as belonging to LFX V2
- `<upstream_service>` - The third-party service providing the secret (e.g., `litellm`, `stripe`, `zoom`)
- `<backend_service>` - The LFX V2 service that will consume the secret (e.g., `pcc`)

**Example tags:**

```yaml
tags: [lfx_v2, litellm, pcc]
```

#### Service Tag Integration

The `<backend_service>` tag is critical as it determines which Kubernetes service accounts can access the
secret. When secrets are deployed to AWS Secrets Manager, they receive a `service` tag that matches your
backend service name (e.g., `service: pcc`). The LFX V2 infrastructure then uses this tag to:

1. **Control Access**: Only the matching service's IAM role can retrieve the secret. I.e the `pcc-sa`
   service account uses role `arn:aws:iam::788942260905:role/k8s-secret-access-pcc` in dev to only
   access specific tagged secrets that match `"aws:ResourceTag/service": "pcc"`
2. **Auto-Discovery**: The External Secrets Operator automatically finds and merges all secrets with the service tag
3. **Environment Variables**: All tagged secrets are made available as environment variables in the service's pods

For the complete technical details of this integration, refer to the
[Service Account Management documentation](https://github.com/linuxfoundation/lfx-v2-opentofu/blob/main/docs/service-accounts.md).

### Vault Requirements

Secrets must first be stored in the appropriate 1Password vault:

| Environment | 1Password Vault |
| ----------- | --------------- |
| Development | `LFX V2 - Development` |
| Staging | `LFX V2 - Staging` |
| Production | `LFX V2 - Production` |

The name of the secret item should be the same in all vaults it is configured for and match the
`source.onepassword.item` field in the configuration.

### AWS Destination Accounts

Secrets are deployed to these AWS accounts by environment:

| Environment | AWS Account |
| ----------- | ----------- |
| Development | `lfx-development` |
| Staging | `lfx-staging` |
| Production | `lfx-production` |

**Note**: Secrets do **not** need to be added to `prdct-*` accounts.

### Region

All secrets are stored in the **us-west-2** region.

### Path Format

Secrets follow this path convention in AWS Secrets Manager:

```text
path: cloud/<3rd_party_service>/<name_or_identifier>
```

They are automatically prefixed with `/cloudops/managed-secrets/` when created.

The name pattern here does not matter as much as the tags do, which determine which secret
store a secret is pulled into.

**Examples:**

- `cloud/LiteLLM/lfx-v2`
- `cloud/zoom/lfx-v2-meeting-service`
- `cloud/supabase/api_key`

## Step-by-Step Process

### 1. Add Secret to 1Password

First, add your secret to the appropriate 1Password vault:

1. Open 1Password and navigate to the correct vault (see [Vault Requirements](#vault-requirements))
2. Create a new item or update an existing one
3. Use type `API Credentials`
4. Include all necessary fields (API keys, tokens, etc.)
5. The chosen field names will be used as keys in the AWS secret in and ultimately in the k8s secret
   too. For example, `LiteLLM LFXv2 Key` has fields `litellm-key-id` and `litellm-api-key` and has
   `source.onepassword.json_fields` set to `['litellm-key-id', 'litellm-api-key']`.
6. Use a descriptive name that identifies the service and purpose.

### 2. Create YAML Configuration

Navigate to the [lfx-secrets-management](https://github.com/linuxfoundation/lfx-secrets-management)
repository and create or update a YAML configuration file in the `secretsmanagement/secrets/` directory.

Do this via a new branch and pull request against the `main` branch.

### 3. Submit Pull Request

Create a pull request with your configuration changes following the repository's contribution guidelines.

### 4. Deploy Secrets

Once approved and merged, the secrets will be deployed via
[**GitHub Actions**](https://github.com/linuxfoundation/lfx-secrets-management/actions/workflows/deploy.yml)

**If the secrets store and external secrets operator are already set up, the secrets will be automatically
available to the service if tagged correctly.**

### 5. Using the Secret in Kubernetes

The secret, once ingested, can be used by your deployments by reference. Here's an example snippet of a deployment manifest:

```yaml
env:
{{- range $name, $config := .Values.environment }}
- name: {{ $name }}
  {{- if $config.value }}
  value: {{ $config.value | quote }}
  {{- else if $config.valueFrom }}
  valueFrom:
    {{- toYaml $config.valueFrom | nindent 14 }}
  {{- end }}
{{- end }}
```

and the in the corresponding `values.yaml`:

```yaml
environment:
  LITELLM_API_KEY:
    valueFrom:
      secretKeyRef:
        name: pcc-secrets
        key: litellm-api-key
  LITELLM_KEY_ID:
    valueFrom:
      secretKeyRef:
        name: pcc-secrets
        key: litellm-key-id
```

The above is dependent on how the helm chart is set up.

> **New service setup:** Steps 1–5 assume the service already has a SecretStore,
> ExternalSecret, and ServiceAccount configured. If this is a **new** service, complete
> Steps 6–12 first — they are a one-time setup spanning three repositories that must be
> done in order. For the full IAM role field reference, see the
> [Service Account Management documentation](https://github.com/linuxfoundation/lfx-v2-opentofu/blob/main/docs/service-accounts.md).

### 6. Register the IAM Service Account Role

Each LFX V2 service needs a dedicated IAM role so Kubernetes can authenticate to AWS on its
behalf via IRSA (IAM Roles for Service Accounts). This role also grants the External Secrets
Operator permission to read secrets tagged with the service identifier.

In the [lfx-v2-opentofu](https://github.com/linuxfoundation/lfx-v2-opentofu) repository, add an
entry to `iam-service-accounts-definitions.yaml`:

```yaml
service_account_roles:
  lfx-v2-myresource-service:
    namespace: "myresource-service"
    service_account: "lfx-v2-myresource-service"
    eso_service_tag: "myresource"   # must match the service tag used in Step 2
```

The `eso_service_tag` must match the `service` tag value set on the AWS Secrets Manager secret
in Step 2.

When this PR is merged and applied by CI, OpenTofu creates per environment:

- An IAM role `lfx-v2-myresource-service` with an OIDC trust policy bound to
  `system:serviceaccount:myresource-service:lfx-v2-myresource-service`
- A tag-scoped `secretsmanager:GetSecretValue` policy granting access only to secrets tagged
  `service = myresource`

The role is created in all three AWS accounts:

| Environment | Account ID | Role ARN |
| ----------- | ---------- | -------- |
| Development | `788942260905` | `arn:aws:iam::788942260905:role/lfx-v2-myresource-service` |
| Staging | `844790888233` | `arn:aws:iam::844790888233:role/lfx-v2-myresource-service` |
| Production | `372256339901` | `arn:aws:iam::372256339901:role/lfx-v2-myresource-service` |

Open a PR against `lfx-v2-opentofu` and have the platform team merge and apply it before
continuing to the chart steps below.

### 7. Add the ServiceAccount Template

In the service's Helm chart, create
`charts/lfx-v2-myresource-service/templates/serviceaccount.yaml`:

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
{{- if .Values.serviceAccount.create }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .Values.serviceAccount.name | default .Chart.Name }}
  namespace: {{ .Release.Namespace }}
  labels:
    app: {{ .Chart.Name }}
  {{- with .Values.serviceAccount.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
{{- end }}
```

The `annotations` block is empty by default and is populated per environment in `lfx-v2-argocd`
with the IRSA role ARN (Step 11).

### 8. Add the SecretStore Template

Create `charts/lfx-v2-myresource-service/templates/secretstore.yaml`:

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
{{ if and .Values.externalSecretsOperator.enabled .Values.global.awsRegion }}
---
apiVersion: external-secrets.io/v1
kind: SecretStore
metadata:
  name: {{ .Chart.Name }}
  namespace: {{ .Release.Namespace }}
spec:
  provider:
    aws:
      service: "SecretsManager"
      region: {{ .Values.global.awsRegion }}
      auth:
        jwt:
          serviceAccountRef:
            name: {{ .Values.serviceAccount.name | default .Chart.Name }}
{{- end }}
```

The `serviceAccountRef` instructs ESO to present the ServiceAccount's OIDC token when calling
AWS. Kubernetes projects this token with the OIDC subject that satisfies the IRSA trust policy
created in Step 6, allowing ESO to assume the role and read tagged secrets.

### 9. Add the ExternalSecret Template

Create `charts/lfx-v2-myresource-service/templates/externalsecret.yaml` using one of two
patterns.

#### Tag discovery (recommended)

Use this pattern when secrets are managed through `lfx-secrets-management`. ESO automatically
discovers every secret tagged `service: myresource` (or `service-myresource: enabled`) and
merges their fields into a single Kubernetes Secret. Adding future secrets in Steps 1–4 requires
no further chart changes — the tag that controls IAM access (Step 6) and the tag that drives
discovery are the same value.

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
{{ if and .Values.externalSecretsOperator.enabled .Values.global.awsRegion }}
---
apiVersion: external-secrets.io/v1
kind: ExternalSecret
metadata:
  name: {{ .Chart.Name }}
  namespace: {{ .Release.Namespace }}
spec:
  refreshInterval: "{{ .Values.externalSecretsOperator.externalSecret.refreshInterval }}"
  secretStoreRef:
    name: {{ .Chart.Name }}
    kind: SecretStore
  target:
    name: {{ .Chart.Name }}
    creationPolicy: Owner
    deletionPolicy: Retain
  dataFrom:
    - find:
        tags:
          service: {{ .Values.externalSecretsOperator.externalSecret.serviceTag }}
      rewrite:
        - merge: {}
    - find:
        tags:
          service-{{ .Values.externalSecretsOperator.externalSecret.serviceTag }}: "enabled"
      rewrite:
        - merge: {}
{{- end }}
```

Both `find` blocks are required to support the two tagging conventions used across the
infrastructure (the tag-scoped IAM policy created in Step 6 accepts both).

#### Explicit data (alternative)

Use this pattern when secrets come from cloudops-managed paths that do not carry the service tag,
when you need custom Kubernetes key names, or when you want to sync only specific fields from a
secret. Adding a new secret requires extending the `data` list in `lfx-v2-argocd`.

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
{{ if and .Values.externalSecretsOperator.enabled .Values.global.awsRegion }}
---
apiVersion: external-secrets.io/v1
kind: ExternalSecret
metadata:
  name: {{ .Chart.Name }}
  namespace: {{ .Release.Namespace }}
spec:
  refreshInterval: "{{ .Values.externalSecretsOperator.externalSecret.refreshInterval }}"
  secretStoreRef:
    name: {{ .Chart.Name }}
    kind: SecretStore
  target:
    name: {{ .Chart.Name }}
    creationPolicy: Owner
    deletionPolicy: Retain
  {{- if .Values.externalSecretsOperator.externalSecret.data }}
  data:
    {{- range .Values.externalSecretsOperator.externalSecret.data }}
    - secretKey: {{ .secretKey }}
      remoteRef:
        key: {{ .remoteRef.key }}
        {{- if .remoteRef.property }}
        property: {{ .remoteRef.property }}
        {{- end }}
    {{- end }}
  {{- end }}
{{- end }}
```

Each entry maps a Kubernetes Secret key (`secretKey`) to a specific AWS Secrets Manager path
(`remoteRef.key`) and field (`remoteRef.property`).

### 10. Add the values.yaml Defaults

Add these blocks to `charts/lfx-v2-myresource-service/values.yaml`:

```yaml
serviceAccount:
  create: true
  name: "lfx-v2-myresource-service"
  annotations: {}

global:
  awsRegion: ""

externalSecretsOperator:
  enabled: false
  externalSecret:
    refreshInterval: "10m"
    # For tag discovery (recommended):
    serviceTag: "myresource"
    # For explicit data (alternative), define a data list instead:
    # data:
    #   - secretKey: my-key
    #     remoteRef:
    #       key: /cloudops/managed-secrets/cloud/myresource/my-group
    #       property: my_field
```

Three values gate the ESO resources — all default to off so local development does not attempt
to reach AWS:

- **`serviceAccount.create`** — controls `serviceaccount.yaml`
- **`global.awsRegion`** — gates both `secretstore.yaml` and `externalsecret.yaml`
- **`externalSecretsOperator.enabled`** — gates both `secretstore.yaml` and `externalsecret.yaml`

### 11. Configure Per-Environment Values

In the [lfx-v2-argocd](https://github.com/linuxfoundation/lfx-v2-argocd) repository, create
values files for the service.

**`values/global/lfx-v2-myresource-service.yaml`** — enable ESO for all deployed environments:

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
---

externalSecretsOperator:
  enabled: true
```

**`values/dev/lfx-v2-myresource-service.yaml`** (repeat for staging and prod with the
matching account ID) — set the AWS region and the IRSA role ARN:

```yaml
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
---

global:
  awsRegion: "us-west-2"

serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::788942260905:role/lfx-v2-myresource-service
```

Account IDs and files per environment:

| Environment | Account ID | `values/` path |
| ----------- | ---------- | -------------- |
| Development | `788942260905` | `values/dev/lfx-v2-myresource-service.yaml` |
| Staging | `844790888233` | `values/staging/lfx-v2-myresource-service.yaml` |
| Production | `372256339901` | `values/prod/lfx-v2-myresource-service.yaml` |

### 12. Configure Local Development

The External Secrets Operator cannot reach AWS from a local development cluster. Add a block to
`lfx-v2-helm/charts/lfx-platform/values.yaml` to keep ESO disabled locally:

```yaml
lfx-v2-myresource-service:
  # External Secrets Operator is disabled because we can't get secrets from
  # AWS Secrets Manager in local development.
  # Instead, create the Kubernetes secret manually. See the chart README.
  externalSecretsOperator:
    enabled: false
```

Create the Kubernetes Secret manually in your local cluster to supply values during development.
See the service chart `README.md` for the exact `kubectl create secret` command.

**PR order:** Merge changes in this sequence — (1) `lfx-v2-opentofu` to create the IAM role,
(2) the service chart to add the templates, (3) `lfx-v2-argocd` to activate ESO per environment.

Once wired, future secrets only require Steps 1–4 (tag discovery) or Steps 1–4 plus extending
the `externalSecretsOperator.externalSecret.data` list in the argocd values (explicit data).

## Configuration Example

Based on the [LiteLLM implementation](https://github.com/linuxfoundation/lfx-secrets-management/pull/96),
here's a complete configuration example:

```yaml
LiteLLM API key for LFXv2:
  tags: [lfx_v2, litellm, pcc]
  environments: [development, staging, production]
  source:
    onepassword:
      vaults:
        development: LFX V2 - Development
        staging: LFX V2 - Staging
        production: LFX V2 - Production
      item: LiteLLM LFXv2 Key
      json_fields:
        - litellm-key-id
        - litellm-api-key
  destinations:
    - aws_secretsmanager:
        tags:
          service: pcc
        accounts:
          development: lfx-development
          staging: lfx-staging
          production: lfx-production
        regions:
          - us-west-2
        path: cloud/litellm/lfx-v2
```

### Configuration Breakdown

- **Name**: Descriptive title for the secret configuration
- **Tags**: Must include `lfx_v2`, upstream service (`litellm`), and backend service (`pcc`)
- **Environments**: List of environments where this secret should be deployed
- **Source**: 1Password configuration with vault mappings and item details
- **source.onepassword.json_fields**: Specifies which fields from the 1Password item to include in the secret
- **source.onepassword.vaults**: Maps environments to their respective 1Password vaults
- **source.onepassword.item**: Name of the 1Password item containing the secret, must match exactly in all vaults
- **Destinations**: AWS Secrets Manager configuration with accounts, regions, and path
- **destinations.aws_secretsmanager.tags.service**: Tags the secret and determines if a secret gets
  pulled into a service's secret store
- **destinations.aws_secretsmanager.accounts**: Maps environments to AWS accounts, always use lfx-*
  accounts for LFX V2

## Deployment Methods

### Using GitHub Actions

1. Navigate to the
   [Deploy Workflow](https://github.com/linuxfoundation/lfx-secrets-management/actions/workflows/deploy.yml)
2. Click "Run workflow"
3. Enter space-separated values:
   - **Tags**: Include your tags (e.g., `lfx_v2 litellm pcc`)
   - **Environments**: Specify target environments (e.g., `development staging production`)

## Validation and Testing

### Verify Deployment

After deployment, verify the secret exists in AWS Secrets Manager:

1. Access the AWS Console for the target environment
2. Navigate to AWS Secrets Manager in the us-west-2 region
3. Search for your secret using the configured path
4. Verify the secret contains the expected values

### Test Service Access

Ensure your LFX V2 services can access the secret:

1. Update service configuration to reference the secret path
2. Test the service in a non-production environment
3. Monitor logs for any access issues
4. Verify the secret values are correctly retrieved

### Audit Secrets

Use the secrets management repository's audit functionality:

```bash
# Audit all secrets
make audit

# Audit specific environment
make audit ENVS="development"

# Audit specific AWS accounts
make audit-aws ACCOUNTS="lfx-development"
```

## Best Practices

### Security Considerations

- **Never store secrets in plain text** in any repository or configuration file
- **Use environment-specific vaults** to prevent accidental cross-environment access
- **Rotate secrets regularly** according to your organization's security policy
- **Monitor secret access** using AWS CloudTrail and 1Password audit logs

### Naming Conventions

- Use descriptive names that clearly identify the service and purpose
- Include version numbers or identifiers when managing multiple versions
- Follow consistent naming patterns across all LFX V2 secrets

### Environment Management

- **Always test in development first** before deploying to staging or production
- **Use the plan mode** (`--plan` flag) to verify changes before deployment
- **Deploy incrementally** (dev → staging → production) to catch issues early

### Change Management

- **Create pull requests** for all secret configuration changes
- **Include detailed descriptions** of what secrets are being added or modified
- **Tag relevant team members** for review, especially for production secrets
- **Document any service configuration changes** required to use the new secrets

## Troubleshooting

### Common Issues

**Secret not found in AWS:**

- Verify the secret was successfully deployed using the audit commands
- Check that the path matches your service configuration exactly
- Ensure the secret was deployed to the correct AWS account and region

**1Password access errors:**

- Verify you have access to the correct vault for your environment
- Check that the item name and field names match your configuration exactly
- Ensure the 1Password service account has appropriate permissions

**Deployment failures:**

- Review the deployment logs in GitHub Actions or local output
- Verify your configuration YAML syntax is correct
- Check that all required tags and fields are present

### Getting Help

For additional support:

- Review the [lfx-secrets-management repository documentation](https://github.com/linuxfoundation/lfx-secrets-management/blob/main/README.md)
- Check existing issues and discussions in the repository
- Contact the LFX engineering team through established channels
- Reference the
  [services documentation](https://github.com/linuxfoundation/lfx-secrets-management/blob/main/secretsmanagement/services/README.md)
  for service-specific guidance
