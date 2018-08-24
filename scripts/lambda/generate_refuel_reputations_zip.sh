#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/refuel_reputations/
GOOS=linux go build -o main
zip refuel_reputations.zip main

aws lambda update-function-code \
  --function-name refuel_reputations_v2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/refuel_reputations/refuel_reputations.zip \
  --publish

mv refuel_reputations.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
