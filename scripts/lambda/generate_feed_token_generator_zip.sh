#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_token_generator/
GOOS=linux go build -o main
zip feed_token_generator.zip main

aws lambda update-function-code \
  --function-name feed_token \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_token_generator/feed_token_generator.zip \
  --publish

mv feed_token_generator.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
