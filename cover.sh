#!/usr/bin/env bash
GOEXPERIMENT=coverageredesign go test ./... -cover -coverpkg ./... -coverprofile cover.cov
#go tool cover -func cover.cov
#go tool cover -html cover.cov
