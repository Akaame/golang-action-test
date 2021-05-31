# golang-action-test
Test Repository for trying out various Golang Github Actions

# Things I have done so far

Create the repository and a simple function with its tests
Go to various golang actions by official repo and the community and add them to this repository using the marketplace

https://github.com/actions/setup-go
https://github.com/mvdan/github-actions-golang

First pipeline failed because I had an error in yml
https://github.com/Akaame/golang-action-test/actions/runs/871607998/workflow

pre-commit and yaml-lint
https://yamllint.readthedocs.io/en/stable/integration.html#integration-with-pre-commit

```shell
pre-commit install
pre-commit sample-config > .pre-commit-config.yaml
```

```yml
-   repo: https://github.com/adrienverge/yamllint.git
    rev: v1.26.1
    hooks:
    -   id: yamllint
        args: [-c=/usr/local/bin/.yamllint]
```

## Add revive and revive config toml

Add Revive Lint step to Github Action Workflow
By simply getting revive binary and running `make lint`

## Cassandra Setup

https://github.com/marketplace/actions/setup-cassandra-action
https://github.com/gocql/gocql
https://pkg.go.dev/github.com/gocql/gocql#pkg-examples

# Taskfile Installation

https://taskfile.dev/#/
brew install go-task/tap/go-task

# Articles
https://presstige.io/p/Using-GitHub-Actions-with-Go-2ca9744b531f4f21bdae9976d1ccbb58

# What to do

- [x] Caching dependencies
- [x] Staticcheck
- [x] Running revive checks
- [ ] Running integration tests within pipeline
- [ ] Running workflows on tag pushes and generating artifacts for releases
- [ ] Creating docker images within actions
- [ ] Connecting to GCP Registry within Actions
- [ ] Using helm, tf, pulumi within pipelines
- [ ] Adding private repositories to goproxy settings
- [ ] Using buf build tool within github actions
- [ ] Using scripts within actions
- [x] Using makefiles within actions (Done in revive linting)
- [x] Using taskfiles within actions
- [x] Using ephemeral services/containers for testing
- [ ] Add codeql analysis
- [ ] Try using ko
- [ ] Add docker hub login and image upload
- [ ] Try using goreleaser
- [ ] Add CQL sanity checks
- [x] Add testcontainers for local testing

# Things to read about

vscode://vscode.github-authentication/
