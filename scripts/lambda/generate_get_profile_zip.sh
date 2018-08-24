#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_profile/
GOOS=linux go build -o main
zip get_profile.zip main

aws lambda update-function-code \
  --function-name get_profile \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_profile/get_profile.zip \
  --publish

mv get_profile.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
