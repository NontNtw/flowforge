# FlowForge

FlowForge is a production-oriented workflow automation platform built to demonstrate backend engineering, reliability, distributed systems concepts, and eventually a polished full-stack product.

The project is inspired by workflow automation platforms such as n8n, Zapier, and Make, while intentionally keeping the implementation scope appropriate for a portfolio project.

## Goals

FlowForge will allow users to:

- Define versioned workflows.
- Trigger workflows through webhooks and APIs.
- Execute workflow nodes asynchronously.
- Call external HTTP APIs.
- Transform and branch data.
- Delay or schedule execution.
- Retry failed operations safely.
- Prevent duplicate processing.
- Inspect workflow and node execution history.

## Architecture

FlowForge starts as a modular monolith.

The HTTP API and background worker are separate processes while sharing domain and infrastructure packages within a single monorepo.

```text
HTTP API
   |
   v
Application / Domain
   |
   +------> PostgreSQL
   |
   +------> Transactional Outbox
                    |
                    v
              NATS JetStream
                    |
                    v
                  Worker

cat > .github/workflows/ci.yml <<'EOF'
name: CI

on:
  pull_request:
    branches:
      - main
  push:
    branches:
      - main

permissions:
  contents: read

concurrency:
  group: ci-${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  repository:
    name: Repository Validation
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Validate repository files
        run: |
          test -f README.md
          test -f CONTRIBUTING.md
          test -f Makefile
          test -f .gitignore
          test -f .github/pull_request_template.md

      - name: Validate repository structure
        run: |
          test -d apps/api
          test -d apps/worker
          test -d apps/web
          test -d internal
          test -d platform
          test -d migrations
          test -d sql/queries
          test -d sql/schema
          test -d api
          test -d deployments
          test -d docs/architecture
          test -d docs/adr
