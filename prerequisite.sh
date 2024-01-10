#! /bin/bash

# Install pre-commit
brew install pre-commit
# OR , pip install pre-commit

# Install prerequisite packages

go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/cmd/goimports@latest
