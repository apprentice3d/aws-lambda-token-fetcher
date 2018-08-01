#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o main
zip main.zip main
rm main

aws lambda create-function --function-name forge-token-fetcher \
                           --runtime go1.x \
                           --role arn:aws:iam::766924750857:role/token_fetcher-executor \
                           --handler main \
                           --zip-file fileb://main.zip

# check the created lambda function
aws lambda invoke --function-name forge-token-fetcher \
                  token.json

# pretty print the content of the lambda call output
less ./token.json | python -mjson.tool


