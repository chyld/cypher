#!/bin/bash

golint .
gofmt -w .
go build
