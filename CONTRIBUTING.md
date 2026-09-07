# Contributing to FlowForge

FlowForge follows production-oriented engineering practices even though it is primarily a portfolio project.

## Branching Strategy

The repository uses a lightweight GitHub Flow / trunk-oriented strategy.

The only permanent branch is:

```text
main
mkdir -p .github
cat > .github/pull_request_template.md <<'EOF'
## What

<!-- Briefly describe what this PR changes. -->

## Why

<!-- Explain the problem or requirement being addressed. -->

## Changes

- 
- 
- 

## Testing

- [ ] Unit tests
- [ ] Integration tests
- [ ] Manual verification
- [ ] Not applicable

## Notes

<!-- Architectural decisions, migrations, limitations, risks, or follow-up work. -->

Closes #
