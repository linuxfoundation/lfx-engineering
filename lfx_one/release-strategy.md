# LFX v2 Release Strategy

## Overview

This document outlines the release strategy for LFX v2 applications using ArgoCD for deployment management
and Helm for package management.

## Key Repositories

- **Argo Repository**: [lfx-v2-argocd](https://github.com/linuxfoundation/lfx-v2-argocd)
- **Helm Repository**: [lfx-v2-helm](https://github.com/linuxfoundation/lfx-v2-helm)

## Release Workflows

### Development Environment

For continuous development and testing:

#### Process

1. **Automatic Development Builds**
   - Merges to `main` branch automatically build a `development` tag of the application
   - No formal release process required for development deployments

2. **ArgoCD Image Updater Configuration**
   - Argo deployment for development environment uses ArgoCD Image Updater
   - Strategy: `digest` - automatically consumes the latest version of the `development` tag
   - Enables continuous deployment of development changes without manual intervention

### Existing Service Releases

For services that already exist in the deployment pipeline:

#### Prerequisites

- Helm configuration uses `latest` tags for all components

#### Existing Service Process

1. **Release Service Version**
   - Release your service with version tag (e.g., `v0.1.2`)

2. **Update Staging Environment**
   - Merge changes to argo repository
   - Update staging configuration to reference new service version (`v0.1.2`)

3. **Update Production Environment**
   - Merge changes to argo repository
   - Update production configuration to reference new service version (`v0.1.2`)

### New Component Releases

For new services being added to the deployment pipeline:

#### New Component Process

1. **Release Service Version**
   - Release your service with version tag (e.g., `v0.1.2`)

2. **Update Helm Configuration**
   - Release umbrella helm chart with new subchart for your service
   - Update relevant tags/dependencies for production dependencies
   - Use "latest" tag for new components

3. **Coordinate with SysOps**
   - **CRITICAL**: Work with SysOps team to ensure all required secrets are present in the target cluster
   - Verify secret availability before proceeding with deployment configuration

4. **Update Staging Environment**
   - Merge changes to argo repository
   - Add new ApplicationSet (AppSet) for your service
   - Configure staging instance/folder with service version (`v0.1.2`)

5. **Update Production Environment**
   - Merge changes to argo repository
   - Add new ApplicationSet (AppSet) for your service
   - Configure production instance/folder with service version (`v0.1.2`)

## Important Notes

- **Development Automation**: Development environment automatically deploys latest changes from `main` branch
- **Helm Strategy**: All Helm configurations assume `latest` tags for dependencies
- **Secret Management**: New components require coordination with SysOps for secret provisioning
- **Environment Progression**: Changes flow from development → staging → production
- **Version Consistency**: Both staging and production should reference the same service version after deployment
- **Image Updater**: Development environment uses digest strategy for automatic updates

## Best Practices

1. Use development environment for continuous integration and initial testing
2. Always test in staging environment before promoting to production
3. Coordinate with SysOps early in the process for new components
4. Ensure proper version tagging for traceability
5. Verify secret availability before deployment configuration
6. Follow the established workflow sequence to maintain deployment consistency
7. Monitor development deployments for automatic updates via ArgoCD Image Updater
