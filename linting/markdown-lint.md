# Markdown Lint

## Command Line Linting

If developers have a container runtime installed on their local machine, they
can use the markdownlint-cli2 docker image to lint markdown files.

```bash
docker run --rm -v "$(pwd):/workdir" davidanson/markdownlint-cli2:latest "**/*.md"
```

Alternatively, developers can install the `markdownlint-cli2` package globally
and run the linting command from the command line.

An example of how to install the markdownlint-cli2 package globally is shown
below:

```bash
npm install -g markdownlint-cli2
```

Then, developers can run the linting command from the command line.

```bash
markdownlint "**/*.md"
```

## GitHub Actions

To lint markdown files in a GitHub Actions workflow, you can use the following example
YAML file and place it in the `.github/workflows` directory. A suggested name
for the file is `markdown-lint.yml`.

```yaml
---
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

name: Markdown Lint

on:
  pull_request:
    paths:
      - '**/*.md'
      - '.markdownlint.json'

jobs:
  markdown-lint:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: markdownlint-cli2-action
      uses: DavidAnson/markdownlint-cli2-action@v9
```

### GitHub Actions Local Testing

To test the GitHub Actions workflow locally, you can use the following command:

```bash
act -W .github/workflows/markdown-lint.yml
```

More information on the `act` command [can be found on the GitHub page](https://github.com/nektos/act).
