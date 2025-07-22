# lfx-engineering

Repository to store a collection of artifacts and notes for engineers.

## Development

### Markdown Linting

This repository uses
[markdownlint](https://www.npmjs.com/package/markdownlint-cli) to ensure
consistent markdown formatting.

#### Setup

Install dependencies:

```bash
npm install
```

#### Running the Linter

Check markdown files:

```bash
npm run lint:md
```

Auto-fix markdown issues:

```bash
npm run lint:md:fix
```

#### Configuration

Markdown linting rules are configured in `.markdownlint.json`. The configuration:

- Uses ATX-style headers (# ## ###)
- Sets line length limit to 120 characters
- Allows nested headers with same content
- Permits HTML tags in markdown
- Disables first-line-h1 requirement for flexibility

## Running GitHub Workflows Locally

This document explains how to run the license header check workflow locally
without pushing to GitHub.

### Using Act (GitHub Actions Locally)

To run the actual GitHub workflow locally, you can use `act`:

#### Install Act

```bash
# On macOS with Homebrew
brew install act

# Or download from https://github.com/nektos/act
```

#### Run the Workflow

```bash
# Run all workflows
act

# Run only the pull_request event workflows
act pull_request

# Run a specific workflow
act -W .github/workflows/license-header-check.yml

# Run with verbose output
act -v -W .github/workflows/license-header-check.yml

# Run the mega-linter workflow
act -W .github/workflows/mega-linter.yml
```

### Expected License Header Format

All source files should include this header in the first 4 lines:

```python
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT
```

For Python files, it can also be in a docstring:

```python
"""
Copyright The Linux Foundation and each contributor to LFX.
SPDX-License-Identifier: MIT
"""
```

### File Types Checked

The file patterns defined in the [common GitHub Actions license header check
workflow](https://github.com/linuxfoundation/lfx-public-workflows/blob/main/.github/workflows/license-header-check.yml)
are checked.

### Excluded Patterns

The excluded file patterns are defined in the [common GitHub Actions license
header check
workflow](https://github.com/linuxfoundation/lfx-public-workflows/blob/main/.github/workflows/license-header-check.yml).

### Troubleshooting

If the check fails:

1. Add the required header to the beginning of the file
2. Make sure it appears in the first 4 lines
3. Use the exact text: "Copyright The Linux Foundation"
4. Include the SPDX identifier

Example fix for a Python file:

```python
#!/usr/bin/env python3
# Copyright The Linux Foundation and each contributor to LFX.
# SPDX-License-Identifier: MIT

# rest of your code...
```
