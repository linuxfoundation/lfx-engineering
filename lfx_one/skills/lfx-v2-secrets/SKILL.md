---
name: lfx-v2-secrets
description: >
  Guide an agent through wiring up secrets for LFX V2 microservices using External Secrets
  Operator (ESO) + IRSA. Supports two modes: (1) full setup for new services touching
  lfx-v2-opentofu, lfx-secrets-management, the service Helm chart, and lfx-v2-argocd;
  (2) adding secrets to existing services already configured with ESO. Use this skill
  whenever someone says "set up secrets", "wire up ESO", "add a secret to this service",
  "IRSA configuration", "External Secrets for V2", or any mention of AWS Secrets Manager
  integration with Kubernetes for LFX V2 services.
allowed-tools: Bash, Read, Glob, Grep, AskUserQuestion, WebFetch
---

<!-- Copyright The Linux Foundation and each contributor to LFX. -->
<!-- SPDX-License-Identifier: MIT -->
<!-- Tool names in this file use Claude Code vocabulary. See docs/tool-mapping.md for other platforms. -->

# LFX V2 Secrets Setup Guide

Secrets for LFX V2 microservices are managed through **External Secrets Operator (ESO)**
combined with **IAM Roles for Service Accounts (IRSA)** on AWS. This provides a secure,
GitOps-driven way to sync secrets from AWS Secrets Manager into Kubernetes.

> **For the AI**: This skill has two modes. Mode 1 (new service) touches four repos
> and requires coordinated changes across infrastructure layers. Mode 2 (existing service)
> is much smaller. Always ask the user which applies before proceeding.

---

## Understanding the Architecture

### How It Works

1. A **Kubernetes ServiceAccount** is annotated with an **IRSA role ARN**
2. ESO's **SecretStore** uses that ServiceAccount's JWT token to authenticate to AWS
3. ESO watches **ExternalSecret** manifests and syncs matching secrets from AWS SM into K8s Secrets
4. Application deployments reference the K8s Secret via environment variable or volume mount
5. Local development skips ESO entirely and injects secret values directly via `extraEnv` in values

### Key Constants

These values are fixed and apply across all V2 services:

| Item | Value |
|------|-------|
| AWS Region | `us-west-2` |
| K8s Secret name | `{{ .Chart.Name }}` (e.g., `lfx-v2-invite-service`) |
| SecretStore name | `{{ .Chart.Name }}` |
| IAM account — dev | `788942260905` |
| IAM account — staging | `844790888233` |
| IAM account — prod | `372256339901` |
| IRSA role ARN pattern | `arn:aws:iam::<account-id>:role/lfx-v2-<service>` |
| AWS SM path pattern | `cloud/<service-short-name>/<secret-group>` |
| ServiceAccount annotation key | `eks.amazonaws.com/role-arn` |
| ESO JWT auth field | `spec.provider.aws.auth.jwt.serviceAccountRef` |

---

## Mode 1: New Service (Full Setup)

Use this mode when a brand-new V2 service needs secrets wired up end-to-end from scratch.
This touches **four repos** and requires changes in a specific order.

### Step 1: Prepare Information

Ask the user to collect:

1. **Service name** — short form (e.g., `invite-service`, `email-service`)
2. **Service short name** — lowercase, dash-separated (e.g., `invite`, `email`)
3. **List of secrets** — each should have:
   - **Secret label** (e.g., "JWT Secret", "SMTP Credentials")
   - **Secret group** (e.g., `jwt`, `smtp`) — used in AWS SM path
   - **Field names in 1Password** — the exact field names as they appear in the source vault
4. **Which environments need this secret** — usually `development`, `staging`, `production`

Example:

```text
Service: invite-service (short: invite)
Secrets:
  - JWT Secret (group: jwt, field: secret_key) — all envs
  - Database password (group: database, field: password) — all envs
```

### Step 2: Create IAM Service Account in `lfx-v2-opentofu`

