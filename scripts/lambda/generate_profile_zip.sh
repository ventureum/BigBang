#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/profile/
GOOS=linux go build -o main
zip profile.zip main

aws lambda update-function-code \
  --function-name profile \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/profile/profile.zip \
  --publish

mv profile.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
