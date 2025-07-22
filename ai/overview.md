# AI Offerings Overview

## Table of Contents

- [Overview](#overview)
- [Available AI Providers](#available-ai-providers)
  - [Google Gemini (Corporate Account)](#google-gemini)
  - [AWS Bedrock (AWS Account)](#aws-bedrock)
  - [Cursor (Team Agreement)](#cursor-team-agreement)
  - [GitHub Copilot Enterprise](#github-copilot-enterprise)
  - [ChatGPT (Personal Account)](#chatgpt-personal-account)
  - [Anthropic Claude (Personal Account)](#anthropic-claude-personal-account)
- [Usage Summary](#usage-summary)

## Overview

This document provides an overview of the AI offerings available to the Linux Foundation engineering team, including
models, capabilities, and usage limits for each provider.

## Available AI Providers

### Google Gemini

**Account Type**: The Linux Foundation offers a corporate Google Workspace account for all LF staff and contractors.

**Available Models**:

1. Gemini 2.5 Pro - the current state-of-the-art thinking model, capable of
   reasoning over complex problems in code, math, and STEM, as well as analyzing
   large datasets, codebases, and documents using long context.
2. Gemini 2.5 Flash - best model in terms of price-performance, offering
   well-rounded capabilities. 2.5 Flash is best for large scale processing,
   low-latency, high volume tasks that require thinking, and agentic use cases.
3. Gemini 2.5 Flash Lite - optimized for cost efficiency and low latency.

A full comprehensive list of models is available [on Google's Gemini API
documentation website](https://ai.google.dev/gemini-api/docs/models).

**Usage Limits**:

- Rate limits are [documented on the Gemini API documentation
website](https://ai.google.dev/gemini-api/docs/rate-limits).

**Access Method**:

- [A Web interface is available for the Chat interface](https://gemini.google.com/)
- API access is available after setting up Google Project and API keys.
  Documentation for configuring this [is available on the Gemini API documentation
  website](https://ai.google.dev/gemini-api/docs/api-key).

### AWS Bedrock

**Account Type**: The Linux Foundation offers AWS Bedrock models through a
*corporate AWS account for all LF staff and contractors.

**Available Models**:

See our AWS Bedrock [documentation notes for the full list of
models](aws-bedrock-models.md).

**Usage Limits**:

AWS has a document that [describes the usage limits for the Bedrock
models](https://docs.aws.amazon.com/bedrock/latest/userguide/quotas.html).

**Access Method**:

- AWS SDKs (Python, Java, etc.)

### Cursor (Team Agreement)

**Account Type**: The Linux Foundation has a team agreement with Cursor for all
*LF staff and contractors.

**Available Models**:

Cursor publishes a list of models that are available to use [on their
website](https://docs.cursor.com/models).

**Usage Limits**:

- Usage limits are [documented on the Cursor
  website](https://docs.cursor.com/account/pricing) under the pricing plan
  section.
- The Linux Foundation has very limited $180.00 spending team limit per month.
- Number of seats: As of 2025-07-22, the Linux Foundation has 49 seats filled
  with typically around 35 active users.

**Access Method**:

- [Cursor IDE application](https://www.cursor.com/download)
- API [more information is available on their
  website](https://docs.cursor.com/settings/api-keys).

### GitHub Copilot Enterprise

**Account Type**: The Linux Foundation has an Enterprise Agreement with GitHub
for all LF staff and contractors.

**Available Models**:

A full list of models is available [on GitHub's
website](https://github.com/marketplace?type=models).

**Usage Limits**:

There are no known usage limits for the GitHub Copilot Enterprise models.

**Access Method**:

- For IDE integration and extensions (VS Code, JetBrains, Vim, etc.) use your GitHub
account credentials to authenticate.

### ChatGPT (Personal Account)

**Account Type**: Currently, the Linux Foundation does not have a team or
*enterprise agreement with OpenAPI. Some users have used their personal accounts
*to access the OpenAPI models.

**Available Models**:

A full list of models is available [on OpenAI's website](https://platform.openai.com/docs/models/chatgpt).

**Usage Limits**:

While the pricing tiers is abit convoluted, the usage limits are [documented on
OpenAI's pricing page](https://platform.openai.com/docs/pricing). The rate limits
defined [on OpenAI's website](https://platform.openai.com/docs/guides/rate-limits).

**Access Method**:

- [ChatGPT Web interface](https://chatgpt.com/)
- [API access](https://platform.openai.com/docs/api-reference)
- Standalone applications for Window and Mac are available. Additionally a number of browser add-ons are also available.
- OpenAI includes a mobile application for Android and Apple devices

### Anthropic Claude (Personal Account)

**Account Type**: Currently, the Linux Foundation does not have a team or
*enterprise agreement with Anthropic. Some users have used their personal accounts
*to access the Anthropic models.

**Available Models**:

A full list of models is available [on Anthropic's website](https://docs.anthropic.com/en/docs/about-claude/models/overview).

**Usage Limits**:

The Anthropic models are priced based (generally) on the number of tokens used.
An updated pricing list and limits can be found [on their
webiste](https://docs.anthropic.com/en/docs/about-claude/pricing).

**Access Method**:

- [Claude Chat Web interface](https://claude.ai)
- [Claude Code/CLI Quickstart Guide](https://docs.anthropic.com/en/docs/claude-code/quickstart)
- [Anthropic API](https://docs.anthropic.com/en/api/overview)

## Usage Summary

| Provider         | Account Type | Model Variety | Primary Use Case             |
|------------------|--------------|---------------|------------------------------|
| Google Gemini    | Corporate    | Medium        | General AI assistance        |
| AWS Bedrock      | Corporate    | High          | Scalable AI applications     |
| Cursor           | Team         | High          | AI-powered coding            |
| GitHub Copilot   | Enterprise   | Medium        | Code completion & assistance |
| ChatGPT          | Personal     | High          | General AI assistance        |
| Anthropic Claude | Personal     | Medium        | Advanced reasoning tasks     |
