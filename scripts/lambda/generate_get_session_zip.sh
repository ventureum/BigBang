#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_session/
GOOS=linux go build -o main
zip get_session.zip main

aws lambda update-function-code \
  --function-name get_session \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_session/get_session.zip \
  --publish

mv get_session.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
