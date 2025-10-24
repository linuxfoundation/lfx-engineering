# LFX v2 Local Development Setup

## Overview

This guide walks you through setting up a local development environment for LFX v2 applications using OrbStack and Helm.

## Prerequisites

### Required Tools

1. **OrbStack**
   - Install OrbStack following the [official installation guide](https://docs.orbstack.dev/install)
   - OrbStack provides a lightweight Docker Desktop alternative with Kubernetes support

2. **Helm**
   - Install Helm package manager for Kubernetes
   - Follow installation instructions at [helm.sh](https://helm.sh/docs/intro/install/)

## Platform Setup

### Install LFX Platform

1. **Clone the Helm Repository**

   ```bash
   git clone https://github.com/linuxfoundation/lfx-v2-helm
   cd lfx-v2-helm
   ```

2. **Navigate to Charts Directory**

   ```bash
   cd charts/lfx-platform
   ```

3. **Install the Platform**

   ```bash
   helm install -n lfx --create-namespace lfx-platform . -f values.yaml
   ```

   This command will:
   - Create the `lfx` namespace if it doesn't exist (`--create-namespace`)
   - Install all required platform components
   - Configure services according to the values.yaml file

## Access Points

Once the platform is successfully installed, you'll have access to:

### API Documentation

- **URL**: <http://lfx-api.k8s.orb.local/docs/>
- Provides interactive API documentation for all LFX v2 services
- Includes endpoint specifications, request/response schemas, and testing capabilities

### Authentication (Authelia)

- **URL**: <https://auth.k8s.orb.local/>
- Central authentication service for the local development environment
- Manages user authentication and authorization

## User Management

### Adding Users

Users are configured through the Helm values file:

- Edit the `values.yaml` file in the helm chart
- Add user configurations in the appropriate section
- Update the deployment: `helm upgrade -n lfx lfx-platform . -f values.yaml`

### Retrieving User Passwords

Passwords are stored in Kubernetes secrets and can be retrieved using kubectl:

```bash
# Get password for project_super_admin user
kubectl get secret/authelia-users -n lfx -o jsonpath='{.data.project_super_admin}' | base64 --decode
```

**General pattern for any user:**

```bash
kubectl get secret/authelia-users -n lfx -o jsonpath='{.data.USERNAME}' | base64 --decode
```

Replace `USERNAME` with the actual username configured in your values file.

### Generating Authentication Tokens

#### Local Environment Tokens (Authelia)

For local development, Authelia tokens can be generated using the token helper script:

- **Script**: [token_helper.py](https://github.com/linuxfoundation/lfx-architecture-scratch/blob/main/2024-12%20ReBAC%20Demo/token_helper/token_helper.py)
- **Note**: The token URL may need to be updated to use `https` (see [line 221](https://github.com/linuxfoundation/lfx-architecture-scratch/blob/main/2024-12%20ReBAC%20Demo/token_helper/token_helper.py#L221))

#### Remote Environment Tokens (Dev/Staging/Production)

For development, staging, and production environments, use the Auth0 token generation script:

- **Script**: [get_auth0_token.py](https://github.com/linuxfoundation-it/itx-misc/blob/main/libs/get_auth0_token.py)
- Supports token generation for all remote environments

This is typically used as follows:

```bash
export TOKEN=$(get_auth0_token.py --api=lfxv2 --stage=dev)
```

Which will open a browser window for authentication and set the `TOKEN` environment variable
to the generated token when complete.

## Development Workflow

1. **Platform Setup**: Complete the installation steps above
2. **Service Development**: Develop your individual services locally
3. **Integration Testing**: Use the local platform to test service integration
4. **API Testing**: Utilize the API docs at <http://lfx-api.k8s.orb.local/docs/>
5. **Authentication Testing**: Test authentication flows via <https://auth.k8s.orb.local/>

## Troubleshooting

### Common Issues

- **DNS Resolution**: Ensure OrbStack DNS is properly configured for `.orb.local` domains
- **Port Conflicts**: Check that required ports are not in use by other services
- **Resource Limits**: Verify sufficient system resources for the platform components

### Helpful Commands

```bash
# Check platform status
kubectl get pods -n lfx

# View platform services
kubectl get services -n lfx

# Check platform logs
kubectl logs -n lfx -l app=your-service-name

# Reset platform (if needed)
helm uninstall -n lfx lfx-platform
```

## Next Steps

After completing the local development setup:

1. Review the [Release Strategy](./release-strategy.md) for deployment workflows
2. Configure your individual services to integrate with the local platform
3. Test authentication flows and API integrations
4. Follow established development practices for the LFX v2 ecosystem
