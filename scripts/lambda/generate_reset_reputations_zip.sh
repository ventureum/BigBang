#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/reset_reputations/
GOOS=linux go build -o main
zip reset_reputations.zip main

aws lambda update-function-code \
  --function-name reset_reputations \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/reset_reputations/reset_reputations.zip \
  --publish

mv reset_reputations.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
