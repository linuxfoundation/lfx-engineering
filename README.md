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
