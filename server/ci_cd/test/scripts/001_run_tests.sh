#!/bin/bash

export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin

go install github.com/jstemmer/go-junit-report/v2@latest

cd server

go test -v ./... | go-junit-report > report.xml

TEST_EXIT_CODE=${PIPESTATUS[0]}

exit $TEST_EXIT_CODE