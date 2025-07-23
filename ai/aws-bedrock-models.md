# AWS Bedrock Models

The LFX engineering teams supports a family of AI models via AWS Bedrock.

The following models are available as of 2025-07-15:

1. amazon.titan-embed-text-v2:0
2. amazon.nova-pro-v1:0
3. amazon.nova-premier-v1:0:8k
4. amazon.nova-premier-v1:0:20k
5. amazon.nova-premier-v1:0:1000k
6. amazon.nova-premier-v1:0:mm
7. amazon.nova-premier-v1:0
8. amazon.nova-lite-v1:0
9. amazon.nova-micro-v1:0
10. anthropic.claude-3-5-sonnet-20240620-v1:0
11. anthropic.claude-3-7-sonnet-20250219-v1:0
12. anthropic.claude-3-haiku-20240307-v1:0:200k
13. anthropic.claude-3-haiku-20240307-v1:0
14. anthropic.claude-3-5-sonnet-20241022-v2:0
15. anthropic.claude-3-5-haiku-20241022-v1:0
16. anthropic.claude-opus-4-20250514-v1:0
17. anthropic.claude-sonnet-4-20250514-v1:0
18. deepseek.r1-v1:0
19. mistral.pixtral-large-2502-v1:0
20. meta.llama3-1-8b-instruct-v1:0:128k
21. meta.llama3-1-8b-instruct-v1:0
22. meta.llama3-1-70b-instruct-v1:0:128k
23. meta.llama3-1-70b-instruct-v1:0
24. meta.llama3-1-405b-instruct-v1:0
25. meta.llama3-2-11b-instruct-v1:0
26. meta.llama3-2-90b-instruct-v1:0
27. meta.llama3-2-1b-instruct-v1:0
28. meta.llama3-2-3b-instruct-v1:0
29. meta.llama3-3-70b-instruct-v1:0
30. meta.llama4-scout-17b-instruct-v1:0
31. meta.llama4-maverick-17b-instruct-v1:0

To get a complete list of models, you can use the following command:

```bash
# Note: LFX Product (Production) AWS account
aws bedrock list-foundation-models --profile lfx-product-production | jq -r '.modelSummaries | .[] | .modelId'
```

## AWS Model Usage

To leverage the AWS Bedrock models, reach out to the CTO office for access and guidance.
