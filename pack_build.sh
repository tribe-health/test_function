#!/bin/bash

GOOS=linux GOARCH=amd64 pack build func-testfunction-go --builder openfunction/builder-go:v2.4.0-1.17 --env FUNC_NAME="TestFunction"  --env FUNC_CLEAR_SOURCE=true