In the [lfx-v2-opentofu](https://github.com/linuxfoundation/lfx-v2-opentofu) repo,
edit `iam-service-accounts-definitions.yaml` and add:

```yaml
service_account_roles:
  lfx-v2-<service>:
    namespace: "<service-short-name>-service"
    service_account: "lfx-v2-<service>"
    eso_service_tag: "<service-short-name>"
```

Example for invite service:

```yaml
service_account_roles:
  lfx-v2-invite-service:
    namespace: "invite-service"
    service_account: "lfx-v2-invite-service"
    eso_service_tag: "invite"
```

> **Note**: The `eso_service_tag` is used in lfx-secrets-management to tag all related secrets.
> The namespace and service account names follow the LFX V2 convention.

### Step 3: Create Sync Entries in `lfx-secrets-management`

In the [lfx-secrets-management](https://github.com/linuxfoundation/lfx-secrets-management) repo,
edit `secrets/lfx/cloud.yml` and add an entry for each secret:

```yaml
LFX V2 <Service> <Secret Label>:
  tags: [lfx_v2, <service_tag>, <type_tag>]
  envs: [development, staging, production]
  source:
    onepassword:
      vaults:
        development: LFX V2 - Development
        staging: LFX V2 - Staging
        production: LFX V2 - Production
      item: "LFX V2 <Service> - <Secret Label>"
      fields: <field_name>
  destinations:
    - aws_secretsmanager:
        regions: us-west-2
        path: "cloud/<service-short-name>/<secret-group>"
```

Example for invite service JWT secret:

```yaml
LFX V2 Invite Service JWT Secret:
  tags: [lfx_v2, invite, jwt]
  envs: [development, staging, production]
  source:
    onepassword:
      vaults:
        development: LFX V2 - Development
        staging: LFX V2 - Staging
        production: LFX V2 - Production
      item: "LFX V2 Invite Service - JWT Secret"
      fields: secret_key
  destinations:
    - aws_secretsmanager:
        regions: us-west-2
        path: "cloud/invite/jwt"
```

> **Tips**:
>
> - Each secret in the lfx-secrets-management source becomes a separate AWS SM path entry
> - The `path` convention is `cloud/<service-short-name>/<secret-group>`
> - Use the `envs` list to sync to all three environments in parallel
> - The `source.onepassword.item` should match exactly the name in 1Password vaults

### Step 4: Create Helm Chart Files in the Service Repo

In the service's Helm chart repo (e.g., `lfx-v2-invite-service`), create three files
under `charts/lfx-v2-<service>/templates/`:

#### `serviceaccount.yaml`

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

#### `secretstore.yaml`

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

#### `externalsecret.yaml`

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

#### Update `values.yaml`

Add the following blocks to the `values.yaml` file:

```yaml
serviceAccount:
  create: true
  name: "lfx-v2-<service>"
  annotations: {}

global:
  awsRegion: ""

externalSecretsOperator:
  enabled: false
  externalSecret:
    refreshInterval: "10m"
    data:
      - secretKey: secret_key
        remoteRef:
          key: "cloud/<service-short-name>/jwt"
          property: field_name
```

#### Wire into `deployment.yaml`

In your deployment template, reference the secret with a guard clause:

```yaml
env:
  {{- if .Values.app.jwtSecretName }}
  - name: JWT_SECRET
    valueFrom:
      secretKeyRef:
        name: {{ .Values.app.jwtSecretName }}
        key: secret_key
  {{- else if .Values.app.extraEnv }}
    # use extraEnv for local dev
  {{- end }}
  {{- with .Values.app.extraEnv }}
  {{- toYaml . | nindent 2 }}
  {{- end }}
```

Add to `values.yaml`:

```yaml
app:
  jwtSecretName: ""  # set by lfx-v2-argocd for non-local-dev envs
  extraEnv: []       # local dev can set this directly
```

### Step 5: Update Values in `lfx-v2-argocd`

In the [lfx-v2-argocd](https://github.com/linuxfoundation/lfx-v2-argocd) repo,
update or create value files for the new service.

#### `values/global/lfx-v2-<service>.yaml`

Create a new file (or append if updating an existing service) with base configuration:

```yaml
---
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

serviceAccount:
  create: true
  name: "lfx-v2-<service>"
  annotations: {}

externalSecretsOperator:
  enabled: true
  externalSecret:
    refreshInterval: "10m"
    data:
      - key: secret_key
        path: "cloud/<service-short-name>/jwt"
      - key: database_password
        path: "cloud/<service-short-name>/database"

app:
  jwtSecretName: "lfx-v2-<service>"
  extraEnv:
    - name: HOST_IP
      valueFrom:
        fieldRef:
          fieldPath: status.hostIP
```

#### `values/dev/lfx-v2-<service>.yaml`

```yaml
---
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

global:
  awsRegion: "us-west-2"

serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: "arn:aws:iam::788942260905:role/lfx-v2-<service>"
```

#### `values/staging/lfx-v2-<service>.yaml`

```yaml
---
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

global:
  awsRegion: "us-west-2"

serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: "arn:aws:iam::844790888233:role/lfx-v2-<service>"
```

#### `values/prod/lfx-v2-<service>.yaml`

```yaml
---
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

global:
  awsRegion: "us-west-2"

serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: "arn:aws:iam::372256339901:role/lfx-v2-<service>"
```

#### Local Development Values (in `lfx-v2-helm` umbrella chart)

In `lfx-v2-helm/charts/lfx-platform/values.yaml`, add a local-dev block:

```yaml
lfx-v2-<service>:
  externalSecretsOperator:
    enabled: false
  app:
    jwtSecretName: ""
    extraEnv:
      - name: JWT_SECRET
        value: "local-dev-jwt-secret-change-me"
      - name: HOST_IP
        value: "127.0.0.1"
```

> **Note**: Local dev never uses `secretKeyRef`. It injects via `extraEnv` with placeholder
> values that developers replace when needed.

---

## Mode 2: Add Secret to Existing Service

Use this mode when adding a new secret to a service that already has ESO + IRSA configured.

### Step 1: Add Entry to `lfx-secrets-management`

In the [lfx-secrets-management](https://github.com/linuxfoundation/lfx-secrets-management) repo,
edit `secrets/lfx/cloud.yml` and add a new entry for the new secret (same pattern as Mode 1, Step 3):

```yaml
LFX V2 <Service> <New Secret Label>:
  tags: [lfx_v2, <service_tag>, <type_tag>]
  envs: [development, staging, production]
  source:
    onepassword:
      vaults:
        development: LFX V2 - Development
        staging: LFX V2 - Staging
        production: LFX V2 - Production
      item: "LFX V2 <Service> - <New Secret Label>"
      fields: <field_name>
  destinations:
    - aws_secretsmanager:
        regions: us-west-2
        path: "cloud/<service-short-name>/<new-secret-group>"
```

### Step 2: Update ExternalSecret Data in `lfx-v2-argocd`

In `values/global/lfx-v2-<service>.yaml`, add to the `externalSecretsOperator.externalSecret.data` list:

```yaml
externalSecretsOperator:
  externalSecret:
    data:
      - key: secret_key
        path: "cloud/<service-short-name>/jwt"
      - key: new_secret_key           # add this
        path: "cloud/<service-short-name>/<new-secret-group>"
```

> **Note**: Only update per-environment values files if the AWS SM path differs
> per environment (rare). The global file provides the mapping template.

### Step 3: Wire into Service Deployment

In the service's `charts/lfx-v2-<service>/templates/deployment.yaml`,
add a new environment variable referencing the same K8s Secret (just a different key):

```yaml
- name: NEW_SECRET_KEY
  valueFrom:
    secretKeyRef:
      name: {{ .Chart.Name }}
      key: new_secret_key
```

---

## Verification Checklist

After completing either mode, verify the setup:

**File checklist — all repos involved:**

- [ ] `lfx-v2-opentofu`: `iam-service-accounts-definitions.yaml` updated with service entry
- [ ] `lfx-secrets-management`: `secrets/lfx/cloud.yml` has sync entries for all secrets
- [ ] Service chart:
  - [ ] `templates/serviceaccount.yaml` created
  - [ ] `templates/secretstore.yaml` created
  - [ ] `templates/externalsecret.yaml` created
  - [ ] `values.yaml` has `serviceAccount`, `global.awsRegion`, `externalSecretsOperator` blocks
  - [ ] `templates/deployment.yaml` wires secrets via `secretKeyRef` (non-local) and `extraEnv` (local)
- [ ] `lfx-v2-argocd`:
  - [ ] `values/global/lfx-v2-<service>.yaml` has ESO config and secret data mappings
  - [ ] `values/dev/lfx-v2-<service>.yaml` has dev account IRSA role ARN
  - [ ] `values/staging/lfx-v2-<service>.yaml` has staging account IRSA role ARN
  - [ ] `values/prod/lfx-v2-<service>.yaml` has prod account IRSA role ARN
  - [ ] (optional) `lfx-v2-helm` umbrella chart has local-dev overrides

**Configuration checks:**

- [ ] IRSA role ARN format is correct: `arn:aws:iam::<account-id>:role/lfx-v2-<service>`
- [ ] AWS SM paths follow pattern: `cloud/<service-short-name>/<secret-group>`
- [ ] All account IDs are correct (dev=788942260905, staging=844790888233, prod=372256339901)
- [ ] K8s Secret name matches `{{ .Chart.Name }}`
- [ ] SecretStore uses `spec.provider.aws.auth.jwt.serviceAccountRef` (not static credentials)

**1Password setup:**

- [ ] 1Password items named exactly as referenced in lfx-secrets-management sync entries
- [ ] Items live in the correct vaults (LFX V2 - Development/Staging/Production)
- [ ] Field names match exactly what's in lfx-secrets-management `fields:` values

---

## Reference Implementations

Real examples in the codebase:

| Service | What It Added | References |
|---------|---------------|-----------|
| Email Service | SMTP credentials | `lfx-v2-email-service` chart, lfx-v2-argocd values |
| Invite Service | JWT secret | `lfx-v2-invite-service` chart (LFXV2-1783), lfx-v2-argocd values |

Check these repos for the exact file structure and conventions used in production.

---

## Common Workflows

### Adding a JWT secret to a new service

1. User asks: "Set up JWT secret for invite-service"
2. Collect: service name, secret name in 1Password, environments
3. Update 4 repos following Mode 1 steps in order
4. Verify using the checklist above
5. Open a coordinated PR across repos (one PR per repo, cross-referenced)

### Adding SMTP credentials to an existing service

1. User asks: "Add SMTP secret to email-service"
2. Add sync entry in lfx-secrets-management
3. Add to ExternalSecret data in lfx-v2-argocd
4. Wire into deployment in service chart
5. Verify using the checklist above
6. Submit PRs (one PR per repo)

### Debugging: "Pods can't read the secret"

Check in order:

1. **Pod events** — `kubectl describe pod <pod>` to see if the SecretStore mounted
2. **ESO logs** — `kubectl logs -n external-secrets-system deployment/external-secrets`
3. **AWS SM permissions** — verify IRSA role has `SecretsManager:GetSecretValue` on the path
4. **Secret exists in AWS SM** — lfx-secrets-management automation has synced the secret
5. **ExternalSecret status** — `kubectl describe externalsecret <name>` shows sync status
6. **Topology/firewalling** — pod can reach AWS API endpoint (check SecurityGroup, NACL, DNS)

---

## Communication Style

This skill serves both platform engineers and application developers:

- **For experienced infrastructure engineers**: Use technical terms freely (IRSA, JWT auth, ESO operator).
- **For application developers touching secrets for the first time**: Explain what ESO is
  (*"it automatically copies secrets from AWS into Kubernetes"*) and IRSA (*"it proves
  your pod is who it claims to be when talking to AWS"*).
- **For non-technical users**: Avoid "Kubernetes", "IRSA", "operator", "manifest". Instead say
  "cloud setup", "permissions", "secure secret storage", "automated sync".

Always finish with the verification checklist so the user can confirm everything is wired correctly.
