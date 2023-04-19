#!/bin/bash

GOOS=linux GOARCH=amd64 pack build "tribehealth/test-function:v2" --builder openfunction/builder-go:v2.4.0-1.17 --env FUNC_NAME="TestFunction"  --env FUNC_CLEAR_SOURCE=true