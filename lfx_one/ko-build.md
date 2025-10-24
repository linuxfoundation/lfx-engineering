# Ko Build for LFX V2 Services

This document explains how to use [ko](https://ko.build/) to build container images for Go-based services in the LFX v2 platform.

## What is Ko?

Ko is a simple, fast container image builder for Go applications. It builds images by directly inserting
the Go binary into a base image, without requiring a Dockerfile. This makes it ideal for building lightweight,
secure container images for Go-based microservices.

By default ko builds and pushes images in the same command, but for local development you can use the
`--local` flag to build images locally without pushing to a registry.

## Prerequisites

1. **Install Ko**: Follow the [official ko installation guide](https://ko.build/install/)
2. **Docker or compatible container runtime**: Ko needs a local container runtime
3. **Go development environment**: Ensure Go is installed and properly configured
4. **Access to the service repository**: Clone the LFX v2 service repository you want to build

## Basic Usage

### Building a Local Container Image

To build a container image locally for an LFX v2 service, for example the project service API, run:

```bash
> cd ~/git/linuxfoundation/lfx-v2-project-service
> ko build --local -B github.com/linuxfoundation/lfx-v2-project-service/cmd/project-api

2025/08/28 09:20:18 Using base cgr.dev/chainguard/static:latest@sha256:6a4b683f4708f1f167ba218e31fcac0b7515d94c33c3acf223c36d5c6acd3783
  for github.com/linuxfoundation/lfx-v2-project-service/cmd/project-api
2025/08/28 09:20:18 git is in a dirty state
Please check in your pipeline what can be changing the following files:
 M internal/middleware/request_logger.go
?? catalog-info.yml

2025/08/28 09:20:18 Building github.com/linuxfoundation/lfx-v2-project-service/cmd/project-api for linux/amd64
2025/08/28 09:20:22 Loading ko.local/project-api:fefcb3f5390f5eb1f2d221c437531b20edf12db0e902d020432a651cf5b1882e
2025/08/28 09:20:23 Loaded ko.local/project-api:fefcb3f5390f5eb1f2d221c437531b20edf12db0e902d020432a651cf5b1882e
2025/08/28 09:20:23 Adding tag latest
2025/08/28 09:20:23 Added tag latest
```

ko will warn you if the git checkout is in a dirty state with uncommitted changes. You can also see that it
adds a `latest` tag by default. This will not be added if other tags are specified.

You can build multiple images in the same repository by specifying other package paths. A further example
for the project service is:

```bash
> ko build --local -B github.com/linuxfoundation/lfx-v2-project-service/scripts/root-project-setup
```

### Command Breakdown

- `ko build`: The main ko build command
- `--local`: Build the image locally instead of pushing to a registry
- `-B`: base import paths: Whether to use the base path without MD5 hash
- `github.com/linuxfoundation/lfx-v2-project-service/cmd/projects-api`: The Go package path to build

### Tags

You can also specify a specific tag using the `-t` flag. This can be useful for building the same tag
during local testing & development:

```bash
ko build --local -B -t latest github.com/linuxfoundation/lfx-v2-project-service/cmd/project-api
```

### Usage with Helm

When using ko-built images with Helm charts, you can reference the locally built image in your Helm values. For example:

```yaml
  lfx-v2-project-service:
    image:
      repository: ko.local/project-api
      tag: latest
```

## Configuration

### Environment Variables

Ko respects several environment variables:

- `KO_DATA_PATH`: Set the path for ko data storage
- `KO_DOCKER_REPO`: Set the default registry for pushed images
- `CGO_ENABLED=0`: Recommended for static binaries

### KO_DATA_PATH

If there is a directory called `kodata` in the build path, ko will automatically include this in the image and set the `KO_DATA_PATH`
environment variable to point to this directory. This is useful for including static assets.

We use this for including the OpenAPI spec files in the service containers. It works automatically for ko
build, but if you run the services locally via `go run` or `go build` you will need to set the environment
variable manually.

## Integration with LFX V2 Development Workflow

### Local Development

1. **Build the service locally**:

   ```bash
   ko build --local -B <service-package-path>
   ```

2. **Use with OrbStack/Helm**: The locally built image can be referenced in your Helm values for local development

3. **Testing**: Run the container locally to test your changes before committing

### GitHub Actions Integration

ko can be easily integrated into GitHub Actions and is the expected method for building go service images.

Examples:

- [LFX v2 Project Service main branch](https://github.com/linuxfoundation/lfx-v2-project-service/blob/main/.github/workflows/ko-build-main.yaml)
- [LFX v2 Project Service Tagged Releases](https://github.com/linuxfoundation/lfx-v2-project-service/blob/main/.github/workflows/ko-build-tag.yaml)

### Viewing Build Context

Check what ko includes in the build:

```bash
ko build --local -B --tarball=output.tar <package-path>
```

## Additional Resources

- [Ko Official Documentation](https://ko.build/)
- [Ko GitHub Repository](https://github.com/ko-build/ko)
- [LFX v2 Helm Documentation](https://github.com/linuxfoundation/lfx-v2-helm/blob/main/README.md)
- [Local Development Setup](./local-development.md)
