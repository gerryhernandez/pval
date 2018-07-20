#!/bin/bash
set -e
go run generate/generate.go > pval.go
go fmt pval.go
go vet pval.go
go build pval.go