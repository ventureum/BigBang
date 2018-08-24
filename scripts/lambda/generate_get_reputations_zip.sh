#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_reputations/
GOOS=linux go build -o main
zip get_reputations.zip main

aws lambda update-function-code \
  --function-name get_reputations_v2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_reputations/get_reputations.zip \
  --publish

mv get_reputations.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
